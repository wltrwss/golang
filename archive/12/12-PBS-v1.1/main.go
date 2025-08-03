package main

import (
	"fmt"
	"time"
)

type Account struct {
	Owner   string
	Balance float64
}

func main() {
	var funcMarker rune = ' '
	user := &Account{Owner: "Eblan", Balance: 1000000.0}
	welcomeCLI()
	userChoise := CLI(&funcMarker)
	DefOperation(&userChoise, *user, &funcMarker)
}

func welcomeCLI() {
	clearTermial()
	fmt.Println("\n\t=== Банковская система 0.1 ===")
}

func requestCLI(funcMarker *rune, user Account) {
	var userChoise string
	fmt.Println("\n Выберите следующее действие:")
	fmt.Println("  1. Повторить действие;")
	fmt.Println("  2. Главная;")
	fmt.Println("  3. Покинуть программу")
	fmt.Println(" Выберите операцию:")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			CLI(funcMarker)
		}
	}()
	_, err := fmt.Scan(&userChoise)
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
		panic("НЕКОРРЕКТНЫЙ ВВОД!")
	}
	switch userChoise {
	case "1":
		switch *funcMarker {
		case 'D':
			clearTermial()
			user.Deposit(funcMarker)
		case 'P':
			clearTermial()
			user.PrintAccount(funcMarker)
		case 'W':
			clearTermial()
			user.Withdraw(funcMarker)
		}
	case "2":
		clearTermial()
		DefOperation(&userChoise, user, funcMarker)
	case "3":
		clearTermial()
		fmt.Print("Выход из программы...")
	default:
		fmt.Print("Вы ввели неверное начение!")
		clearTermial()
		CLI(funcMarker)
	}
}

func CLI(funcMarker *rune) string {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			CLI(funcMarker)
		}
	}()
	var userChoise string
	fmt.Println(" Выберите действие, которое хотите совершить:")
	fmt.Println("  1. Пополнить баланс;")
	fmt.Println("  2. Снять денежные средства;")
	fmt.Println("  3. Отобразить состояние банковского счёта.")
	fmt.Println(" Выберите операцию:")
	_, err := fmt.Scan(&userChoise)
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
		panic("НЕКОРРЕКТНЫЙ ВВОД! ФУНКЦИЯ-CLI")
	}
	return userChoise
}

func DefOperation(userChoise *string, user Account, funcMarker *rune) {
	if *funcMarker != ' ' {
		CLI(funcMarker)
	}
	switch *userChoise {
	case "1":
		clearTermial()
		fmt.Print(" Главная / Пополнение баланса\n")
		user.Deposit(funcMarker)
	case "2":
		clearTermial()
		fmt.Print(" Главная / Снятие денежных средств\n")
		user.Withdraw(funcMarker)
	case "3":
		clearTermial()
		fmt.Print(" Главная / Состояние лицевого счета на момент: ", time.Now())
		user.PrintAccount(funcMarker)
	default:
		fmt.Print("Вы ввели неверное начение!")
		clearTermial()
		CLI(funcMarker)
	}
}

func (user *Account) Deposit(funcMarker *rune) float64 {
	*funcMarker = 'D'
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			user.Deposit(funcMarker)
		}
	}()
	fmt.Print(" Сумма к пополнению: ")
	var userSum float64
	_, err := fmt.Scan(&userSum)
	if err != nil {
		panic("НЕКОРРЕКТНЫЙ ВВОД!")
	}
	if userSum < 0 {
		userSum *= -1
	}
	fmt.Print(" Указанная сумма успешно внесена на Ваш счёт!")
	user.Balance += userSum
	user.PrintAccount(funcMarker)
	return user.Balance
}

func (user *Account) Withdraw(funcMarker *rune) float64 {
	*funcMarker = 'W'
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			user.Withdraw(funcMarker)
		}
	}()
	fmt.Println("  Сумма к списанию:")
	var userSum float64
	_, err := fmt.Scan(&userSum)
	if (err != nil) || (userSum > user.Balance) {
		panic("НЕКОРРЕКТНЫЙ ВВОД СУММЫ! Сумма введенная Вами привышает сумму депозита. Пожалуйста, попробуйте снова.")
	}
	fmt.Print(" Указанная сумма успешно списана с Вашего счёта!")
	user.Balance -= userSum
	user.PrintAccount(funcMarker)
	return user.Balance
}

func (user *Account) PrintAccount(funcMarker *rune) {
	if *funcMarker == ' ' {
		*funcMarker = 'P'
	}
	fmt.Printf("\n\n  Владелец счета: %20s", user.Owner)
	fmt.Printf("\n  Баланс:         %20.2f\n\n", user.Balance)
	requestCLI(funcMarker, *user)
}

func clearTermial() {
	fmt.Print("\033[H\033[2J")
}
