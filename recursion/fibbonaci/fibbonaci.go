package main

import "fmt"

func main() {
	var in int
	for {
		fmt.Println("Enter a number")
		_, err := fmt.Scan(&in)
		if err != nil {
			fmt.Println("Error reading int from stdin")
			continue
		}

		if in == 0 {
			fmt.Println("Invalid input. Input shd be > 0")
		}

		fibs := fib(in)
		fmt.Println("Fibs --> ", fibs)
	}
}

func fib(i int) []int {
	if i == 1 {
		return []int{0}
	}
	if i == 2 {
		return []int{0, 1}
	}

	fibs := make([]int, i)
	fibs[0] = 0
	fibs[1] = 1

	for j := 2; j <= i-1; j++ {
		fibs[j] = fibs[j-2] + fibs[j-1]
	}

	return fibs
}
