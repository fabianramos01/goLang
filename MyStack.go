package main

import "fmt"

type Node struct {
	info string
	next *Node
}

type Stack struct {
	head *Node
}

func createNode(info string) *Node {
	return &Node{info, nil}
}

func (stack *Stack) push(node *Node) {
	if stack.head != nil {
		node.next = stack.head
		stack.head = node
	} else {
		stack.head = node
	}
}

func (stack *Stack) pop() *Node {
	var actual *Node = stack.head
	if actual != nil {
		stack.head = stack.head.next
	}
	return actual
}

func (stack *Stack) peek() *Node {
	return stack.head
}

func (stack *Stack) size() int {
	var actual *Node = stack.head
	var count int
	for actual != nil {
		count++
		actual = actual.next
	}
	return count
}

func (stack *Stack) toArray() []string {
	var actual *Node = stack.head
	var list []string
	for actual != nil {
		list = append(list, actual.info)
		actual = actual.next
	}
	return list
}

func (stack *Stack) print() {
	var actual *Node = stack.head
	for actual != nil {
		fmt.Println(actual.info)
		actual = actual.next
	}
}

func main() {
	var stack = Stack{}
	stack.push(createNode("1"))
	stack.push(createNode("2"))
	stack.push(createNode("3"))
	stack.push(createNode("4"))
	stack.print()
	stack.pop()
	fmt.Println("-----")
	stack.print()
	fmt.Println("-----")
	fmt.Println(stack.size())
}
