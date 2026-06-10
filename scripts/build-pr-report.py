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

    if not response.strip():
        response = "_No response was returned by Pi._"

    # Keep body safely below GitHub comment max size.
    max_response_len = 50000
    if len(response) > max_response_len:
        response = response[:max_response_len] + "\n\n... _truncated_"

    body = f"""{marker}
## CLI Skill Report ({command})

| Metric | Value |
|---|---|
| Success | {success or 'unknown'} |
| Duration (s) | {duration or 'n/a'} |
| Cost (USD) | {cost or 'n/a'} |
| Tokens (in/out) | {input_tokens or 'n/a'} / {output_tokens or 'n/a'} |
| Workflow run | [view run]({run_url}) |

{response}
"""

    report_path = Path("pi-report-comment.md")
    report_path.write_text(body.strip() + "\n", encoding="utf-8")
    print(f"Wrote {report_path}")


if __name__ == "__main__":
    main()
