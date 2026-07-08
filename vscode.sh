#!/usr/bin/env bash
set -euo pipefail

# Extract workshop name from workshop.yaml
WORKSHOP=$(grep "^name:" workshop.yaml | cut -d: -f2 | xargs)
PROJECT_PATH="/project"
IP=`workshop info $WORKSHOP | yq .hostname`
echo "Opening VSCode remote session at $IP..."
code --folder-uri "vscode-remote://ssh-remote+workshop@$IP$PROJECT_PATH"
