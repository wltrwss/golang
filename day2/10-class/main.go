// ------------- Задание -----------//
// Создать CLI-приложение которое может создавать базу данных студентов
// с заполнением их ФИО и группы.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//------ создаем пустую карту данных которую будем заполнять ------//
	var mpFCsGroup map[string]string = make(map[string]string)
	var FCs string
	var Group string
	mpFCsGroup[FCs] = Group

	welcomeCLI()
	userInteraction(mpFCsGroup, &FCs, &Group)

}

// ------CLI-интерфейс------//
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

func userInteraction(mpFCsGroup map[string]string, FCs *string, Group *string) {

	var userChoice string
	fmt.Print("   Ваш выбор:")
	fmt.Scanln(&userChoice)

	switch userChoice {
	case "1":
		fmt.Print("   Вы выбрали добавить запись. ")
		newEntry(mpFCsGroup, FCs, Group)
	case "2":
		fmt.Print("   Вы выбрали удалить запись. ")
		//deleteEntry(mpFCsGroup, FCs, Group)
	case "3":
		fmt.Print("   Вы выбрали изменить запись. ")
		//changeEntry(mpFCsGroup, FCs, Group)
	case "4":
		dbDemonstration(mpFCsGroup)
	case "5":
		fmt.Print("   Выход из программы...")
	default:
		fmt.Println("   ОШИБКА ВВОДА: Вы ввели неверное значение, попробуйте ещё раз.")
		welcomeCLI()
		userInteraction(mpFCsGroup, FCs, Group)
	}
}

func newEntry(mpFCsGroup map[string]string, FCs *string, Group *string) {
	err := loadFromFile(mpFCsGroup)
	if err != nil {
		fmt.Print("\nФАМИЛИЯ:")
		fmt.Scanln(FCs)
		fmt.Println("")
		fmt.Print("\nГРУППА:")
		fmt.Scanln(Group)
		mpFCsGroup[*FCs] = *Group

		saveToFile(mpFCsGroup)
		welcomeCLI()
		userInteraction(mpFCsGroup, FCs, Group)
	} else {
		fmt.Println("ФАМИЛИЯ:")
		fmt.Scanln(FCs)
		fmt.Println("ГРУППА:")
		fmt.Scanln(Group)
		mpFCsGroup[*FCs] = *Group

		saveToFile(mpFCsGroup)
		welcomeCLI()
		userInteraction(mpFCsGroup, FCs, Group)
	}
}

//func deleteEntry() {

//}

//func changeEntry() {

//}

func loadFromFile(mpFCsGroup map[string]string) error {

	file, err := os.Open("dbStudent.json")
	if err != nil {
		fmt.Println("ОШИБКА:", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&mpFCsGroup)
	if err != nil {
		fmt.Println("ОШИБКА:", err)
		return err
	}

	fmt.Println("Режим БД - Штатный")
	return nil
}

func saveToFile(mpFCsGroup map[string]string) error {

	file, err := os.Create("dbStudent.json")
	if err != nil {
		fmt.Println("ОШИБКА:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&mpFCsGroup)

	fmt.Println("ИМЕНЕНИЯ ЗАПИСАНЫ В БАЗУ ДАННЫХ")
	return nil
}

func dbDemonstration(mpFCsGroup map[string]string) {
	loadFromFile(mpFCsGroup)
	for value := range mpFCsGroup {
		fmt.Println(mpFCsGroup["key"], mpFCsGroup[value])
	}

}
