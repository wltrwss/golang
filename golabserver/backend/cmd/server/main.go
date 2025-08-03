package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w - это объект, в который мы пишем ответ клиенту.
	// r - это объект с данными запроса(метод, заголовк, URL и тд.)
	fmt.Println(w, "Hello, World!")
}

func main(){
		http.HandleFunc("/", handler)

		fmt.Println("Сервер запущен на http://localhost:8080")
		http.ListenAndServe(":8080", nil)
}