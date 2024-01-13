package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	lastOne := &Node{Value: 1}
	elemNode := &Node{2, lastOne}
	head := &Node{3, elemNode}

	Node := Reverse(head)
	fmt.Println(Node)

}

func Reverse(list *Node) *Node {
	pre := new(Node)
	pre = nil
	for list != nil {
		tmp := list
		list = list.Next

		tmp.Next = pre
		pre = tmp
	}

	return pre
}
