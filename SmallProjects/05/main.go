package main

import (
	"fmt"
	"math/rand"
)

func main() {
	systemNumOne, systemNumTwo := rand.Intn(100), rand.Intn(100)
	var arrResult [4]int

	fmt.Println("Числа из памяти:\t", systemNumOne, ",", systemNumTwo)
	if (systemNumOne == 0) || (systemNumTwo == 0) {
		fmt.Println("Одно из значений равно нулю, деление на ноль невозможно, поэтому программа будет выполненна только в части сложения, вычитания и умножения!")
		arrResult[0] = systemNumOne + systemNumTwo
		arrResult[1] = systemNumOne - systemNumTwo
		arrResult[2] = systemNumOne * systemNumTwo
	} else {
		arrResult[0] = systemNumOne + systemNumTwo
		arrResult[1] = systemNumOne - systemNumTwo
		arrResult[2] = systemNumOne * systemNumTwo
		arrResult[3] = systemNumOne / systemNumTwo
	}
	for i := 0; i < len(arrResult); i++ {
		fmt.Println(arrResult[i])
	}
}
