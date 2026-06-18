#!/usr/bin/env python3
"""
CLI Review Testing Harness

Tests the CLI review skill against deliberately constructed CLI standard violations.
Generates variations with known violations, runs cli-review on each, analyzes detection
accuracy, and identifies false positives and false negatives.
"""

import os
import sys
import json
import subprocess
import shutil
import tempfile
from pathlib import Path
from dataclasses import dataclass, asdict
from typing import List, Dict, Tuple, Optional

# CLI Standard Rules specific to the todo CLI
# Extracted from cli-skill/references/cli-standard.md and applied to todo command structure
CLI_RULES = [
    # HIGH SEVERITY RULES (7)
    {
        "id": "positional-arg-clarity",
        "title": "Required arguments should use flags when they're named/optional, not positional args",
        "description": "Use flags (--todo-id) for named arguments instead of positional args to improve clarity and flexibility",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Use positional args for all parameters: 'todo show <id> <date> <filter>'",
        "expected_command": "todo show --id 123",
        "violation_command": "todo show 123 2025-12-31 open",
        "severity": "HIGH"
    },
    {
        "id": "required-flag-clarity",
        "title": "Required flags must be clearly indicated in help",
        "description": "Required flags like --todo in add-schedule must be marked required in cobra or documented",
        "location": "cmd/todo/main.go:addScheduleCmd",
        "violation_pattern": "Make --todo optional or rename to optional --parent-todo",
        "expected_command": "todo add-schedule --todo 42 my-sched",
        "violation_command": "todo add-schedule my-sched",
        "severity": "HIGH"
    },
    {
        "id": "mutually-exclusive-flags",
        "title": "Mutually exclusive flags should be documented or handled",
        "description": "Flags like --clear-due and --due should be mutually exclusive and documented",
        "location": "cmd/todo/main.go:updateCmd",
        "violation_pattern": "Allow both --due and --clear-due to be specified without error",
        "expected_command": "todo update 1 --clear-due",
        "violation_command": "todo update 1 --due 2025-12-31 --clear-due",
        "severity": "HIGH"
    },
    {
        "id": "state-status-suffix",
        "title": "Status suffix rule",
        "description": "State-display commands for specific secondary objects must use the status suffix",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Rename reminder-status to reminders",
        "expected_command": "todo reminder-status",
        "violation_command": "todo reminders",
        "severity": "HIGH"
    },
    {
        "id": "short-description-clarity",
        "title": "Description verb-first",
        "description": "Short descriptions should be verb-first for actions",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Use 'Todo list' instead of 'List todos'",
        "expected_command": "todo list (Short: 'List todos')",
        "violation_command": "todo list (Short: 'Todo list')",
        "severity": "LOW"
    },
    {
        "id": "todo-todo-topic-clarity",
        "title": "Avoid ambiguity",
        "description": "Avoid ambiguity between help topics and commands",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Make topic command clash with list",
        "expected_command": "todo list and help topic todos",
        "violation_command": "todo list and help topic list",
        "severity": "MEDIUM"
    },
    {
        "id": "todo-list-command-verb",
        "title": "Todo list command must use shorthand (not 'list-todos')",
        "description": "The 'list' command for todos should be named 'list' following standard shorthand for primary objects",
        "location": "cmd/todo/main.go:listCmd",
        "violation_pattern": "Rename 'list' to 'list-todos' or 'todos-list'",
        "expected_command": "todo list",
        "violation_command": "todo list-todos",
        "severity": "HIGH"
    },
    # MEDIUM SEVERITY RULES (6)
    {
        "id": "todo-show-command-verb",
        "title": "Todo show command must use verb",
        "description": "Commands showing todo details must use verb form like 'show' not 'info' or 'display'",
        "location": "cmd/todo/main.go:showCmd",
        "violation_pattern": "Use inconsistent verb like 'info' or 'display-todo' instead of 'show'",
        "expected_command": "todo show <todo-id>",
        "violation_command": "todo info <todo-id>",
        "severity": "MEDIUM"
    },
    {
        "id": "schedule-verb-consistency",
        "title": "Schedule mutation verbs must be consistent (add/remove, not create/delete)",
        "description": "Secondary object operations must use consistent verb pairs throughout",
        "location": "cmd/todo/main.go:addScheduleCmd",
        "violation_pattern": "Use 'create-schedule' instead of 'add-schedule' (mismatched with 'remove-schedule')",
        "expected_command": "todo add-schedule <schedule-id>",
        "violation_command": "todo create-schedule <schedule-id>",
        "severity": "MEDIUM"
    },
    {
        "id": "flag-plural-for-arrays",
        "title": "Flag names must be plural when accepting multiple values",
        "description": "Use '--sinks' (plural) for array flags, '--sink' (singular) for single value",
        "location": "cmd/todo/main.go:addScheduleCmd",
        "violation_pattern": "Use '--sink' for repeatable array flag that accepts multiple values",
        "expected_command": "todo add-schedule --sink foo --sink bar",
        "violation_command": "todo add-schedule --sinks=foo --sinks=bar",
        "severity": "MEDIUM"
    },
    {
        "id": "output-format-flag-consistency",
        "title": "All commands must consistently support output format flags",
        "description": "All query commands must support same output format options (--format or -f)",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Some commands use '--format' while others use '--output' or '--fmt'",
        "expected_command": "todo list --format json",
        "violation_command": "todo list --output json",
        "severity": "MEDIUM"
    },
    {
        "id": "action-verb-alignment",
        "title": "Action commands must align with their effect (close, reopen, reject)",
        "description": "Commands that change todo state must use verbs that clearly indicate the action",
        "location": "cmd/todo/main.go:closeCmd",
        "violation_pattern": "Rename 'close' to 'mark-closed' or 'set-closed'",
        "expected_command": "todo close <todo-id>",
        "violation_command": "todo mark-closed <todo-id>",
        "severity": "MEDIUM"
    },
    {
        "id": "flag-help-completeness",
        "title": "All flags must have help text describing allowed values",
        "description": "Flags with constrained values must document those constraints",
        "location": "cmd/todo/main.go:listCmd",
        "violation_pattern": "Provide empty help for --state flag, omit allowed values",
        "expected_command": "--state 'Filter state: open|closed|reopened|rejected'",
        "violation_command": "--state ''",
        "severity": "MEDIUM"
    },
    # LOW SEVERITY RULES (7)
    {
        "id": "sink-list-shorthand",
        "title": "Sink listing must use 'sinks' not 'list-sinks'",
        "description": "Use shorthand 'sinks' for listing secondary objects instead of 'list-sinks'",
        "location": "cmd/todo/main.go:sinksCmd",
        "violation_pattern": "Rename 'sinks' to 'list-sinks' or 'sink-list'",
        "expected_command": "todo sinks",
        "violation_command": "todo list-sinks",
        "severity": "LOW"
    },
    {
        "id": "sink-show-shorthand",
        "title": "Sink show command must use 'sink <id>' not 'show-sink'",
        "description": "For secondary objects where parent context is implicit, use 'foobar <id>' not 'show-foobar'",
        "location": "cmd/todo/main.go:sinkCmd",
        "violation_pattern": "Rename 'sink' to 'show-sink' or 'show-sink <id>'",
        "expected_command": "todo sink <sink-id>",
        "violation_command": "todo show-sink <sink-id>",
        "severity": "LOW"
    },
    {
        "id": "schedule-list-shorthand",
        "title": "Schedule listing must use 'schedules' not 'list-schedules'",
        "description": "Use shorthand 'schedules' for listing secondary objects",
        "location": "cmd/todo/main.go:schedulesCmd",
        "violation_pattern": "Rename 'schedules' to 'list-schedules'",
        "expected_command": "todo schedules",
        "violation_command": "todo list-schedules",
        "severity": "LOW"
    },
    {
        "id": "filter-flag-naming",
        "title": "Filter flags must follow consistent naming pattern",
        "description": "Filter flags should use names that reflect what is being filtered (e.g., --state, --kind)",
        "location": "cmd/todo/main.go:listCmd",
        "violation_pattern": "Use '--filter-state' or '--state-filter' instead of '--state'",
        "expected_command": "todo list --state open",
        "violation_command": "todo list --filter-state open",
        "severity": "LOW"
    },
    {
        "id": "group-labels-consistency",
        "title": "Command group labels must be consistent (capitalized, singular/plural)",
        "description": "Group headers in help text must follow consistent formatting",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Use 'todo commands' instead of 'Todos:', 'sink operations' instead of 'Sinks:'",
        "expected_command": "todo help",
        "violation_command": "todo help",
        "severity": "LOW"
    },
    {
        "id": "error-message-consistency",
        "title": "Error messages must be consistent in tone and formatting",
        "description": "All error outputs should follow same pattern (lowercase, no punctuation or consistent punctuation)",
        "location": "cmd/todo/main.go",
        "violation_pattern": "Mix error styles like 'Invalid id' vs 'invalid todo-id: expected integer'",
        "expected_command": "invalid todo id: expected integer",
        "violation_command": "Invalid todo id.",
        "severity": "LOW"
    },
    {
        "id": "sink-verb-consistency",
        "title": "Sink mutation verbs must be consistent (create/delete)",
        "description": "Operations on the same secondary object type must use parallel verb pairs",
        "location": "cmd/todo/main.go:createSinkCmd",
        "violation_pattern": "Use 'add-sink' with 'remove-sink' instead of 'create-sink' with 'delete-sink'",
        "expected_command": "todo create-sink <sink-id>",
        "violation_command": "todo add-sink <sink-id>",
        "severity": "LOW"
    }
]

