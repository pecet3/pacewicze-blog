package controllers

import (
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := models.GetAllPosts()
	res, err := json.Marshal(posts)

	if err != nil {
		log.Fatalln("error get all posts: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
