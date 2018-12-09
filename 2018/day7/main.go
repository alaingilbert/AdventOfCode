package main

import (
	"AdventOfCode/2018/day7/builder"
	"AdventOfCode/2018/day7/elfManager"
	"AdventOfCode/2018/day7/manual"
	"fmt"
)

func part1() {
	m := manual.New("2018/day7/small.txt")
	fmt.Println("Part1: " + builder.New(m).GetBuildOrder())
}

func part2() {
	secs := 0
	manual := manual.New("2018/day7/data.txt")
	builder := builder.New(manual)
	elfManager := elfManager.New(builder, 5)
	for {

		for {
			worker := elfManager.GetElf()
			if builder.HasTaskAvailable() && worker != nil {
				task := builder.GetNextTask()
				elfManager.AssignTask(task, worker)
			} else {
				break
			}
		}

		fmt.Printf("% 3d | ", secs)
		for _, elf := range elfManager.Elfs {
			elf.Update()
			if elf.IsBusy() {
				fmt.Print(elf.Task.Name, " | ")
			} else {
				fmt.Print(".", " | ")
			}
		}
		for _, t := range builder.Completed {
			fmt.Print(t.Name)
		}
		fmt.Print("\n")

		elfManager.Update()
		if builder.IsDone() {
			break
		}
		secs++
	}

	fmt.Print("Part2: ", secs+1)
}

func main() {
	part1()
	part2()
}
