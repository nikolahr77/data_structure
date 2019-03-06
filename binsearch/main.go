package main

import "fmt"

// Node is a node of a binary tree
type Node struct {
	Value      int   // Value is the value of the current node
	LeftChild  *Node // LeftChild is a new node that contains the value that is less than the parent value
	RightChild *Node // RightChild is a new node that contains the value that is more than the parent value
}

// NewNode creates a new node
func NewNode(v int) *Node {
	return &Node{Value: v}
}

// AddNode adds a node
func (n *Node) AddNode(v int) {
	if v < n.Value {
		if n.LeftChild == nil {
			n.LeftChild = NewNode(v)
			return
		}
		n.LeftChild.AddNode(v)
		return
	}
	if n.RightChild == nil {
		n.RightChild = NewNode(v)
		return
	}
	n.RightChild.AddNode(v)
}

// PrintNodes prints all of the nodes in order
func (s *Node) PrintNodes() {
	if s != nil {
		s.LeftChild.PrintNodes()
		fmt.Println(s.Value)
		s.RightChild.PrintNodes()
	}
}

func main() {
	s := NewNode(5)
	s.AddNode(8)
	s.AddNode(1)
	s.AddNode(3)
	s.AddNode(10)
	s.PrintNodes()
	fmt.Printf("%#v", s)
}
