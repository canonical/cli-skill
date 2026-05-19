#!/usr/bin/env python3
"""
Convert markdown tables in a file to JSON.

Usage:
    python3 mdtables_to_json.py input.md [output.json]

If output is not specified, writes to input.json (same directory, same basename).
Supports multiple tables per file — each table becomes an array of objects.
If multiple tables exist, the output is a dict mapping "table_N" to each array.
"""

import sys
import re
import json
import os

def parse_markdown_table(lines):
    """Parse a markdown table from a list of lines. Returns list of dicts or None."""
    # Find table rows (lines starting with |)
    table_lines = []
    in_table = False
    for line in lines:
        stripped = line.strip()
        if stripped.startswith('|'):
            table_lines.append(stripped)
            in_table = True
        elif in_table and not stripped:
            break
        elif in_table and not stripped.startswith('|'):
            break

    if len(table_lines) < 2:
        return None

    # Remove separator line (second line with only |, -, :, spaces)
    if re.match(r'^\|[\s\-:|]+\|$', table_lines[1]) or re.match(r'^\|[\s\-:|]+\|$', table_lines[1].replace(' ', '')):
        header_line = table_lines[0]
        data_lines = table_lines[2:]
    else:
        header_line = table_lines[0]
        data_lines = table_lines[1:]

    def split_cells(line):
        # Remove leading/trailing |, then split by |
        inner = line.strip()
        if inner.startswith('|'):
            inner = inner[1:]
        if inner.endswith('|'):
            inner = inner[:-1]
        return [cell.strip() for cell in inner.split('|')]

    headers = split_cells(header_line)
    rows = []
    for line in data_lines:
        cells = split_cells(line)
        if len(cells) < len(headers):
            cells.extend([''] * (len(headers) - len(cells)))
        row = {}
        for i, h in enumerate(headers):
            # Clean up markdown formatting in headers and cells
            clean_header = re.sub(r'[*_`]', '', h).strip()
            clean_cell = re.sub(r'[*_`]', '', cells[i]).strip() if i < len(cells) else ''
            row[clean_header] = clean_cell
        rows.append(row)

    return rows


def extract_tables(md_text):
    """Extract all markdown tables from text. Returns list of (title, rows)."""
    lines = md_text.split('\n')
    tables = []
    i = 0
    current_title = None

    while i < len(lines):
        line = lines[i]
        # Check for heading before table
        if line.startswith('#'):
            current_title = line.lstrip('#').strip()

        if line.strip().startswith('|'):
            # Find table extent
            start = i
            while i < len(lines) and lines[i].strip().startswith('|'):
                i += 1
            table_lines = lines[start:i]
            rows = parse_markdown_table(table_lines)
            if rows:
                title = current_title or f"table_{len(tables) + 1}"
                tables.append((title, rows))
            continue
        i += 1

    return tables


def main():
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} input.md [output.json]")
        sys.exit(1)

    input_path = sys.argv[1]
    if len(sys.argv) >= 3:
        output_path = sys.argv[2]
    else:
        base, _ = os.path.splitext(input_path)
        output_path = base + '.json'

    with open(input_path, 'r', encoding='utf-8') as f:
        md_text = f.read()

    tables = extract_tables(md_text)

    if not tables:
        print(f"No tables found in {input_path}")
        # Write empty object
        with open(output_path, 'w', encoding='utf-8') as f:
            json.dump({}, f, indent=2)
        return

    if len(tables) == 1:
        output = tables[0][1]
    else:
        output = {}
        seen = {}
        for title, rows in tables:
            # slugify title
            key = re.sub(r'[^\w\-]', '_', title).lower()
            key = re.sub(r'_+', '_', key).strip('_')
            if not key:
                key = f"table_{len(output) + 1}"
            # deduplicate
            orig = key
            counter = 1
            while key in seen:
                key = f"{orig}_{counter}"
                counter += 1
            seen[key] = True
            output[key] = rows

    with open(output_path, 'w', encoding='utf-8') as f:
        json.dump(output, f, indent=2, ensure_ascii=False)

    print(f"Wrote {len(tables)} table(s) to {output_path}")


if __name__ == '__main__':
    main()
