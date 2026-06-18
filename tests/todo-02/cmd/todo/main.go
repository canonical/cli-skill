package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
	root := &cobra.Command{
		Use:   "todo",
		Short: "Todo client",
		Long:  "Todo client for managing tasks with deadlines, delivery sinks, and schedules.",
	}

	// Add command groups for help organization
	root.AddGroup(&cobra.Group{ID: "todos", Title: "Todos:"})
	root.AddGroup(&cobra.Group{ID: "sinks", Title: "Sinks:"})
	root.AddGroup(&cobra.Group{ID: "schedules", Title: "Schedules:"})
	root.AddGroup(&cobra.Group{ID: "other", Title: "Other:"})
	root.SetHelpCommandGroupID("other")
	root.SetCompletionCommandGroupID("other")

	// Set custom help function with color formatting
	root.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(colorizedHelp(cmd))
	})

	todosTopicCmd := &cobra.Command{
		Use:   "todos",
		Short: "What todos mean in this app",
		Long: `Todos are the main things you are tracking in the app.

A todo represents a task, deadline, or commitment you want to keep visible until it is done, rejected, or no longer relevant.

From a user perspective:
- create a todo when you want to track a single piece of work
- give it a due date when timing matters
- update it as details change
- close it when the work is finished
- reject it when you decide it should not be done

Schedules and sinks build on top of todos. A schedule decides when reminders should happen, and a sink decides where those reminders should go.`,
	}

	newClient := func(cmd *cobra.Command) *client.Client {
		return client.New(10 * time.Second)
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List todos",
		RunE: func(cmd *cobra.Command, args []string) error {
			state, _ := cmd.Flags().GetString("state")
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			ctx := context.Background()
			todos, err := cli.ListTodos(ctx, state)
			if err != nil {
				return err
			}
			return printTodos(todos, fmt, rfc)
		},
	}
	listCmd.Flags().String("state", "", "Filter state: open|closed|reopened|rejected")
	addOutputFlags(listCmd)

	showCmd := &cobra.Command{
		Use:   "show <todo-id>",
		Short: "Show a todo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			todo, err := cli.ShowTodo(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, fmt, rfc)
		},
	}
	addOutputFlags(showCmd)

	createCmd := &cobra.Command{
		Use:   "create <title>",
		Short: "Create a todo",
		Args:  cobra.MinimumNArgs(1),
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
				Title: strings.Join(args, " "),
				DueAt: dueAt,
			}
			for i, raw := range scheduleFlags {
				spec, err := parseInlineSchedule(raw)
				if err != nil {
					return fmt.Errorf("invalid --schedule at #%d: %w", i+1, err)
				}
				req.Schedule = append(req.Schedule, spec)
			}
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			todo, err := cli.CreateTodo(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, fmt, rfc)
		},
	}
	createCmd.Flags().String("due", "", "Optional due date (RFC3339 or human-readable)")
	createCmd.Flags().StringArray("schedule", nil, "Optional schedule spec (upcoming[:before=24h]:motd|sink=<id>, overdue[:every=1d]:motd|sink=<id>)")
	addOutputFlags(createCmd)

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
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			todo, err := cli.UpdateTodo(context.Background(), args[0], req)
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, fmt, rfc)
		},
	}
	updateCmd.Flags().String("title", "", "New title")
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date")
	addOutputFlags(updateCmd)

	closeCmd := todoActionCmd("close <todo-id>", "Close a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.CloseTodo(context.Background(), id)
	}, newClient)
	reopenCmd := todoActionCmd("reopen <todo-id>", "Reopen a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.ReopenTodo(context.Background(), id)
	}, newClient)
	rejectCmd := todoActionCmd("reject <todo-id>", "Reject a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.RejectTodo(context.Background(), id)
	}, newClient)

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show todo system status",
		RunE: func(cmd *cobra.Command, args []string) error {
			outFmt, _ := outputFlags(cmd)
			cli := newClient(cmd)
			payload, err := cli.Status(context.Background())
			if err != nil {
				return err
			}
			if outFmt == "json" {
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
	statusCmd.Flags().String("format", "table", "Output format: table|json")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show client version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	reminderStatusCmd := &cobra.Command{
		Use:   "reminder-status",
		Short: "Print pending MOTD reminder messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient(cmd)
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
		Use:   "list-sinks",
		Short: "List sinks",
		Long: `Sinks are reminder destinations.

In the todo app, a sink is where a reminder gets delivered outside the local terminal experience.

From a user perspective:
- create a sink when you want reminders sent to another system
- point the sink at a webhook URL
- subscribe it to upcoming or overdue events
- attach schedules to one or more sinks to decide where notifications go

Use the sinks command to inspect the sinks you have configured.`,
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
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			sinks, err := cli.ListSinks(context.Background(), enabled, event)
			if err != nil {
				return err
			}
			return printJSONOrTable(sinks, fmt, rfc)
		},
	}
	sinksCmd.Flags().String("enabled", "", "Filter enabled: true|false")
	sinksCmd.Flags().String("event", "", "Filter event: upcoming|overdue")
	addOutputFlags(sinksCmd)

	sinkCmd := &cobra.Command{
		Use:   "sink <sink-id>",
		Short: "Show a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			sink, err := cli.ShowSink(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, fmt, rfc)
		},
	}
	addOutputFlags(sinkCmd)

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
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			sink, err := cli.CreateSink(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(sink, fmt, rfc)
		},
	}
	createSinkCmd.Flags().String("url", "", "Webhook URL")
	createSinkCmd.Flags().StringArray("event", nil, "Event subscriptions (upcoming|overdue)")
	addOutputFlags(createSinkCmd)

	deleteSinkCmd := &cobra.Command{
		Use:   "delete-sink <sink-id>",
		Short: "Delete a sink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient(cmd)
			if err := cli.DeleteSink(context.Background(), args[0]); err != nil {
				return err
			}
			fmt.Println("sink deleted")
			return nil
		},
	}

	// Schedule commands
	schedulesCmd := &cobra.Command{
		Use:   "schedules",
		Short: "List schedules",
		Long: `Schedules define when reminders should be sent for a todo.

In the todo app, a schedule connects a todo to one or more reminder times and destinations.

From a user perspective:
- use an upcoming schedule to remind yourself before a due date
- use an overdue schedule to keep reminding after a due date has passed
- send reminders to MOTD, to sinks, or both

Schedules make todos active over time instead of being just stored records.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			todoID, _ := cmd.Flags().GetString("todo")
			kind, _ := cmd.Flags().GetString("kind")
			status, _ := cmd.Flags().GetString("status")
			target, _ := cmd.Flags().GetString("target")
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			schedules, err := cli.ListSchedules(context.Background(), todoID, kind, status, target)
			if err != nil {
				return err
			}
			return printJSONOrTable(schedules, fmt, rfc)
		},
	}
	schedulesCmd.Flags().String("todo", "", "Optional todo id filter")
	schedulesCmd.Flags().String("kind", "", "Filter kind: upcoming|overdue")
	schedulesCmd.Flags().String("status", "", "Filter status: active|sent")
	schedulesCmd.Flags().String("target", "", "Filter target: sink|motd")
	addOutputFlags(schedulesCmd)

	scheduleCmd := &cobra.Command{
		Use:   "schedule <schedule-id>",
		Short: "Show a schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			sc, err := cli.ShowSchedule(context.Background(), args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(sc, fmt, rfc)
		},
	}
	addOutputFlags(scheduleCmd)

	addScheduleCmd := &cobra.Command{
		Use:   "add-schedule <schedule-id>",
		Short: "Add an immutable schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			todoIDStr, _ := cmd.Flags().GetString("todo")
			kind, _ := cmd.Flags().GetString("kind")
			before, _ := cmd.Flags().GetString("before")
			every, _ := cmd.Flags().GetString("every")
			sinks, _ := cmd.Flags().GetStringArray("sink")
			motd, _ := cmd.Flags().GetBool("motd")
			if strings.TrimSpace(todoIDStr) == "" {
				return fmt.Errorf("--todo is required")
			}
			todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid todo id: %w", err)
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
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			sc, err := cli.AddSchedule(context.Background(), req)
			if err != nil {
				return err
			}
			return printJSONOrTable(sc, fmt, rfc)
		},
	}
	addScheduleCmd.Flags().String("todo", "", "Todo id")
	addScheduleCmd.Flags().String("kind", "", "Schedule kind: upcoming|overdue")
	addScheduleCmd.Flags().String("before", "", "Upcoming schedule offset before due date, default 24h")
	addScheduleCmd.Flags().String("every", "", "Overdue reminder frequency, default 1d")
	addScheduleCmd.Flags().StringArray("sink", nil, "Optional sink id (repeatable)")
	addScheduleCmd.Flags().Bool("motd", false, "Deliver via shell MOTD channel")
	addOutputFlags(addScheduleCmd)

	removeScheduleCmd := &cobra.Command{
		Use:   "remove-schedule <schedule-id>",
		Short: "Remove a schedule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := newClient(cmd)
			if err := cli.RemoveSchedule(context.Background(), args[0]); err != nil {
				return err
			}
			fmt.Println("schedule removed")
			return nil
		},
	}

	// Set command groups for better help organization
	listCmd.GroupID = "todos"
	showCmd.GroupID = "todos"
	createCmd.GroupID = "todos"
	updateCmd.GroupID = "todos"
	closeCmd.GroupID = "todos"
	reopenCmd.GroupID = "todos"
	rejectCmd.GroupID = "todos"

	sinksCmd.GroupID = "sinks"
	sinkCmd.GroupID = "sinks"
	createSinkCmd.GroupID = "sinks"
	deleteSinkCmd.GroupID = "sinks"

	schedulesCmd.GroupID = "schedules"
	scheduleCmd.GroupID = "schedules"
	addScheduleCmd.GroupID = "schedules"
	removeScheduleCmd.GroupID = "schedules"

	reminderStatusCmd.GroupID = "other"
	statusCmd.GroupID = "other"
	versionCmd.GroupID = "other"
	root.AddCommand(listCmd, showCmd, createCmd, updateCmd, closeCmd, reopenCmd, rejectCmd)
	root.AddCommand(sinksCmd, sinkCmd, createSinkCmd, deleteSinkCmd)
	root.AddCommand(schedulesCmd, scheduleCmd, addScheduleCmd, removeScheduleCmd)
	root.AddCommand(reminderStatusCmd, statusCmd, versionCmd)
	root.AddCommand(todosTopicCmd)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func todoActionCmd(use, short string, fn func(*client.Client, string) (model.Todo, error), newClient func(*cobra.Command) *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt, rfc := outputFlags(cmd)
			cli := newClient(cmd)
			todo, err := fn(cli, args[0])
			if err != nil {
				return err
			}
			return printJSONOrTable(todo, fmt, rfc)
		},
	}
	addOutputFlags(cmd)
	return cmd
}

// addOutputFlags registers --format and --rfc3339 on a command.
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("format", "table", "Output format: table|json")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}

// outputFlags reads --format and --rfc3339 from a command.
// json and yaml formats imply rfc3339.
func outputFlags(cmd *cobra.Command) (format string, rfc3339 bool) {
	format, _ = cmd.Flags().GetString("format")
	rfc3339, _ = cmd.Flags().GetBool("rfc3339")
	if format == "json" || format == "yaml" {
		rfc3339 = true
	}
	return
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
	
	// Build rows with formatted data
	type row struct {
		id    string
		state string
		due   string
		title string
	}
	var rows []row
	for _, t := range todos {
		var dueStr string
		if rfc3339 {
			dueStr = common.FormatTime(t.DueAt, true)
		} else {
			dueStr = common.FormatRelativeTime(t.DueAt)
		}
		rows = append(rows, row{
			id:    fmt.Sprintf("%d", t.ID),
			state: t.State,
			due:   dueStr,
			title: t.Title,
		})
	}
	
	// Calculate column widths (minimum width for header)
	idWidth := len("ID")
	stateWidth := len("STATE")
	dueWidth := len("DUE")
	
	for _, r := range rows {
		if len(r.id) > idWidth {
			idWidth = len(r.id)
		}
		if len(r.state) > stateWidth {
			stateWidth = len(r.state)
		}
		if len(r.due) > dueWidth {
			dueWidth = len(r.due)
		}
	}
	
	// Print header with 2-space separator
	sep := "  "
	fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
	
	// Print rows
	for _, r := range rows {
		fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, r.id, sep, stateWidth, r.state, sep, dueWidth, r.due, sep, r.title)
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
		fmt.Printf("id: %d\n", t.ID)
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

// colorizedHelp takes a cobra command and returns its help text with color formatting
func colorizedHelp(cmd *cobra.Command) string {
	if cmd.Parent() == nil {
		return colorizeSections(rootHelp(cmd))
	}

	var b strings.Builder
	b.WriteString("<b>")
	b.WriteString(cmd.Root().Name())
	b.WriteString(" ")
	b.WriteString(version)
	b.WriteString("</b>\n\n")
	if cmd.Long != "" {
		b.WriteString(cmd.Long)
		b.WriteString("\n\n")
	} else if cmd.Short != "" {
		b.WriteString(cmd.Short)
		b.WriteString("\n\n")
	}
	b.WriteString(cmd.UsageString())
	help := b.String()
	if cmd.Parent() == nil {
		help = strings.Replace(help, "\nFlags:\n", "\nGlobal options:\n", 1)
	}

	return colorizeSections(help)
}

func rootHelp(cmd *cobra.Command) string {
	var b strings.Builder
	b.WriteString("<b>")
	b.WriteString(cmd.Root().Name())
	b.WriteString(" ")
	b.WriteString(version)
	b.WriteString("</b>\n\n")
	if cmd.Long != "" {
		b.WriteString(cmd.Long)
		b.WriteString("\n\n")
	} else if cmd.Short != "" {
		b.WriteString(cmd.Short)
		b.WriteString("\n\n")
	}
	b.WriteString("Usage:\n")
	b.WriteString("  todo [command]\n\n")

	writeCommandSection(&b, cmd, "Todos:", "todos")
	b.WriteString("\n")
	writeCommandSection(&b, cmd, "Schedules:", "schedules")
	b.WriteString("\n")
	b.WriteString("Other:\n")
	b.WriteString(formatOtherSection(cmd))
	b.WriteString("\nGlobal options:\n")
	b.WriteString("  -h, --help   help for todo\n\n")
	b.WriteString("Use \"todo [command] --help\" for more information about a command.\n")
	b.WriteString("Use \"todo help <topic>\" for more information about a concept, for example \"todo help todos\".\n")
	return b.String()
}

func writeCommandSection(b *strings.Builder, cmd *cobra.Command, title, groupID string) {
	b.WriteString(title)
	b.WriteString("\n")
	entries := sectionEntries(cmd, groupID)
	width := longestName(entries)
	if width < 15 {
		width = 15
	}
	for _, entry := range entries {
		b.WriteString(fmt.Sprintf("  %-*s %s\n", width, entry.Name, entry.Short))
	}
}

type helpEntry struct {
	Name  string
	Short string
}

func sectionEntries(cmd *cobra.Command, groupID string) []helpEntry {
	entries := make([]helpEntry, 0)
	for _, sub := range cmd.Commands() {
		if !sub.IsAvailableCommand() || sub.IsAdditionalHelpTopicCommand() {
			continue
		}
		if sub.GroupID == groupID {
			entries = append(entries, helpEntry{Name: sub.Name(), Short: sub.Short})
		}
	}
	return entries
}

func longestName(entries []helpEntry) int {
	width := 0
	for _, entry := range entries {
		if len(entry.Name) > width {
			width = len(entry.Name)
		}
	}
	return width
}

func formatOtherSection(cmd *cobra.Command) string {
	var b strings.Builder
	width := 15
	b.WriteString(fmt.Sprintf("  %-*s %s\n", width, "Sinks", "create-sink, delete-sink, sink, sinks"))
	for _, name := range []string{"completion", "help", "reminder-status", "status", "version"} {
		entry, ok := findEntry(cmd, name)
		if !ok {
			if name == "help" {
				b.WriteString(fmt.Sprintf("  %-*s %s\n", width, "help", "Help about any command"))
			}
			continue
		}
		b.WriteString(fmt.Sprintf("  %-*s %s\n", width, entry.Name, entry.Short))
	}
	return b.String()
}

func findEntry(cmd *cobra.Command, name string) (helpEntry, bool) {
	for _, sub := range cmd.Commands() {
		if !sub.IsAvailableCommand() {
			continue
		}
		if sub.Name() == name {
			return helpEntry{Name: sub.Name(), Short: sub.Short}, true
		}
	}
	return helpEntry{}, false
}

func colorizeSections(help string) string {
	// Protect overlapping labels before generic replacements.
	help = strings.ReplaceAll(help, "Global Flags:", "__TODO_GLOBAL_FLAGS__")

	// Apply color formatting to section headers
	sections := []string{
		"Usage:",
		"Summary:",
		"Todos:",
		"Sinks:",
		"Schedules:",
		"Other:",
		"Available Commands:",
		"Additional Commands:",
		"Global options:",
		"Flags:",
		"Examples:",
		"Related commands:",
	}

	// Apply color formatting to each section header
	for _, section := range sections {
		colored := common.ColorSection(section)
		help = strings.ReplaceAll(help, section, colored)
	}

	help = strings.ReplaceAll(help, "__TODO_GLOBAL_FLAGS__", common.ColorSection("Global Flags:"))
	return common.RenderInlineTags(help)
}
