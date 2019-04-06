package tree

import "fmt"

func (node *Node) PreTraverse() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.Traverse()
	node.Right.Traverse()
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) BackTraverse() {
	if node == nil {
		return
	}
	node.Right.Traverse()
	node.Left.Traverse()
	node.Print()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) NewTraverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}
