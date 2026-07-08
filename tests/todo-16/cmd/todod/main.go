package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"

	"todo/internal/daemon"
	"todo/internal/store"
)

const version = "0.1.0"

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var dbPath string
	var format string

	root := &cobra.Command{
		Use:   "todod",
		Short: "Todo daemon",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := dbPath
			if db == "" {
				db = store.DefaultDBPath(home)
			}
			srv, err := daemon.NewServer(db)
			if err != nil {
				return err
			}
			return srv.Run(context.Background())
		},
	}
	startCmd.Flags().StringVar(&dbPath, "db", "", "SQLite database path")

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show daemon status",
		RunE: func(cmd *cobra.Command, args []string) error {
			sock := daemon.SocketPath()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			client := newSocketHTTPClient(sock, 5*time.Second)
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://unix/status", nil)
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode >= 300 {
				return fmt.Errorf("status request failed: %s", resp.Status)
			}
			var payload map[string]any
			if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
				return err
			}
			if format == "json" {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				return enc.Encode(payload)
			}
			fmt.Printf("todod: %s\n", sock)
			fmt.Printf("database: %v\n", payload["database_path"])
			fmt.Printf("open todos: %v\n", payload["active_todo_count"])
			fmt.Printf("active schedules: %v\n", payload["active_schedule_count"])
			fmt.Printf("enabled sinks: %v\n", payload["enabled_sink_count"])
			return nil
		},
	}
	statusCmd.Flags().StringVar(&format, "format", "table", "Output format: table|json")

	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			sock := daemon.SocketPath()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			client := newSocketHTTPClient(sock, 5*time.Second)
			req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://unix/shutdown", nil)
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode >= 300 {
				return fmt.Errorf("shutdown request failed: %s", resp.Status)
			}
			fmt.Println("todod stop requested")
			return nil
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show daemon version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	root.AddCommand(startCmd, stopCmd, statusCmd, versionCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newSocketHTTPClient(socketPath string, timeout time.Duration) *http.Client {
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			var dialer net.Dialer
			return dialer.DialContext(ctx, "unix", socketPath)
		},
	}
	return &http.Client{Timeout: timeout, Transport: transport}
}
