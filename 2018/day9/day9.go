package main

import "fmt"

type Player struct {
	Points int
}

type Marble struct {
	Value int
	Next  *Marble
	Prev  *Marble
}

type Circle struct {
	Initial *Marble
	Current *Marble
}

func (c *Circle) PlaceMarble(m *Marble) {
	if c.Current == nil {
		c.Current = m
		c.Current.Next = m
		c.Current.Prev = m
		c.Initial = m
	} else {
		next := c.Current.Next
		m.Next = next.Next
		m.Prev = next
		next.Next.Prev = m
		next.Next = m
		c.Current = m
	}
}

func (c *Circle) SpecialCase() *Marble {
	for i := 0; i < 7; i++ {
		c.Current = c.Current.Prev
	}
	toRemove := c.Current
	toRemove.Prev.Next = toRemove.Next
	c.Current = toRemove.Next
	return toRemove
}

func (c *Circle) Print() {
	it := c.Initial
	for {
		if it == c.Current {
			fmt.Printf("(%d) ", it.Value)
		} else {
			fmt.Printf("%d ", it.Value)
		}
		it = it.Next
		if it == c.Initial {
			break
		}
	}
	fmt.Print("\n")
}

func getScore(nbPlayers, totalMarble int) int {
	players := make([]*Player, nbPlayers)
	for i := 0; i < nbPlayers; i++ {
		players[i] = &Player{}
	}
	circle := Circle{}
	circle.PlaceMarble(&Marble{Value: 0})
	//fmt.Printf("[-] ")
	//circle.Print()
	for i := 1; i <= totalMarble; i++ {
		playerID := (i - 1) % nbPlayers
		marble := &Marble{Value: i}
		if marble.Value%23 == 0 {
			player := players[playerID]
			player.Points += marble.Value
			removedMarble := circle.SpecialCase()
			player.Points += removedMarble.Value
		} else {
			circle.PlaceMarble(marble)
		}
		//fmt.Printf("[%d] ", playerID)
		//circle.Print()
	}
	maxPoints := 0
	for _, p := range players {
		if p.Points > maxPoints {
			maxPoints = p.Points
		}
	}
	return maxPoints
}

func part1() {
	pts := getScore(465, 71940)
	fmt.Println("Part1: ", pts)
}

func part2() {
	pts := getScore(465, 7194000)
	fmt.Println("Part2: ", pts)
}

func main() {
	part1()
	part2()
}
