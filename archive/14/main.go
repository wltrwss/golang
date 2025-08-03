package main

import "fmt"

type Wallet struct {
	Balance int
}

type User struct {
	Name string
	Wallet
}

type Shop struct {
	Title string
	Price int
}

func (u *User) Spend(s *Shop) {
	u.Balance -= s.Price
	fmt.Println("Your balance after buy =", u.Balance-s.Price)

}

func (s *Shop) SellTo(u *User) {
	u.Spend(s)
}

func main() {
	u := User{
		"Max",
		Wallet{
			Balance: 1000,
		},
	}
	s := Shop{"Maybach", 450}
	fmt.Println("Your balance before buy =", u.Balance)
	s.SellTo(&u)
}
