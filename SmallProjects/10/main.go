package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	mpFCsGroup, err := loadFromFile()
	if err != nil {
		mpFCsGroup = make(map[string]string)
	}

	for {
		welcomeCLI()
		if !userInteraction(mpFCsGroup) {
			break
		}
	}
}

func welcomeCLI() {
	fmt.Println("\n|-----------------------------------------------------|")
	fmt.Println(" ==== ДОБРО ПОЖАЛОВАТЬ В СТУДЕНЧЕСКУЮ БАЗУ ДАННЫХ ====")
	fmt.Println("|-----------------------------------------------------|")
	fmt.Println("\tВЫБЕРИТЕ ДЕЙСТВИЕ КОТОРОЕ ХОТИТЕ ВЫПОЛНИТЬ:")
	fmt.Println("|-----------------------------------------------------|")
	fmt.Println("  1. Добавить запись;")
	fmt.Println("  2. Удалить запись;")
	fmt.Println("  3. Изменить запись;")
	fmt.Println("  4. Визуализировать базу данных;")
	fmt.Println("  5. Выход из программы.")
	fmt.Println("|-----------------------------------------------------|")
}

func userInteraction(mpFCsGroup map[string]string) bool {
	var userChoice string
	fmt.Print("   Ваш выбор: ")
	fmt.Scanln(&userChoice)

	switch userChoice {
	case "1":
		newEntry(mpFCsGroup)
	case "2":
		deleteEntry(mpFCsGroup)
	case "3":
		changeEntry(mpFCsGroup)
	case "4":
		dbDemonstration(mpFCsGroup)
	case "5":
		fmt.Println("   Выход из программы...")
		return false
	default:
		fmt.Println("   ОШИБКА ВВОДА: Вы ввели неверное значение, попробуйте ещё раз.")
	}
	return true
}

func newEntry(mpFCsGroup map[string]string) {
	var FCs, Group string
	fmt.Print("ФАМИЛИЯ: ")
	fmt.Scanln(&FCs)
	fmt.Print("ГРУППА: ")
	fmt.Scanln(&Group)

	if FCs == "" || Group == "" {
		fmt.Println("   ОШИБКА: Фамилия и группа не могут быть пустыми.")
		return
	}

	mpFCsGroup[FCs] = Group
	saveToFile(mpFCsGroup)
}

func deleteEntry(mpFCsGroup map[string]string) {
	var deleteKey string
	dbDemonstration(mpFCsGroup)
	fmt.Print("Введите фамилию студента, которого хотите удалить: ")
	fmt.Scanln(&deleteKey)

	if _, exists := mpFCsGroup[deleteKey]; exists {
		delete(mpFCsGroup, deleteKey)
		saveToFile(mpFCsGroup)
		fmt.Println("Запись удалена.")
	} else {
		fmt.Println("   Студент не найден.")
	}
}

func changeEntry(mpFCsGroup map[string]string) {
	var key string
	dbDemonstration(mpFCsGroup)
	fmt.Print("Введите фамилию студента, запись о котором хотите изменить: ")
	fmt.Scanln(&key)

	if _, ok := mpFCsGroup[key]; !ok {
		fmt.Println("   Студент не найден.")
		return
	}

	fmt.Println("Выберите, что хотите изменить:")
	fmt.Println("  1. Фамилию;")
	fmt.Println("  2. Группу.")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		var newKey string
		fmt.Print("Новая фамилия: ")
		fmt.Scanln(&newKey)
		mpFCsGroup[newKey] = mpFCsGroup[key]
		delete(mpFCsGroup, key)
		saveToFile(mpFCsGroup)
	case "2":
		var newGroup string
		fmt.Print("Новая группа: ")
		fmt.Scanln(&newGroup)
		mpFCsGroup[key] = newGroup
		saveToFile(mpFCsGroup)
	default:
		fmt.Println("Неверный выбор.")
	}
}

func dbDemonstration(mpFCsGroup map[string]string) {
	if len(mpFCsGroup) == 0 {
		fmt.Println("База данных пуста.")
		return
	}

	fmt.Printf("\n%-20s | %-10s\n", "Фамилия", "Группа")
	fmt.Println(strings.Repeat("-", 33))

	for k, v := range mpFCsGroup {
		fmt.Printf("%-20s | %-10s\n", k, v)
	}
	fmt.Println()
}

func loadFromFile() (map[string]string, error) {
	file, err := os.Open("dbStudent.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string]string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	fmt.Println("Режим БД — Штатный")
	return data, nil
}

func saveToFile(mpFCsGroup map[string]string) {
	file, err := os.Create("dbStudent.json")
	if err != nil {
		fmt.Println("ОШИБКА СОХРАНЕНИЯ:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(mpFCsGroup)
	if err != nil {
		fmt.Println("ОШИБКА ЗАПИСИ:", err)
		return
	}

	fmt.Println("Изменения сохранены в базу данных.")
}
