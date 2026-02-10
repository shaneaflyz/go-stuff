package main

import (
	"cmp"
	"encoding/json"
	"errors"
	"os"
	"slices"
	"time"
)

const tasksFile = "tasks.json"

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(tasksFile); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	content, err := os.ReadFile(tasksFile)

	if err != nil {
		return nil, err
	}

	var tasks []Task

	if len(content) == 0 {
		return []Task{}, nil
	}

	if err := json.Unmarshal(content, &tasks); err != nil {
		return nil, err
	}

	lenCmp := func(a, b Task) int {
		return cmp.Compare(a.ID, b.ID)
	}

	slices.SortFunc(tasks, lenCmp)

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	b, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	tmp := tasksFile + ".tmp"

	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, tasksFile)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func findByID(tasks []Task, id int) (int, *Task) {
	for i := range tasks {
		if tasks[i].ID == id {
			return i, &tasks[i]
		}
	}
	return -1, nil
}

func now() time.Time { return time.Now().UTC() }
