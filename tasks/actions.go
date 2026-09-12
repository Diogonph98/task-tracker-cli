package tasks

func AddTask(task string) []Task {
	var newTask Task
	newTask.ID = 1
	newTask.Title = task
	newTask.IsCompleted = false

	// When appending an element to a list, it creates a new list with the appended element
	// So, we need to assign the new list back to allTasks
	allTasks = append(allTasks, newTask)
	println("You added a new task!")
	return allTasks
}

func ListTasks() {
	for _, task := range allTasks {
		println(task.ID, task.Title, task.IsCompleted)
	}
}
