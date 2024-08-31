package cli

import (
	"task-tracker-cli/src/data"
	"task-tracker-cli/src/tasks"
)

// Create CLI command that creates a new task with description and save it to
// the data source
func CreateTask(description string) {
	dataSource := data.NewDataSourceJSON[tasks.TaskList]("tasks.json")

	taskList := tasks.NewTaskList([]tasks.Task{})

	dataSource.Read(taskList)

	newTask := tasks.NewTask(taskList.GenId(), description)

	taskList.Append(newTask)

	dataSource.Write(taskList)
}
