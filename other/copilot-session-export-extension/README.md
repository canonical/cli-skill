# Copilot Session Export

A VS Code extension that exports all your GitHub Copilot chat sessions to a formatted Markdown file.

## Features

- Finds all Copilot chat transcripts across all workspace storage folders
- Extracts user prompts (stripping internal XML metadata)
- Generates a clean Markdown file with:
  - Table of contents
  - Session title (derived from first prompt)
  - Date and session ID
  - Summary (from first assistant response)
  - Numbered list of all prompts

## Usage

1. Open the Command Palette (`Ctrl+Shift+P` / `Cmd+Shift+P`)
2. Run **"Export Copilot Sessions to Markdown"**
3. Choose where to save the file
4. The generated Markdown file opens automatically

## Install from source

```bash
cd copilot-export-extension
npm install
npm run compile
```

Then press `F5` in VS Code to launch an Extension Development Host with the extension loaded.

## Package as VSIX

```bash
npm install -g @vscode/vsce
vsce package
```

Then install the `.vsix` file via **Extensions > ... > Install from VSIX**.

## How it works

Copilot chat stores transcripts as JSONL files in:
```
~/.vscode/data/User/workspaceStorage/<id>/GitHub.copilot-chat/transcripts/*.jsonl
```

Each line is a JSON event (`session.start`, `user.message`, `assistant.message`, etc.). The extension reads all these files, parses the user prompts and assistant responses, and formats them into a readable Markdown document.
