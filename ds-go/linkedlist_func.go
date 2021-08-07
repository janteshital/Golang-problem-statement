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
	fmt.Println("Insert Node at last: 2")
	linklist.insertAtTail(2)
	fmt.Println("Insert Node at last: 200")
	linklist.insertAtTail(200)
	fmt.Println("Insert Node at front: 15")
	linklist.insertAtFront(15)
	fmt.Println("Insert Node at front: 100")
	linklist.insertAtFront(100)
	fmt.Println("Insert Node at location 3")
	linklist.insertAtGivenLocation(100, 3)
	// linklist.search(50)
	linklist.display()
	// linklist.sortList()
	// linklist.display()
	// linklist.removeDuplicate()
	linklist.removeFromList(200)
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
	fmt.Println()
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

func (l *linkedList) length() int{
	current := l.head.next
	tmp :=0
	for current != nil {
		tmp ++
		current = current.next
	}
	return tmp
}

func (l *linkedList) sortList() {
	for i:=0;i<(l.length()-1);i++ {
		current := l.head.next
		for current.next != nil{
			if current.value > current.next.value {
				tmp := current.value
				current.value = current.next.value
				current.next.value = tmp
			}
			current = current.next
		}
	}
}

func (l *linkedList) removeFromList(num int) {
	current := l.head.next
	flag := false
	if current.value == num{
		l.head.next = current.next
	}else{
		current = current.next
		// for current.next != nil{
		// 	if current.value == num{
		// 		current.value = current.next.value
		// 		current.next = current.next.next
		// 		flag = true
		// 		fmt.Print(("check"))
		// 		break
		// 	}
		// 	current = current.next
		// }
		if flag != true{
			// current = l.head.next
			for current.next.next != nil{
				fmt.Println("value", current)
				if current.next.value == num && current.next.next == nil {
					current.next = nil
					break
				}else if current.next.value == num && current.next.next != nil {
					current.next = current.next.next
					break
				}
				fmt.Println("flow")
				current = current.next
			}
		}
	}

}

func (l *linkedList) removeDuplicate() {
	current := l.head.next
	for current.next != nil{
		nextnode := current.next
		// fmt.Println("current", current.value, "next", nextnode.value)
		for i:=1;i<(l.length()-i);i++{
			fmt.Println("current", current.value, "next", nextnode.value)
			if current.value == nextnode.value{
				fmt.Println("list have duplicate number")
				current.next = current.next.next
				break
			}
			if nextnode.next != nil{
				nextnode = nextnode.next
			}
		}
		current = current.next	
	}
}