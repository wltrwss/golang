package main

import "fmt"

//Напишите программу, которая
//находит самый наименьший элемент в этом
//списке:

func main() {
	slc := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	comparison(slc)
}

/* ТЕСТ НА ПРОВЕРКУ ПЕРЕДАЧИ В ФУНКЦИЮ ДЛИННЫ МАССИВА
func lenfn(slc_f []int) int {
	lenvar := len(slc_f)
	fmt.Println(lenvar)
	return lenvar
}
*/
func comparison(slc_f []int) int {
	var num_1, num_2, num_3, num_4 int = 0, 0, 0, 0

	for i, counter := 0, 1; i < len(slc_f); i++ {
		num_1 = slc_f[i]
		i++
		num_2 = slc_f[i]
		fmt.Println("Итерация №", counter, ".", "Сравниваются числа:", num_1, "и", num_2)
		if num_1 < num_2 {
			num_3 = num_1
		} else if num_2 < num_1 {
			num_3 = num_2
		}

		if i < 2 {
			num_4 = num_3
		} else if num_3 < num_4 {
			num_4 = num_3
		}
		fmt.Println("Наименьшее число: ", num_4)
		counter++
	}
	fmt.Println(num_4)
	return num_4
}
