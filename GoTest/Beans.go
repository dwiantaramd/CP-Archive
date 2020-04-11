package main

import "fmt"

const NMax = 7
type seeds [NMax]int

func inputData(d *seeds, n int) {
	var idx int

	fmt.Println("Observation on day 3")

	for idx = 0	; idx <= n-1; idx++ {
		fmt.Print("Height seed ", idx+1, " : ")
		fmt.Scanln(&d[idx])
	}
}

func outputData(d seeds, n int) {
	var idx int

	for idx = 0; idx <= n-1; idx++ {
		fmt.Println("Height seed", idx+1, ":", d[idx])
	}
}

func isOdd(h int) bool {
	return h % 2 != 0
}

func isDivisible3(h int) bool {
	return h % 3 == 0
}

func updateData(d *seeds, n int) {
	var idx int

	for idx = 0; idx <= n-1; idx++ {
		if d[idx] != 0 {
			if isDivisible3(d[idx]) {
				d[idx] += 2
			} else if isOdd(d[idx]) {
				d[idx] += 3
			} else {
				d[idx] += 1
			}
		}
	}
}

func main() {
	var N, day, idx int
	var height seeds


	fmt.Print("How many seeds? ")
	fmt.Scanln(&N)
	inputData(&height, N)

	fmt.Print("Prediction days : ")
	fmt.Scanln(&day)

	for idx = 1; idx <= day; idx++ {
		updateData(&height, N)
		fmt.Println("Day ", idx+3)
		outputData(height, N)
	}

}