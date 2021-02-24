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

func createList(input string) *Node {
	v, _ := strconv.Atoi(string(input[0]))
	head := &Node{Val: v}
	prev := head
	var newNode *Node
	for _, c := range input[1:] {
		v, _ := strconv.Atoi(string(c))
		newNode = &Node{Val: v}
		prev.Next = newNode
		prev = newNode
	}
	newNode.Next = head
	return head
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

func removeThree(currNode *Node) (removed *Node) {
	removed = currNode.Next
	currNode.Next = currNode.Next.Next.Next.Next
	removed.Next.Next.Next = nil
	return
}

func findDestination(currNode, removed *Node) (dest *Node) {
	newVal := currNode.Val
	for {
		newVal--
		if newVal == 0 {
			newVal = 9
		}
		if !isRemoved(removed, newVal) {
			break
		}
	}
	dest = currNode
	for dest.Val != newVal {
		dest = dest.Next
	}
	return
}

func addRemoved(dest, removed *Node) {
	tmp := dest.Next
	dest.Next = removed
	removed.Next.Next.Next = tmp
}

func part1() *Node {
	input := "853192647"
	currNode := createList(input)
	for i := 0; i < 100; i++ {
		removed := removeThree(currNode)
		destNode := findDestination(currNode, removed)
		addRemoved(destNode, removed)
		currNode = currNode.Next
	}
	for currNode.Val != 1 {
		currNode = currNode.Next
	}
	printList(currNode)
	return currNode
}

func part2(currNode *Node) {
	newList := &Node{Val: 10}
	head := newList
	for i := 11; i <= 100; i++ {
		newList.Next = &Node{Val: i}
		newList = newList.Next
	}
	printList(head)
}

func main() {
	currNode := part1()
	part2(currNode)
}
