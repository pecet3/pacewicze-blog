package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
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

	if method == "POST" {
		post := &models.Post{}
		utils.ParseJsonBody(r, post)
		post, err := post.CreateAPost()
		if err != nil {
			log.Fatalln("error create post: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
