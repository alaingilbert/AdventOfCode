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
	head := &Node{Val: char2Int(input[0])}
	curr := head
	var newNode *Node
	for i := range input[1:] {
		newNode = &Node{Val: char2Int(input[1:][i])}
		curr.Next = newNode
		curr = newNode
	}
	newNode.Next = head
	return head
}

func extendList(head *Node) {
	newList := &Node{Val: 10}
	newHead := newList
	for i := 11; i <= 1_000_000; i++ {
		newList.Next = &Node{Val: i}
		newList = newList.Next
	}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newHead
	newList.Next = head
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
	input := "853192647"
	llist := createList(input)
	cache := buildCache(llist)
	nbMoves := 100
	mixCups(llist, nbMoves, cache)
	// Find node with label 1
	nodeLabelOne := cache[1]
	printList(nodeLabelOne)
}

func part2() {
	input := "853192647"
	llist := createList(input)
	extendList(llist) // extend the list to 1M
	cache := buildCache(llist)
	nbMoves := 10_000_000
	mixCups(llist, nbMoves, cache)
	// Find node with label 1
	nodeLabelOne := cache[1]
	// Multiply next two nodes
	fmt.Println(nodeLabelOne.Next.Val * nodeLabelOne.Next.Next.Val)
}

func main() {
	part1()
	part2()
}
