package main

import (
	"fmt"
	"net/http"
)

func startServer(port string) (err error) {
	fmt.Printf("Сервер успешно запущен! Порт%s\n", port)
	err = http.ListenAndServe(port, nil)
	return
}

func slashHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "/")
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ping")
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello")
}

func main() {
	var port string = ": 1111"

	// ОБРАБОТЧИК "/"
	http.HandleFunc("/", slashHandle)

	// ОБРАБОТЧИК "/ping"
	http.HandleFunc("/ping", pingHandle)

	// ОБРАБОТЧИК "/hello"
	http.HandleFunc("/hello", helloHandle)

	// ЗАПУСК СЕРВЕРА
	err := startServer(port)
	if err != nil {
		fmt.Println("ОШИБКА:", err)
	}
}
