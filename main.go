package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        string `json:"id"`        //remap the fields to theior lowercase version while encoding
	Title     string `json:"title"`     //so that the task.jsonl file is consistent
	Completed bool   `json:"completed"` //so that the task.jsonl file is consistent
}
type TaskList []Task

const fileName string = "tasks.jsonl"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goTrack <command>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: goTrack add <task>")
			return
		}
		tasks, err := getAllTasks(fileName)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		taskId := strconv.Itoa(len(tasks) + 1)

		title := strings.Join(os.Args[2:], " ")
		task := Task{
			ID:        taskId,
			Title:     title,
			Completed: false,
		}
		err = appendTask(fileName, task)
		if err != nil {
			fmt.Println("Error appending task:", err)
			return
		}
		fmt.Println("Added task:", task)
	case "list":
		tasks, err := getAllTasks(fileName)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		for _, task := range tasks {
			fmt.Println(task)
		}
	default:
		fmt.Println("Unknown command")
	}
}
