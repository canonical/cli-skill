package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"todo/internal/common"
	"todo/internal/model"
	"todo/internal/store"
)

type Service struct {
	store *store.Store
}

func New(s *store.Store) *Service {
	return &Service{store: s}
}

type CreateTodoRequest struct {
	ID        string
	Title     string
	DueAt     *time.Time
	Schedules []ScheduleSpec
}

type UpdateTodoRequest struct {
	Title    *string
	DueAt    *time.Time
	ClearDue bool
}

type CreateSinkRequest struct {
	ID      string
	URL     string
	Events  []string
	Enabled bool
}

type ScheduleSpec struct {
	ID       string
	TodoID   string
	Kind     string
	Before   string
	Every    string
	SinkIDs  []string
	UseMOTD  bool
	Implicit bool
}

type StatusResponse struct {
	Now                      time.Time `json:"now"`
	DatabasePath             string    `json:"database_path"`
	NeedsMOTDLoginScriptHint bool      `json:"needs_motd_login_script_hint"`
	MOTDLoginScriptHint      string    `json:"motd_login_script_hint"`
	MOTDLoginScriptCheckPath string    `json:"motd_login_script_check_path"`
	ActiveTodoCount          int       `json:"active_todo_count"`
	ActiveScheduleCount      int       `json:"active_schedule_count"`
	EnabledSinkCount         int       `json:"enabled_sink_count"`
}

func (s *Service) CreateTodo(ctx context.Context, req CreateTodoRequest) (model.Todo, error) {
	if strings.TrimSpace(req.Title) == "" {
		return model.Todo{}, fmt.Errorf("title is required")
	}
	id := strings.TrimSpace(req.ID)
	var todoID int64
	if id == "" {
		// Create todo and server generates ID
		now := time.Now().UTC()
		todo := model.Todo{
			Title:     strings.TrimSpace(req.Title),
			DueAt:     req.DueAt,
			State:     model.TodoStateOpen,
			CreatedAt: now,
			UpdatedAt: now,
		}
		var err error
		todoID, err = s.store.CreateTodo(ctx, todo)
		if err != nil {
			return model.Todo{}, err
		}
	} else {
		// User provided explicit ID - not supported with int64
		return model.Todo{}, fmt.Errorf("explicit todo ids are no longer supported")
	}
	todo, err := s.store.GetTodo(ctx, todoID)
	if err != nil {
		return model.Todo{}, err
	}
	for i := range req.Schedules {
		sched := req.Schedules[i]
		sched.TodoID = fmt.Sprintf("%d", todo.ID)
		if _, err := s.AddSchedule(ctx, sched); err != nil {
			return model.Todo{}, err
		}
	}
	return todo, nil
}

func (s *Service) UpdateTodo(ctx context.Context, idStr string, req UpdateTodoRequest) (model.Todo, error) {
	if strings.TrimSpace(idStr) == "" {
		return model.Todo{}, fmt.Errorf("todo id is required")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.Todo{}, fmt.Errorf("invalid todo id")
	}
	if req.Title == nil && req.DueAt == nil && !req.ClearDue {
		return model.Todo{}, fmt.Errorf("nothing to update")
	}
	if err := s.store.UpdateTodo(ctx, id, req.Title, req.DueAt, req.ClearDue); err != nil {
		return model.Todo{}, err
	}
	return s.store.GetTodo(ctx, id)
}

func (s *Service) CloseTodo(ctx context.Context, idStr string) (model.Todo, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.Todo{}, fmt.Errorf("invalid todo id")
	}
	if err := s.store.TransitionTodo(ctx, id, model.TodoStateClosed); err != nil {
		return model.Todo{}, err
	}
	return s.store.GetTodo(ctx, id)
}

func (s *Service) ReopenTodo(ctx context.Context, idStr string) (model.Todo, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.Todo{}, fmt.Errorf("invalid todo id")
	}
	if err := s.store.TransitionTodo(ctx, id, model.TodoStateReopened); err != nil {
		return model.Todo{}, err
	}
	return s.store.GetTodo(ctx, id)
}

