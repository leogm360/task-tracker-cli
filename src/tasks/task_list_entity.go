package tasks

// Represents a TaskList entity containing a list of tasks
type TaskList struct {
	NextId uint   `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

func NewTaskList(tasks []Task) *TaskList {
	return &TaskList{
		NextId: 1,
		Tasks:  tasks,
	}
}

func (tl *TaskList) Get(id uint) (*Task, int, bool) {
	for i, t := range tl.Tasks {
		if id == t.Id {
			return &t, i, true
		}
	}

	return nil, 0, false
}

func (tl *TaskList) Set(t *Task) bool {
	_, i, ok := tl.Get(t.Id)

	if !ok {
		return ok
	}

	tl.Tasks[i] = *t

	return true
}

func (tl *TaskList) Append(t *Task) {
	tl.Tasks = append(tl.Tasks, *t)
}

func (tl *TaskList) GenId() uint {
	id := tl.NextId

	tl.NextId += 1

	return id
}
