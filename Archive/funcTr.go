package main

import (
	"fmt"
)

func main() {
	var firstValue, secondValue, result float64
	var user_Action, userExpression string

	fmt.Println("Enter your expression:")
	fmt.Scan(&userExpression)

	switch user_Action {
	case "+":
		result = adder(firstValue, secondValue)
		fmt.Println("Result is:", result)

	case "-":
		result = subtractor(firstValue, secondValue)
		fmt.Println("Result is:", result)

	case "*":
		result = multiplier(firstValue, secondValue)
		fmt.Println("Result is:", result)

	case "/":
		result = divider(firstValue, secondValue)
		fmt.Println("Result is:", result)

	default:
		fmt.Println("Ебаный ты выблядок! Ты сделал что-то не так и моя программа пошла по пизде!")
	}
	return
}

func adder(x, y float64) float64 {
	return x + y
}

func subtractor(x, y float64) float64 {
	return x - y
}

func multiplier(x, y float64) float64 {
	return x * y
}

func divider(x, y float64) float64 {
	return x / y
}
