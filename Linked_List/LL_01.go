//Linked List Insertion

package main1

import "fmt"

// create structure
type Node struct {
	data int
	next *Node
}

// Constructor
type LinkedList struct {
	head *Node
}

// Insert At Head

func (l *LinkedList) insertHead(data int) {
	// Create A newNode
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}

}

//Print Function
func (l *LinkedList) print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Print()
}

func main1() {
	list := LinkedList{}
	list.insertHead(20)
	list.insertHead(30)
	list.insertHead(40)
	list.insertHead(50)
	list.print()

}
