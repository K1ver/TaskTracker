package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Status string

const (
	StatusToDo       Status = "todo"
	StatusDone              = "done"
	StatusInProgress        = "in-progress"
)

type Task struct {
	Id          int
	Description string
	Status      Status
	CreatedAt   string
	UpdatedAt   string
}

type Storage struct {
	Tasks []Task `json:"tasks"`
}

func NewStorage() *Storage {
	var s Storage
	exePath, _ := os.Executable()
	bolB, _ := os.ReadFile(filepath.Join(filepath.Dir(exePath), "tasks.json"))
	err := json.Unmarshal(bolB, &s)
	if err == nil {
		return &s
	}
	return &Storage{}
}

func (s *Storage) AddTask(description string) {
	s.Tasks = append(s.Tasks, Task{Id: len(s.Tasks) + 1, Description: description, Status: StatusToDo, CreatedAt: time.Now().Format("2006-01-02"), UpdatedAt: time.Now().Format("2006-01-02")})
	s.saveTask()
	fmt.Printf("Task added successfully (ID: %d)", len(s.Tasks))
}

func (s *Storage) UpdateDescriptionTask(id int, description string) {
	task := s.Tasks[id]
	task.Description = description
	task.UpdatedAt = time.Now().Format("2006-01-02")
	s.Tasks[id] = task
	s.saveTask()
}

func (s *Storage) UpdateStatusTask(id int, status Status) {
	task := s.Tasks[id]
	task.Status = status
	task.UpdatedAt = time.Now().Format("2006-01-02")
	s.Tasks[id] = task
	s.saveTask()
}

func (s *Storage) RemoveTask(id int) {
	copy(s.Tasks[id:], s.Tasks[id+1:])
	s.Tasks[len(s.Tasks)-1] = Task{}
	s.Tasks = s.Tasks[:len(s.Tasks)-1]
	for i := id; i < len(s.Tasks); i++ {
		s.Tasks[i].Id--
	}
	s.saveTask()
}

func (s *Storage) saveTask() {
	bolB, _ := json.Marshal(s)
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("error getting executable path:", err)
		return
	}
	path := filepath.Join(filepath.Dir(exePath), "tasks.json")
	err = os.WriteFile(path, bolB, 0644)
	if err != nil {
		fmt.Println("error writing tasks.json", err)
		return
	}
}

func (s *Storage) ShowAllTasks(status string) {
	for _, task := range s.Tasks {
		if status == "" {
			fmt.Printf("Id: %d, Description: %s, Status: %v, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		} else if task.Status == Status(status) {
			fmt.Printf("Id: %d, Description: %s, Status: %v, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}
