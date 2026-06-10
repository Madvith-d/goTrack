# goTrack

A simple Command Line Interface (CLI) task tracker built in Go (Golang).

> [!NOTE]
> This is a **learning project** designed to explore Go's core concepts, CLI argument handling, file I/O operations, JSON encoding/decoding, and basic Go project structures.

---

## Features

- **Sequential Task ID Generation**: Automatically increments task IDs based on list size.
- **Task Management**:
  - `add`: Add new tasks to the tracker.
  - `list`: View all tasks with their completion status.
  - `done`: Mark a specific task as completed.
  - `del`: Delete a task (automatically shifts subsequent task IDs to maintain clean sequential order).
- **JSON Lines Persistence**: Stores tasks in a local `tasks.jsonl` file.

---

## Installation & Setup

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.26.4 or later recommended)

### 1. Running in Development
To run the CLI during development without building a binary:
```bash
# Run the entire package/directory (required due to multi-file structure)
go run . <command> [arguments]
```

### 2. Building Locally
To compile the project into a local executable:
```bash
# Build the binary
go build -o goTrack

# Run the compiled binary
./goTrack <command> [arguments]
```

### 3. Installing Locally
To install the binary globally into your `$GOPATH/bin` or `$GOBIN`:
```bash
go install
```
After installation, you can run the command from anywhere:
```bash
goTrack <command> [arguments]
```

### 4. Installing from GitHub Releases
If you have configured GitHub Releases via GoReleaser:
1. Navigate to the **Releases** page of this repository.
2. Download the appropriate pre-compiled archive for your OS and architecture.
3. Extract the binary and place it in your system's PATH.

---

## Usage

Here are the commands you can use:

### Add a Task
```bash
go run . add "Buy groceries"
go run . add "Learn Go concepts"
```

### List Tasks
```bash
go run . list
```
**Output Example:**
```
[ ] 1 Buy groceries
[ ] 2 Learn Go concepts
```

### Complete a Task
```bash
go run . done 1
```
**Output Example:**
```
Completed task: {1 Buy groceries true}
```

### Delete a Task
```bash
go run . del 1
```
**Output Example:**
```
Deleted task: {1 Buy groceries true}
```
> [!TIP]
> Deleting a task will shift the IDs of all tasks after it (e.g., Task `2` becomes Task `1`) to keep the list sequentially numbered without gaps.

---

## Roadmap

- [ ] **Sequential ID Generation Fix**: Transition to database-like stateful auto-incrementing IDs so shifting isn't necessary upon deletion.
- [ ] **CLI Argument Parsing**: Migrate from basic manual `os.Args` manipulation to Go's standard library `flag` package or `cobra` for flags and clean help screens.
- [ ] **User Experience Improvements**: Implement pretty-printed tables or color-coded terminal text to indicate completed vs. active tasks.