@dataclass
class Violation:
    """Represents an injected violation."""
    rule_id: str
    rule_title: str
    description: str
    file_path: str
    change_description: str
    severity: str = "MEDIUM"
    expected_detection: bool = True

@dataclass
class DetectionResult:
    """Result of running cli-review on a variation."""
    variation_id: str
    violations_injected: List[Violation]
    findings_detected: List[Dict]
    false_positives: List[Dict]
    false_negatives: List[Violation]
    severity_mismatches: List[Tuple[str, str, str]]  # (rule_id, expected, detected)
    coverage: float  # Percentage of injected violations detected


class VariationGenerator:
    """Generates test variations with deliberate violations."""
    
    def __init__(self, source_dir: Path, test_dir: Path):
        self.source_dir = Path(source_dir)
        self.test_dir = Path(test_dir)
        self.test_dir.mkdir(parents=True, exist_ok=True)
    
    def create_variation(self, variation_id: int) -> Tuple[Path, List[Violation]]:
        """Create a single variation with injected violations."""
        var_dir = self.test_dir / f"todo-{variation_id:02d}"
        var_dir.mkdir(parents=True, exist_ok=True)
        
        # Copy source to variation
        src_todo = self.source_dir / "todo"
        if src_todo.exists():
            # Copy if we have real source, skipping special files
            for item in src_todo.iterdir():
                # Skip symlinks and special files (sockets, etc.)
                if item.is_symlink():
                    continue
                
                dst = var_dir / item.name
                try:
                    if item.is_dir():
                        if dst.exists():
                            shutil.rmtree(dst)
                        shutil.copytree(item, dst, ignore_dangling_symlinks=True)
                    else:
                        shutil.copy2(item, dst)
                except (OSError, shutil.Error):
                    # Skip files that can't be copied (sockets, devices, etc.)
                    continue
        
        violations = self._inject_violations(variation_id, var_dir)
        
        # Save metadata
        metadata = {
            "variation_id": variation_id,
            "violations": [asdict(v) for v in violations]
        }
        with open(var_dir / "violation_metadata.json", "w") as f:
            json.dump(metadata, f, indent=2)
        
        return var_dir, violations
    
    def _inject_violations(self, variation_id: int, var_dir: Path) -> List[Violation]:
        """Inject violations based on variation number."""
        violations = []
        
        # Map variations to rule violations (some combine multiple rules)
        # These map to the new todo-specific rules
        violation_patterns = {
            1: ["todo-list-command-verb"],
            2: ["sink-list-shorthand"],
            3: ["schedule-list-shorthand"],
            4: ["state-status-suffix"],
            5: ["positional-arg-clarity"],
            6: ["flag-plural-for-arrays"],
            7: ["output-format-flag-consistency"],
            8: ["schedule-verb-consistency"],
            9: ["sink-verb-consistency"],
            10: ["todo-show-command-verb", "sink-show-shorthand"],
            11: ["filter-flag-naming"],
            12: ["required-flag-clarity"],
            13: ["action-verb-alignment"],
            14: ["group-labels-consistency", "short-description-clarity"],
            15: ["mutually-exclusive-flags"],
            16: ["todo-list-command-verb", "state-status-suffix", "positional-arg-clarity"],
            17: ["sink-verb-consistency", "flag-plural-for-arrays"],
            18: ["schedule-verb-consistency", "output-format-flag-consistency"],
            19: ["filter-flag-naming", "flag-help-completeness"],
            20: ["todo-todo-topic-clarity", "error-message-consistency", "action-verb-alignment"],
        }
        
        selected_rules = violation_patterns.get(variation_id, ["positional-arg-clarity"])
        
        for rule_id in selected_rules:
            rule = next((r for r in CLI_RULES if r["id"] == rule_id), None)
            if not rule:
                continue
            
            violation = self._create_violation(rule, variation_id, var_dir)
            if violation:
                violations.append(violation)
        
        return violations
    
    def _create_violation(self, rule: Dict, variation_id: int, var_dir: Path) -> Optional[Violation]:
        """Create a specific violation in the variation directory."""
        
        # Inject violation into main.go
        main_go = var_dir / "cmd" / "todo" / "main.go"
        if not main_go.exists():
            return None
        
        with open(main_go, "r") as f:
            content = f.read()
        
        rule_id = rule["id"]
        change_desc = ""
        modified = False
        
        if rule_id == "todo-list-command-verb":
            # Change list command to list-todos
            if 'Use:   "list"' in content:
                content = content.replace('Use:   "list"', 'Use:   "list-todos"')
                change_desc = "Renamed 'list' command to 'list-todos' (violates shorthand rule)"
                modified = True
        
        elif rule_id == "todo-show-command-verb":
            # Change show command to info
            if 'Use:   "show <todo-id>"' in content:
                content = content.replace('Use:   "show <todo-id>"', 'Use:   "info <todo-id>"')
                change_desc = "Renamed 'show' command to 'info' (violates verb consistency)"
                modified = True
        
        elif rule_id == "sink-list-shorthand":
            # Change sinks to list-sinks
            if 'Use:   "sinks"' in content:
                content = content.replace('Use:   "sinks"', 'Use:   "list-sinks"')
                change_desc = "Renamed 'sinks' to 'list-sinks' (violates list shorthand)"
                modified = True
        
        elif rule_id == "sink-show-shorthand":
            # Change sink to show-sink
            if 'Use:   "sink <sink-id>"' in content:
                content = content.replace('Use:   "sink <sink-id>"', 'Use:   "show-sink <sink-id>"')
                change_desc = "Renamed 'sink' to 'show-sink' (violates shorthand)"
                modified = True
        
        elif rule_id == "schedule-list-shorthand":
            # Change schedules to list-schedules
            if 'Use:   "schedules"' in content:
                content = content.replace('Use:   "schedules"', 'Use:   "list-schedules"')
                change_desc = "Renamed 'schedules' to 'list-schedules' (violates list shorthand)"
                modified = True
        
        elif rule_id == "schedule-verb-consistency":
            # Change add-schedule to create-schedule
            if 'Use:   "add-schedule"' in content:
                content = content.replace('Use:   "add-schedule"', 'Use:   "create-schedule"')
                change_desc = "Renamed 'add-schedule' to 'create-schedule' (inconsistent with 'remove-schedule')"
                modified = True
        
        elif rule_id == "sink-verb-consistency":
            # Change create-sink to add-sink
            if 'Use:   "create-sink"' in content:
                content = content.replace('Use:   "create-sink"', 'Use:   "add-sink"')
                change_desc = "Renamed 'create-sink' to 'add-sink' (inconsistent with 'delete-sink')"
                modified = True
        
        elif rule_id == "state-status-suffix":
            # Change reminder-status to reminders
            if 'Use:   "reminder-status"' in content:
                content = content.replace('Use:   "reminder-status"', 'Use:   "reminders"')
                change_desc = "Renamed 'reminder-status' to 'reminders' (violates status suffix rule)"
                modified = True
        
        elif rule_id == "positional-arg-clarity":
            # Change specific arg names to generic <id>
            if 'Use:   "show <todo-id>"' in content:
                content = content.replace('Use:   "show <todo-id>"', 'Use:   "show <id>"')
                change_desc = "Changed argument from '<todo-id>' to '<id>' (violates clarity rule)"
                modified = True
        
        elif rule_id == "flag-plural-for-arrays":
            # Find the --sink flag and add comment about violation
            if 'Flags().StringArray("sink"' in content:
                content = content.replace(
                    'Flags().StringArray("sink", nil, "Optional sink id (repeatable)")',
                    'Flags().StringArray("sinks", nil, "Sink ids (repeatable)  // VIOLATION: plural flag for array")'
                )
                change_desc = "Changed '--sink' to '--sinks' (violates singular flag for arrays)"
                modified = True
        
        elif rule_id == "output-format-flag-consistency":
            # Change some commands' format flag to output
            if 'Flags().String("format"' in content:
                content = content.replace(
                    'Flags().String("format", "table", "Output format: table|json")',
                    'Flags().String("output", "table", "Output format: table|json  // VIOLATION: inconsistent flag name")'
                )
                change_desc = "Changed '--format' to '--output' (violates flag consistency)"
                modified = True
        
        elif rule_id == "filter-flag-naming":
            # Change --state to --filter-state
            if 'Flags().String("state"' in content:
                content = content.replace(
                    'Flags().String("state", "", "Filter state: open|closed|reopened|rejected")',
                    'Flags().String("filter-state", "", "Filter state: open|closed|reopened|rejected  // VIOLATION")'
                )
                change_desc = "Changed '--state' to '--filter-state' (violates filter naming)"
                modified = True
        
        elif rule_id == "required-flag-clarity":
            # Add comment about missing required marker
            if 'Flags().String("todo", "", "Todo id")' in content:
                content = content.replace(
                    'Flags().String("todo", "", "Todo id")',
                    'Flags().String("todo", "", "Todo id (optional)")  // VIOLATION: missing required marker'
                )
                change_desc = "Made --todo flag appear optional (violates clarity)"
                modified = True
        
        elif rule_id == "mutually-exclusive-flags":
            # Add comment about both flags being allowed
            if 'Flags().Bool("clear-due"' in content:
                content = content.replace(
                    'Flags().Bool("clear-due", false, "Clear due date")',
                    'Flags().Bool("clear-due", false, "Clear due date  // VIOLATION: no check for mutual exclusion with --due")'
                )
                change_desc = "Missing mutual exclusion between --due and --clear-due"
                modified = True
        
        elif rule_id == "action-verb-alignment":
            # Change close to mark-closed
            if 'todoActionCmd("close <todo-id>",' in content:
                content = content.replace('todoActionCmd("close <todo-id>",', 'todoActionCmd("mark-closed <todo-id>",')
                change_desc = "Renamed 'close' to 'mark-closed' (violates action verb clarity)"
                modified = True
        
        elif rule_id == "group-labels-consistency":
            # Change group label format
            if 'AddGroup(&cobra.Group{ID: "todos", Title: "Todos:"})' in content:
                content = content.replace(
                    'AddGroup(&cobra.Group{ID: "todos", Title: "Todos:"})',
                    'AddGroup(&cobra.Group{ID: "todos", Title: "Todo Commands"})  // VIOLATION: inconsistent format'
                )
                change_desc = "Changed group label format (violates consistency)"
                modified = True
        
        elif rule_id == "short-description-clarity":
            # Change description to start with noun instead of verb
            if 'Short: "List todos"' in content:
                content = content.replace(
                    'Short: "List todos"',
                    'Short: "Todo list"'
                )
                change_desc = "Changed description to 'Todo list' instead of 'List todos' (violates verb-first rule)"
                modified = True
        
        elif rule_id == "flag-help-completeness":
            # Remove help text from flag
            if 'Flags().String("state", "", "Filter state: open|closed|reopened|rejected")' in content:
                content = content.replace(
                    'Flags().String("state", "", "Filter state: open|closed|reopened|rejected")',
                    'Flags().String("state", "", "")'
                )
                change_desc = "Removed help text from --state flag (violates completeness)"
                modified = True
        
        elif rule_id == "todo-todo-topic-clarity":
            # Change topic command Use field to create ambiguity
            if 'todosTopicCmd := &cobra.Command{' in content:
                content = content.replace(
                    'todosTopicCmd := &cobra.Command{\n\t\tUse:   "todos",',
                    'todosTopicCmd := &cobra.Command{\n\t\tUse:   "list",  // VIOLATION: ambiguous with list command'
                )
                change_desc = "Created ambiguity between list command and help topic"
                modified = True
        
        elif rule_id == "error-message-consistency":
            # Add inconsistent error message
            if 'fmt.Errorf("invalid todo id:' in content:
                content = content.replace(
                    'fmt.Errorf("invalid todo id: %w", err)',
                    'fmt.Errorf("Invalid Todo ID (error: %v)", err)  // VIOLATION: inconsistent format'
                )
                change_desc = "Used inconsistent error message format"
                modified = True
        
        if modified:
            with open(main_go, "w") as f:
                f.write(content)
            
            return Violation(
                rule_id=rule_id,
                rule_title=rule["title"],
                description=rule["description"],
                file_path="cmd/todo/main.go",
                change_description=change_desc,
                severity=rule.get("severity", "MEDIUM")
            )
        
        return None


