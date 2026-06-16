package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"todo/internal/daemon"
	"todo/internal/model"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func New(addr string, timeout time.Duration) *Client {
	if addr == "" {
		addr = daemon.ParseAddr("127.0.0.1", 44180)
	}
	return &Client{
		baseURL: "http://" + addr,
		http: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) do(ctx context.Context, method, path string, reqBody any, respBody any) error {
	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return err
	}
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		raw, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return errors.New(strings.TrimSpace(string(raw)))
	}
	if respBody != nil {
		return json.NewDecoder(resp.Body).Decode(respBody)
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	return nil
}

func (c *Client) CreateTodo(ctx context.Context, req daemon.CreateTodoRequest) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodPost, "/todos", req, &out)
	return out, err
}

func (c *Client) ListTodos(ctx context.Context, state string) ([]model.Todo, error) {
	q := url.Values{}
	if state != "" {
		q.Set("state", state)
	}
	path := "/todos"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var out struct {
		Todos []model.Todo `json:"todos"`
	}
	err := c.do(ctx, http.MethodGet, path, nil, &out)
	return out.Todos, err
}

func (c *Client) ShowTodo(ctx context.Context, id string) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodGet, "/todos/"+id, nil, &out)
	return out, err
}

func (c *Client) UpdateTodo(ctx context.Context, id string, req daemon.UpdateTodoRequest) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodPatch, "/todos/"+id, req, &out)
	return out, err
}

func (c *Client) CloseTodo(ctx context.Context, id string) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodPost, "/todos/"+id+"/close", nil, &out)
	return out, err
}

func (c *Client) ReopenTodo(ctx context.Context, id string) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodPost, "/todos/"+id+"/reopen", nil, &out)
	return out, err
}

func (c *Client) RejectTodo(ctx context.Context, id string) (model.Todo, error) {
	var out model.Todo
	err := c.do(ctx, http.MethodPost, "/todos/"+id+"/reject", nil, &out)
	return out, err
}

func (c *Client) CreateSink(ctx context.Context, req daemon.CreateSinkRequest) (model.Sink, error) {
	var out model.Sink
	err := c.do(ctx, http.MethodPost, "/sinks", req, &out)
	return out, err
}

func (c *Client) ListSinks(ctx context.Context, enabled *bool, event string) ([]model.Sink, error) {
	q := url.Values{}
	if enabled != nil {
		if *enabled {
			q.Set("enabled", "true")
		} else {
			q.Set("enabled", "false")
		}
	}
	if event != "" {
		q.Set("event", event)
	}
	path := "/sinks"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var out struct {
		Sinks []model.Sink `json:"sinks"`
	}
	err := c.do(ctx, http.MethodGet, path, nil, &out)
	return out.Sinks, err
}

func (c *Client) ShowSink(ctx context.Context, id string) (model.Sink, error) {
	var out model.Sink
	err := c.do(ctx, http.MethodGet, "/sinks/"+id, nil, &out)
	return out, err
}

func (c *Client) UpdateSink(ctx context.Context, id string, req daemon.UpdateSinkRequest) (model.Sink, error) {
	var out model.Sink
	err := c.do(ctx, http.MethodPatch, "/sinks/"+id, req, &out)
	return out, err
}

func (c *Client) DeleteSink(ctx context.Context, id string) error {
	return c.do(ctx, http.MethodDelete, "/sinks/"+id, nil, nil)
}

func (c *Client) EnableSink(ctx context.Context, id string) (model.Sink, error) {
	var out model.Sink
	err := c.do(ctx, http.MethodPost, "/sinks/"+id+"/enable", nil, &out)
	return out, err
}

func (c *Client) DisableSink(ctx context.Context, id string) (model.Sink, error) {
	var out model.Sink
	err := c.do(ctx, http.MethodPost, "/sinks/"+id+"/disable", nil, &out)
	return out, err
}

func (c *Client) AddSchedule(ctx context.Context, req daemon.AddScheduleRequest) (model.Schedule, error) {
	var out model.Schedule
	err := c.do(ctx, http.MethodPost, "/schedules", req, &out)
	return out, err
}

func (c *Client) ListSchedules(ctx context.Context, todoID, kind, status, target string) ([]model.Schedule, error) {
	q := url.Values{}
	if todoID != "" {
		q.Set("todo", todoID)
	}
	if kind != "" {
		q.Set("kind", kind)
	}
	if status != "" {
		q.Set("status", status)
	}
	if target != "" {
		q.Set("target", target)
	}
	path := "/schedules"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var out struct {
		Schedules []model.Schedule `json:"schedules"`
	}
	err := c.do(ctx, http.MethodGet, path, nil, &out)
	return out.Schedules, err
}

func (c *Client) ShowSchedule(ctx context.Context, id string) (model.Schedule, error) {
	var out model.Schedule
	err := c.do(ctx, http.MethodGet, "/schedules/"+id, nil, &out)
	return out, err
}

func (c *Client) RemoveSchedule(ctx context.Context, id string) error {
	return c.do(ctx, http.MethodDelete, "/schedules/"+id, nil, nil)
}

func (c *Client) MOTDMessage(ctx context.Context) ([]string, error) {
	var out struct {
		Messages []string `json:"messages"`
	}
	err := c.do(ctx, http.MethodGet, "/motd/message", nil, &out)
	return out.Messages, err
}

func (c *Client) Status(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodGet, "/status", nil, &out)
	return out, err
}
