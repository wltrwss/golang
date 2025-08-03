package main

import (
	"fmt"
	"project/utiles"
)

func main() {
	a, b := utiles.GetUserNums()
	var act rune
	var result int
	var err bool

	for {
		act = utiles.GetUserAct()
		result, err = utiles.UserAction(act, a, b)
		if !err {
			break
		}
		fmt.Println("Ошибка ввода действия. Попробуйте снова.")
	}

	fmt.Println("Результат:", result)
}
