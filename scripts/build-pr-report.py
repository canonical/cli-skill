#!/usr/bin/env python3
import json
import os
import sys
from pathlib import Path


def format_duration(seconds: int) -> str:
    """Format seconds to human readable format (Xm Ys)."""
    if not seconds or seconds == 0:
        return "unknown"
    minutes = seconds // 60
    secs = seconds % 60
    return f"{minutes}m {secs}s"


def format_tokens(input_tokens: int, output_tokens: int) -> str:
    """Format tokens with commas for readability."""
    if not input_tokens or not output_tokens:
        return "unknown"
    total = input_tokens + output_tokens
    # Format with commas, convert large numbers to M notation
    if total >= 1_000_000:
        return f"{total / 1_000_000:.1f}M"
    elif total >= 1_000:
        return f"{total / 1_000:.1f}K"
    else:
        return str(total)


def extract_metrics_from_session(session_jsonl_path: str) -> dict:
    """Extract metrics (duration, tokens, cost, model) from pi session JSONL file."""
    metrics = {
        "duration": 0,
        "input_tokens": 0,
        "output_tokens": 0,
        "cost": 0.0,
        "model": "",
    }

    if not Path(session_jsonl_path).exists():
        return metrics

    try:
        with open(session_jsonl_path, "r") as f:
            for line in f:
                try:
                    entry = json.loads(line.strip())
                    
                    # Extract duration from metadata
                    if "metadata" in entry and "duration" in entry["metadata"]:
                        metrics["duration"] = entry["metadata"]["duration"]
                    
                    # Extract tokens from usage
                    if "usage" in entry:
                        if "input_tokens" in entry["usage"]:
                            metrics["input_tokens"] = entry["usage"]["input_tokens"]
                        if "output_tokens" in entry["usage"]:
                            metrics["output_tokens"] = entry["usage"]["output_tokens"]
                    
                    # Extract cost from metadata
                    if "metadata" in entry and "cost" in entry["metadata"]:
                        metrics["cost"] = entry["metadata"]["cost"]
                    
                    # Extract model from metadata
                    if "metadata" in entry and "model" in entry["metadata"]:
                        metrics["model"] = entry["metadata"]["model"]
                except json.JSONDecodeError:
                    continue
    except Exception as e:
        print(f"Error reading session file: {e}", file=sys.stderr)

    return metrics


def main() -> None:
    if len(sys.argv) < 2:
        print("Usage: build-pr-report.py <session_jsonl_path>", file=sys.stderr)
        sys.exit(1)

    session_jsonl_path = sys.argv[1]
    marker = os.environ.get("MARKER", "<!-- cli-review-report -->")
    command = os.environ.get("COMMAND", "/cli-review")
    response = os.environ.get("PI_RESPONSE", "")
    pi_version = os.environ.get("PI_VERSION", "")
    pi_thinking = os.environ.get("PI_THINKING", "")
    pi_model = os.environ.get("PI_MODEL", "unknown")
    gh_run_id = os.environ.get("GH_RUN_ID", "")
    gh_repo = os.environ.get("GH_REPO", "")
    gh_server_url = os.environ.get("GH_SERVER_URL", "https://github.com")

    # Extract metrics from session file
    metrics = extract_metrics_from_session(session_jsonl_path)

    if not response.strip():
        response = "_No response was returned by Pi._"

    # Keep body safely below GitHub comment max size.
    max_response_len = 50000
    if len(response) > max_response_len:
        response = response[:max_response_len] + "\n\n... _truncated_"

    # Format metrics for footer
    duration_str = format_duration(metrics["duration"])
    tokens_str = format_tokens(metrics["input_tokens"], metrics["output_tokens"])
    cost_str = f"${metrics['cost']:.2f}" if metrics["cost"] > 0 else "unknown"
    
    # Use provided model or extract from session
    model_display = pi_model if pi_model != "unknown" else metrics.get("model", "unknown")

    # Build action run link
    action_run_link = ""
    if gh_run_id and gh_repo and gh_server_url:
        action_run_link = f"[View action run]({gh_server_url}/{gh_repo}/actions/runs/{gh_run_id})"
    else:
        action_run_link = "Action run"

    # Build footer with optional thinking level and SDK version
    footer_parts = [action_run_link, f"Model: {model_display}"]
    if pi_thinking:
        footer_parts[1] += f" (thinking: {pi_thinking})"
    footer_parts.extend([f"Time: {duration_str}", f"Tokens: {tokens_str}", f"Cost: {cost_str}"])
    if pi_version:
        footer_parts.append(f"Pi SDK {pi_version}")
    footer_line = " | ".join(footer_parts)

    body = f"""{marker}
## CLI Skill Report ({command})

{response}

---

{footer_line}
"""

    report_path = Path("pr_report.md")
    report_path.write_text(body.strip() + "\n", encoding="utf-8")
    print(f"Wrote {report_path}")


if __name__ == "__main__":
    main()
