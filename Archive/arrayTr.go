package main

import (
	"fmt"
	"math/rand"
)

// Напишите программу для поиска
// наибольшего элемента в заданном массиве целых чисел.
func main() {
	var arr [10]int
	var num_One, num_Two, num_Three, res int = 0, 0, 0, 0
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
		fmt.Println(i+1, " - ", arr[i])
	}
	/*-------------Реализация Сортировщика-------------*/
	for i := 0; i < len(arr); i++ {

		num_One = arr[i]
		i++
		num_Two = arr[i]
		/*-------------Сравниваю два числа и передаю большее в третью переменную-------------*/
		if num_One > num_Two {
			num_Three = num_One
		} else {
			num_Three = num_Two
		}
		/*-------------Сохраняю после первой итерации переменную для последующего сравнения с другими-------------*/
		/*-------------В переменной res остается только самое большое число-------------*/
		if i == 1 {
			res = num_Three
		} else if res > num_Three {
			num_Three = 0
		} else {
			res = num_Three
			num_Three = 0
		}

		if i == len(arr)-1 {
			fmt.Println("Самое большое число массива = ", res)
		}
	}

}
