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
	fmt.Println("Insert Node at last: 10")
	linklist.insertAtTail(10)
	fmt.Println("Insert Node at last: 20")
	linklist.insertAtTail(20)
	fmt.Println("Insert Node at front: 5")
	linklist.insertAtFront(5)
	fmt.Println("Insert Node at location 3")
	linklist.insertAtFront(50)
	linklist.insertAtGivenLocation(100, 3)
	linklist.search(50)
	linklist.display()

}

func (l *linkedList) insertAtTail(num int) {
	n := node{value: num}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &n
}

func (l *linkedList) insertAtFront(num int) {
	n := node{value: num}
	tmp := l.head.next
	l.head.next = &n
	l.head.next.next = tmp
}

func (l *linkedList) display() {
	current := l.head.next
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
}

func (l *linkedList) insertAtGivenLocation(num, location int) {
	n := node{value: num}
	current := l.head
	for i := 1; i < location; i++ {
		current = current.next
	}
	tmp := current.next
	current.next = &n
	current.next.next = tmp
}

func (l *linkedList) search(num int) {
	current := l.head.next
	for current != nil {
		if current.value == num {
			fmt.Println("found number")
			break
		}
		current = current.next
	}
}
