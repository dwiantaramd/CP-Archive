package main

import "fmt"

func main() {
	var number, idx, idx2 int
	var binary [64]int
	
	fmt.Scanln(&number)
	idx = 0
	for number > 0 {
		binary[idx] = number % 2
		number /= 2
		idx++
	}

	for idx2 = idx - 1; idx2 >= 0; idx2-- {
		fmt.Print(binary[idx2])
	}
}