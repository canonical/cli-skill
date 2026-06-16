package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestParseInlineSchedule(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantKind  string
		wantBefore string
		wantEvery string
		wantMOTD  bool
		wantSinks []string
		wantErr   bool
	}{
		{
			name:    "rejects missing target segment",
			input:   "upcoming",
			wantErr: true,
		},
		{
			name:       "upcoming motd before",
			input:      "upcoming:before=24h:motd",
			wantKind:   "upcoming",
			wantBefore: "24h",
			wantMOTD:   true,
		},
		{
			name:      "overdue every with multiple sinks",
			input:     "overdue:every=1d:sink=s1:sink=s2",
			wantKind:  "overdue",
			wantEvery: "1d",
			wantMOTD:  false,
			wantSinks: []string{"s1", "s2"},
		},
		{
			name:      "defaults to motd when no targets",
			input:     "upcoming:",
			wantKind:  "upcoming",
			wantMOTD:  true,
			wantSinks: nil,
		},
		{
			name:      "sink target disables motd default",
			input:     "upcoming:sink=abc",
			wantKind:  "upcoming",
			wantMOTD:  false,
			wantSinks: []string{"abc"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseInlineSchedule(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for input %q", tc.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Kind != tc.wantKind {
				t.Fatalf("kind mismatch: got %q want %q", got.Kind, tc.wantKind)
			}
			if got.Before != tc.wantBefore {
				t.Fatalf("before mismatch: got %q want %q", got.Before, tc.wantBefore)
			}
			if got.Every != tc.wantEvery {
				t.Fatalf("every mismatch: got %q want %q", got.Every, tc.wantEvery)
			}
			if got.MOTD != tc.wantMOTD {
				t.Fatalf("motd mismatch: got %v want %v", got.MOTD, tc.wantMOTD)
			}
			if strings.Join(got.SinkID, ",") != strings.Join(tc.wantSinks, ",") {
				t.Fatalf("sink ids mismatch: got %v want %v", got.SinkID, tc.wantSinks)
			}
		})
	}
}

func TestColorizedHelp(t *testing.T) {
	newRoot := func() *cobra.Command {
		root := &cobra.Command{
			Use:   "todo",
			Short: "Todo client",
			Long:  "Todo client for managing tasks.",
		}
		root.AddGroup(&cobra.Group{ID: "todos", Title: "Todos:"})
		root.AddGroup(&cobra.Group{ID: "sinks", Title: "Sinks:"})
		root.AddGroup(&cobra.Group{ID: "schedules", Title: "Schedules:"})
		root.AddGroup(&cobra.Group{ID: "other", Title: "Other:"})
		root.SetHelpCommandGroupID("other")
		root.SetCompletionCommandGroupID("other")

		root.PersistentFlags().Bool("verbose", false, "Verbose output")

		listCmd := &cobra.Command{Use: "list", Short: "List todos", Run: func(cmd *cobra.Command, args []string) {}}
		listCmd.GroupID = "todos"
		scheduleCmd := &cobra.Command{Use: "schedule", Short: "Show a schedule", Run: func(cmd *cobra.Command, args []string) {}}
		scheduleCmd.GroupID = "schedules"
		sinksCmd := &cobra.Command{Use: "sinks", Short: "List sinks", Run: func(cmd *cobra.Command, args []string) {}}
		sinksCmd.GroupID = "sinks"
		sinkCmd := &cobra.Command{Use: "sink", Short: "Show a sink", Run: func(cmd *cobra.Command, args []string) {}}
		sinkCmd.GroupID = "sinks"
		createSinkCmd := &cobra.Command{Use: "create-sink", Short: "Create a sink", Run: func(cmd *cobra.Command, args []string) {}}
		createSinkCmd.GroupID = "sinks"
		deleteSinkCmd := &cobra.Command{Use: "delete-sink", Short: "Delete a sink", Run: func(cmd *cobra.Command, args []string) {}}
		deleteSinkCmd.GroupID = "sinks"
		versionCmd := &cobra.Command{Use: "version", Short: "Show version", Run: func(cmd *cobra.Command, args []string) {}}
		versionCmd.GroupID = "other"
		motdCmd := &cobra.Command{Use: "motd-message", Short: "Print pending MOTD reminder messages", Run: func(cmd *cobra.Command, args []string) {}}
		motdCmd.GroupID = "other"
		statusCmd := &cobra.Command{Use: "status", Short: "Show todo system status", Run: func(cmd *cobra.Command, args []string) {}}
		statusCmd.GroupID = "other"
		root.AddCommand(listCmd, scheduleCmd, sinksCmd, sinkCmd, createSinkCmd, deleteSinkCmd, versionCmd, motdCmd, statusCmd)
		return root
	}

	t.Run("uses color for sections when color is supported", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")
		t.Setenv("COLORFGBG", "15;0") // explicit dark background

		root := newRoot()
		got := colorizedHelp(root)

		if !strings.Contains(got, "\033[37mUsage:\033[0m") {
			t.Fatalf("expected colored Usage section, got: %q", got)
		}
		if !strings.Contains(got, "\033[37mGlobal options:\033[0m") {
			t.Fatalf("expected colored Global options section, got: %q", got)
		}
		if strings.Contains(got, "\nFlags:\n") {
			t.Fatalf("root help should rename Flags to Global options")
		}
	})

	t.Run("root help uses sinks row inside other", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		t.Setenv("TERM", "xterm-256color")

		root := newRoot()
		got := colorizedHelp(root)
		if strings.Contains(got, "Sinks:\n") {
			t.Fatalf("did not expect standalone Sinks section: %q", got)
		}
		if !strings.Contains(got, "Other:") {
			t.Fatalf("expected Other section: %q", got)
		}
		if !strings.Contains(got, "Sinks") || !strings.Contains(got, "create-sink, delete-sink, sink, sinks") {
			t.Fatalf("expected compact sinks row in Other section: %q", got)
		}
	})

	t.Run("uses bold fallback with NO_COLOR", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		t.Setenv("TERM", "xterm-256color")

		root := newRoot()
		got := colorizedHelp(root)

		if !strings.Contains(got, "\033[1mUsage:\033[0m") {
			t.Fatalf("expected bold Usage fallback, got: %q", got)
		}
		if strings.Contains(got, "\033[37mUsage:\033[0m") || strings.Contains(got, "\033[30mUsage:\033[0m") {
			t.Fatalf("did not expect colored Usage with NO_COLOR")
		}
	})

	t.Run("keeps global flags label for subcommands", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")

		root := newRoot()
		cmd := &cobra.Command{Use: "create", Short: "Create todo"}
		cmd.Flags().String("title", "", "Todo title")
		root.AddCommand(cmd)

		got := colorizedHelp(cmd)
		if !strings.Contains(got, "Global") || !strings.Contains(got, "Flags") {
			t.Fatalf("expected Global Flags section in subcommand help, got: %q", got)
		}
	})
}
