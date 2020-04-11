package main

import "fmt"

func main() {
	var weight int
	var kg, gram, cost, Addcost int

	fmt.Print("Parcel Weight (in grams) : ")
	fmt.Scanln(&weight)

	kg = weight / 1000
	gram = weight % 1000

	fmt.Println("Weight Breakdown :", kg, "kgs and", gram, "grams") 

	cost = kg * 10000
	if gram < 500 {
		Addcost = 15 * gram
	} else {
		Addcost = 5 * gram
	}
	fmt.Println("Cost Breakdown : Rp.", cost, "+ Rp.", Addcost)
	if weight > 10000 {
		fmt.Println("Total cost :", cost)
	} else {
		fmt.Println("Total cost :", cost + Addcost)
	}
	

}