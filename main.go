package main

import (
	"fmt"
	"net/http"
)

/*
Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
w http.ResponseWriter - куда будем писать ответ
req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

func main() {
	/*
		Собстенный ServeMux
	*/
	router := http.NewServeMux()
	/*
		Обработчик
	*/
	router.HandleFunc("/hello", hello) // если первм агрументом ничего не передавать, будут обрабатываться толкько GET запросы

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")

	/*
		Случшает все коннекшены по tcp по определенном у порту и адресу и передавать
		их по определенному роутингу, который мы так же можем задавать
	*/
	server.ListenAndServe() //ипользуем дефолтный ServerMux

}
