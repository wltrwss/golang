/*
Создать карту (map), в которой:
  - ключ — это имя человека (строка),
  - значение — структура, содержащая:
  - возраст (int),
  - вес (int).

Заполнить карту минимум тремя записями (например: Вася, Петя, Маша).

После заполнения:
 1. Вывести информацию о каждом человеке: имя, возраст, вес.
 2. Найти и вывести имя самого младшего человека.
 3. Найти и вывести имя самого тяжёлого человека.
*/
package main

import "fmt"

func main() {
	var mapStudent map[string]int = make(map[string]int)

	mapStudent["A"] = 3415
	mapStudent["B"] = 467
	mapStudent["C"] = 204
	mapStudent["D"] = 78

	fmt.Println("====== ПРОИНИЦИАЛИЗИРОВАННАЯ КАРТА ======")

	for key := range mapStudent {
		fmt.Println(mapStudent[key])
	}

	fmt.Println("====== УДАЛЕНИЕ КЛЮЧА \"C\" ======")

	delete(mapStudent, "C")

	for key := range mapStudent {
		fmt.Println(mapStudent[key])
	}

	fmt.Println("====== ПРОВЕРКА НАЛИЧИЯ КЛЮЧА \"C\" ======")

	value, exists := mapStudent["C"]
	if exists {
		fmt.Println("Элемент найден", value)
	} else {
		fmt.Println("Элемент не найден", value)
	}

	fmt.Println("====== ДОБАВЛЕНИЕ НОВОГО ЭЛЕМЕНТА ======")

	mapStudent["New element"] = 43

	for key := range mapStudent {
		fmt.Println(mapStudent[key])
	}

}
