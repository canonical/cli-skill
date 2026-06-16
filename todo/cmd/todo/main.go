package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"todo/internal/client"
	"todo/internal/common"
	"todo/internal/daemon"
	"todo/internal/model"
)

const version = "0.1.0"

func main() {
	var host string
	var port int
	var timeout time.Duration
	var format string
	var rfc3339 bool

	root := &cobra.Command{
		Use:   "todo",
		Short: "Todo client",
	}
	root.PersistentFlags().StringVar(&host, "host", "127.0.0.1", "Daemon host")
	root.PersistentFlags().IntVar(&port, "port", 44180, "Daemon port")
	root.PersistentFlags().DurationVar(&timeout, "timeout", 10*time.Second, "Request timeout")
	root.PersistentFlags().StringVar(&format, "format", "table", "Output format: table|json")
	root.PersistentFlags().BoolVar(&rfc3339, "rfc3339", false, "Force RFC3339 date output")

	newClient := func() *client.Client {
		return client.New(daemon.ParseAddr(host, port), timeout)
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List todos",
		RunE: func(cmd *cobra.Command, args []string) error {
			state, _ := cmd.Flags().GetString("state")
			cli := newClient()
			ctx := context.Background()
			todos, err := cli.ListTodos(ctx, state)
			if err != nil {
				return err
			}
			return printTodos(todos, format, rfc3339)
		},
	}
	listCmd.Flags().String("state", "", "Filter state: open|closed|reopened|rejected")

	showCmd := &cobra.Command{
		Use:   "show <todo-id>",
		Short: "Show a todo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			todo, err := cli.ShowTodo(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, format, rfc3339)
		},
	}

	createCmd := &cobra.Command{
		Use:   "create <todo-id> <title>",
		Short: "Create a todo",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			dueInput, _ := cmd.Flags().GetString("due")
			scheduleFlags, _ := cmd.Flags().GetStringArray("schedule")
			var dueAt *time.Time
			if strings.TrimSpace(dueInput) != "" {
				t, err := common.ParseDateTime(dueInput)
				if err != nil {
					return err
				}
				t = t.UTC()
				dueAt = &t
			}
			req := daemon.CreateTodoRequest{
				ID:    args[0],
				Title: strings.Join(args[1:], " "),
				DueAt: dueAt,
			}
			for i, raw := range scheduleFlags {
				spec, err := parseInlineSchedule(raw)
				if err != nil {
					return fmt.Errorf("invalid --schedule at #%d: %w", i+1, err)
				}
				req.Schedule = append(req.Schedule, spec)
			}
			cli := newClient()
			todo, err := cli.CreateTodo(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, format, rfc3339)
		},
	}
	createCmd.Flags().String("due", "", "Optional due date (RFC3339 or human-readable)")
	createCmd.Flags().StringArray("schedule", nil, "Optional schedule spec (upcoming[:before=24h]:motd|sink=<id>, overdue[:every=1d]:motd|sink=<id>)")

	updateCmd := &cobra.Command{
		Use:   "update <todo-id>",
		Short: "Update a todo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			dueInput, _ := cmd.Flags().GetString("due")
			clearDue, _ := cmd.Flags().GetBool("clear-due")
			var req daemon.UpdateTodoRequest
			if strings.TrimSpace(title) != "" {
				req.Title = &title
			}
			if strings.TrimSpace(dueInput) != "" {
				t, err := common.ParseDateTime(dueInput)
				if err != nil {
					return err
				}
				t = t.UTC()
				req.DueAt = &t
			}
			req.ClearDue = clearDue
			cli := newClient()
			todo, err := cli.UpdateTodo(context.Background(), args[0], req)
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, format, rfc3339)
		},
	}
	updateCmd.Flags().String("title", "", "New title")
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date")

	closeCmd := todoActionCmd("close <todo-id>", "Close a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.CloseTodo(context.Background(), id)
	}, &format, &rfc3339, newClient)
	reopenCmd := todoActionCmd("reopen <todo-id>", "Reopen a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.ReopenTodo(context.Background(), id)
	}, &format, &rfc3339, newClient)
	rejectCmd := todoActionCmd("reject <todo-id>", "Reject a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.RejectTodo(context.Background(), id)
	}, &format, &rfc3339, newClient)

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show todo system status",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			payload, err := cli.Status(context.Background())
			if err != nil {
				return err
			}
			if format == "json" {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				return enc.Encode(payload)
			}
			fmt.Printf("now: %v\n", payload["now"])
			fmt.Printf("database: %v\n", payload["database_path"])
			fmt.Printf("open todos: %v\n", payload["active_todo_count"])
			fmt.Printf("active schedules: %v\n", payload["active_schedule_count"])
			fmt.Printf("enabled sinks: %v\n", payload["enabled_sink_count"])
			if need, ok := payload["needs_motd_login_script_hint"].(bool); ok && need {
				if hint, ok := payload["motd_login_script_hint"].(string); ok {
					fmt.Println(hint)
				}
			}
			return nil
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show client version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	motdMessageCmd := &cobra.Command{
		Use:   "motd-message",
		Short: "Print pending MOTD reminder messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			msgs, err := cli.MOTDMessage(context.Background())
			if err != nil {
				return err
			}
			for _, m := range msgs {
				fmt.Println(m)
			}
			return nil
		},
	}

	// Sink commands
	sinksCmd := &cobra.Command{
		Use:   "sinks",
		Short: "List sinks",
		RunE: func(cmd *cobra.Command, args []string) error {
			var enabled *bool
			enabledStr, _ := cmd.Flags().GetString("enabled")
			if enabledStr == "true" {
				v := true
				enabled = &v
			} else if enabledStr == "false" {
				v := false
				enabled = &v
			}
			event, _ := cmd.Flags().GetString("event")
			cli := newClient()
			sinks, err := cli.ListSinks(context.Background(), enabled, event)
			if err != nil {
				return err
			}
			return printJSONOrTable(sinks, format, rfc3339)
		},
	}
	sinksCmd.Flags().String("enabled", "", "Filter enabled: true|false")
	sinksCmd.Flags().String("event", "", "Filter event: upcoming|overdue")

	sinkCmd := &cobra.Command{
		Use:   "sink <sink-id>",
		Short: "Show a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			sink, err := cli.ShowSink(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, format, rfc3339)
		},
	}

	createSinkCmd := &cobra.Command{
		Use:   "create-sink <sink-id>",
		Short: "Create a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			events, _ := cmd.Flags().GetStringArray("event")
			if strings.TrimSpace(url) == "" {
				return fmt.Errorf("--url is required")
			}
			req := daemon.CreateSinkRequest{ID: args[0], URL: url, Events: events}
			cli := newClient()
			sink, err := cli.CreateSink(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, format, rfc3339)
		},
	}
	createSinkCmd.Flags().String("url", "", "Webhook URL")
	createSinkCmd.Flags().StringArray("event", nil, "Event subscriptions (upcoming|overdue)")

	updateSinkCmd := &cobra.Command{
		Use:   "update-sink <sink-id>",
		Short: "Update a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			events, _ := cmd.Flags().GetStringArray("event")
			clearEvents, _ := cmd.Flags().GetBool("clear-events")
			var urlPtr *string
			if strings.TrimSpace(url) != "" {
				urlPtr = &url
			}
			req := daemon.UpdateSinkRequest{URL: urlPtr, Events: events, ClearEvents: clearEvents}
			cli := newClient()
			sink, err := cli.UpdateSink(context.Background(), args[0], req)
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, format, rfc3339)
		},
	}
	updateSinkCmd.Flags().String("url", "", "New webhook URL")
	updateSinkCmd.Flags().StringArray("event", nil, "Replace events (upcoming|overdue)")
	updateSinkCmd.Flags().Bool("clear-events", false, "Clear event subscriptions")

	deleteSinkCmd := &cobra.Command{
		Use:   "delete-sink <sink-id>",
		Short: "Delete a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			if err := cli.DeleteSink(context.Background(), args[0]); err != nil {
				return err
			}
			fmt.Println("sink deleted")
			return nil
		},
	}

	enableSinkCmd := &cobra.Command{
		Use:   "enable-sink <sink-id>",
		Short: "Enable a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			sink, err := cli.EnableSink(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, format, rfc3339)
		},
	}

	disableSinkCmd := &cobra.Command{
		Use:   "disable-sink <sink-id>",
		Short: "Disable a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			sink, err := cli.DisableSink(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, format, rfc3339)
		},
	}

	// Schedule commands
	schedulesCmd := &cobra.Command{
		Use:   "schedules",
		Short: "List schedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			todoID, _ := cmd.Flags().GetString("todo")
			kind, _ := cmd.Flags().GetString("kind")
			status, _ := cmd.Flags().GetString("status")
			target, _ := cmd.Flags().GetString("target")
			cli := newClient()
			schedules, err := cli.ListSchedules(context.Background(), todoID, kind, status, target)
			if err != nil {
				return err
			}
			return printJSONOrTable(schedules, format, rfc3339)
		},
	}
	schedulesCmd.Flags().String("todo", "", "Optional todo id filter")
	schedulesCmd.Flags().String("kind", "", "Filter kind: upcoming|overdue")
	schedulesCmd.Flags().String("status", "", "Filter status: active|sent")
	schedulesCmd.Flags().String("target", "", "Filter target: sink|motd")

	scheduleCmd := &cobra.Command{
		Use:   "schedule <schedule-id>",
		Short: "Show a schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			sc, err := cli.ShowSchedule(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sc, format, rfc3339)
		},
	}

	addScheduleCmd := &cobra.Command{
		Use:   "add-schedule <schedule-id>",
		Short: "Add an immutable schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			todoID, _ := cmd.Flags().GetString("todo")
			kind, _ := cmd.Flags().GetString("kind")
			before, _ := cmd.Flags().GetString("before")
			every, _ := cmd.Flags().GetString("every")
			sinks, _ := cmd.Flags().GetStringArray("sink")
			motd, _ := cmd.Flags().GetBool("motd")
			if strings.TrimSpace(todoID) == "" {
				return fmt.Errorf("--todo is required")
			}
			if strings.TrimSpace(kind) == "" {
				return fmt.Errorf("--kind is required")
			}
			req := daemon.AddScheduleRequest{
				ID:     args[0],
				TodoID: todoID,
				Kind:   kind,
				Before: before,
				Every:  every,
				MOTD:   motd,
				SinkID: sinks,
			}
			cli := newClient()
			sc, err := cli.AddSchedule(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(sc, format, rfc3339)
		},
	}
	addScheduleCmd.Flags().String("todo", "", "Todo id")
	addScheduleCmd.Flags().String("kind", "", "Schedule kind: upcoming|overdue")
	addScheduleCmd.Flags().String("before", "", "Upcoming schedule offset before due date, default 24h")
	addScheduleCmd.Flags().String("every", "", "Overdue reminder frequency, default 1d")
	addScheduleCmd.Flags().StringArray("sink", nil, "Optional sink id (repeatable)")
	addScheduleCmd.Flags().Bool("motd", false, "Deliver via shell MOTD channel")

	removeScheduleCmd := &cobra.Command{
		Use:   "remove-schedule <schedule-id>",
		Short: "Remove a schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			if err := cli.RemoveSchedule(context.Background(), args[0]); err != nil {
				return err
			}
			fmt.Println("schedule removed")
			return nil
		},
	}

	root.AddCommand(listCmd, showCmd, createCmd, updateCmd, closeCmd, reopenCmd, rejectCmd)
	root.AddCommand(sinksCmd, sinkCmd, createSinkCmd, updateSinkCmd, deleteSinkCmd, enableSinkCmd, disableSinkCmd)
	root.AddCommand(schedulesCmd, scheduleCmd, addScheduleCmd, removeScheduleCmd)
	root.AddCommand(motdMessageCmd, statusCmd, versionCmd)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func todoActionCmd(use, short string, fn func(*client.Client, string) (model.Todo, error), format *string, rfc3339 *bool, newClient func() *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient()
			todo, err := fn(cli, args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, *format, *rfc3339)
		},
	}
}

