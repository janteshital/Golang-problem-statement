
package main

import "fmt"

func main() {
	example := []int{20, 5, 1, 9, 3}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {

			if example[i] > example[j] {

				example[i], example[j] = example[j], example[i]
			}

		}
	}
	fmt.Println(example)
	calculateMaxSum(example)
	calculateMinSum(example)
}

func calculateMaxSum(e []int) {
	sum := 0
	for i := 1; i < len(e); i++ {
		sum += e[i]
	}
	fmt.Printf("Max Sum of any 4 numbers is: %v\n", sum)
}

func calculateMinSum(e []int) {
	sum := 0
	for i := 0; i < len(e)-1; i++ {
		sum += e[i]
	}
	fmt.Printf("Min Sum of any 4 numbers is: %v\n", sum)
}
