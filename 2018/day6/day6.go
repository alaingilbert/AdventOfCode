package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) Distance(v Coord) int {
	return int(math.Abs(float64(v.X-c.X)) + math.Abs(float64(v.Y-c.Y)))
}

func loadCoords() (coords []Coord, width, height int) {
	by, _ := ioutil.ReadFile("2018/day6/data.txt")
	minX, minY := math.MaxInt32, math.MaxInt32
	width, height = -1, -1
	for _, line := range strings.Split(string(by), "\n") {
		var coord Coord
		fmt.Sscanf(line, "%d, %d", &coord.X, &coord.Y)
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > width {
			width = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > height {
			height = coord.Y
		}
		coords = append(coords, coord)
	}

	// Normalize coordinates
	for i := range coords {
		coords[i].X -= minX
		coords[i].Y -= minY
	}

	width -= minX
	height -= minY
	width++
	height++

	return
}

type Map struct {
	width  int
	height int
	data   []int
	coords []Coord
}

func NewMap(width, height int) *Map {
	m := new(Map)
	m.width = width
	m.height = height
	m.data = make([]int, m.width*m.height)
	return m
}

func (m Map) idx2Coord(idx int) Coord {
	y := idx / m.width
	return Coord{X: idx - (y * m.width), Y: y}
}

func (m *Map) SetCoords(coords []Coord) *Map {
	m.coords = coords
	return m
}

func (m *Map) Calc() *Map {
	for i := range m.data {
		cellCoord := m.idx2Coord(i)
		minDist := math.MaxInt32
		minIdx := 0
		equal := false
		for coordIdx, coord := range m.coords {
			dist := cellCoord.Distance(coord)
			if dist < minDist {
				minDist = dist
				minIdx = coordIdx + 1
				equal = false
			} else if dist == minDist {
				equal = true
			}
		}
		if equal {
			minDist = -1
			minIdx = -1
		}
		m.data[i] = minIdx
	}
	return m
}

func (m *Map) Display() *Map {
	for i := 0; i < len(m.data); i++ {
		fmt.Printf("%02d ", m.data[i])
		if (i+1)%m.width == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
	return m
}

// Return either or not an index is part of the contour
func (m *Map) isContour(idx int) bool {
	coord := m.idx2Coord(idx)
	return coord.Y == 0 || coord.Y == m.height-1 || coord.X == 0 || coord.X == m.width-1
}

func (m *Map) MaxArea() *Map {
	// Build a map of infinite area
	infinite := make(map[int]bool)
	infinite[-1] = true
	for i, v := range m.data {
		if m.isContour(i) {
			infinite[v] = true
		}
	}
	// Build a map of areaID -> size
	areas := make(map[int]int)
	for _, v := range m.data {
		if _, ok := infinite[v]; !ok {
			areas[v]++
		}
	}
	// Find largest area
	maxSize := 0
	maxID := 0
	for areaID, areaSize := range areas {
		if areaSize > maxSize {
			maxSize = areaSize
			maxID = areaID
		}
	}
	fmt.Println("Max area is: ", maxID, "with", maxSize)
	return m
}

func p1() {
	coords, width, height := loadCoords()
	NewMap(width, height).SetCoords(coords).Calc().Display().MaxArea()
}

func main() {
	p1()
}
