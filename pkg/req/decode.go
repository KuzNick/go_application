package req

import (
	"encoding/json"
	"io"
)

func Decode[T any] (body io.ReadCloser) (T, error) {
	/*
		Объявляем структуру с телом запроса
	*/
	var payload T
	/*
		Мы декодируем полученный body (req.Body) и передаём его в итоговый payload
	*/
	err := json.NewDecoder(body).Decode(&payload)
	/*
		Проверяем на возможные ошибки, например тело не в формате Json
	*/
	if err != nil {
		return payload, err
	}
	return payload, nil
}