package daemon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"

	"todo/internal/common"
	"todo/internal/model"
	"todo/internal/store"
)

type Service struct {
	store  *store.Store
	client *http.Client
}

func NewService(st *store.Store) *Service {
	return &Service{
		store: st,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type CreateTodoRequest struct {
	ID       string         `json:"id,omitempty"`
	Title    string         `json:"title"`
	DueAt    *time.Time     `json:"due_at,omitempty"`
	Schedule []ScheduleSpec `json:"schedule,omitempty"`
}

type UpdateTodoRequest struct {
	Title    *string    `json:"title,omitempty"`
	DueAt    *time.Time `json:"due_at,omitempty"`
	ClearDue bool       `json:"clear_due"`
}

type CreateSinkRequest struct {
	ID     string   `json:"id"`
	URL    string   `json:"url"`
	Events []string `json:"events,omitempty"`
}

type ScheduleSpec struct {
	Kind   string   `json:"kind"`
	Before string   `json:"before,omitempty"`
	Every  string   `json:"every,omitempty"`
	MOTD   bool     `json:"motd"`
	SinkID []string `json:"sink_id,omitempty"`
}

type AddScheduleRequest struct {
	ID     string   `json:"id"`
	TodoID string   `json:"todo_id"`
	Kind   string   `json:"kind"`
	Before string   `json:"before,omitempty"`
	Every  string   `json:"every,omitempty"`
	MOTD   bool     `json:"motd"`
	SinkID []string `json:"sink_id,omitempty"`
}

func (s *Service) CreateTodo(ctx context.Context, req CreateTodoRequest) (model.Todo, error) {
	if strings.TrimSpace(req.Title) == "" {
		return model.Todo{}, fmt.Errorf("title is required")
	}
	id := strings.TrimSpace(req.ID)
	if id == "" {
		id = uuid.NewString()
	}
	now := time.Now().UTC()
	todo := model.Todo{
		ID:        id,
		Title:     req.Title,
		DueAt:     req.DueAt,
		State:     model.TodoStateOpen,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.store.CreateTodo(ctx, todo); err != nil {
		return model.Todo{}, err
	}
	for i := range req.Schedule {
		sp := req.Schedule[i]
		scReq := AddScheduleRequest{
			ID:     fmt.Sprintf("%s-sch-%d", todo.ID, i+1),
			TodoID: todo.ID,
			Kind:   sp.Kind,
			Before: sp.Before,
			Every:  sp.Every,
			MOTD:   sp.MOTD,
			SinkID: sp.SinkID,
		}
		if _, err := s.AddSchedule(ctx, scReq); err != nil {
			return model.Todo{}, err
		}
	}
	return todo, nil
}

func (s *Service) UpdateTodo(ctx context.Context, id string, req UpdateTodoRequest) error {
	if req.Title == nil && req.DueAt == nil && !req.ClearDue {
		return fmt.Errorf("at least one field must be changed")
	}
	return s.store.UpdateTodo(ctx, id, req.Title, req.DueAt, req.ClearDue)
}

func (s *Service) CloseTodo(ctx context.Context, id string) error {
	return s.store.TransitionTodo(ctx, id, model.TodoStateClosed)
}

func (s *Service) ReopenTodo(ctx context.Context, id string) error {
	return s.store.TransitionTodo(ctx, id, model.TodoStateReopened)
}

func (s *Service) RejectTodo(ctx context.Context, id string) error {
	return s.store.TransitionTodo(ctx, id, model.TodoStateRejected)
}

func (s *Service) CreateSink(ctx context.Context, req CreateSinkRequest) (model.Sink, error) {
	if strings.TrimSpace(req.ID) == "" {
		return model.Sink{}, fmt.Errorf("sink id is required")
	}
	if strings.TrimSpace(req.URL) == "" {
		return model.Sink{}, fmt.Errorf("sink url is required")
	}
	events := req.Events
	if len(events) == 0 {
		events = []string{"upcoming", "overdue"}
	}
	for _, e := range events {
		if e != model.ScheduleKindUpcoming && e != model.ScheduleKindOverdue {
			return model.Sink{}, fmt.Errorf("invalid sink event %q", e)
		}
	}
	now := time.Now().UTC()
	sink := model.Sink{
		ID:        req.ID,
		URL:       req.URL,
		Events:    events,
		Enabled:   true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.store.CreateSink(ctx, sink); err != nil {
		return model.Sink{}, err
	}
	return sink, nil
}

func (s *Service) DeleteSink(ctx context.Context, id string) error {
	return s.store.DeleteSink(ctx, id)
}

func (s *Service) AddSchedule(ctx context.Context, req AddScheduleRequest) (model.Schedule, error) {
	if strings.TrimSpace(req.ID) == "" {
		return model.Schedule{}, fmt.Errorf("schedule id is required")
	}
	if strings.TrimSpace(req.TodoID) == "" {
		return model.Schedule{}, fmt.Errorf("todo id is required")
	}
	if req.Kind != model.ScheduleKindUpcoming && req.Kind != model.ScheduleKindOverdue {
		return model.Schedule{}, fmt.Errorf("kind must be upcoming or overdue")
	}
	if req.Before != "" && req.Every != "" {
		return model.Schedule{}, fmt.Errorf("before and every are mutually exclusive")
	}
	if req.Kind == model.ScheduleKindUpcoming {
		if req.Every != "" {
			return model.Schedule{}, fmt.Errorf("every only applies to overdue schedules")
		}
		if req.Before == "" {
			req.Before = "24h"
		}
		if _, err := common.ParseHumanDuration(req.Before); err != nil {
			return model.Schedule{}, fmt.Errorf("invalid before duration: %w", err)
		}
	}
	if req.Kind == model.ScheduleKindOverdue {
		if req.Before != "" {
			return model.Schedule{}, fmt.Errorf("before only applies to upcoming schedules")
		}
		if req.Every == "" {
			req.Every = "24h"
		}
		if _, err := common.ParseHumanDuration(req.Every); err != nil {
			return model.Schedule{}, fmt.Errorf("invalid every duration: %w", err)
		}
	}
	if !req.MOTD && len(req.SinkID) == 0 {
		req.MOTD = true
	}
	if req.MOTD {
		for _, sinkID := range req.SinkID {
			if strings.TrimSpace(sinkID) == "" {
				continue
			}
			if _, err := s.store.GetSink(ctx, sinkID); err != nil {
				return model.Schedule{}, fmt.Errorf("sink %q not found: %w", sinkID, err)
			}
		}
	}
	todo, err := s.store.GetTodo(ctx, req.TodoID)
	if err != nil {
		return model.Schedule{}, err
	}
	if todo.DueAt == nil {
		return model.Schedule{}, fmt.Errorf("todo %q has no due date", req.TodoID)
	}

	now := time.Now().UTC()
	sc := model.Schedule{
		ID:         req.ID,
		TodoID:     req.TodoID,
		Kind:       req.Kind,
		Before:     req.Before,
		Every:      req.Every,
		Status:     model.ScheduleStatusActive,
		TargetMOTD: req.MOTD,
		SinkIDs:    req.SinkID,
		CreatedAt:  now,
	}
	if err := s.store.CreateSchedule(ctx, sc); err != nil {
		return model.Schedule{}, err
	}
	return sc, nil
}

func (s *Service) RemoveSchedule(ctx context.Context, id string) error {
	return s.store.DeleteSchedule(ctx, id)
}

func (s *Service) RunScheduler(ctx context.Context) error {
	schedules, todos, err := s.store.ActiveSchedulesWithTodos(ctx)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	for idx, sc := range schedules {
		todo := todos[idx]
		if todo.DueAt == nil {
			continue
		}
		plannedAt, shouldSend, err := computePlannedAtAndEligibility(sc, todo, now)
		if err != nil {
			continue
		}
		if !shouldSend {
			continue
		}
		allDelivered := true
		if sc.TargetMOTD {
			delivered, err := s.sendToMOTD(ctx, sc, todo, plannedAt)
			if err != nil {
				allDelivered = false
			}
			if !delivered {
				allDelivered = false
			}
		}
		sinkTargets, err := s.resolveSinkTargets(ctx, sc)
		if err != nil {
			allDelivered = false
		}
		for _, sink := range sinkTargets {
			delivered, err := s.sendToSink(ctx, sc, todo, sink, plannedAt)
			if err != nil {
				allDelivered = false
			}
			if !delivered {
				allDelivered = false
			}
		}
		if sc.Kind == model.ScheduleKindUpcoming && allDelivered {
			_ = s.store.MarkScheduleSent(ctx, sc.ID)
		}
	}
	return nil
}

func (s *Service) resolveSinkTargets(ctx context.Context, sc model.Schedule) ([]model.Sink, error) {
	allSinks, err := s.store.ListSinks(ctx, boolPtr(true), sc.Kind)
	if err != nil {
		return nil, err
	}
	if len(sc.SinkIDs) == 0 {
		return allSinks, nil
	}
	out := make([]model.Sink, 0)
	for _, sink := range allSinks {
		if slices.Contains(sc.SinkIDs, sink.ID) {
			out = append(out, sink)
		}
	}
	return out, nil
}

func (s *Service) sendToMOTD(ctx context.Context, sc model.Schedule, todo model.Todo, plannedAt time.Time) (bool, error) {
	const targetType = "motd"
	const targetID = "local"
	shouldAttempt, err := s.store.ShouldAttemptDelivery(ctx, sc.ID, targetType, targetID, plannedAt)
	if err != nil {
		return false, err
	}
	if !shouldAttempt {
		delivered, err := s.store.IsTargetDelivered(ctx, sc.ID, targetType, targetID, plannedAt)
		return delivered, err
	}
	message := buildReminderMessage(sc, todo)
	if err := s.store.QueueMOTDMessage(ctx, todo.ID, sc.ID, message); err != nil {
		_ = s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, targetID, plannedAt, false, err.Error())
		return false, err
	}
	if err := s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, targetID, plannedAt, true, ""); err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) sendToSink(ctx context.Context, sc model.Schedule, todo model.Todo, sink model.Sink, plannedAt time.Time) (bool, error) {
	const targetType = "sink"
	shouldAttempt, err := s.store.ShouldAttemptDelivery(ctx, sc.ID, targetType, sink.ID, plannedAt)
	if err != nil {
		return false, err
	}
	if !shouldAttempt {
		delivered, err := s.store.IsTargetDelivered(ctx, sc.ID, targetType, sink.ID, plannedAt)
		return delivered, err
	}
	payload := map[string]any{
		"todo_id":     todo.ID,
		"title":       todo.Title,
		"state":       todo.State,
		"due_at":      todo.DueAt,
		"schedule_id": sc.ID,
		"kind":        sc.Kind,
		"planned_at":  plannedAt,
		"message":     buildReminderMessage(sc, todo),
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, sink.URL, bytes.NewReader(body))
	if err != nil {
		_ = s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, sink.ID, plannedAt, false, err.Error())
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		_ = s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, sink.ID, plannedAt, false, err.Error())
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		raw, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		errText := fmt.Sprintf("status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(raw)))
		_ = s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, sink.ID, plannedAt, false, errText)
		return false, nil
	}
	if err := s.store.UpsertDeliveryResult(ctx, sc.ID, targetType, sink.ID, plannedAt, true, ""); err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) PullMOTDMessages(ctx context.Context) ([]string, error) {
	return s.store.PullMOTDMessages(ctx)
}

