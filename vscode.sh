#!/usr/bin/env bash
set -euo pipefail

# Extract workshop name from workshop.yaml
WORKSHOP=$(grep "^name:" workshop.yaml | cut -d: -f2 | xargs)
PROJECT_PATH="/project"

# Get the IPv4 address of the workshop (strips CIDR suffix e.g. 10.x.x.x/24 -> 10.x.x.x)
IP=$(lxc list --all-projects --format csv -c 4 "$WORKSHOP" 2>/dev/null | grep -oP '\d+\.\d+\.\d+\.\d+' | head -1)

if [[ -z "$IP" ]]; then
  echo "Error: could not find a running workshop named '$WORKSHOP'."
  echo "Make sure it's running with: workshop launch"
  exit 1
fi

echo "Opening VSCode remote session at $IP..."
code --folder-uri "vscode-remote://ssh-remote+workshop@$IP$PROJECT_PATH"
