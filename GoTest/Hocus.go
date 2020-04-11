package main

import "fmt"

func main() {
	var nTest, idx int
	var O, C, B int = 0, 0, 0
	var temp rune
	fmt.Scanln(&nTest)

	for idx = 1; idx <= nTest; idx++ {
		fmt.Scanf("%c\n", &temp)
		if temp == 'O' {
			O++
		} else if temp == 'C' {
			C++
		} else if temp == 'B' {
			B++
		}
	}

	fmt.Println("Letter O =", O)
	fmt.Println("Letter B =", B)
	fmt.Println("Letter C =", C)
	fmt.Print("Experiment result :")
	if (O == B) && (B == C) {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}
}
