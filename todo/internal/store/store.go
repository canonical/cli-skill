package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"todo/internal/model"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

type TodoFilter struct {
	State     string
	DueBefore *time.Time
	DueAfter  *time.Time
	Overdue   bool
}

type ScheduleFilter struct {
	TodoID string
	Kind   string
	Status string
	Target string
}

func DefaultDBPath(home string) string {
	return filepath.Join(home, ".local", "share", "todo", "todo.db")
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", "file:"+path)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	if err := s.initSchema(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) initSchema() error {
	const schema = `
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	due_at TEXT,
	state TEXT NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,
	closed_at TEXT,
	rejected_at TEXT
);

CREATE TABLE IF NOT EXISTS sinks (
	id TEXT PRIMARY KEY,
	url TEXT NOT NULL,
	events_json TEXT NOT NULL,
	enabled INTEGER NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS schedules (
	id TEXT PRIMARY KEY,
	todo_id INTEGER NOT NULL,
	kind TEXT NOT NULL,
	before_dur TEXT,
	every_dur TEXT,
	status TEXT NOT NULL,
	target_motd INTEGER NOT NULL,
	created_at TEXT NOT NULL,
	FOREIGN KEY(todo_id) REFERENCES todos(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS schedule_sinks (
	schedule_id TEXT NOT NULL,
	sink_id TEXT NOT NULL,
	PRIMARY KEY (schedule_id, sink_id),
	FOREIGN KEY(schedule_id) REFERENCES schedules(id) ON DELETE CASCADE,
	FOREIGN KEY(sink_id) REFERENCES sinks(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS deliveries (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	schedule_id TEXT NOT NULL,
	target_type TEXT NOT NULL,
	target_id TEXT NOT NULL,
	planned_at TEXT NOT NULL,
	sent_at TEXT,
	status TEXT NOT NULL,
	error TEXT,
	next_attempt_at TEXT,
	attempts INTEGER NOT NULL DEFAULT 0,
	UNIQUE(schedule_id, target_type, target_id, planned_at)
);

CREATE TABLE IF NOT EXISTS motd_messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	todo_id TEXT NOT NULL,
	schedule_id TEXT NOT NULL,
	message TEXT NOT NULL,
	created_at TEXT NOT NULL,
	shown INTEGER NOT NULL DEFAULT 0
);
`
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) CreateTodo(ctx context.Context, t model.Todo) (int64, error) {
	result, err := s.db.ExecContext(ctx, `
INSERT INTO todos(title, due_at, state, created_at, updated_at, closed_at, rejected_at)
VALUES(?, ?, ?, ?, ?, ?, ?)`,
		t.Title,
		timePtrToDB(t.DueAt),
		t.State,
		t.CreatedAt.Format(time.RFC3339),
		t.UpdatedAt.Format(time.RFC3339),
		timePtrToDB(t.ClosedAt),
		timePtrToDB(t.RejectedAt),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func (s *Store) UpdateTodo(ctx context.Context, id int64, title *string, dueAt *time.Time, clearDue bool) error {
	todo, err := s.GetTodo(ctx, id)
	if err != nil {
		return err
	}
	if title != nil {
		todo.Title = *title
	}
	if clearDue {
		todo.DueAt = nil
	} else if dueAt != nil {
		todo.DueAt = dueAt
	}
	todo.UpdatedAt = time.Now().UTC()
	_, err = s.db.ExecContext(ctx, `
UPDATE todos SET title=?, due_at=?, updated_at=? WHERE id=?`,
		todo.Title, timePtrToDB(todo.DueAt), todo.UpdatedAt.Format(time.RFC3339), id)
	return err
}

func (s *Store) TransitionTodo(ctx context.Context, id int64, state string) error {
	if !slices.Contains([]string{model.TodoStateClosed, model.TodoStateReopened, model.TodoStateRejected}, state) {
		return fmt.Errorf("unsupported state transition %q", state)
	}
	now := time.Now().UTC().Format(time.RFC3339)
	closedAt := any(nil)
	rejectedAt := any(nil)
	if state == model.TodoStateClosed {
		closedAt = now
	}
	if state == model.TodoStateRejected {
		rejectedAt = now
	}
	_, err := s.db.ExecContext(ctx, `
UPDATE todos
SET state=?, updated_at=?, closed_at=COALESCE(?, closed_at), rejected_at=COALESCE(?, rejected_at)
WHERE id=?`,
		state,
		now,
		closedAt,
		rejectedAt,
		id,
	)
	return err
}

func (s *Store) GetTodo(ctx context.Context, id int64) (model.Todo, error) {
	var row struct {
		ID         int64
		Title      string
		DueAt      sql.NullString
		State      string
		CreatedAt  string
		UpdatedAt  string
		ClosedAt   sql.NullString
		RejectedAt sql.NullString
	}
	err := s.db.QueryRowContext(ctx, `
SELECT id, title, due_at, state, created_at, updated_at, closed_at, rejected_at
FROM todos WHERE id=?`, id).Scan(
		&row.ID,
		&row.Title,
		&row.DueAt,
		&row.State,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.ClosedAt,
		&row.RejectedAt,
	)
	if err != nil {
		return model.Todo{}, err
	}
	return scanTodo(row.ID, row.Title, row.DueAt, row.State, row.CreatedAt, row.UpdatedAt, row.ClosedAt, row.RejectedAt)
}

func (s *Store) ListTodos(ctx context.Context, f TodoFilter) ([]model.Todo, error) {
	query := `SELECT id, title, due_at, state, created_at, updated_at, closed_at, rejected_at FROM todos WHERE 1=1`
	args := make([]any, 0)
	if f.State != "" {
		query += ` AND state = ?`
		args = append(args, f.State)
	}
	if f.DueBefore != nil {
		query += ` AND due_at IS NOT NULL AND due_at < ?`
		args = append(args, f.DueBefore.UTC().Format(time.RFC3339))
	}
	if f.DueAfter != nil {
		query += ` AND due_at IS NOT NULL AND due_at > ?`
		args = append(args, f.DueAfter.UTC().Format(time.RFC3339))
	}
	if f.Overdue {
		query += ` AND due_at IS NOT NULL AND due_at < ?`
		args = append(args, time.Now().UTC().Format(time.RFC3339))
	}
	query += ` ORDER BY created_at DESC`

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Todo, 0)
	for rows.Next() {
		var id int64
		var title, state, createdAt, updatedAt string
		var dueAt, closedAt, rejectedAt sql.NullString
		if err := rows.Scan(&id, &title, &dueAt, &state, &createdAt, &updatedAt, &closedAt, &rejectedAt); err != nil {
			return nil, err
		}
		t, err := scanTodo(id, title, dueAt, state, createdAt, updatedAt, closedAt, rejectedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}

func (s *Store) CreateSink(ctx context.Context, sink model.Sink) error {
	eventsJSON, _ := json.Marshal(sink.Events)
	_, err := s.db.ExecContext(ctx, `
INSERT INTO sinks(id, url, events_json, enabled, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?)`,
		sink.ID,
		sink.URL,
		string(eventsJSON),
		boolToInt(sink.Enabled),
		sink.CreatedAt.UTC().Format(time.RFC3339),
		sink.UpdatedAt.UTC().Format(time.RFC3339),
	)
	return err
}

func (s *Store) DeleteSink(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM sinks WHERE id=?`, id)
	return err
}

func (s *Store) GetSink(ctx context.Context, id string) (model.Sink, error) {
	var sink model.Sink
	var eventsJSON string
	var enabled int
	var createdAt, updatedAt string
	err := s.db.QueryRowContext(ctx, `
SELECT id, url, events_json, enabled, created_at, updated_at FROM sinks WHERE id=?`, id).Scan(
		&sink.ID,
		&sink.URL,
		&eventsJSON,
		&enabled,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return model.Sink{}, err
	}
	if err := json.Unmarshal([]byte(eventsJSON), &sink.Events); err != nil {
		return model.Sink{}, err
	}
	sink.Enabled = enabled == 1
	if sink.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
		return model.Sink{}, err
	}
	if sink.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt); err != nil {
		return model.Sink{}, err
	}
	return sink, nil
}

func (s *Store) ListSinks(ctx context.Context, enabled *bool, event string) ([]model.Sink, error) {
	query := `SELECT id, url, events_json, enabled, created_at, updated_at FROM sinks WHERE 1=1`
	args := make([]any, 0)
	if enabled != nil {
		query += ` AND enabled=?`
		args = append(args, boolToInt(*enabled))
	}
	query += ` ORDER BY created_at DESC`
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.Sink, 0)
	for rows.Next() {
		var sID, url, eventsJSON, createdAt, updatedAt string
		var en int
		if err := rows.Scan(&sID, &url, &eventsJSON, &en, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		sink := model.Sink{ID: sID, URL: url, Enabled: en == 1}
		if err := json.Unmarshal([]byte(eventsJSON), &sink.Events); err != nil {
			return nil, err
		}
		if event != "" && !slices.Contains(sink.Events, event) {
			continue
		}
		if sink.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
			return nil, err
		}
		if sink.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt); err != nil {
			return nil, err
		}
		out = append(out, sink)
	}
	return out, rows.Err()
}

func (s *Store) CreateSchedule(ctx context.Context, sc model.Schedule) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, `
INSERT INTO schedules(id, todo_id, kind, before_dur, every_dur, status, target_motd, created_at)
VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		sc.ID,
		sc.TodoID,
		sc.Kind,
		nullIfEmpty(sc.Before),
		nullIfEmpty(sc.Every),
		sc.Status,
		boolToInt(sc.TargetMOTD),
		sc.CreatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return err
	}
	for _, sinkID := range sc.SinkIDs {
		_, err = tx.ExecContext(ctx, `INSERT INTO schedule_sinks(schedule_id, sink_id) VALUES(?, ?)`, sc.ID, sinkID)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}

func (s *Store) DeleteSchedule(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM schedules WHERE id=?`, id)
	return err
}

func (s *Store) GetSchedule(ctx context.Context, id string) (model.Schedule, error) {
	var sc model.Schedule
	var before, every sql.NullString
	var targetMOTD int
	var createdAt string
	err := s.db.QueryRowContext(ctx, `
SELECT id, todo_id, kind, before_dur, every_dur, status, target_motd, created_at
FROM schedules WHERE id=?`, id).Scan(
		&sc.ID,
		&sc.TodoID,
		&sc.Kind,
		&before,
		&every,
		&sc.Status,
		&targetMOTD,
		&createdAt,
	)
	if err != nil {
		return model.Schedule{}, err
	}
	if before.Valid {
		sc.Before = before.String
	}
	if every.Valid {
		sc.Every = every.String
	}
	sc.TargetMOTD = targetMOTD == 1
	if sc.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
		return model.Schedule{}, err
	}
	sc.SinkIDs, err = s.listScheduleSinks(ctx, sc.ID)
	if err != nil {
		return model.Schedule{}, err
	}
	return sc, nil
}

func (s *Store) ListSchedules(ctx context.Context, f ScheduleFilter) ([]model.Schedule, error) {
	query := `SELECT id, todo_id, kind, before_dur, every_dur, status, target_motd, created_at FROM schedules WHERE 1=1`
	args := make([]any, 0)
	if f.TodoID != "" {
		query += ` AND todo_id=?`
		args = append(args, f.TodoID)
	}
	if f.Kind != "" {
		query += ` AND kind=?`
		args = append(args, f.Kind)
	}
	if f.Status != "" {
		query += ` AND status=?`
		args = append(args, f.Status)
	}
	if f.Target == "motd" {
		query += ` AND target_motd=1`
	}
	if f.Target == "sink" {
		query += ` AND id IN (SELECT DISTINCT schedule_id FROM schedule_sinks)`
	}
	query += ` ORDER BY created_at DESC`

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.Schedule, 0)
	for rows.Next() {
		var id, todoID, kind, status, createdAt string
		var before, every sql.NullString
		var targetMOTD int
		if err := rows.Scan(&id, &todoID, &kind, &before, &every, &status, &targetMOTD, &createdAt); err != nil {
			return nil, err
		}
		sc := model.Schedule{ID: id, TodoID: todoID, Kind: kind, Status: status, TargetMOTD: targetMOTD == 1}
		if before.Valid {
			sc.Before = before.String
		}
		if every.Valid {
			sc.Every = every.String
		}
		if sc.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
			return nil, err
		}
		sc.SinkIDs, err = s.listScheduleSinks(ctx, id)
		if err != nil {
			return nil, err
		}
		out = append(out, sc)
	}
	return out, rows.Err()
}

func (s *Store) listScheduleSinks(ctx context.Context, scheduleID string) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT sink_id FROM schedule_sinks WHERE schedule_id=?`, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]string, 0)
	for rows.Next() {
		var sinkID string
		if err := rows.Scan(&sinkID); err != nil {
			return nil, err
		}
		out = append(out, sinkID)
	}
	return out, rows.Err()
}

func (s *Store) QueueMOTDMessage(ctx context.Context, todoID, scheduleID, message string) error {
	_, err := s.db.ExecContext(ctx, `
INSERT INTO motd_messages(todo_id, schedule_id, message, created_at, shown)
VALUES(?, ?, ?, ?, 0)
`, todoID, scheduleID, message, time.Now().UTC().Format(time.RFC3339))
	return err
}

func (s *Store) PullMOTDMessages(ctx context.Context) ([]string, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, `SELECT id, message FROM motd_messages WHERE shown=0 ORDER BY created_at ASC`)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	ids := make([]int64, 0)
	msgs := make([]string, 0)
	for rows.Next() {
		var id int64
		var m string
		if err := rows.Scan(&id, &m); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		ids = append(ids, id)
		msgs = append(msgs, m)
	}
	if err := rows.Err(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	for _, id := range ids {
		if _, err := tx.ExecContext(ctx, `UPDATE motd_messages SET shown=1 WHERE id=?`, id); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return msgs, nil
}

type DeliveryRecord struct {
	ScheduleID    string
	TargetType    string
	TargetID      string
	PlannedAt     time.Time
	SentAt        *time.Time
	Status        string
	Error         string
	NextAttemptAt *time.Time
	Attempts      int
}

func (s *Store) GetDelivery(ctx context.Context, scheduleID, targetType, targetID string, plannedAt time.Time) (DeliveryRecord, error) {
	var rec DeliveryRecord
	var sentAt, nextAttemptAt sql.NullString
	err := s.db.QueryRowContext(ctx, `
SELECT schedule_id, target_type, target_id, planned_at, sent_at, status, error, next_attempt_at, attempts
FROM deliveries WHERE schedule_id=? AND target_type=? AND target_id=? AND planned_at=?`,
		scheduleID,
		targetType,
		targetID,
		plannedAt.UTC().Format(time.RFC3339),
	).Scan(
		&rec.ScheduleID,
		&rec.TargetType,
		&rec.TargetID,
		&rec.PlannedAt,
		&sentAt,
		&rec.Status,
		&rec.Error,
		&nextAttemptAt,
		&rec.Attempts,
	)
	if err != nil {
		return DeliveryRecord{}, err
	}
	if sentAt.Valid {
		t, err := time.Parse(time.RFC3339, sentAt.String)
		if err == nil {
			rec.SentAt = &t
		}
	}
	if nextAttemptAt.Valid {
		t, err := time.Parse(time.RFC3339, nextAttemptAt.String)
		if err == nil {
			rec.NextAttemptAt = &t
		}
	}
	return rec, nil
}

func (s *Store) UpsertDeliveryResult(ctx context.Context, scheduleID, targetType, targetID string, plannedAt time.Time, success bool, errText string) error {
	status := "failed"
	var sentAt any
	var nextAttempt any
	attempts := 1
	if success {
		status = "sent"
		now := time.Now().UTC()
		sentAt = now.Format(time.RFC3339)
		nextAttempt = nil
		errText = ""
	} else {
		prev, err := s.GetDelivery(ctx, scheduleID, targetType, targetID, plannedAt)
		if err == nil {
			attempts = prev.Attempts + 1
		}
		backoff := time.Minute * time.Duration(minInt(60, 1<<minInt(6, attempts)))
		nextAttempt = time.Now().UTC().Add(backoff).Format(time.RFC3339)
	}

	_, err := s.db.ExecContext(ctx, `
INSERT INTO deliveries(schedule_id, target_type, target_id, planned_at, sent_at, status, error, next_attempt_at, attempts)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(schedule_id, target_type, target_id, planned_at)
DO UPDATE SET sent_at=excluded.sent_at, status=excluded.status, error=excluded.error, next_attempt_at=excluded.next_attempt_at, attempts=excluded.attempts
`, scheduleID, targetType, targetID, plannedAt.UTC().Format(time.RFC3339), sentAt, status, errText, nextAttempt, attempts)
	return err
}

func (s *Store) IsTargetDelivered(ctx context.Context, scheduleID, targetType, targetID string, plannedAt time.Time) (bool, error) {
	var count int
	err := s.db.QueryRowContext(ctx, `
SELECT COUNT(1) FROM deliveries WHERE schedule_id=? AND target_type=? AND target_id=? AND planned_at=? AND status='sent'`,
		scheduleID, targetType, targetID, plannedAt.UTC().Format(time.RFC3339)).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *Store) ShouldAttemptDelivery(ctx context.Context, scheduleID, targetType, targetID string, plannedAt time.Time) (bool, error) {
	var status string
	var nextAttemptAt sql.NullString
	err := s.db.QueryRowContext(ctx, `
SELECT status, next_attempt_at FROM deliveries WHERE schedule_id=? AND target_type=? AND target_id=? AND planned_at=?`,
		scheduleID, targetType, targetID, plannedAt.UTC().Format(time.RFC3339)).Scan(&status, &nextAttemptAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}
	if status == "sent" {
		return false, nil
	}
	if !nextAttemptAt.Valid {
		return true, nil
	}
	nextAt, err := time.Parse(time.RFC3339, nextAttemptAt.String)
	if err != nil {
		return true, nil
	}
	return !time.Now().UTC().Before(nextAt), nil
}

func (s *Store) MarkScheduleSent(ctx context.Context, scheduleID string) error {
	_, err := s.db.ExecContext(ctx, `UPDATE schedules SET status=? WHERE id=?`, model.ScheduleStatusSent, scheduleID)
	return err
}

func (s *Store) listActiveSchedulesWithTodos(ctx context.Context) ([]model.Schedule, []model.Todo, error) {
	rows, err := s.db.QueryContext(ctx, `
SELECT s.id, s.todo_id, s.kind, s.before_dur, s.every_dur, s.status, s.target_motd, s.created_at,
       t.id, t.title, t.due_at, t.state, t.created_at, t.updated_at, t.closed_at, t.rejected_at
FROM schedules s
JOIN todos t ON t.id=s.todo_id
WHERE s.status=? AND t.due_at IS NOT NULL
`, model.ScheduleStatusActive)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	schedules := make([]model.Schedule, 0)
	todos := make([]model.Todo, 0)
	for rows.Next() {
		var sc model.Schedule
		var before, every sql.NullString
		var targetMOTD int
		var createdAt string
		var todoID int64
		var title, state, todoCreated, todoUpdated string
		var dueAt, closedAt, rejectedAt sql.NullString
		if err := rows.Scan(
			&sc.ID, &sc.TodoID, &sc.Kind, &before, &every, &sc.Status, &targetMOTD, &createdAt,
			&todoID, &title, &dueAt, &state, &todoCreated, &todoUpdated, &closedAt, &rejectedAt,
		); err != nil {
			return nil, nil, err
		}
		if before.Valid {
			sc.Before = before.String
		}
		if every.Valid {
			sc.Every = every.String
		}
		sc.TargetMOTD = targetMOTD == 1
		if sc.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
			return nil, nil, err
		}
		sc.SinkIDs, err = s.listScheduleSinks(ctx, sc.ID)
		if err != nil {
			return nil, nil, err
		}
		todo, err := scanTodo(todoID, title, dueAt, state, todoCreated, todoUpdated, closedAt, rejectedAt)
		if err != nil {
			return nil, nil, err
		}
		schedules = append(schedules, sc)
		todos = append(todos, todo)
	}
	return schedules, todos, rows.Err()
}

func (s *Store) ActiveSchedulesWithTodos(ctx context.Context) ([]model.Schedule, []model.Todo, error) {
	return s.listActiveSchedulesWithTodos(ctx)
}

func scanTodo(id int64, title string, dueAt sql.NullString, state, createdAt, updatedAt string, closedAt, rejectedAt sql.NullString) (model.Todo, error) {
	var t model.Todo
	t.ID = id
	t.Title = title
	t.State = state
	var err error
	if t.CreatedAt, err = time.Parse(time.RFC3339, createdAt); err != nil {
		return model.Todo{}, err
	}
	if t.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt); err != nil {
		return model.Todo{}, err
	}
	if dueAt.Valid {
		d, err := time.Parse(time.RFC3339, dueAt.String)
		if err != nil {
			return model.Todo{}, err
		}
		t.DueAt = &d
	}
	if closedAt.Valid {
		d, err := time.Parse(time.RFC3339, closedAt.String)
		if err != nil {
			return model.Todo{}, err
		}
		t.ClosedAt = &d
	}
	if rejectedAt.Valid {
		d, err := time.Parse(time.RFC3339, rejectedAt.String)
		if err != nil {
			return model.Todo{}, err
		}
		t.RejectedAt = &d
	}
	return t, nil
}

func timePtrToDB(t *time.Time) any {
	if t == nil {
		return nil
	}
	return t.UTC().Format(time.RFC3339)
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func nullIfEmpty(v string) any {
	if strings.TrimSpace(v) == "" {
		return nil
	}
	return v
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
