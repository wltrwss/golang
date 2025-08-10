package main

import "fmt"

type userInt float64
type userIntP = string


func main() {
	//=========ПЕРЕМЕННЫЕ=========//
	var var1 int = 10
	var2 := 10

	var var3 float64 = 10.5
	var4 := userInt(10.5)

	var var5 string = "EXAMPLE"
	var6 := userIntP("EXAMPLE")

	const var7 bool = true
	const var8 = true

	//=========МАССИВ=========//

	var arr1 [5]int                                // полное объявление массива без инициализации (все значения = 0)
	var arr2 [5]rune = [5]rune{1, 2, 3, 4, 5}      // полное объявление и полная инициализация массива
	arr3 := [5]int{1, 3, 5, 7, 3}                  // сокращённое объявление и полная инициализация
	arr4 := [100]int{1, 2, 5, 7, 8}                // частичная инициализация: остальные 95 элементов = 0
	arr5 := [5]int{0: 1, 1: 43, 2: 5, 3: 6, 4: 76} // инициализация с явным указанием индексов
	arr6 := [...]int{1, 2, 34, 5, 7, 98, 564}      // инициализация с автоматическим определением длины

	//=========СЛАЙС==========//

	var slc1 []rune
	var slc2 []rune = []rune{'r', '0', 'd', 'k', 'h', 'f'}
	slc3 := []rune{'q', 'w', 'e', 'r', 't', 'y'}
	slc4 := []rune{0: 'q', 1: 'w', 2: '4'}

	slc5 := arr2[1:3]
	slc6 := make([]rune, 5)
	slc6_1 := make([]rune, 5, 10)

	slc7 := []int{1, 2, 3, 5, 6}
	slc7 = append(slc7, 7, 4, 6, 7, 8, 9, 0, 5, 3)

	slc8 := []int{3, 3, 3, 3, 3}
	slc9 := []int{1, 1, 1, 1, 1}

	copy(slc8, slc9)

	slc11 := []int{34, 45, 7, 67, 8, 34, 53, 4, 7456, 7, 45, 8, 97, 35, 56, 3, 23, 4}
	slc12 := slc11[5:10]


	allData := []interface{}{var1,var2,var3,var4,var5,var6,var7,var8, arr1,arr2,arr3,arr4,arr5,arr6,slc1,slc2,slc3,slc4,slc5,slc6,slc7,slc8,slc9,slc6_1,slc11,slc12}
	for i, _ := range allData {
		fmt.Println(allData[i])
	}

}
