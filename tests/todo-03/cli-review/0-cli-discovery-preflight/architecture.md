# Architecture

## Tech Stack
- **Language**: Go (Golang)
- **CLI Framework**: Cobra (`github.com/spf13/cobra`)
- **Terminal Styling**: termenv (`github.com/muesli/termenv`)
- **Database**: SQLite (used by the daemon for persistence)

## Architecture Style
- **Primary Style**: Client-server CLI
- **Description**: The application is split into a CLI client (`todo`) and a background daemon (`todod`). The daemon manages the state and persistence via SQLite, while the client interacts with the daemon using a UNIX domain socket and a custom HTTP client helper.
