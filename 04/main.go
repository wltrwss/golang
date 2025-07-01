//для внесенния ищменений на GitHub
//git status - показывает произведенные изменения
//git -A - добавление изменений во все ищззмененные файлы
//git commit -m "описание изменений краткое"
//git push - отправка изменений на GitHub

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [10]int
	var sum int = 0
	var minValue int = 0

	fmt.Print("\n\t\t=СОДЕРЖИМОЕ МАССИВА ARR=\t\t\n\n")

	for index := range arr {
		arr[index] = rand.Intn(100)
		sum += arr[index]
		fmt.Printf("Элемент массива  №%d - значение: %d\n", index+1, arr[index])
	}
	fmt.Printf("\nСумма всех элементов массива: %d\n\n", sum)

	/*----------Наименьшее число массива----------*/
	minValue = arr[0]
	for i := 1; i < len(arr); i++ {
		if minValue < arr[i] {
			minValue = arr[i]
		}
		if i == len(arr)-1 {
			fmt.Printf("Наименьшее число массива:%d\n\n", minValue)
		}
	}
}
