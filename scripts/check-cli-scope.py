#!/usr/bin/env python3
import fnmatch
import os
import pathlib
import re
import subprocess
import sys


NO_CLI_PATTERNS = [
    r"\bthis\s+(project|repository|repo)\s+has\s+no\s+cli\b",
    r"\bthere\s+is\s+no\s+cli\b",
    r"\bdoes\s+not\s+(provide|expose)\s+a\s+cli\b",
    r"\bno\s+command[- ]line\s+interface\b",
    r"\bnot\s+a\s+cli\s+project\b",
]


def fail(message: str) -> None:
    print(message, file=sys.stderr)
    sys.exit(1)


def write_proceed(output_path: pathlib.Path, proceed: bool) -> None:
    with output_path.open("a", encoding="utf-8") as file_handle:
        file_handle.write(f"proceed={'true' if proceed else 'false'}\n")


def parse_pattern_list(value: str) -> list[str]:
    if not value:
        return []

    parts = []
    for line in value.splitlines():
        for item in line.split(","):
            stripped = item.strip()
            if stripped:
                parts.append(stripped)
    return parts


def find_cli_doc(root: pathlib.Path, enforce_metadata: bool) -> pathlib.Path | None:
    kb_dir = root / "knowledge_base"
    if not kb_dir.exists():
        if not enforce_metadata:
            return None
        fail(
            "Missing knowledge_base/CLI.md. This workflow only runs for repositories that declare CLI scope metadata. "
            "Either add knowledge_base/CLI.md (with 'paths: [...]' for CLI repositories or 'has_cli: false' for non-CLI repositories) "
            "or provide path filters via CLI_PATHS_INCLUDE."
        )

    cli_candidates = sorted(
        path for path in kb_dir.iterdir() if path.is_file() and path.name.lower() == "cli.md"
    )
    if not cli_candidates:
        if not enforce_metadata:
            return None
        fail(
            "Missing knowledge_base/CLI.md. This workflow only runs for repositories that declare CLI scope metadata. "
            "Either add knowledge_base/CLI.md (with 'paths: [...]' for CLI repositories or 'has_cli: false' for non-CLI repositories) "
            "or provide path filters via CLI_PATHS_INCLUDE."
        )
    return cli_candidates[0]


def has_cli(text: str) -> bool | None:
    match = re.search(r"(?im)^has_cli\s*:\s*(true|false)\s*$", text)
    if match:
        return match.group(1).lower() == "true"

    normalized = re.sub(r"\s+", " ", text.lower())
    if any(re.search(pattern, normalized) for pattern in NO_CLI_PATTERNS):
        return False

    return None


def parse_paths(text: str) -> list[str]:
    inline_match = re.search(r"(?im)^paths\s*:\s*\[(.*?)\]\s*$", text)
    if inline_match:
        raw = inline_match.group(1)
        return [item.strip().strip('"\'') for item in raw.split(",") if item.strip()]

    block_match = re.search(r"(?ims)^paths\s*:\s*\n((?:[ \t]*-[^\n]*\n)+)", text)
    if block_match:
        values = []
        for line in block_match.group(1).splitlines():
            line = line.strip()
            if line.startswith("-"):
                value = line[1:].strip().strip('"\'')
                if value:
                    values.append(value)
        return values

    return []


def matches(spec: str, changed_file: str) -> bool:
    spec = spec.strip()
    if not spec:
        return False
    if any(char in spec for char in "*?[]"):
        return fnmatch.fnmatch(changed_file, spec)
    if spec.endswith("/"):
        return changed_file.startswith(spec)
    return changed_file == spec or changed_file.startswith(spec + "/")


def select_matching(
    changed: list[str], include_specs: list[str], exclude_specs: list[str]
) -> list[str]:
    selected = [item for item in changed if any(matches(spec, item) for spec in include_specs)]
    if not exclude_specs:
        return selected
    return [item for item in selected if not any(matches(spec, item) for spec in exclude_specs)]


def changed_files() -> list[str]:
    base_sha = os.environ.get("BASE_SHA", "").strip()
    head_sha = os.environ.get("HEAD_SHA", "").strip()
    if not base_sha or not head_sha:
        fail(
            "BASE_SHA and HEAD_SHA must be set. This CLI scope gate only supports pull request "
            "comparisons between the target branch and the source branch."
        )

    try:
        merge_base = subprocess.check_output(
            ["git", "merge-base", base_sha, head_sha], text=True
        ).strip()
    except subprocess.CalledProcessError as exc:
        fail(f"Unable to determine merge base for pull request diff: {exc}")

    diff_cmd = ["git", "diff", "--name-only", f"{merge_base}..{head_sha}"]

    try:
        return subprocess.check_output(diff_cmd, text=True).splitlines()
    except subprocess.CalledProcessError as exc:
        fail(f"Unable to determine changed files for CLI scope gate: {exc}")


def main() -> None:
    root = pathlib.Path.cwd()
    output_env = os.environ.get("GITHUB_OUTPUT")
    if not output_env:
        fail("GITHUB_OUTPUT is not set.")
    output_path = pathlib.Path(output_env)

    base_sha = os.environ.get("BASE_SHA", "").strip()
    head_sha = os.environ.get("HEAD_SHA", "").strip()
    if not base_sha or not head_sha:
        write_proceed(output_path, True)
        print("Success: no pull request diff context; proceeding by default")
        return

    include_specs = parse_pattern_list(os.environ.get("CLI_PATHS_INCLUDE", ""))
    exclude_specs = parse_pattern_list(os.environ.get("CLI_PATHS_EXCLUDE", ""))
    enforce_metadata = os.environ.get("ENFORCE_CLI_METADATA", "false").lower() == "true"
    changed = changed_files()

    if include_specs:
        matched = select_matching(changed, include_specs, exclude_specs)
        write_proceed(output_path, bool(matched))
        if matched:
            print("CLI-relevant changes detected via input path filters:")
            for item in matched:
                print(f"- {item}")
        else:
            print("Success: no CLI changes matched input path filters")
        return

    cli_doc = find_cli_doc(root, enforce_metadata=enforce_metadata)
    if cli_doc is None:
        write_proceed(output_path, True)
        print(
            "Success: no CLI metadata found; proceeding by default. "
            "Set ENFORCE_CLI_METADATA=true to require knowledge_base/CLI.md."
        )
        return

    text = cli_doc.read_text(encoding="utf-8")
    cli_present = has_cli(text)

    if cli_present is False:
        write_proceed(output_path, False)
        print("Success: no CLI found")
        return

    paths = parse_paths(text)
    if not paths:
        if not enforce_metadata:
            write_proceed(output_path, True)
            print(
                "Success: CLI metadata does not declare paths; proceeding by default. "
                "Set ENFORCE_CLI_METADATA=true to require a YAML 'paths' entry."
            )
            return
        fail(
            f"{cli_doc} does not declare a YAML 'paths' entry. For CLI repositories, add a line like "
            "'paths: [cmd/, internal/cli/, docs/cli/]' so the workflow can determine whether a pull request changes CLI-relevant files. "
            "If the repository has no CLI, declare 'has_cli: false' or include a clear sentence such as 'This repository has no CLI.'"
        )

    matched = select_matching(changed, paths, exclude_specs)

    write_proceed(output_path, bool(matched))

    if matched:
        print("CLI-relevant changes detected:")
        for item in matched:
            print(f"- {item}")
    else:
        print("Success: no CLI changes")


if __name__ == "__main__":
    main()
