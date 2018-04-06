package main

import "fmt"

type Node struct {
	info string
	next *Node
}

type Queue struct {
	head *Node
}

func createNode(info string) *Node {
	return &Node{info, nil}
}

func (queue *Queue) enqueue(node *Node) {
	var actual *Node = queue.head
	if actual != nil {
		for actual.next != nil {
			actual = actual.next
		}
		actual.next = node
	} else {
		queue.head = node
	}
}

func (queue *Queue) dequeue() *Node {
	var actual *Node = queue.head
	if actual != nil {
		queue.head = queue.head.next
	}
	return actual
}

func (queue *Queue) peek() *Node {
	return queue.head
}

func (queue *Queue) size() int {
	var actual *Node = queue.head
	var count int
	for actual != nil {
		count++
		actual = actual.next
	}
	return count
}

func (queue *Queue) toArray() []string {
	var actual *Node = queue.head
	var list []string
	for actual != nil {
		list = append(list, actual.info)
		actual = actual.next
	}
	return list
}

func (queue *Queue) print() {
	var actual *Node = queue.head
	for actual != nil {
		fmt.Println(actual.info)
		actual = actual.next
	}
}

func main() {
	var queueList = new(Queue)
	queueList.enqueue(createNode("1"))
	queueList.enqueue(createNode("2"))
	queueList.enqueue(createNode("3"))
	queueList.enqueue(createNode("4"))
	queueList.print()
	queueList.dequeue()
	fmt.Println("-----")
	queueList.print()
	fmt.Println("-----")
	fmt.Println(queueList.size())
}
