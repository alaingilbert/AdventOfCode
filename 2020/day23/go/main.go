package main

import (
	"fmt"
	"strconv"
)

// Node ...
type Node struct {
	Next *Node
	Val  int
}

func char2Int(c byte) int {
	v, _ := strconv.Atoi(string(c))
	return v
}

func createList(input string) *Node {
	first, rest := input[0], input[1:]
	head := &Node{Val: char2Int(first)}
	curr := head
	for i := range rest {
		curr.Next = &Node{Val: char2Int(rest[i])}
		curr = curr.Next
	}
	curr.Next = head
	return head
}

func extendList(llist *Node) {
	head := &Node{Val: 10}
	curr := head
	for i := 11; i <= 1_000_000; i++ {
		curr.Next = &Node{Val: i}
		curr = curr.Next
	}
	curr.Next = llist
	llist.Next.Next.Next.Next.Next.Next.Next.Next.Next = head
}

func printList(n *Node) {
	tmp := n
	for {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
		if tmp == nil || tmp == n {
			break
		}
	}
	fmt.Println()
}

func removeThree(currNode *Node) (removed *Node) {
	removed = currNode.Next
	currNode.Next = currNode.Next.Next.Next.Next
	removed.Next.Next.Next = nil
	return
}

func addRemoved(dest, removed *Node) {
	tmp := dest.Next
	dest.Next = removed
	removed.Next.Next.Next = tmp
}

func isRemoved(n *Node, v int) (isRemoved bool) {
	tmp := n
	for tmp != nil {
		if tmp.Val == v {
			return true
		}
		tmp = tmp.Next
	}
	return
}

func findDestinationUsingCache(currNode, removed *Node, cache map[int]*Node) (dest *Node) {
	newVal := currNode.Val
	for {
		newVal--
		if newVal == 0 {
			newVal = len(cache)
		}
		if !isRemoved(removed, newVal) {
			break
		}
	}
	dest = cache[newVal]
	return
}

func buildCache(head *Node) map[int]*Node {
	out := make(map[int]*Node)
	tmp := head
	for {
		out[tmp.Val] = tmp
		tmp = tmp.Next
		if tmp == head {
			break
		}
	}
	return out
}

func mixCups(currNode *Node, nbMoves int, cache map[int]*Node) {
	for i := 0; i < nbMoves; i++ {
		removed := removeThree(currNode)
		destNode := findDestinationUsingCache(currNode, removed, cache)
		addRemoved(destNode, removed)
		currNode = currNode.Next
	}
}

func part1() {
	llist := createList(puzzleInput)
	cache := buildCache(llist)
	nbMoves := 100
	mixCups(llist, nbMoves, cache)
	// Find node with label 1
	nodeLabelOne := cache[1]
	printList(nodeLabelOne)
}

func part2() {
	llist := createList(puzzleInput)
	extendList(llist) // extend the list to 1M
	cache := buildCache(llist)
	nbMoves := 10_000_000
	mixCups(llist, nbMoves, cache)
	// Find node with label 1
	nodeLabelOne := cache[1]
	// Multiply next two nodes
	fmt.Println(nodeLabelOne.Next.Val * nodeLabelOne.Next.Next.Val)
}

var puzzleInput = "853192647"

func main() {
	part1()
	part2()
}
