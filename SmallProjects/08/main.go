package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var userArray [10]int

	for index := range userArray {
		userArray[index] = rand.Intn(100)
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, userArray[index])
	}

	var minValue int = userArray[0]
	var maxValue int = userArray[0]

	for i := 1; i < len(userArray); i++ {
		if minValue > userArray[i] {
			minValue = userArray[i]
		}
		if maxValue < userArray[i] {
			maxValue = userArray[i]
		}
	}
	fmt.Println("Минимальное число массива:", minValue)
	fmt.Println("Максимальное число массива:", maxValue)
}
