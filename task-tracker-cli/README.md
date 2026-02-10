# Task CLI

Simple command-line task manager written in Go. Stores tasks in `tasks.json` in the current directory.

Build:

```bash
go build -o task-cli
```

Examples:

```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1

./task-cli mark-in-progress 1
./task-cli mark-done 1

./task-cli list
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

Commands: `add`, `update`, `delete`, `mark-in-progress`, `mark-done`, `list`.
