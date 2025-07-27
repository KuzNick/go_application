package auth

import (
	"encoding/json"
	"fmt"
	"go_apllication/configs"
	"go_apllication/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
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
	router.HandleFunc("POST /auth/login", handler.Login())       // если первым агрументом ничего не передавать, будут обрабатываться только GET запросы
	router.HandleFunc("POST /auth/register", handler.Register()) // если первым агрументом ничего не передавать, будут обрабатываться только GET запросы
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	/*
		Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
		w http.ResponseWriter - куда будем писать ответ
		req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
	*/
	return func(w http.ResponseWriter, req *http.Request) {
		/*
			Объявляем структуру с телом запроса
		*/
		var payload LoginRequest
		/*
			Мы декодируем полученный body (req.Body) и передаём его в итоговый payload
		*/
		err := json.NewDecoder(req.Body).Decode(&payload)
		/*
			Проверяем на возможные ошибки, например тело не в формате Json
		*/
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}

		/*
			Блоки с простейшей валидацией полей Email и Password
		*/
		if payload.Email == "" {
			res.Json(w, "Email required", 402)
			return
		}
		if payload.Password == "" {
			res.Json(w, "Password required", 402)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}

		/*
			Подтягиваем функцию из пакета "go_apllication/pkg/res" по формированию ответа
		*/
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	/*
		Обязательные 2 аргумента: (w http.ResponseWriter, req *http.Request)
		w http.ResponseWriter - куда будем писать ответ
		req *http.Request - указатель на hhtp request (запрос, который пришёл изначально)
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