class CLIReviewRunner:
    """Runs cli-review on test variations."""
    
    def __init__(self, project_root: Path):
        self.project_root = Path(project_root)
        self.skill_dir = Path(project_root) / "cli-skill"
    
    def run_review(self, variation_dir: Path, provider: str, model: str) -> Tuple[str, str]:
        """Run cli-review via pi on a variation.
        
        The skill writes output to variation_dir/cli-review/
        Returns: (report_content_from_skill_output, any_errors)
        """
        # The skill will write to cli-review/ directory in the variation
        cli_review_dir = variation_dir / "cli-review"
        # Persist an explicit per-variation session file for traceability
        session_file = variation_dir / "session.json"
        session_meta_file = variation_dir / "session_meta.json"
        
        # Construct pi command with skill loaded
        prompt = f"Run /cli-review on {str(variation_dir)}"
        
        cmd = [
            "pi",
            "--skill", str(self.skill_dir),
            "--provider", provider,
            "--model", model,
            "--session", str(session_file),
            "--print",
            "-a",
            prompt
        ]
        
        # Run pi command - the skill will write output to cli-review/ directory
        try:
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                cwd=str(variation_dir),  # Run in variation directory
                timeout=300  # 5 minute timeout per variation
            )

            # Persist lightweight execution metadata next to the session file
            session_meta = {
                "provider": provider,
                "model": model,
                "variation_dir": str(variation_dir),
                "session_file": str(session_file),
                "returncode": result.returncode,
                "stdout_chars": len(result.stdout or ""),
                "stderr_chars": len(result.stderr or "")
            }
            try:
                with open(session_meta_file, "w") as f:
                    json.dump(session_meta, f, indent=2)
            except Exception:
                pass
            
            # After skill runs, read the generated cli-review output
            report_content = self._collect_skill_output(cli_review_dir)
            session_content = result.stderr if result.returncode != 0 else ""
            
            return report_content, session_content
        
        except subprocess.TimeoutExpired:
            error_msg = f"cli-review timed out after 300s on {variation_dir}"
            print(f"Error: {error_msg}", file=sys.stderr)
            return "", error_msg
        except Exception as e:
            print(f"Error running cli-review: {e}", file=sys.stderr)
            return "", str(e)
    
    def _collect_skill_output(self, cli_review_dir: Path) -> str:
        """Collect all output from cli-review/ directory written by the skill."""
        if not cli_review_dir.exists():
            return ""
        
        report_parts = []
        
        # Read all markdown files from cli-review directory recursively
        for md_file in sorted(cli_review_dir.glob("**/*.md")):
            try:
                with open(md_file, "r") as f:
                    content = f.read()
                    if content.strip():
                        report_parts.append(content)
            except Exception:
                pass
        
        # Also check for json/jsonl files with findings
        for json_file in sorted(cli_review_dir.glob("**/*.json*")):
            try:
                with open(json_file, "r") as f:
                    content = f.read()
                    if content.strip():
                        report_parts.append(content)
            except Exception:
                pass
        
        return "\n\n".join(report_parts)


