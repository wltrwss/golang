package main

import (
	"fmt"
)

type Account struct {
	Owner   string
	Balance float64
}

func main() {
	account := &Account{Owner: "Eblan", Balance: 1000000.0}
	for {
		clearTerminal()
		showWelcome()
		choice := askMainMenu()

		switch choice {
		case "1":
			account.Deposit()
		case "2":
			account.Withdraw()
		case "3":
			account.PrintAccount()
		case "4":
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
			pause()
		}
	}
}

func showWelcome() {
	fmt.Println("\n=== Банковская система 0.1 ===")
}

func askMainMenu() string {
	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Пополнить баланс")
	fmt.Println("2. Снять средства")
	fmt.Println("3. Просмотреть счёт")
	fmt.Println("4. Выход")
	fmt.Print("Ваш выбор: ")

	var input string
	fmt.Scan(&input)
	return input
}

func (a *Account) Deposit() {
	clearTerminal()
	defer handlePanic(a.Deposit)

	fmt.Print("Введите сумму пополнения: ")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil || amount <= 0 {
		panic("Некорректная сумма!")
	}

	a.Balance += amount
	fmt.Println("Счёт пополнен.")
	a.PrintAccount()
	pause()
}

func (a *Account) Withdraw() {
	clearTerminal()
	defer handlePanic(a.Withdraw)

	fmt.Print("Введите сумму для снятия: ")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil || amount <= 0 || amount > a.Balance {
		panic("Некорректная сумма для снятия!")
	}

	a.Balance -= amount
	fmt.Println("Средства сняты.")
	a.PrintAccount()
	pause()
}

func (a *Account) PrintAccount() {
	fmt.Printf("\nВладелец счета: %s\n", a.Owner)
	fmt.Printf("Баланс: %.2f\n", a.Balance)
	pause()
}

func handlePanic(retryFunc func()) {
	if r := recover(); r != nil {
		fmt.Println("Ошибка:", r)
		pause()
		retryFunc()
	}
}

func pause() {
	fmt.Print("\nНажмите Enter для продолжения...")
	fmt.Scanln()
	fmt.Scanln()
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
