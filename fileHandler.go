package main

import (
	"encoding/json"
	"os"
)

// This will append the tasks to the tasks.jsonl file
func appendTask(filename string, task ...Task) error {
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

func getAllTasks(filename string) (TaskList, error) {
	file, err := os.Open(filename) //open in read mode
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks TaskList
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var task Task
		if err := decoder.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func writeTasks(filename string, tasks TaskList) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) //open in write/truncate mode
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, t := range tasks {
		if err := encoder.Encode(t); err != nil {
			return err
		}
	}
	return nil
}

