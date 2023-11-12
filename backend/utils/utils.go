package utils

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

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
