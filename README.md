# goTrack

A simple Command Line Interface (CLI) task tracker built in Go (Golang).

> [!NOTE]
> This is a **learning project** designed to explore Go's core concepts, CLI argument handling, file I/O operations, JSON encoding/decoding, and basic Go project structures.

---

## What is Implemented

Currently, `goTrack` supports the following:

- **Command Route Architecture**: Basic command-line routing using a switch-case pattern on terminal arguments.
- **Task Creation (`add`)**:
  - Command: `go run main.go add "<task name>"`
  - Generates a cryptographically secure random ID (from 0 to 999) using Go's `crypto/rand`.
- **Structured Persistence**: Appends tasks to a local `tasks.jsonl` file in JSON Lines format.
- **Consistent Data Mapping**: Uses custom Go struct tags to ensure tasks are serialized and deserialized with consistent lowercase JSON keys (`id`, `title`, and `completed`).

---

## What is Left / Future Roadmap

As this is a learning project, there are several key features and improvements planned:

- [ ] **List Tasks (`list`)**:
  - The `list` command is currently a stub that only prints `Listing tasks`. It needs to read and parse the JSON Lines from `tasks.jsonl` and output them to the terminal.
- [ ] **Complete Tasks (`complete`)**:
  - Add a command to mark a task as completed (e.g., `goTrack complete <id>`).
- [ ] **Delete Tasks (`delete`)**:
  - Add a command to remove a task from the list.
- [ ] **File Updating Logic**:
  - Create standard load-update-save helper logic to modify task statuses and delete records in the `tasks.jsonl` file.
- [ ] **Better ID Generation**:
  - Transition from random 0-999 integers (which have high risk of collision) to sequential IDs or UUIDs.
- [ ] **CLI Argument Parsing**:
  - Migrate from basic manual `os.Args` manipulation to Go's standard library `flag` package or a library like `cobra` to allow flags and better help screens.
- [ ] **User Experience Improvements**:
  - Implement pretty-printed tables or color-coded terminal text to indicate completed vs. active tasks.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.26.4 or later recommended)

### How to Use

1. **Clone the repository** (or navigate to the project folder):
   ```bash
   cd goTrack
   ```

2. **Add a task**:
   ```bash
   go run main.go add "Finish writing the project README"
   ```

3. **Check the records**:
   See your added tasks in `tasks.jsonl`.
