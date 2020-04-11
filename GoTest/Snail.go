package main

import "fmt"

func main() {
	var height int
	var night int = 0
	var time string

	fmt.Scanln(&height , &time)
	if time == "nighttime" {
		height++
	}

	if height % 3 != 0 {
		night = height / 3 + 1
	} else {
		night = height / 3
	}
	fmt.Println("Time to reach the top : ", night, "night")
}