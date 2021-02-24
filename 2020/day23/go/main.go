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

func findDestination2(currNode, removed *Node, cache map[int]*Node) (dest *Node) {
	newVal := currNode.Val
	for {
		newVal--
		if newVal == 0 {
			newVal = 1000000
		}
		if !isRemoved(removed, newVal) {
			break
		}
	}
	dest = cache[newVal]
	return
}

func addRemoved(dest, removed *Node) {
	tmp := dest.Next
	dest.Next = removed
	removed.Next.Next.Next = tmp
}

func part1() {
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

func part2() {
	input := "853192647"
	head := createList(input)
	newList := &Node{Val: 10}
	newHead := newList
	for i := 11; i <= 1000000; i++ {
		newList.Next = &Node{Val: i}
		newList = newList.Next
	}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newHead
	newList.Next = head

	cache := buildCache(head)

	currNode := head
	for i := 0; i < 10000000; i++ {
		removed := removeThree(currNode)
		destNode := findDestination2(currNode, removed, cache)
		addRemoved(destNode, removed)
		currNode = currNode.Next
	}
	for currNode.Val != 1 {
		currNode = currNode.Next
	}
	fmt.Println(currNode.Next.Val * currNode.Next.Next.Val)
}

func main() {
	part1()
	part2()
}
