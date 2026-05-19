#!/usr/bin/env python3
"""
Validate JSON table exports against source markdown tables.

Usage:
    python3 validate_json_tables.py <source.md> <exported.json>

Checks:
1. JSON is valid and parseable
2. Number of tables matches between markdown and JSON
3. Table row counts match
4. All column headers from markdown appear in JSON objects
5. No JSON objects have empty or all-whitespace values for required fields
"""

import sys
import json
import re

def extract_markdown_tables(md_text):
    """Extract tables from markdown, return list of (heading, headers, rows)."""
    lines = md_text.split('\n')
    tables = []
    i = 0
    current_heading = None

    while i < len(lines):
        line = lines[i]
        if line.startswith('#'):
            current_heading = line.lstrip('#').strip()

        if line.strip().startswith('|'):
            start = i
            while i < len(lines) and lines[i].strip().startswith('|'):
                i += 1
            table_lines = lines[start:i]

            # Skip separator
            if len(table_lines) >= 2 and re.match(r'^\|[\s\-:|]+\|$', table_lines[1].replace(' ', '')):
                header_line = table_lines[0]
                data_lines = table_lines[2:]
            else:
                header_line = table_lines[0]
                data_lines = table_lines[1:]

            def split_cells(l):
                inner = l.strip()
                if inner.startswith('|'): inner = inner[1:]
                if inner.endswith('|'): inner = inner[:-1]
                return [c.strip() for c in inner.split('|')]

            headers = split_cells(header_line)
            rows = []
            for dl in data_lines:
                cells = split_cells(dl)
                row = {}
                for idx, h in enumerate(headers):
                    clean_h = re.sub(r'[*_`]', '', h).strip()
                    row[clean_h] = cells[idx] if idx < len(cells) else ''
                rows.append(row)

            title = current_heading or f"table_{len(tables)+1}"
            tables.append((title, headers, rows))
            continue
        i += 1

    return tables


def validate_json_against_markdown(md_path, json_path):
    errors = []
    warnings = []

    # Load markdown tables
    with open(md_path, 'r', encoding='utf-8') as f:
        md_text = f.read()
    md_tables = extract_markdown_tables(md_text)

    # Load JSON
    try:
        with open(json_path, 'r', encoding='utf-8') as f:
            json_data = json.load(f)
    except json.JSONDecodeError as e:
        return False, [f"Invalid JSON: {e}"], []
    except Exception as e:
        return False, [f"Cannot read JSON: {e}"], []

    # Determine JSON structure
    if isinstance(json_data, list):
        json_tables = [("table_1", json_data)]
    elif isinstance(json_data, dict):
        json_tables = []
        for key, value in json_data.items():
            if isinstance(value, list):
                json_tables.append((key, value))
    else:
        return False, ["JSON root must be an array or object"], []

    # Check table count
    if len(md_tables) != len(json_tables):
        errors.append(
            f"Table count mismatch: markdown has {len(md_tables)}, JSON has {len(json_tables)}"
        )

    # Validate each table
    for idx, (md_title, md_headers, md_rows) in enumerate(md_tables):
        if idx >= len(json_tables):
            errors.append(f"Missing JSON table for markdown table '{md_title}'")
            continue

        json_title, json_rows = json_tables[idx]

        # Row count
        if len(md_rows) != len(json_rows):
            errors.append(
                f"Table '{md_title}': row count mismatch (markdown: {len(md_rows)}, JSON: {len(json_rows)})"
            )

        # Header presence
        clean_md_headers = [re.sub(r'[*_`]', '', h).strip() for h in md_headers]
        if json_rows:
            json_keys = list(json_rows[0].keys())
            missing_in_json = [h for h in clean_md_headers if h not in json_keys]
            missing_in_md = [k for k in json_keys if k not in clean_md_headers]
            if missing_in_json:
                errors.append(
                    f"Table '{md_title}': headers missing in JSON: {missing_in_json}"
                )
            if missing_in_md:
                warnings.append(
                    f"Table '{md_title}': extra keys in JSON not in markdown: {missing_in_md}"
                )

        # Empty value check
        for row_idx, row in enumerate(json_rows):
            empty_keys = [k for k, v in row.items() if not v or not str(v).strip()]
            if empty_keys:
                warnings.append(
                    f"Table '{json_title}', row {row_idx+1}: empty values for keys: {empty_keys}"
                )

    return len(errors) == 0, errors, warnings


def main():
    if len(sys.argv) < 3:
        print(f"Usage: {sys.argv[0]} <source.md> <exported.json>")
        sys.exit(1)

    md_path = sys.argv[1]
    json_path = sys.argv[2]

    valid, errors, warnings = validate_json_against_markdown(md_path, json_path)

    print(f"Validation: {'PASS' if valid else 'FAIL'}")
    print(f"Markdown tables found: {len(extract_markdown_tables(open(md_path).read()))}")

    if errors:
        print("\nERRORS:")
        for e in errors:
            print(f"  ❌ {e}")

    if warnings:
        print("\nWARNINGS:")
        for w in warnings:
            print(f"  ⚠️  {w}")

    if valid and not warnings:
        print("  ✅ All checks passed")

    sys.exit(0 if valid else 1)


if __name__ == '__main__':
    main()
