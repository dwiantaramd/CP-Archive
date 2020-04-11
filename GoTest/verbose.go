package main

import "fmt"

const MaxWords = 97
type WordList [MaxWords]string

func readWords(w *WordList, n int) {
	var idx int
	fmt.Println("Enter the words, seperated by spaces :")
	for idx = 0; idx <= n-1; idx++ {
		fmt.Scanln(&w[idx])
	}
}

func printWords(w WordList, n int, f string) {
	var idx int
	fmt.Print("Not matched words : ")
	for idx = 0; idx <= n-1; idx++{
		if w[idx] != f {
			fmt.Print(w[idx], " ")
		}
 	}
}

func main() {
	var word WordList
	var N int
	var temp string

	fmt.Print("How many words? ")
	fmt.Scanln(&N)
	readWords(&word, N)

	fmt.Print("Enter the word to search : ")
	fmt.Scanln(&temp)
	printWords(word, N, temp)

}