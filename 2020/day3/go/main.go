package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getTrees(lines []string, right, down int) int {
	lineLen := len(lines[0])
	var row, col, trees int
	for i := 0; i < len(lines); i++ {
		if row > len(lines) {
			break
		}
		if lines[row][col%lineLen] == '#' {
			trees += 1
		}
		row += down
		col += right
	}
	return trees
}

func part1() {
	by, _ := ioutil.ReadFile("2020/day3/input.txt")
	lines := strings.Split(string(by), "\n")
	fmt.Println(getTrees(lines, 3, 1))
}

func part2() {
	by, _ := ioutil.ReadFile("2020/day3/input.txt")
	lines := strings.Split(string(by), "\n")
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	res := 1
	for _, slope := range slopes {
		trees := getTrees(lines, slope[0], slope[1])
		res *= trees
	}
	fmt.Println(res)
}

func main() {
	part1()
	part2()
}
