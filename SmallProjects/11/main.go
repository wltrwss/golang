package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Pages  int
}

func main() {
	var bookOne Book = Book{"Война и мир", "Толстой Л.Н.", 1869, 1360}
	var bookTwo Book = Book{"Мертвые души", "Гоголь Н.В.", 1842, 352}
	var bookThree Book = Book{"Преступление и наказание", "Достоевский Ф.М.", 1866, 672}
	var bookFour Book = Book{"Вишнёвый сад", "Чехов А.П.", 1903, 64}

	var bookSlice []Book = make([]Book, 4)
	bookSlice[0] = bookOne
	bookSlice[1] = bookTwo
	bookSlice[2] = bookThree
	bookSlice[3] = bookFour

	printBook(bookSlice)
}

func printBook(bookSlice []Book) {
	fmt.Printf("\n %-45s %-55s \n", " ", "= СПИСОК КНИГ =")
	var tempVariable Book = Book{"", "", 0, 0}
	var line string = "-"
	fmt.Println(strings.Repeat(line, 110))
	fmt.Printf("%-30s %3s %-20s %3s %-20s %3s %-20s \n", "Название", "|", "Автор", "|", "Год издания", "|", "Колличество страниц")
	fmt.Println(strings.Repeat(line, 110))
	for value := range bookSlice {
		tempVariable = bookSlice[value]
		fmt.Printf("%-30s %3s %-20s %3s %-20d %3s %-20d \n", tempVariable.Title, "|", tempVariable.Author, "|", tempVariable.Year, "|", tempVariable.Pages)
	}
	fmt.Println(strings.Repeat(line, 110))
}
