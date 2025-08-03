package utiles

import "fmt"

func GetUserAct() rune {
	var act rune
	fmt.Println("Введите действие, которое хотите выполнить:")
	fmt.Scanf("%c\n", &act)
	return act
}
