package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	Name       rune
	NbNodes    int
	NbMetadata int
	Metadata   []int
	Nodes      []Node
}

func (n *Node) GetValue() int {
	val := 0
	if len(n.Nodes) == 0 {
		for _, m := range n.Metadata {
			val += m
		}
	} else {
		for _, m := range n.Metadata {
			if m == 0 {
				continue
			}
			if m-1 >= len(n.Nodes) {
				continue
			}
			val += n.Nodes[m-1].GetValue()
		}
	}
	return val
}

func ParseNode(scanner *bufio.Scanner) Node {
	node := Node{}
	scanner.Scan()
	node.NbNodes, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	node.NbMetadata, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < node.NbNodes; i++ {
		node.Nodes = append(node.Nodes, ParseNode(scanner))
	}
	for i := 0; i < node.NbMetadata; i++ {
		scanner.Scan()
		meta, _ := strconv.Atoi(scanner.Text())
		node.Metadata = append(node.Metadata, meta)
	}
	return node
}

func CalcMetadata(node Node) int {
	res := 0
	q := []Node{node}
	for len(q) > 0 {
		node, q = q[0], q[1:]
		q = append(q, node.Nodes...)
		for _, m := range node.Metadata {
			res += m
		}
	}
	return res
}

func main() {
	by, _ := ioutil.ReadFile("2018/day8/data.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(by)))
	scanner.Split(bufio.ScanWords)
	root := ParseNode(scanner)
	fmt.Println("Part1: ", CalcMetadata(root))
	fmt.Println("Part2: ", root.GetValue())
}
