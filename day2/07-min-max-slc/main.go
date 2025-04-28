package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var mainArray [20]int
	var mainSlice = mainArray[:]

	randSlice(mainSlice)

	var minValue, maxValue = 0, 0
	minValue, maxValue = sliceMinMax(mainSlice, minValue, maxValue)

	fmt.Printf("Наименьшее число слайса: %d\n", minValue)
	fmt.Printf("Наибольшее число слайса: %d\n", maxValue)
}

func randSlice(a []int) {
	for index := range a {
		a[index] = rand.Intn(100)
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, a[index])
	}

}

func sliceMinMax(a []int, min int, max int) (int, int) {
	min = a[0]
	max = a[0]

	for i := 1; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	return min, max
}
