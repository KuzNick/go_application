package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}


func NewAuthHandler(router *http.ServeMux) {
	/*
		Создали инстанс для AuthHandler
	*/
	handler := &AuthHandler{}
	/*
		Обработчик
	*/
	router.HandleFunc("POST /auth/login", handler.Login()) // если первым агрументом ничего не передавать, будут обрабатываться только GET запросы
	router.HandleFunc("POST /auth/register", handler.Register()) // если первым агрументом ничего не передавать, будут обрабатываться только GET запросы
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	/*
		Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
		w http.ResponseWriter - куда будем писать ответ
		req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
	*/
	return  func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	/*
		Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
		w http.ResponseWriter - куда будем писать ответ
		req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
	*/
	return  func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}