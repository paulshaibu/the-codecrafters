package main

import (
	"fmt"
)

func main() {

	var number1, number2 float64
	var operator string
	for {
		fmt.Println("\nWelcome to My Paul calculator")
		fmt.Print("Enter the first number: ")
		if _, err := fmt.Scan(&number1); err != nil {
			fmt.Println("Invalid Input")
			return
		}

		fmt.Print("Enter the Second Number:")
		if _, err := fmt.Scan(&number2); err != nil {
			fmt.Println("Invalid Input")
			return
		}
		fmt.Print("Enter the Operator: (+ - * / help quit) ")
		if _, err := fmt.Scan(&operator); err != nil {
			fmt.Println("Invalid operator")
			return
		}

		switch operator {
		case "help":
			fmt.Println("supported commands: add / sub / div / mul / help / quit  ")
		case "quit":
			fmt.Println("GoodBye Paul")
			return
		case "+":
			fmt.Printf("Output = %v %s %v = %v", number1, operator, number2, number1+number2)
			fmt.Println()
		case "-":
			fmt.Printf("Output = %v %s %v = %v", number1, operator, number2, number1-number2)
			fmt.Println()
		case "*":
			fmt.Printf("Output = %v %s %v = %v", number1, operator, number2, number1*number2)
			fmt.Println()
		case "/":
			if number2 == 0 {
				fmt.Print("Leave Here: cannot divide by 0")
				fmt.Println()
			} else {
				fmt.Printf("Output = %v %s %v = %0.0v", number1, operator, number2, number1/number2)
				fmt.Println("GoodBye")
			}
		default:
			fmt.Print("Invalid Input.. Leave Here")
		}

	}
}
