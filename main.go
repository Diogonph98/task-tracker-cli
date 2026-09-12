package main

import (
	"os"
	"task-tracker-cli/tasks"
)

func main() {
	if len(os.Args) < 2 {
		println("You didn't add a valid command! Please try again.")
		return
	}

	actionArg := os.Args[1]

	switch actionArg {
	case "add":
		tasks.AddTask(os.Args[2])

	case "check":
		tasks.ListTasks()
	}

}
