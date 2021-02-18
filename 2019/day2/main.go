package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	p1()
	p2()
}

func p1() {
	by, _ := ioutil.ReadFile("2019/day2/input.txt")
	arrStr := strings.Split(string(by), ",")
	arr := make([]int64, len(arrStr))
	for idx, el := range arrStr {
		arr[idx], _ = strconv.ParseInt(el, 10, 64)
	}
	arr[1] = 12
	arr[2] = 2
	idx := 0
	for {
		op := arr[idx+0]
		if op == 99 {
			break
		} else if op == 1 {
			arr[arr[idx+3]] = arr[arr[idx+1]] + arr[arr[idx+2]]
		} else if op == 2 {
			arr[arr[idx+3]] = arr[arr[idx+1]] * arr[arr[idx+2]]
		} else {
			fmt.Println("ERROR")
		}
		idx += 4
	}
	fmt.Println("Res:", arr[0])
}

func p2() {
	by, _ := ioutil.ReadFile("2019/day2/input.txt")
	arrStr := strings.Split(string(by), ",")
	var noun, verb int64
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			arr := make([]int64, len(arrStr))
			for idx, el := range arrStr {
				arr[idx], _ = strconv.ParseInt(el, 10, 64)
			}
			arr[1] = noun
			arr[2] = verb
			idx := 0
			for {
				op := arr[idx+0]
				if op == 99 {
					break
				} else if op == 1 {
					arr[arr[idx+3]] = arr[arr[idx+1]] + arr[arr[idx+2]]
				} else if op == 2 {
					arr[arr[idx+3]] = arr[arr[idx+1]] * arr[arr[idx+2]]
				} else {
					fmt.Println("ERROR")
				}
				idx += 4
			}
			if arr[0] == 19690720 {
				fmt.Println("Res:", noun*100+verb)
			}
		}
	}
}