func (s *Service) RejectTodo(ctx context.Context, idStr string) (model.Todo, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.Todo{}, fmt.Errorf("invalid todo id")
	}
	if err := s.store.TransitionTodo(ctx, id, model.TodoStateRejected); err != nil {
		return model.Todo{}, err
	}
	return s.store.GetTodo(ctx, id)
}

func (s *Service) ShowTodo(ctx context.Context, idStr string) (model.Todo, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.Todo{}, fmt.Errorf("invalid todo id")
	}
	return s.store.GetTodo(ctx, id)
}

func (s *Service) ListTodos(ctx context.Context, f store.TodoFilter) ([]model.Todo, error) {
	return s.store.ListTodos(ctx, f)
}

func normalizeEvents(events []string) []string {
	if len(events) == 0 {
		return []string{model.ScheduleKindUpcoming, model.ScheduleKindOverdue}
	}
	out := make([]string, 0, len(events))
	for _, ev := range events {
		ev = strings.TrimSpace(strings.ToLower(ev))
		if ev == "" {
			continue
		}
		if ev != model.ScheduleKindUpcoming && ev != model.ScheduleKindOverdue {
			continue
		}
		if !slices.Contains(out, ev) {
			out = append(out, ev)
		}
	}
	if len(out) == 0 {
		return []string{model.ScheduleKindUpcoming, model.ScheduleKindOverdue}
	}
	return out
}