func computePlannedAtAndEligibility(sc model.Schedule, todo model.Todo, now time.Time) (time.Time, bool, error) {
	due := *todo.DueAt
	if sc.Kind == model.ScheduleKindUpcoming {
		before, err := common.ParseHumanDuration(sc.Before)
		if err != nil {
			return time.Time{}, false, err
		}
		planned := due.Add(-before)
		return planned, !now.Before(planned), nil
	}
	if sc.Kind == model.ScheduleKindOverdue {
		every, err := common.ParseHumanDuration(sc.Every)
		if err != nil {
			return time.Time{}, false, err
		}
		if now.Before(due) {
			return due, false, nil
		}
		elapsed := now.Sub(due)
		steps := int(elapsed / every)
		planned := due.Add(time.Duration(steps) * every)
		return planned, true, nil
	}
	return time.Time{}, false, fmt.Errorf("unknown kind %q", sc.Kind)
}

func buildReminderMessage(sc model.Schedule, todo model.Todo) string {
	due := ""
	if todo.DueAt != nil {
		due = todo.DueAt.Local().Format("2006-01-02 15:04 MST")
	}
	if sc.Kind == model.ScheduleKindUpcoming {
		return fmt.Sprintf("Upcoming todo: %s (%s) due at %s", todo.Title, todo.ID, due)
	}
	return fmt.Sprintf("Overdue todo: %s (%s) was due at %s", todo.Title, todo.ID, due)
}

func boolPtr(v bool) *bool {
	return &v
}
