package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [10]int
	var sum int = 0
	var numOne, numTwo int = 0, 0

	fmt.Println("\n\t\t=СОДЕРЖИМОЕ МАССИВА ARR=\t\t\n")

	for index, value := range arr {
		value = rand.Intn(100)
		arr[index] = value
		sum += arr[index]
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, arr[index])
	}
	fmt.Printf("\nСумма всех элементов массива: %d\n\n", sum)

	/*----------Наименьшее число массива----------*/
	for i := 0; i < len(arr); i++ {
		numOne = arr[i]
		i++
		numTwo = arr[i]

		if numOne < numTwo {
			if numOne < sum {
				sum = numOne
			}
		} else if numOne > numTwo {
			if numTwo < sum {
				sum = numTwo
			}
		}
		if i == len(arr)-1 {
			fmt.Printf("Наименьшее число массива:%d\n\n", sum)
		}
	}
}
