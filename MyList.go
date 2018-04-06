package main

import "fmt"

type Node struct {
	info string
	next *Node
}

type MyList struct {
	head *Node
}

func createNode(info string) *Node {
	return &Node{info, nil}
}

func (list *MyList) addLast(node *Node) {
	var actual *Node = list.head
	if actual != nil {
		for actual.next != nil {
			actual = actual.next
		}
		actual.next = node
	} else {
		list.head = node
	}
}

func (list *MyList) addFirst(node *Node) {
	if list.head != nil {
		node.next = list.head
		list.head = node
	} else {
		list.head = node
	}
}

func (list *MyList) getFirst() *Node {
	var actual *Node = list.head
	if actual != nil {
		list.head = list.head.next
	}
	return actual
}

func (list *MyList) peek() *Node {
	return list.head
}

func (list *MyList) size() int {
	var actual *Node = list.head
	var count int
	for actual != nil {
		count++
		actual = actual.next
	}
	return count
}

func (list *MyList) toArray() []string {
	var actual *Node = list.head
	var dataList []string
	for actual != nil {
		dataList = append(dataList, actual.info)
		actual = actual.next
	}
	return dataList
}

func (list *MyList) print() {
	var actual *Node = list.head
	for actual != nil {
		fmt.Println(actual.info)
		actual = actual.next
	}
}

func main() {
	var list = new(MyList)
	list.addFirst(createNode("1"))
	list.addFirst(createNode("2"))
	list.addLast(createNode("3"))
	list.addLast(createNode("4"))
	list.print()
	list.getFirst()
	fmt.Println("-----")
	list.print()
	fmt.Println("-----")
	fmt.Println(list.peek().info)
}