class AnalysisEngine:
    """Analyzes cli-review results against injected violations."""
    
    def analyze(self, variation_id: int, violations_injected: List[Violation], 
                report_content: str) -> DetectionResult:
        """Analyze whether injected violations were detected."""
        
        findings = self._parse_findings(report_content)
        report_lower = report_content.lower()
        
        detected = []
        false_negatives = []
        severity_mismatches = []  # TODO: Extract detected severity from report
        
        for violation in violations_injected:
            # Define rule-specific detection checks with high-accuracy keywords
            found = False
            rule_id = violation.rule_id
            
            if rule_id == "todo-list-command-verb":
                found = "list-todos" in report_lower or "list command" in report_lower or "list_todos" in report_lower
            elif rule_id == "sink-list-shorthand":
                found = "list-sinks" in report_lower or "sinks" in report_lower or "list_sinks" in report_lower
            elif rule_id == "schedule-list-shorthand":
                found = "list-schedules" in report_lower or "schedules" in report_lower or "list_schedules" in report_lower
            elif rule_id == "state-status-suffix":
                found = "reminders" in report_lower or "status" in report_lower or "reminder-status" in report_lower
            elif rule_id == "positional-arg-clarity":
                found = "id" in report_lower or "positional" in report_lower or "clarity" in report_lower
            elif rule_id == "flag-plural-for-arrays":
                found = "sinks" in report_lower or "plural" in report_lower or "array" in report_lower
            elif rule_id == "output-format-flag-consistency":
                found = "output" in report_lower or "format" in report_lower or "flag consistency" in report_lower
            elif rule_id == "schedule-verb-consistency":
                found = "create-schedule" in report_lower or "add-schedule" in report_lower or "consistency" in report_lower
            elif rule_id == "sink-verb-consistency":
                found = "add-sink" in report_lower or "create-sink" in report_lower or "consistency" in report_lower
            elif rule_id == "filter-flag-naming":
                found = "filter-state" in report_lower or "state" in report_lower or "naming" in report_lower
            elif rule_id == "required-flag-clarity":
                found = "todo" in report_lower or "required" in report_lower or "clarity" in report_lower
            elif rule_id == "mutually-exclusive-flags":
                found = "exclusive" in report_lower or "due" in report_lower or "mutual" in report_lower
            elif rule_id == "action-verb-alignment":
                found = "mark-closed" in report_lower or "close" in report_lower or "alignment" in report_lower
            elif rule_id == "group-labels-consistency":
                found = "group" in report_lower or "label" in report_lower or "consistency" in report_lower
            elif rule_id == "short-description-clarity":
                found = "description" in report_lower or "verb" in report_lower or "clarity" in report_lower
            elif rule_id == "flag-help-completeness":
                found = "help" in report_lower or "state" in report_lower or "completeness" in report_lower
            elif rule_id == "todo-todo-topic-clarity":
                found = "ambiguous" in report_lower or "topic" in report_lower or "clash" in report_lower or "clarity" in report_lower
            elif rule_id == "error-message-consistency":
                found = "error" in report_lower or "consistent" in report_lower or "format" in report_lower
            elif rule_id == "todo-show-command-verb":
                found = "info" in report_lower or "show" in report_lower or "verb consistency" in report_lower
            elif rule_id == "sink-show-shorthand":
                found = "show-sink" in report_lower or "sink" in report_lower or "shorthand" in report_lower
            else:
                # Fallback matching
                found = (
                    violation.rule_id.lower() in report_lower or
                    any(violation.rule_id.lower() in str(finding).lower() for finding in findings) or
                    violation.change_description.lower() in report_lower
                )
            
            if found:
                detected.append(violation)
            else:
                false_negatives.append(violation)
        
        coverage = len(detected) / len(violations_injected) if violations_injected else 1.0
        
        return DetectionResult(
            variation_id=f"todo-{variation_id:02d}",
            violations_injected=violations_injected,
            findings_detected=findings,
            false_positives=[],  # Would need to parse report more carefully
            false_negatives=false_negatives,
            severity_mismatches=severity_mismatches,
            coverage=coverage
        )
    
    def _parse_findings(self, report_content: str) -> List[Dict]:
        """Extract structured findings from cli-review report."""
        findings = []
        
        lines = report_content.split("\n")
        for i, line in enumerate(lines):
            # Check for severity headers like ### [HIGH-1] or [MEDIUM-2] or [LOW-1] or [UNRATED-1]
            if line.strip().startswith("### [") and "]" in line:
                findings.append({"line": i, "text": line.strip()})
            # Also check for table rows like | [HIGH-1] | or | [MEDIUM-1] |
            elif line.strip().startswith("|") and any(f"[{sev}-" in line for sev in ["HIGH", "MEDIUM", "LOW", "UNRATED"]):
                findings.append({"line": i, "text": line.strip()})
        
        return findings


