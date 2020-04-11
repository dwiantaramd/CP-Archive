package main

import "fmt"

func add(op1, op2 float64, result *float64) {
	*result = op1 + op2
}

func subtract(op1, op2 float64, result *float64) {
	*result = op1 - op2
}

func multiplication(op1, op2 float64) float64 {
	return op1*op2
}

func division(op1, op2 float64) float64 {
	return op1 / op2
}

func idiv(op1, op2 float64) int {
	return int(op1) / int(op2)
}

func imod(op1, op2 float64) float64 {
	return float64(int(op1) % int(op2))
}

func main() {
	var num1, num2, result float64
	var str string = ""
	fmt.Print("Enter Command : ")
	fmt.Scanln(&str, &num1, &num2)
	for str != "quit" {
		switch str {
		case "add" :
			add(num1, num2, &result)
			fmt.Println(num1, " + ", num2, " = ",result)
		case "sub" :
			subtract(num1, num2, &result)
			fmt.Println(num1, " - ", num2, " = ",result)
		case "mul" :
			fmt.Println(num1, " x ", num2, " = ",multiplication(num1, num2))
		case "div" :
			fmt.Print(num1, " : ", num2, " = ")
			fmt.Printf("%0.2f", division(num1, num2))
			fmt.Println("")
		case "idiv" :
			fmt.Println(num1, " div ", num2, " = ",idiv(num1, num2))
		case "imod" :
			fmt.Println(num1, " mod ", num2, " = ",imod(num1, num2))
		default :
			fmt.Println("Command ", str, " is unknown")
		}

		fmt.Print("Enter Command : ")
		fmt.Scanln(&str, &num1, &num2)

	}
	fmt.Println("Calculator Ends")

}
