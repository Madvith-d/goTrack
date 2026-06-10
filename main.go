package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			status := "[ ]"
			if task.Completed {
				status = "[x]"
			}

			fmt.Printf("%s %s %s\n", status, task.ID, task.Title)
		}
	case "del":
		if len(os.Args) != 3 {
			fmt.Println("Usage : goTrack del <taskID>")
			return
		}
		tasks, err := getAllTasks(fileName)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		taskId := os.Args[2]
		for i, task := range tasks {
			if task.ID == taskId {
				for j := i + 1; j < len(tasks); j++ {
					id, _ := strconv.Atoi(tasks[j].ID)
					tasks[j].ID = strconv.Itoa(id - 1)
				}
				tasks = append(tasks[:i], tasks[i+1:]...)
				err = writeTasks(fileName, tasks)
				if err != nil {
					fmt.Println("Error deleting task:", err)
					return
				}
				fmt.Println("Deleted task:", task)
				return
			}
		}
		fmt.Println("Task not found")
	case "done":
		if len(os.Args) != 3 {
			fmt.Println("Usage : goTrack done <taskID>")
			return
		}
		tasks, err := getAllTasks(fileName)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		taskId := os.Args[2]
		for i := 0; i < len(tasks); i++ {
			if tasks[i].ID == taskId {
				tasks[i].Completed = true
				err = writeTasks(fileName, tasks)
				if err != nil {
					fmt.Println("Error completing task:", err)
					return
				}
				fmt.Println("Completed task:", tasks[i])
				return
			}
		}
		fmt.Println("Task not found")
	default:
		fmt.Println("Unknown command")
	}
}
