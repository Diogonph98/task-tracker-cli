package tasks

type Task struct {
	ID          int
	Title       string
	IsCompleted bool
}

var allTasks []Task
