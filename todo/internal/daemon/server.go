package daemon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"todo/internal/store"
)

type Server struct {
	st       *store.Store
	svc      *Service
	dbPath   string
	addr     string
	httpSrv  *http.Server
	stopOnce sync.Once
	stopCh   chan struct{}
}

func NewServer(dbPath, addr string) (*Server, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, err
	}
	st, err := store.Open(dbPath)
	if err != nil {
		return nil, err
	}
	s := &Server{
		st:     st,
		svc:    NewService(st),
		dbPath: dbPath,
		addr:   addr,
		stopCh: make(chan struct{}),
	}
	mux := http.NewServeMux()
	s.registerRoutes(mux)
	s.httpSrv = &http.Server{
		Addr:              addr,
		Handler:           loggingMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}
	return s, nil
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/status", s.handleStatus)
	mux.HandleFunc("/shutdown", s.handleShutdown)
	mux.HandleFunc("/motd/message", s.handleMOTDMessage)
	mux.HandleFunc("/todos", s.handleTodos)
	mux.HandleFunc("/todos/", s.handleTodoByID)
	mux.HandleFunc("/sinks", s.handleSinks)
	mux.HandleFunc("/sinks/", s.handleSinkByID)
	mux.HandleFunc("/schedules", s.handleSchedules)
	mux.HandleFunc("/schedules/", s.handleScheduleByID)
}

func (s *Server) handleShutdown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	go s.Stop()
	writeJSON(w, http.StatusOK, map[string]any{"status": "stopping"})
}

func (s *Server) Run(ctx context.Context) error {
	go s.schedulerLoop(ctx)
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
		case <-sigCh:
			s.Stop()
		}
	}()

	log.Printf("todod listening on %s", s.addr)
	err := s.httpSrv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Stop() {
	s.stopOnce.Do(func() {
		close(s.stopCh)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = s.httpSrv.Shutdown(ctx)
		_ = s.st.Close()
	})
}

func (s *Server) schedulerLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.stopCh:
			return
		case <-ticker.C:
			if err := s.svc.RunScheduler(ctx); err != nil {
				log.Printf("scheduler error: %v", err)
			}
		}
	}
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	todos, err := s.st.ListTodos(r.Context(), store.TodoFilter{State: "open"})
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	schedules, err := s.st.ListSchedules(r.Context(), store.ScheduleFilter{Status: "active"})
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	trueVal := true
	sinks, err := s.st.ListSinks(r.Context(), &trueVal, "")
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	home, _ := os.UserHomeDir()
	profile := filepath.Join(home, ".profile")
	b, _ := os.ReadFile(profile)
	hasMotd := strings.Contains(strings.ToLower(string(b)), "todo motd-message")
	writeJSON(w, http.StatusOK, map[string]any{
		"now":                            time.Now().UTC().Format(time.RFC3339),
		"database_path":                  s.dbPath,
		"active_todo_count":              len(todos),
		"active_schedule_count":          len(schedules),
		"enabled_sink_count":             len(sinks),
		"needs_motd_login_script_hint":   !hasMotd,
		"motd_login_script_hint":         "To show todo reminders on login, run: echo 'todo motd-message' >> ~/.profile",
		"motd_login_script_check_path":   profile,
	})
}

func (s *Server) handleMOTDMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	msgs, err := s.svc.PullMOTDMessages(r.Context())
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"messages": msgs})
}

func (s *Server) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req CreateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid json")
			return
		}
		todo, err := s.svc.CreateTodo(r.Context(), req)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, todo)
	case http.MethodGet:
		filter := store.TodoFilter{State: r.URL.Query().Get("state")}
		todos, err := s.st.ListTodos(r.Context(), filter)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"todos": todos})
	default:
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/todos/")
	parts := strings.Split(path, "/")
	id := parts[0]
	if id == "" {
		writeErr(w, http.StatusBadRequest, "todo id is required")
		return
	}
	if len(parts) == 1 {
		switch r.Method {
		case http.MethodGet:
			todo, err := s.st.GetTodo(r.Context(), id)
			if err != nil {
				writeErr(w, http.StatusNotFound, err.Error())
				return
			}
			writeJSON(w, http.StatusOK, todo)
		case http.MethodPatch:
			var req UpdateTodoRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeErr(w, http.StatusBadRequest, "invalid json")
				return
			}
			if err := s.svc.UpdateTodo(r.Context(), id, req); err != nil {
				writeErr(w, http.StatusBadRequest, err.Error())
				return
			}
			todo, _ := s.st.GetTodo(r.Context(), id)
			writeJSON(w, http.StatusOK, todo)
		default:
			writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		}
		return
	}
	action := parts[1]
	if r.Method != http.MethodPost {
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var err error
	switch action {
	case "close":
		err = s.svc.CloseTodo(r.Context(), id)
	case "reopen":
		err = s.svc.ReopenTodo(r.Context(), id)
	case "reject":
		err = s.svc.RejectTodo(r.Context(), id)
	default:
		writeErr(w, http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	todo, _ := s.st.GetTodo(r.Context(), id)
	writeJSON(w, http.StatusOK, todo)
}

func (s *Server) handleSinks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req CreateSinkRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid json")
			return
		}
		sink, err := s.svc.CreateSink(r.Context(), req)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, sink)
	case http.MethodGet:
		q := r.URL.Query()
		event := q.Get("event")
		var enabled *bool
		if q.Get("enabled") != "" {
			v := q.Get("enabled") == "true"
			enabled = &v
		}
		sinks, err := s.st.ListSinks(r.Context(), enabled, event)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"sinks": sinks})
	default:
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) handleSinkByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/sinks/")
	parts := strings.Split(path, "/")
	id := parts[0]
	if id == "" {
		writeErr(w, http.StatusBadRequest, "sink id is required")
		return
	}
	if len(parts) == 1 {
		switch r.Method {
		case http.MethodGet:
			sink, err := s.st.GetSink(r.Context(), id)
			if err != nil {
				writeErr(w, http.StatusNotFound, err.Error())
				return
			}
			writeJSON(w, http.StatusOK, sink)
		case http.MethodDelete:
			if err := s.svc.DeleteSink(r.Context(), id); err != nil {
				writeErr(w, http.StatusBadRequest, err.Error())
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		}
		return
	}
	writeErr(w, http.StatusNotFound, "not found")
}

func (s *Server) handleSchedules(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()
		schedules, err := s.st.ListSchedules(r.Context(), store.ScheduleFilter{
			TodoID: q.Get("todo"),
			Kind:   q.Get("kind"),
			Status: q.Get("status"),
			Target: q.Get("target"),
		})
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"schedules": schedules})
	case http.MethodPost:
		var req AddScheduleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid json")
			return
		}
		sc, err := s.svc.AddSchedule(r.Context(), req)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, sc)
	default:
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) handleScheduleByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/schedules/")
	if id == "" {
		writeErr(w, http.StatusBadRequest, "schedule id is required")
		return
	}
	switch r.Method {
	case http.MethodGet:
		sc, err := s.st.GetSchedule(r.Context(), id)
		if err != nil {
			writeErr(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, sc)
	case http.MethodDelete:
		if err := s.svc.RemoveSchedule(r.Context(), id); err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"error": msg})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start).Round(time.Millisecond))
	})
}

func ParseAddr(host string, port int) string {
	if strings.TrimSpace(host) == "" {
		host = "127.0.0.1"
	}
	if port <= 0 {
		port = 44180
	}
	return fmt.Sprintf("%s:%d", host, port)
}
