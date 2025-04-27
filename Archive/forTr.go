package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var arr [100]float64
	var total float64 = 0.0
	/*-------Заполняю массив случайными дробными числами-------*/
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Float64() * 100
	}
	/*-------Вывожу массив на экран подписывая номер значения и округляю значения до сотых-------*/
	for i := 0; i < len(arr); i++ {
		fmt.Print(i+1, " - ")
		fmt.Println(math.Ceil(arr[i]*100) / 100)
	}
	/*-------Складываю все числа массива между собой и помещаю в переменную total-------*/
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	/*-------Вывожу результат на экран-------*/
	fmt.Println("Average rating is:", math.Ceil((total/100)*100)/100)
}

//FizzBuzz
/*func main() {
for i := 1; i < 100; i++ {
	if (i%3 == 0) && (i%5 == 0) {
		fmt.Println(i, " - FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println(i, " - Fizz")
	} else if i%5 == 0 {
		fmt.Println(i, " - Buzz")
	} else {
		fmt.Println(i)
	}
}*/
