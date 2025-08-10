package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Инициализация роутера
	router := http.NewServeMux()

	// Создание сценариев
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "VOID")
	})

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ping")
	})

	router.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello iam server")
	})

	router.HandleFunc("/link", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "link")
	})

	// Запуск
	err := http.ListenAndServe(":1488", router)
	if err != nil {
		fmt.Println("ERROR", err)
		panic(err)
	}

}
