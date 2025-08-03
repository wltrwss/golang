package utiles

import (
	"project/calculator"
)

func UserAction(act rune, a, b int) (int, bool) {
	switch act {
	case '+':
		return calculator.Plus(a, b), false
	case '-':
		return calculator.Minus(a, b), false
	case '*':
		return calculator.Multiply(a, b), false
	case '/':
		return calculator.Division(a, b), false
	default:
		return 0, true // ошибка: неизвестное действие
	}
}
