package main

import (
	"AdventOfCode/2018/day7/manual"
	"fmt"
	"sort"
)

type Task struct {
}

type Builder struct {
	manual         *manual.Manual
	availableTasks []string
	progress       map[string]int
}

func NewBuilder(manual *manual.Manual) *Builder {
	b := new(Builder)
	b.manual = manual
	b.availableTasks = make([]string, 0)
	b.progress = make(map[string]int)

	for _, step := range b.manual.GetSteps() {
		requirements := b.manual.GetRequirements(step)
		b.progress[step] += len(requirements)
		if len(requirements) == 0 {
			b.availableTasks = append(b.availableTasks, step)
		}
	}

	return b
}

func (b *Builder) GetNextTask() (task string) {
	sort.Strings(b.availableTasks)
	task, b.availableTasks = b.availableTasks[0], b.availableTasks[1:] // pop
	return
}

func (b *Builder) TaskCompleted(task string) {
	followingSteps := b.manual.GetFollowingSteps(task)
	for _, followingStep := range followingSteps {
		b.progress[followingStep]--
		if b.progress[followingStep] == 0 {
			b.UnlockStep(followingStep)
		}
	}
}

func (b *Builder) UnlockStep(step string) {
	b.availableTasks = append(b.availableTasks, step)
}

func (b *Builder) GetBuildOrder() (answer string) {
	for len(b.availableTasks) > 0 {
		task := b.GetNextTask()
		answer += task
		b.TaskCompleted(task)
	}
	return answer
}

func part1() {
	m := manual.New("2018/day7/small.txt")
	fmt.Println("Part1: " + NewBuilder(m).GetBuildOrder())
}

func part2() {

}

func main() {
	part1()
	part2()
}
