package utiles

import "fmt"

func GetUserNums() (int, int) {
	var a, b int
	fmt.Println("Введите первое значение:")
	fmt.Scanln(&a)
	fmt.Println("Введите второе значение:")
	fmt.Scanln(&b)
	return a, b
}
