package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	for (num1 % 2 == 0 && num2 % 2 != 0) || (num1 % 2 != 0 && num2 % 2 == 0) {
		fmt.Println((num1 * num2) % (num1 + num2))
		fmt.Scanln(&num1, &num2)
	}

	fmt.Println("Last pair is odd?", num1 % 2 != 0 && num2 % 2 != 0)
}