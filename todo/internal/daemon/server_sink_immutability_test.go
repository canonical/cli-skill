package daemon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func newTestServer(t *testing.T) *Server {
	t.Helper()
	dbPath := filepath.Join(t.TempDir(), "todo-test.db")
	t.Setenv("TODOD_SOCKET_PATH", filepath.Join(t.TempDir(), "todod-test.sock"))
	srv, err := NewServer(dbPath)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}
	t.Cleanup(func() {
		_ = srv.st.Close()
	})
	return srv
}

func TestSinkRoutesAreImmutable(t *testing.T) {
	srv := newTestServer(t)

	_, err := srv.svc.CreateSink(context.Background(), CreateSinkRequest{
		ID:     "sink-1",
		URL:    "https://example.com/hook",
		Events: []string{"upcoming"},
	})
	if err != nil {
		t.Fatalf("failed to create sink fixture: %v", err)
	}

	t.Run("patch is not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/sinks/sink-1", nil)
		rr := httptest.NewRecorder()

		srv.handleSinkByID(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Fatalf("expected %d, got %d", http.StatusMethodNotAllowed, rr.Code)
		}
	})

	t.Run("enable route is removed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/sinks/sink-1/enable", nil)
		rr := httptest.NewRecorder()

		srv.handleSinkByID(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected %d, got %d", http.StatusNotFound, rr.Code)
		}
	})

	t.Run("disable route is removed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/sinks/sink-1/disable", nil)
		rr := httptest.NewRecorder()

		srv.handleSinkByID(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected %d, got %d", http.StatusNotFound, rr.Code)
		}
	})

	t.Run("delete remains allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/sinks/sink-1", nil)
		rr := httptest.NewRecorder()

		srv.handleSinkByID(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Fatalf("expected %d, got %d", http.StatusNoContent, rr.Code)
		}
	})
}
