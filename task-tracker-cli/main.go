package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: task-cli <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("  add \"Title\"                    Add a new task")
	fmt.Println("  update <id> \"New Title\"         Update task title")
	fmt.Println("  delete <id>                        Delete task")
	fmt.Println("  mark-in-progress <id>              Mark task in-progress")
	fmt.Println("  mark-done <id>                     Mark task done")
	fmt.Println("  list [all|done|todo|in-progress]   List tasks (default: all)")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Title is required")
			return
		}

		title := strings.Join(os.Args[2:], " ")
		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		t := Task{ID: nextID(tasks), Title: title, Status: "todo", CreatedAt: now(), UpdatedAt: now()}

		tasks = append(tasks, t)

		if err := saveTasks(tasks); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", t.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> \"New title\"")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("invalid id")
			return
		}

		title := strings.Join(os.Args[3:], " ")
		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		index, task := findByID(tasks, id)

		if task == nil {
			fmt.Println("task not found")
			return
		}

		task.Title = title
		task.UpdatedAt = now()
		tasks[index] = *task

		if err := saveTasks(tasks); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Task updated successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("invalid id")
			return
		}

		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		idx, _ := findByID(tasks, id)

		if idx == -1 {
			fmt.Println("task not found")
			return
		}

		tasks = append(tasks[:idx], tasks[idx+1:]...)

		if err := saveTasks(tasks); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Task deleted successfully")

	case "mark-in-progress", "mark-inprogress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-in-progress <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("invalid id")
			return
		}

		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		idx, t := findByID(tasks, id)

		if t == nil {
			fmt.Println("task not found")
			return
		}

		t.Status = "inprogress"
		t.UpdatedAt = now()
		tasks[idx] = *t

		if err := saveTasks(tasks); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Task marked in progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-done <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("invalid id")
			return
		}

		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		idx, t := findByID(tasks, id)

		if t == nil {
			fmt.Println("task not found")
			return
		}

		t.Status = "done"
		t.UpdatedAt = now()
		tasks[idx] = *t

		if err := saveTasks(tasks); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Task marked done")

	case "list":
		status := "all"

		if len(os.Args) >= 3 {
			status = strings.ToLower(os.Args[2])

			if status == "in-progress" {
				status = "inprogress"
			}
		}

		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		for _, t := range tasks {
			show := false
			switch status {
			case "all":
				show = true
			case "done":
				show = t.Status == "done"
			case "todo":
				show = t.Status == "todo"
			case "inprogress":
				show = t.Status == "inprogress"
			default:
				fmt.Println("unknown status filter")
				return
			}
			if show {
				fmt.Printf("%d: %s [%s]\n", t.ID, t.Title, t.Status)
				if t.Description != "" {
					fmt.Printf("    %s\n", t.Description)
				}
			}
		}

	default:
		printUsage()
	}
}
