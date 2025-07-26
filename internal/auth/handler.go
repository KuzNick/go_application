package auth

import (
	"fmt"
	"go_apllication/configs"
	"net/http"
)

type AuthHandler struct{
	*configs.Config
}

type AuthHandlerDeps struct{
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	/*
		Создали инстанс для AuthHandler
	*/
	handler := &AuthHandler{
		Config: deps.Config,
	}
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
		fmt.Println(handler.Config.Auth.Secret)
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