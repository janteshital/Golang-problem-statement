package main

import (
	"fmt"
)

type queue []string

func Queue() {
	fmt.Println("Hello, playground")
	var newq queue
	newq.enqueue("10")
	newq.enqueue("20")
	newq.enqueue("30")
	newq.printStack()
	newq.dequeue()
	newq.printStack()
	
}

func (q *queue)enqueue(n string){
	*q = append(*q, n)
}

func (q *queue)printStack(){
	fmt.Println(*q)
}

func (q *queue)dequeue(){
	if len(*q) == 0 {
		fmt.Println("no element")
	}else {
		*q = (*q)[1:] 
	}
	
}
