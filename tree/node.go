package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node ignore")
		return
	}
	node.Value = Value
}

func main() {
	var root Node

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)

	root.Traverse()
	fmt.Println()

	nodes := []Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()

	root.SetValue(4)
	root.Print()

	var pRoot *Node
	pRoot.SetValue(100)
	pRoot = &root
	pRoot.SetValue(200)
	pRoot.Print()

}
