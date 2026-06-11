#!/usr/bin/env python3
import os
from pathlib import Path


def main() -> None:
    marker = os.environ.get("MARKER", "<!-- cli-skill-report -->")
    command = os.environ.get("COMMAND", "/cli-review")
    success = os.environ.get("PI_SUCCESS", "")
    cost = os.environ.get("PI_COST", "")
    duration = os.environ.get("PI_DURATION", "")
    input_tokens = os.environ.get("PI_INPUT_TOKENS", "")
    output_tokens = os.environ.get("PI_OUTPUT_TOKENS", "")
    run_url = os.environ.get("RUN_URL", "")
    response = os.environ.get("PI_RESPONSE", "")
    session_html_artifact_url = os.environ.get("SESSION_HTML_ARTIFACT_URL", "")
    session_jsonl_artifact_url = os.environ.get("SESSION_JSONL_ARTIFACT_URL", "")

    if not response.strip():
        response = "_No response was returned by Pi._"

    # Keep body safely below GitHub comment max size.
    max_response_len = 50000
    if len(response) > max_response_len:
        response = response[:max_response_len] + "\n\n... _truncated_"

    artifact_links: list[str] = []
    if session_html_artifact_url:
        artifact_links.append(f"- Session HTML: [download]({session_html_artifact_url})")
    if session_jsonl_artifact_url:
        artifact_links.append(f"- Session JSONL: [download]({session_jsonl_artifact_url})")

    artifacts_section = ""
    if artifact_links:
        artifacts_section = "\n### Artifacts\n\n" + "\n".join(artifact_links) + "\n"

    body = f"""{marker}
## CLI Skill Report ({command})

| Metric | Value |
|---|---|
| Success | {success or 'unknown'} |
| Duration (s) | {duration or 'n/a'} |
| Cost (USD) | {cost or 'n/a'} |
| Tokens (in/out) | {input_tokens or 'n/a'} / {output_tokens or 'n/a'} |
| Workflow run | [view run]({run_url}) |

{artifacts_section}
{response}
"""

    report_path = Path("pi-report-comment.md")
    report_path.write_text(body.strip() + "\n", encoding="utf-8")
    print(f"Wrote {report_path}")


if __name__ == "__main__":
    main()
