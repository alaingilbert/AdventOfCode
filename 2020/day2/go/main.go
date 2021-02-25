package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	by, _ := ioutil.ReadFile("./2020/day2/input.txt")
	lines := strings.Split(string(by), "\n")
	rgx := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	good := 0
	for _, line := range lines {
		parts := rgx.FindStringSubmatch(line)
		min, _ := strconv.Atoi(parts[1])
		max, _ := strconv.Atoi(parts[2])
		letter := parts[3]
		password := parts[4]
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			good++
		}
	}
	fmt.Println(good)
}

func part2() {
	by, _ := ioutil.ReadFile("./2020/day2/input.txt")
	lines := strings.Split(string(by), "\n")
	rgx := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	good := 0
	for _, line := range lines {
		parts := rgx.FindStringSubmatch(line)
		first, _ := strconv.Atoi(parts[1])
		second, _ := strconv.Atoi(parts[2])
		letter := parts[3]
		password := parts[4]
		if (password[first-1] == letter[0] ||
			password[second-1] == letter[0]) &&
			password[first-1] != password[second-1] {
			good++
		}
	}
	fmt.Println(good)
}

func main() {
	part1()
	part2()
}
