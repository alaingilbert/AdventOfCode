package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func p1() {
	by, err := ioutil.ReadFile("2018/day2/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	count2 := 0
	count3 := 0
	for _, line := range strings.Split(string(by), "\n") {
		m := make(map[int32]int)
		for _, c := range line {
			m[c]++
		}
		for _, v := range m {
			if v == 2 {
				count2++
				break
			}
		}
		for _, v := range m {
			if v == 3 {
				count3++
				break
			}
		}
	}
	fmt.Println(count2 * count3)
}

func p2() {
	by, err := ioutil.ReadFile("2018/day2/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var res []string
	lines := strings.Split(string(by), "\n")
	for i := 0; i < len(lines)-1; i++ {
	LOOP2:
		for j := i + 1; j < len(lines); j++ {
			count := 0
			for ii := range lines[j] {
				if lines[i][ii] != lines[j][ii] {
					count++
					if count > 1 {
						continue LOOP2
					}
					res = strings.Split(lines[i], "")
					res[ii] = "_"
				}
			}
			fmt.Println("CALISS", lines[i], lines[j], strings.Join(res, ""))
		}
	}
}

func main() {
	p1()
	p2()
}
