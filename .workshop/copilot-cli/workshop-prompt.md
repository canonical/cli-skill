# Workshop Container Environment Instructions

## Environment Context

You are operating inside a Workshop — an isolated Ubuntu container environment
where a user's project is mounted `/project`.

**Container constraints:**
- No default access to display servers, GPUs, cameras, or audio devices
- Hardware passthrough requires explicit container configuration
- You cannot modify container configuration

## Required Behavior

Before using any hardware resource (display, GPU, camera, audio):

1. **Detect** — Check for the resource using the appropriate method (env vars, /dev paths)
2. **Validate** — Confirm the resource is functional if detected
3. **Stop** — If required resource is missing, halt execution immediately
4. **Report** — Tell the user exactly what is missing and why it is needed
5. **Wait** — Do not proceed until the user confirms or provides the resource

## Prohibited Actions

- Never assume resource availability without verification
- Never silently fall back to headless or stub implementations
- Never continue execution when required resources are unavailable

Favor deterministic failure over silent degradation.
