package main

import (
	"TaskTracker/storage"
	"fmt"
	"os"
	"strconv"
)

func getID(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return id
}

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Println("Usage: ./task-cli <task> [arg]")
		return
	}
	newStorage := storage.NewStorage()
	switch arg[0] {
	case "add":
		if len(arg) != 2 {
			fmt.Println("Usage: ./task-cli add <Task>")
		}
		newStorage.AddTask(arg[1])
	case "update":
		if len(arg) != 3 {
			fmt.Println("Usage: ./task-cli update <Id> <Task>")
		}
		id := getID(arg[1])
		newStorage.UpdateDescriptionTask(id-1, arg[2])
	case "delete":
		if len(arg) != 2 {
			fmt.Println("Usage: ./task-cli delete <Id>")
		}
		id := getID(arg[1])
		newStorage.RemoveTask(id - 1)
	case "mark-in-progress":
		if len(arg) != 2 {
			fmt.Println("Usage: ./task-cli mark-in-progress <Id>")
		}
		id := getID(arg[1])
		newStorage.UpdateStatusTask(id-1, storage.StatusInProgress)
	case "mark-done":
		if len(arg) != 2 {
			fmt.Println("Usage: ./task-cli mark-done <Id>")
		}
		id := getID(arg[1])
		newStorage.UpdateStatusTask(id-1, storage.StatusDone)
	case "list":
		if len(arg) == 1 {
			newStorage.ShowAllTasks("")
		} else {
			newStorage.ShowAllTasks(arg[1])
		}
	default:
		fmt.Println("Incorrect task")
	}
}