func printTodos(todos []model.Todo, format string, rfc3339 bool) error {
	if format == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(todos)
	}
	if len(todos) == 0 {
		fmt.Fprintln(os.Stderr, "No todos found")
		return nil
	}
	fmt.Printf("%-24s  %-12s  %-16s  %s\n", "ID", "STATE", "DUE", "TITLE")
	for _, t := range todos {
		due := common.FormatTime(t.DueAt, rfc3339)
		fmt.Printf("%-24s  %-12s  %-16s  %s\n", t.ID, t.State, due, t.Title)
	}
	return nil
}

func printJSONOrTable(v any, format string, rfc3339 bool) error {
	if format == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
	switch t := v.(type) {
	case model.Todo:
		fmt.Printf("id: %s\n", t.ID)
		fmt.Printf("title: %s\n", t.Title)
		fmt.Printf("state: %s\n", t.State)
		fmt.Printf("due: %s\n", common.FormatTime(t.DueAt, rfc3339))
		fmt.Printf("created: %s\n", common.FormatTime(&t.CreatedAt, rfc3339))
		fmt.Printf("updated: %s\n", common.FormatTime(&t.UpdatedAt, rfc3339))
		return nil
	default:
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
}

func parseInlineSchedule(raw string) (daemon.ScheduleSpec, error) {
	parts := strings.Split(raw, ":")
	if len(parts) < 2 {
		return daemon.ScheduleSpec{}, fmt.Errorf("schedule must include kind and target")
	}
	out := daemon.ScheduleSpec{Kind: strings.TrimSpace(parts[0])}
	for _, p := range parts[1:] {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if p == "motd" {
			out.MOTD = true
			continue
		}
		if strings.HasPrefix(p, "before=") {
			out.Before = strings.TrimPrefix(p, "before=")
			continue
		}
		if strings.HasPrefix(p, "every=") {
			out.Every = strings.TrimPrefix(p, "every=")
			continue
		}
		if strings.HasPrefix(p, "sink=") {
			out.SinkID = append(out.SinkID, strings.TrimPrefix(p, "sink="))
			continue
		}
	}
	if !out.MOTD && len(out.SinkID) == 0 {
		out.MOTD = true
	}
	return out, nil
}
