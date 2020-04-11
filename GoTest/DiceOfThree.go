package main

import "fmt"

func main() {
	var dice1, dice2, dice3, sum int
	fmt.Scanln(&dice1, &dice2, &dice3)
	sum = dice1 + dice2 + dice3

	fmt.Println("Correct dice values?", (dice1 >= 1 && dice1 <=6) && (dice2 >= 1 && dice2 <=6) && (dice3 >= 1 && dice3 <=6))
	fmt.Println("The sum is", sum)
	fmt.Println("Big?", sum > 10)
	fmt.Println("Even?", sum % 2 == 0)
	fmt.Println("Is it a nine?", sum == 9)
}