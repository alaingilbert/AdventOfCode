package task

type Task struct {
	Name     string
	Duration int
	Progress int
}

func New(name string) *Task {
	t := new(Task)
	t.Name = name
	t.Duration = (int(name[0]) - 64) + 60
	return t
}
