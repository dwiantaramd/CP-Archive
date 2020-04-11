package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var card [13]string = [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	var hand [5]string
	var check [13]bool
	var idx, num int
	var guess string
	var temp bool

	fmt.Println("Game Of Poker")
	fmt.Print("Cards : ")

	for idx = 0; idx <= 12; idx++ {
		fmt.Print(card[idx], " ")
	}

	//User Cards
	rand.Seed(time.Now().Unix())
	fmt.Print("\n\nYour Cards : ")
	idx = 0
	for idx <= 4 {
		var random = rand.Intn(12)
		if !check[random] {
			check[random] = true
			fmt.Print(card[random], " ")
			hand[idx] = card[random]
			idx++
		}
	}
	//All Check back to false
	for idx = 0; idx <= 12; idx++ {
		check[idx] = false
	}

	//Dealer Guess
	fmt.Println("\nPress enter to read dealer's guess")
	fmt.Scanln()
	guess = card[rand.Intn(12)]
	fmt.Println("\nDealer guess : ", guess)
	for idx = 0; idx <= 4; idx++ {
		if guess == hand[idx] {
			temp = true
		}
	}
	if temp {
		fmt.Println("Dealer guessed corectly")
	} else {
		fmt.Println("Dealer made a wrong guess")
	}
	temp = false

	//Dealer Cards
	fmt.Print("\nYour guess(1..13) : ")
	fmt.Scanln(&num)
	guess = card[num-1]
	
	fmt.Print("Dealer Cards : ")
	idx = 0
	for idx <= 4 {
		var random = rand.Intn(12)
		if !check[random] {
			check[random] = true
			fmt.Print(card[random], " ")
			hand[idx] = card[random]
			idx++
		}
	}

	//User Guess Guess
	fmt.Println("")
	for idx = 0; idx <= 4; idx++ {
		if guess == hand[idx] {
			temp = true
		}
	}
	if temp {
		fmt.Println("You guessed corectly")
	} else {
		fmt.Println("You made a wrong guess")
	}
}