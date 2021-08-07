package main

import (
	"fmt"
)

func main() {	
	message := make(chan int)
	for i:=0;i<100;i++{
		go printOdd(i, message)
		go printEven(i, message)
		result := <- message
		fmt.Println(result)
	}
	close(message)
	
}

func printOdd(n int, c chan int){
	if n%2 != 0{
		c <- n
	}
	
}

func printEven(n int, c chan int){
	if n%2 == 0{
		c <- n
	}
}
