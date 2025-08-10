package calculator

import "fmt"

func Plus(a, b int) int { // допущена ошибка
	return a - b
}

func Minus(a, b int) int { // допущена ошибка
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	if b == 0 {
		fmt.Println("Ошибка: деление на ноль")
		return 0
	}
	return a / b
}