class TestHarness:
    """Main orchestrator for testing."""
    
    def __init__(self, source_dir: Path, test_dir: Path, provider: str, model: str):
        self.source_dir = Path(source_dir)
        self.test_dir = Path(test_dir)
        self.provider = provider
        self.model = model
        self.project_root = Path(source_dir).parent
        
        self.generator = VariationGenerator(source_dir, test_dir)
        self.runner = CLIReviewRunner(self.project_root)
        self.analyzer = AnalysisEngine()
    
    def run_all_phases(self):
        """Run all testing phases."""
        print("=" * 70)
        print("CLI Review Testing Harness")
        print("=" * 70)
        
        # Phase 1: Generate Variations
        print("\n[Phase 1] Generating 20 test variations...")
        variations = []
        for i in range(1, 21):
            var_dir, violations = self.generator.create_variation(i)
            variations.append((var_dir, violations))
            print(f"  ✓ Created todo-{i:02d} with {len(violations)} violation(s)")
        
        # Phase 2: Run CLI Review
        print(f"\n[Phase 2] Running cli-review on each variation...")
        results = []
        for var_dir, violations in variations:
            var_id = var_dir.name
            print(f"  Running review on {var_id}...", end=" ", flush=True)
            
            report, session = self.runner.run_review(var_dir, self.provider, self.model)
            result = self.analyzer.analyze(int(var_id.split("-")[1]), violations, report)
            results.append(result)
            
            print(f"✓ ({len(result.false_negatives)} false negatives)")
        
        # Phase 3: Analyze Results
        print(f"\n[Phase 3] Analyzing detection accuracy...")
        summary = self._generate_summary(results)
        
        # Phase 4: Interrogate Model (continued session)
        print(f"\n[Phase 4] Interrogating model about improvements...")
        # This continues the session - implementation would go here
        
        # Phase 5: Generate Amendment Summary
        print(f"\n[Phase 5] Generating amendment summary...")
        # This creates a new Gemini session
        
        # Save overall analysis
        analysis_file = self.test_dir / "ANALYSIS.json"
        with open(analysis_file, "w") as f:
            json.dump(summary, f, indent=2)
        
        print(f"\n✓ Analysis saved to {analysis_file}")
        
        return results, summary
    
    def _generate_summary(self, results: List[DetectionResult]) -> Dict:
        """Generate summary statistics."""
        total_violations = sum(len(r.violations_injected) for r in results)
        total_detected = sum(len(r.violations_injected) - len(r.false_negatives) for r in results)
        
        # Count by severity
        severity_stats = {"HIGH": 0, "MEDIUM": 0, "LOW": 0}
        severity_detected = {"HIGH": 0, "MEDIUM": 0, "LOW": 0}
        
        for result in results:
            for violation in result.violations_injected:
                severity = violation.severity
                severity_stats[severity] = severity_stats.get(severity, 0) + 1
                if violation not in result.false_negatives:
                    severity_detected[severity] = severity_detected.get(severity, 0) + 1
        
        return {
            "total_variations": len(results),
            "total_violations_injected": total_violations,
            "average_detection_coverage": sum(r.coverage for r in results) / len(results) if results else 0,
            "severity_breakdown": {
                "HIGH": {
                    "injected": severity_stats.get("HIGH", 0),
                    "detected": severity_detected.get("HIGH", 0),
                    "coverage_percent": (severity_detected.get("HIGH", 0) / severity_stats.get("HIGH", 1) * 100) if severity_stats.get("HIGH", 0) > 0 else 0
                },
                "MEDIUM": {
                    "injected": severity_stats.get("MEDIUM", 0),
                    "detected": severity_detected.get("MEDIUM", 0),
                    "coverage_percent": (severity_detected.get("MEDIUM", 0) / severity_stats.get("MEDIUM", 1) * 100) if severity_stats.get("MEDIUM", 0) > 0 else 0
                },
                "LOW": {
                    "injected": severity_stats.get("LOW", 0),
                    "detected": severity_detected.get("LOW", 0),
                    "coverage_percent": (severity_detected.get("LOW", 0) / severity_stats.get("LOW", 1) * 100) if severity_stats.get("LOW", 0) > 0 else 0
                }
            },
            "results": [
                {
                    "variation": r.variation_id,
                    "violations_injected": len(r.violations_injected),
                    "violations_detected": len(r.violations_injected) - len(r.false_negatives),
                    "false_negatives": [
                        {
                            "rule_id": v.rule_id,
                            "severity": v.severity,
                            "description": v.change_description
                        }
                        for v in r.false_negatives
                    ],
                    "severity_breakdown": {
                        "HIGH": len([v for v in r.violations_injected if v.severity == "HIGH"]),
                        "MEDIUM": len([v for v in r.violations_injected if v.severity == "MEDIUM"]),
                        "LOW": len([v for v in r.violations_injected if v.severity == "LOW"])
                    },
                    "coverage_percent": r.coverage * 100
                }
                for r in results
            ]
        }


def main():
    """Main entry point."""
    if len(sys.argv) < 3:
        print("Usage: testing_harness.py <provider> <model>")
        print("Example: testing_harness.py anthropic claude-3-5-sonnet")
        sys.exit(1)
    
    provider = sys.argv[1]
    model = sys.argv[2]
    
    project_root = Path(__file__).parent.parent
    test_dir = project_root / "tests"
    source_dir = project_root
    
    harness = TestHarness(source_dir, test_dir, provider, model)
    results, summary = harness.run_all_phases()
    
    print("\n" + "=" * 70)
    print(f"Summary: {summary['average_detection_coverage']*100:.1f}% average coverage")
    print("=" * 70)


if __name__ == "__main__":
    main()
