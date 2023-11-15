package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error_message"`
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
