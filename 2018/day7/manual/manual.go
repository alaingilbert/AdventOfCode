package manual

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

type Manual struct {
	steps        map[string]int
	instructions map[string][]string
	requirements map[string][]string
}

func New(fileName string) *Manual {
	m := new(Manual)
	m.steps = make(map[string]int)
	m.instructions = make(map[string][]string)
	m.requirements = make(map[string][]string)

	by, _ := ioutil.ReadFile(fileName)
	scanner := bufio.NewScanner(strings.NewReader(string(by)))
	var step1, step2 string
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &step1, &step2)
		m.steps[step1] = 0
		m.steps[step2] = 0
		m.instructions[step1] = append(m.instructions[step1], step2)
		m.requirements[step2] = append(m.requirements[step2], step1)
	}

	return m
}

func (m *Manual) GetStepDuration(step string) int {
	return int(step[0]) - 96
}

func (m *Manual) GetSteps() []string {
	steps := make([]string, 0)
	for step := range m.steps {
		steps = append(steps, step)
	}
	return steps
}

func (m *Manual) GetFollowingSteps(step string) []string {
	return m.instructions[step]
}

func (m *Manual) GetRequirements(step string) []string {
	return m.requirements[step]
}
