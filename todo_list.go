package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {

	bytes, _ := os.ReadFile("todos.json")
	
	var tasks []*Task

	json.Unmarshal(bytes, &tasks)

	if len(tasks) == 0 {
		tasks = append(tasks, &Task{Title: "Learn Go", Done: false})
	}

	for _, t := range tasks {
		if t.Title == "Learn Go" {
			t.Done = !t.Done 
			fmt.Println("Updated 'Learn Go' status to:", t.Done)
		}
	}
	newBytes, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile("todos.json", newBytes, 0644)
}
