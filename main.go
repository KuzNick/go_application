package main

import (
	"fmt"
	"net/http"
)




func main() {
	/*
		Собстенный ServeMux
	*/
	router := http.NewServeMux()
	NewHelloHandler(router)

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