func (s *Service) CreateSink(ctx context.Context, req CreateSinkRequest) (model.Sink, error) {
	if strings.TrimSpace(req.ID) == "" {
		return model.Sink{}, fmt.Errorf("sink id is required")
	}
	if strings.TrimSpace(req.URL) == "" {
		return model.Sink{}, fmt.Errorf("sink url is required")
	}
	now := time.Now().UTC()
	sink := model.Sink{
		ID:        req.ID,
		URL:       req.URL,
		Events:    normalizeEvents(req.Events),
		Enabled:   req.Enabled,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.store.CreateSink(ctx, sink); err != nil {
		return model.Sink{}, err
	}
	return s.store.GetSink(ctx, sink.ID)
}

func (s *Service) DeleteSink(ctx context.Context, id string) error {
	return s.store.DeleteSink(ctx, id)
}

func (s *Service) ShowSink(ctx context.Context, id string) (model.Sink, error) {
	return s.store.GetSink(ctx, id)
}

func (s *Service) ListSinks(ctx context.Context, enabled *bool, event string) ([]model.Sink, error) {
	return s.store.ListSinks(ctx, enabled, strings.TrimSpace(strings.ToLower(event)))
}

func (s *Service) AddSchedule(ctx context.Context, req ScheduleSpec) (model.Schedule, error) {
	if strings.TrimSpace(req.ID) == "" {
		return model.Schedule{}, fmt.Errorf("schedule id is required")
	}
	if strings.TrimSpace(req.TodoID) == "" {
		return model.Schedule{}, fmt.Errorf("--todo is required")
	}
	if req.Kind != model.ScheduleKindUpcoming && req.Kind != model.ScheduleKindOverdue {
		return model.Schedule{}, fmt.Errorf("--kind must be upcoming or overdue")
	}
	if req.Kind == model.ScheduleKindUpcoming {
		if strings.TrimSpace(req.Before) == "" {
			req.Before = "24h"
		}
		if _, err := common.ParseHumanDuration(req.Before); err != nil {
			return model.Schedule{}, fmt.Errorf("invalid --before duration: %w", err)
		}
		if req.Every != "" {
			return model.Schedule{}, fmt.Errorf("--every is not valid for upcoming schedules")
		}
	} else {
		if strings.TrimSpace(req.Every) == "" {
			req.Every = "24h"
		}
		if _, err := common.ParseHumanDuration(req.Every); err != nil {
			return model.Schedule{}, fmt.Errorf("invalid --every duration: %w", err)
		}
		if req.Before != "" {
			return model.Schedule{}, fmt.Errorf("--before is not valid for overdue schedules")
		}
	}
	todoID, err := strconv.ParseInt(req.TodoID, 10, 64)
	if err != nil {
		return model.Schedule{}, fmt.Errorf("invalid todo id")
	}
	todo, err := s.store.GetTodo(ctx, todoID)
	if err != nil {
		return model.Schedule{}, err
	}
	if todo.DueAt == nil {
		return model.Schedule{}, fmt.Errorf("todo %q has no due date; schedules require a due date", req.TodoID)
	}
	if !req.UseMOTD && len(req.SinkIDs) == 0 {
		req.UseMOTD = true
	}
	sc := model.Schedule{
		ID:         req.ID,
		TodoID:     req.TodoID,
		Kind:       req.Kind,
		Before:     req.Before,
		Every:      req.Every,
		Status:     model.ScheduleStatusActive,
		TargetMOTD: req.UseMOTD,
		SinkIDs:    dedupe(req.SinkIDs),
		CreatedAt:  time.Now().UTC(),
	}
	if err := s.store.CreateSchedule(ctx, sc); err != nil {
		return model.Schedule{}, err
	}
	return s.store.GetSchedule(ctx, sc.ID)
}

func (s *Service) RemoveSchedule(ctx context.Context, id string) error {
	return s.store.DeleteSchedule(ctx, id)
}

func (s *Service) ShowSchedule(ctx context.Context, id string) (model.Schedule, error) {
	return s.store.GetSchedule(ctx, id)
}

func (s *Service) ListSchedules(ctx context.Context, f store.ScheduleFilter) ([]model.Schedule, error) {
	return s.store.ListSchedules(ctx, f)
}

func (s *Service) MOTDMessage(ctx context.Context) ([]string, error) {
	return s.store.PullMOTDMessages(ctx)
}

func (s *Service) Status(ctx context.Context, dbPath string) (StatusResponse, error) {
	todos, err := s.store.ListTodos(ctx, store.TodoFilter{State: model.TodoStateOpen})
	if err != nil {
		return StatusResponse{}, err
	}
	schedules, err := s.store.ListSchedules(ctx, store.ScheduleFilter{Status: model.ScheduleStatusActive})
	if err != nil {
		return StatusResponse{}, err
	}
	trueVal := true
	sinks, err := s.store.ListSinks(ctx, &trueVal, "")
	if err != nil {
		return StatusResponse{}, err
	}
	home, _ := os.UserHomeDir()
	profile := filepath.Join(home, ".profile")
	needsHint := !profileHasMotdMessage(profile)
	return StatusResponse{
		Now:                      time.Now().UTC(),
		DatabasePath:             dbPath,
		NeedsMOTDLoginScriptHint: needsHint,
		MOTDLoginScriptHint:      "To show todo reminders on login, run: echo 'todo motd-message' >> ~/.profile",
		MOTDLoginScriptCheckPath: profile,
		ActiveTodoCount:          len(todos),
		ActiveScheduleCount:      len(schedules),
		EnabledSinkCount:         len(sinks),
	}, nil
}

func profileHasMotdMessage(path string) bool {
	b, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return bytes.Contains(bytes.ToLower(b), []byte("todo motd-message"))
}

func dedupe(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, v := range in {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func computeSchedulePlannedAt(sc model.Schedule, todo model.Todo, now time.Time) (time.Time, bool, error) {
	if todo.DueAt == nil {
		return time.Time{}, false, nil
	}
	due := todo.DueAt.UTC()
	if sc.Kind == model.ScheduleKindUpcoming {
		d, err := common.ParseHumanDuration(sc.Before)
		if err != nil {
			return time.Time{}, false, err
		}
		planned := due.Add(-d)
		if now.Before(planned) {
			return planned, false, nil
		}
		return planned, true, nil
	}
	freq, err := common.ParseHumanDuration(sc.Every)
	if err != nil {
		return time.Time{}, false, err
	}
	if now.Before(due) {
		return time.Time{}, false, nil
	}
	elapsed := now.Sub(due)
	steps := int(elapsed / freq)
	planned := due.Add(time.Duration(steps) * freq)
	return planned, true, nil
}

func (s *Service) EvaluateAndDispatchSchedules(ctx context.Context) error {
	now := time.Now().UTC()
	schedules, todos, err := s.store.ActiveSchedulesWithTodos(ctx)
	if err != nil {
		return err
	}
	for i, sc := range schedules {
		todo := todos[i]
		plannedAt, dueNow, err := computeSchedulePlannedAt(sc, todo, now)
		if err != nil || !dueNow {
			continue
		}

		totalTargets := 0
		deliveredTargets := 0

		if sc.TargetMOTD {
			totalTargets++
			ok, err := s.store.IsTargetDelivered(ctx, sc.ID, "motd", "local", plannedAt)
			if err != nil {
				continue
			}
			if ok {
				deliveredTargets++
			} else {
				shouldAttempt, err := s.store.ShouldAttemptDelivery(ctx, sc.ID, "motd", "local", plannedAt)
				if err != nil {
					continue
				}
				if shouldAttempt {
					msg := fmt.Sprintf("[%s] %s (due %s)", strings.ToUpper(sc.Kind), todo.Title, todo.DueAt.Local().Format("2006-01-02 15:04 MST"))
					err = s.store.QueueMOTDMessage(ctx, fmt.Sprintf("%d", todo.ID), sc.ID, msg)
					if err != nil {
						_ = s.store.UpsertDeliveryResult(ctx, sc.ID, "motd", "local", plannedAt, false, err.Error())
					} else {
						_ = s.store.UpsertDeliveryResult(ctx, sc.ID, "motd", "local", plannedAt, true, "")
						deliveredTargets++
					}
				}
			}
		}

		for _, sinkID := range sc.SinkIDs {
			totalTargets++
			ok, err := s.store.IsTargetDelivered(ctx, sc.ID, "sink", sinkID, plannedAt)
			if err != nil {
				continue
			}
			if ok {
				deliveredTargets++
				continue
			}
			shouldAttempt, err := s.store.ShouldAttemptDelivery(ctx, sc.ID, "sink", sinkID, plannedAt)
			if err != nil || !shouldAttempt {
				continue
			}
			sink, err := s.store.GetSink(ctx, sinkID)
			if err != nil || !sink.Enabled {
				continue
			}
			if len(sink.Events) > 0 && !slices.Contains(sink.Events, sc.Kind) {
				continue
			}
			payload := fmt.Sprintf(`{"todo_id":%q,"title":%q,"kind":%q,"due_at":%q,"planned_at":%q}`,
				todo.ID,
				todo.Title,
				sc.Kind,
				todo.DueAt.UTC().Format(time.RFC3339),
				plannedAt.UTC().Format(time.RFC3339),
			)
			err = postWebhook(ctx, sink.URL, payload)
			if err != nil {
				_ = s.store.UpsertDeliveryResult(ctx, sc.ID, "sink", sinkID, plannedAt, false, err.Error())
			} else {
				_ = s.store.UpsertDeliveryResult(ctx, sc.ID, "sink", sinkID, plannedAt, true, "")
				deliveredTargets++
			}
		}

		if sc.Kind == model.ScheduleKindUpcoming && totalTargets > 0 && deliveredTargets == totalTargets {
			_ = s.store.MarkScheduleSent(ctx, sc.ID)
		}
	}
	return nil
}

func postWebhook(ctx context.Context, url, payload string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("webhook returned non-2xx")
	}
	return nil
}
