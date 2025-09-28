package main

import "fmt"

// Create Node
type node struct {
	data int
	next *node
}

type LinkList struct {
	head *node
}

// insert Funtion At TAil
func (l *LinkList) insertTail(data int) {
	// create a new node
	newNode := &node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode

}

// Print Funtion
func (l *LinkList) print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
}

func main() {
	list := LinkList{}
	list.insertTail(10)
	list.insertTail(20)
	list.insertTail(30)
	list.insertTail(40)
	list.insertTail(50)

	list.print()
}
