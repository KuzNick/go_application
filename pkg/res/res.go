package res

import (
	"encoding/json"
	"net/http"
)


/*
	Функция, которая формирует ответ (Response)
*/

func Json(w http.ResponseWriter, data any, statuscode int) {
	/*
		Задаём формат ответа Json
	*/
	w.Header().Set("Content-Type", "application/json")
	/*
		Задаём статус код ответа
	*/
	w.WriteHeader(statuscode)
	/*
		Закодировали в http.ResponseWriter наш ответ из переменной res
	*/
	json.NewEncoder(w).Encode(data)
}