package main

import (
	"fmt"
	"task-tracker-cli/src/data"
	tasks "task-tracker-cli/src/domain/task"
)

const DatabasePath = "./database.local.json"

func main() {
	newTask, err := tasks.NewTask("Test description")

	dataSource := data.NewJSONDataSource(DatabasePath)

	if err != nil {
		panic(fmt.Errorf("it was not possible to create the task: %s", err))
	}

	newTask.Persist(dataSource)
}
