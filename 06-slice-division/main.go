package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var mainArray [30]int
	var mainSlice []int = mainArray[:]
	var firstSlice []int = mainSlice[:10]
	var secondSlice []int = mainSlice[10:20]
	var thirdSlice []int = mainSlice[20:30]

	fmt.Println("=========== ЗНАЧЕНИЯ ЗАПИСАННЫЕ В ГЛАВНЫЙ СЛАЙС ===========")
	sliceRand(mainSlice)

	fmt.Println("=========== ПЕРВАЯ ЧАСТЬ СЛАЙСА ===========")
	sliceDiv(firstSlice)
	sliceMinMax(firstSlice)
	fmt.Println("=========== ВТОРАЯ ЧАСТЬ СЛАЙСА ===========")
	sliceDiv(secondSlice)
	sliceMinMax(secondSlice)
	fmt.Println("=========== ТРЕТЬЯ ЧАСТЬ СЛАЙСА ===========")
	sliceDiv(thirdSlice)
	sliceMinMax(thirdSlice)

}

func sliceRand(mainSlice []int) {
	for index := range mainSlice {
		mainSlice[index] = rand.Intn(150)
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, mainSlice[index])
	}
	fmt.Println(" ")
}

func sliceDiv(slicePart []int) {
	for index := range slicePart {
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, slicePart[index])
	}
	fmt.Println(" ")
}

func sliceMinMax(slicePart []int) {

	var minValue, maxValue int = slicePart[0], slicePart[0]
	for i := 1; i < 10; i++ {
		if minValue > slicePart[i] {
			minValue = slicePart[i]
		}
		if maxValue < slicePart[i] {
			maxValue = slicePart[i]
		}
	}

	fmt.Println("Минимальное значение указаной части массива:", minValue)
	fmt.Println("Максимальное значение указаной части массива:", maxValue)
}
