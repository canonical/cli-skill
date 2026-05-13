#!/usr/bin/env bash
set -euo pipefail

# Extract workshop name from workshop.yaml
WORKSHOP=$(grep "^name:" workshop.yaml | cut -d: -f2 | xargs)
PROJECT_PATH="/project"
WORKSHOP_HOST_VM="workshop-host"
SSH_HOST_ALIAS="workshop-via-host"
SSH_CONFIG_FILE="${HOME}/.ssh/config"
MANAGED_BEGIN="# >>> design-workshop managed block >>>"
MANAGED_END="# <<< design-workshop managed block <<<"

# Get the IPv4 address of the workshop-host Multipass VM.
HOST_IP=$(multipass list --format csv 2>/dev/null \
  | awk -F, -v host="$WORKSHOP_HOST_VM" '$1 == host { print $3; exit }' \
  | grep -oE '([0-9]{1,3}\.){3}[0-9]{1,3}' \
  | head -1)

if [[ -z "$HOST_IP" ]]; then
  echo "Error: could not find a running Multipass VM named '$WORKSHOP_HOST_VM'."
  echo "Make sure it's running with: multipass start $WORKSHOP_HOST_VM"
  exit 1
fi

# Get the workshop container IPv4 from inside the workshop-host VM.
IP=$(multipass exec "$WORKSHOP_HOST_VM" -- \
  lxc list --all-projects --format csv -c 4 "$WORKSHOP" 2>/dev/null \
  | grep -oE '([0-9]{1,3}\.){3}[0-9]{1,3}' \
  | head -1)

if [[ -z "$IP" ]]; then
  echo "Error: could not find a running workshop named '$WORKSHOP'."
  echo "Make sure it's running with: workshop launch"
  exit 1
fi

mkdir -p "${HOME}/.ssh"

# Keep a single managed block in ~/.ssh/config for the VS Code SSH target.
if [[ -f "$SSH_CONFIG_FILE" ]]; then
  awk -v begin="$MANAGED_BEGIN" -v end="$MANAGED_END" '
    $0 == begin { skip = 1; next }
    $0 == end { skip = 0; next }
    !skip { print }
  ' "$SSH_CONFIG_FILE" > "${SSH_CONFIG_FILE}.tmp"
else
  : > "${SSH_CONFIG_FILE}.tmp"
fi

cat >> "${SSH_CONFIG_FILE}.tmp" <<EOF
$MANAGED_BEGIN
Host $SSH_HOST_ALIAS
  HostName $IP
  User workshop
  ProxyJump ubuntu@$HOST_IP
  StrictHostKeyChecking accept-new
$MANAGED_END
EOF

mv "${SSH_CONFIG_FILE}.tmp" "$SSH_CONFIG_FILE"
chmod 600 "$SSH_CONFIG_FILE"

echo "Opening VSCode remote session at $IP via jump host ubuntu@$HOST_IP..."
code --folder-uri "vscode-remote://ssh-remote+$SSH_HOST_ALIAS$PROJECT_PATH"
