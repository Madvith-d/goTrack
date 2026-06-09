package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

type Task struct {
	ID        string `json:"id"`        //remap the fields to theior lowercase version while encoding
	Title     string `json:"title"`     //so that the task.jsonl file is consistent
	Completed bool   `json:"completed"` //so that the task.jsonl file is consistent
}
type TaskList []Task

func (t TaskList) Add(task Task) TaskList {
	t = append(t, task)
	return t
}
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
		randId, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			fmt.Println("Error generating random ID:", err)
			return
		}
		task := Task{
			ID:        randId.String(),
			Title:     os.Args[2],
			Completed: false,
		}
		err = AppendTask("tasks.jsonl", task)
		if err != nil {
			fmt.Println("Error appending task:", err)
			return
		}
		fmt.Println("Added task:", task)
	case "list":
		fmt.Println("Listing tasks")
	default:
		fmt.Println("Unknown command")
	}
}

// This will append the tasks to the tasks.jsonl file
func AppendTask(filename string, task ...Task) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) //open in append mode
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file) //new encoder
	defer file.Close()               //close the file after the funtion is done
	for _, t := range task {
		if err := encoder.Encode(t); err != nil {
			return err
		}
	}
	return nil
}
