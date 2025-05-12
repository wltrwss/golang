/*
Задание: Банковский аккаунт с защитой от неверных операций

Создай программу, в которой будет:

1. Структура Account
Поля:
 • Owner (string) — имя владельца;
 • Balance (float64) — баланс счета.

2. Методы
 • Deposit(amount float64) — метод пополнения счёта. Увеличивает баланс.
 • Withdraw(amount float64) — метод снятия средств.
Если на счёте недостаточно средств, вызывает panic("Недостаточно средств").

3. Обёртка вокруг Withdraw с recover
Сделай отдельную функцию, например SafeWithdraw(acc *Account, amount float64), которая вызывает acc.Withdraw(amount) и перехватывает панику с помощью recover, выводя сообщение вместо аварийного завершения.

4. В main()
 • Создай объект Account с каким-то именем и начальными деньгами.
 • Попробуй положить и снять деньги.
 • Сделай попытку снять больше, чем есть, чтобы сработала panic, и она поймалась через recover.
*/

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
	user := &Account{Owner: "Eblan", Balance: 1000000.0}
	welcomeCLI()
	userChoise := CLI()
	DefOperation(&userChoise, *user)
}

func welcomeCLI() {
	clearTermial()
	fmt.Println("\n\t=== Банковская система 0.1 ===")
}

func requestCLI(funcMarker *rune, user Account) {
	clearTermial()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			CLI()
		}
	}()
	var userChoise string
	fmt.Println(" Выберите следующее действие:")
	fmt.Println("  1. Повторить действие;")
	fmt.Println("  2. Главная;")
	fmt.Println("  3. Покинуть программу")
	fmt.Println(" Выберите операцию:")
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
			user.Deposit()
		case 'P':
			clearTermial()
			user.PrintAccount()
		case 'W':
			clearTermial()
			user.Withdraw()
		}
	case "2":
		clearTermial()
		CLI()
	case "3":
		clearTermial()
		fmt.Print("Выход из программы...")
	default:
		fmt.Print("Вы ввели неверное начение!")
		clearTermial()
		CLI()
	}
}

func CLI() string {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			CLI()
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

func DefOperation(userChoise *string, user Account) {
	switch *userChoise {
	case "1":
		clearTermial()
		fmt.Print(" Главная / Пополнение баланса\n")
		user.Deposit()
	case "2":
		clearTermial()
		fmt.Print(" Главная / Снятие денежных средств\n")
		user.Withdraw()
	case "3":
		clearTermial()
		fmt.Print(" Главная / Состояние лицевого счета на момент: ", time.Now())
		user.PrintAccount()
	default:
		fmt.Print("Вы ввели неверное начение!")
		clearTermial()
		CLI()
	}
}

func (user *Account) Deposit() float64 {
	var funcMarker rune = 'D'
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:", p)
			user.Deposit()
		}
	}()
	fmt.Print(" Сумма к пополнению: ")
	var userSum float64
	_, err := fmt.Scan(&userSum)
	if err != nil {
		panic("НЕКОРРЕКТНЫЙ ВВОД!")
	}
	fmt.Print(" Указанная сумма успешно внесена на Ваш счёт!")
	user.Balance += userSum
	user.PrintAccount()
	requestCLI(&funcMarker, *user)
	return user.Balance
}

func (user *Account) Withdraw() float64 {
	var funcMarker rune = 'W'
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Произошла ошибка:")
			user.Withdraw()
		}
	}()

	fmt.Println("  Сумма к списанию:")
	var userSum float64
	_, err := fmt.Scan(&userSum)
	if (err != nil) || (userSum < user.Balance) {
		panic("НЕКОРРЕКТНЫЙ ВВОД СУММЫ! Сумма введенная Вами привышает сумму депозита. Пожалуйста, попробуйте снова.")
	}
	fmt.Print(" Указанная сумма успешно списана с Вашего счёта!")
	user.Balance -= userSum
	user.PrintAccount()
	requestCLI(&funcMarker, *user)
	return user.Balance
}

func (user *Account) PrintAccount() {
	var funcMarker rune = 'P'
	fmt.Printf("\n\n  Владелец счета: %20s", user.Owner)
	fmt.Printf("\n  Баланс:         %20.2f\n\n", user.Balance)
	requestCLI(&funcMarker, *user)
}

func clearTermial() {
	fmt.Print("\033[H\033[2J")
}
