package utils

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error_message"`
}

func GetRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func ParseJsonBody(r *http.Request, b interface{}) {
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), b)
	if err != nil {
		log.Println("error parse json: ", err)
	}
	return
}

func SendErrorJson(w http.ResponseWriter, message string) error {
	errorMessage := ErrorResponse{Error: message}

	res, err := json.Marshal(errorMessage)

	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return nil
}
