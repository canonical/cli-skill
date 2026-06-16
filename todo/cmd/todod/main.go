package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	var host string
	var port int
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
			addr := daemon.ParseAddr(host, port)
			db := dbPath
			if db == "" {
				db = store.DefaultDBPath(home)
			}
			srv, err := daemon.NewServer(db, addr)
			if err != nil {
				return err
			}
			return srv.Run(context.Background())
		},
	}
	startCmd.Flags().StringVar(&host, "host", "127.0.0.1", "HTTP listen host")
	startCmd.Flags().IntVar(&port, "port", 44180, "HTTP listen port")
	startCmd.Flags().StringVar(&dbPath, "db", "", "SQLite database path")

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show daemon status",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := daemon.ParseAddr(host, port)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://"+addr+"/status", nil)
			resp, err := http.DefaultClient.Do(req)
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
			fmt.Printf("todod: %s\n", addr)
			fmt.Printf("database: %v\n", payload["database_path"])
			fmt.Printf("open todos: %v\n", payload["active_todo_count"])
			fmt.Printf("active schedules: %v\n", payload["active_schedule_count"])
			fmt.Printf("enabled sinks: %v\n", payload["enabled_sink_count"])
			return nil
		},
	}
	statusCmd.Flags().StringVar(&host, "host", "127.0.0.1", "HTTP daemon host")
	statusCmd.Flags().IntVar(&port, "port", 44180, "HTTP daemon port")
	statusCmd.Flags().StringVar(&format, "format", "table", "Output format: table|json")

	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := daemon.ParseAddr(host, port)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://"+addr+"/shutdown", nil)
			resp, err := http.DefaultClient.Do(req)
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
	stopCmd.Flags().StringVar(&host, "host", "127.0.0.1", "HTTP daemon host")
	stopCmd.Flags().IntVar(&port, "port", 44180, "HTTP daemon port")

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
