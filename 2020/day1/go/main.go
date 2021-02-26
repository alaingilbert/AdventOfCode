package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sortInsert(arr []int, el int) []int {
	i := sort.SearchInts(arr, el)
	arr = append(arr, 0)
	copy(arr[i+1:], arr[i:])
	arr[i] = el
	return arr
}

func readLines() (out []int) {
	f, _ := os.Open("./2020/day1/go/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		out = sortInsert(out, v)
	}
	return
}

func part1() {
	lines := readLines()
	cache := make(map[int]bool)
	for _, line := range lines {
		target := 2020 - line
		if _, ok := cache[target]; ok {
			fmt.Println(line, target, line*target)
			return
		}
		cache[line] = true
	}
}

func part2() {
	lines := readLines()
	cache := make(map[int]bool)
	for _, line := range lines {
		cache[line] = true
	}
	for i := 0; i < len(lines)-2; i++ {
		first := lines[i]
		for j := i + 1; j < len(lines)-1; j++ {
			second := lines[j]
			third := 2020 - first - second
			if _, ok := cache[third]; ok {
				fmt.Println(first, second, third, first*second*third)
				return
			}
		}
	}
	fmt.Println(lines)
}

func main() {
	part1()
	part2()
}
