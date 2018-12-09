package builder

import (
	"AdventOfCode/2018/day7/manual"
	"AdventOfCode/2018/day7/task"
	"sort"
)

type Builder struct {
	manual         *manual.Manual
	availableTasks []*task.Task
	progress       map[string]int
	Completed      []*task.Task
}

func New(manual *manual.Manual) *Builder {
	b := new(Builder)
	b.manual = manual
	b.availableTasks = make([]*task.Task, 0)
	b.progress = make(map[string]int)
	steps := b.manual.GetSteps()
	for _, step := range steps {
		requirements := b.manual.GetRequirements(step)
		b.progress[step] += len(requirements)
		if len(requirements) == 0 {
			b.availableTasks = append(b.availableTasks, task.New(step))
		}
	}

	return b
}

func (b *Builder) HasTaskAvailable() bool {
	return len(b.availableTasks) > 0
}

func (b *Builder) GetNextTask() (t *task.Task) {
	if len(b.availableTasks) == 0 {
		return nil
	}
	sort.Slice(b.availableTasks, func(i, j int) bool {
		return b.availableTasks[i].Name < b.availableTasks[j].Name
	})
	t, b.availableTasks = b.availableTasks[0], b.availableTasks[1:] // pop
	return
}

func (b *Builder) TaskCompleted(t *task.Task) {
	b.Completed = append(b.Completed, t)
	followingSteps := b.manual.GetFollowingSteps(t.Name)
	for _, followingStep := range followingSteps {
		b.progress[followingStep]--
		if b.progress[followingStep] == 0 {
			b.UnlockStep(task.New(followingStep))
		}
	}
}

func (b *Builder) IsDone() bool {
	return len(b.manual.GetSteps()) == len(b.Completed)
}

func (b *Builder) UnlockStep(t *task.Task) {
	b.availableTasks = append(b.availableTasks, t)
}

func (b *Builder) GetBuildOrder() (answer string) {
	for len(b.availableTasks) > 0 {
		task := b.GetNextTask()
		answer += task.Name
		b.TaskCompleted(task)
	}
	return answer
}
