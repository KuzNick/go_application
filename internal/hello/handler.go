package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}


func NewHelloHandler(router *http.ServeMux) {
	/*
		Создали инстанс для HallowHandler
	*/
	handler := &HelloHandler{}
	/*
		Обработчик
	*/
	router.HandleFunc("/hello", handler.Hello()) // если первым агрументом ничего не передавать, будут обрабатываться только GET запросы
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	/*
		Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
		w http.ResponseWriter - куда будем писать ответ
		req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
	*/
	return  func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
	}
}