package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Claim struct {
	ID      int
	Left    int
	Top     int
	Width   int
	Height  int
	Overlap bool
}

type Fabric struct {
	data    []int
	data1   []*Claim
	width   int
	height  int
	Overlap int
}

func NewFabric(w, h int) *Fabric {
	f := new(Fabric)
	f.width = w
	f.height = h
	f.data = make([]int, w*h)
	f.data1 = make([]*Claim, w*h)
	return f
}

func (f *Fabric) ProcessClaim(c *Claim) {
	for w := 0; w < c.Width; w++ {
		for h := 0; h < c.Height; h++ {
			idx := c.Top*f.width + c.Left + w + h*f.width
			f.data[idx]++
			if f.data[idx] == 2 {
				f.Overlap++
			}
			if f.data1[idx] != nil {
				f.data1[idx].Overlap = true
				c.Overlap = true
			}
			f.data1[idx] = c
		}
	}
}

func part1() {
	by, _ := ioutil.ReadFile("2018/day3/data.txt")
	fabric := NewFabric(1000, 1000)
	claims := make([]*Claim, 0)
	for _, line := range strings.Split(string(by), "\n") {
		var id, left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)
		claim := &Claim{id, left, top, width, height, false}
		claims = append(claims, claim)
		fabric.ProcessClaim(claim)
	}
	fmt.Println("Part1:", fabric.Overlap)
	for _, c := range claims {
		if !c.Overlap {
			fmt.Println("Part2:", c.ID)
		}
	}
}

func main() {
	part1()
}
