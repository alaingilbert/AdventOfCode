package elfManager

import (
	"AdventOfCode/2018/day7/builder"
	"AdventOfCode/2018/day7/task"
	"sync/atomic"
)

type Elf struct {
	busy bool
	Task *task.Task
}

func (e *Elf) IsBusy() bool {
	return e.busy
}

func (e *Elf) Update() {
	if e.Task != nil {
		e.Task.Progress++
	}
}

func (e *Elf) DoTask(t *task.Task) {
	e.busy = true
	e.Task = t
}

type ElfManager struct {
	ElfAvailable int32
	Builder      *builder.Builder
	Elfs         []*Elf
}

func New(b *builder.Builder, nbr int) *ElfManager {
	m := new(ElfManager)
	m.Builder = b
	m.Elfs = make([]*Elf, nbr)
	for i := 0; i < nbr; i++ {
		m.Elfs[i] = new(Elf)
	}
	atomic.StoreInt32(&m.ElfAvailable, int32(nbr))
	return m
}

func (m *ElfManager) GetElf() *Elf {
	for _, elf := range m.Elfs {
		if !elf.IsBusy() {
			return elf
		}
	}
	return nil
}

func (m *ElfManager) HasElfAvailable() bool {
	return atomic.LoadInt32(&m.ElfAvailable) > 0
}

func (m *ElfManager) AssignTask(t *task.Task, elf *Elf) {
	elf.DoTask(t)
}

func (m *ElfManager) Update() {
	for _, elf := range m.Elfs {
		if elf.IsBusy() {
			task := elf.Task
			if task.Progress == task.Duration {
				m.Builder.TaskCompleted(task)
				elf.Task = nil
				elf.busy = false
			}
		}
	}
}
