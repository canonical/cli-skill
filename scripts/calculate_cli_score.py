#!/usr/bin/env python3
"""
Calculate CLI standards compliance score from a JSON issues table.

Input: JSON file with structure:
{
  "commands": <int>,
  "issues": [
    {"severity": "High|Medium|Low|Unrated", "category": str, "message": str},
    ...
  ]
}

Output: JSON with score, pass/fail status, and severity counts.
"""

import json
import sys
from pathlib import Path

def calculate_score(issues_data: dict) -> dict:
    """Calculate compliance score from issues table."""
    issues = issues_data.get("issues", [])
    num_commands = issues_data.get("commands", 0)

    if not issues:
        return {
            "score": 100,
            "passed": True,
            "summary": "No issues found",
            "command_count": num_commands,
            "issue_count": 0,
            "high_count": 0,
            "medium_count": 0,
            "low_count": 0,
        }

    # Count severities
    high_count = sum(1 for issue in issues if issue.get("severity") == "High")
    medium_count = sum(1 for issue in issues if issue.get("severity") == "Medium")
    low_count = sum(1 for issue in issues if issue.get("severity") == "Low")

    # Calculate score: W = 100/N, High: -W, Medium: -0.5W, Low: -0.2W
    if num_commands == 0:
        score = 100.0
    else:
        weight = 100.0 / num_commands
        score = 100.0
        score -= high_count * weight
        score -= medium_count * 0.5 * weight
        score -= low_count * 0.2 * weight

    score = max(0.0, min(100.0, score))  # Clamp to 0-100

    passed = score > 80

    return {
        "score": score,
        "passed": passed,
        "summary": f"{high_count} High, {medium_count} Medium, {low_count} Low",
        "command_count": num_commands,
        "issue_count": len(issues),
        "high_count": high_count,
        "medium_count": medium_count,
        "low_count": low_count,
    }


def main() -> None:
    if len(sys.argv) != 2:
        print("Usage: calculate_cli_score.py <issues.json>", file=sys.stderr)
        sys.exit(1)
    
    json_path = Path(sys.argv[1])
    if not json_path.exists():
        print(f"Error: {json_path} not found", file=sys.stderr)
        sys.exit(1)
    
    try:
        issues_data = json.loads(json_path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        print(f"Error: invalid JSON: {exc}", file=sys.stderr)
        sys.exit(1)
    
    result = calculate_score(issues_data)
    print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()
