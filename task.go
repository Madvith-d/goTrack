package main

type Task struct {
	ID        string `json:"id"`        //remap the fields to theior lowercase version while encoding
	Title     string `json:"title"`     //so that the task.jsonl file is consistent
	Completed bool   `json:"completed"` //so that the task.jsonl file is consistent
}
type TaskList []Task
