package models

type Task struct {
	Description string
	ID          int
	Complete    bool
}

func NewTask(desc string, id int) Task {
	return Task{
		Description: desc,
		ID:          id,
		Complete:    false,
	}
}

func ValueExists(task *Task, description string) bool {
	if task.Description == description {
		return true
	}
	return false
}
