package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
}

func main() {
	linklist := linkedList{
		&node{next: nil}}
	linklist.insertAtTail(10)
	linklist.display()
	linklist.insertAtTail(20)
	linklist.display()

}

func (linklist *linkedList) insertAtTail(num int) {
	n := node{value: num}
	current := linklist.head
	for current.next != nil {
		current = current.next
	}
	current.next = &n
	fmt.Println("Inserted Node")
}

func (linklist *linkedList) display() {
	current := linklist.head.next
	for current != nil {
		fmt.Print(current.value)
		current = current.next
	}
	fmt.Println()
}
