package main

import (
	"fmt"
	"go_apllication/configs"
	"go_apllication/internal/auth"
	"net/http"
)




func main() {
	/*
		Прокидываем получаение конфига из модуля Configs
	*/
	conf := configs.LoadConfig()

	/*
		Собстенный ServeMux
	*/
	router := http.NewServeMux()

	/*
		СПередаём конфигурацию явно
	*/
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

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
