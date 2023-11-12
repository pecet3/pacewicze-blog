package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Post(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		posts, err := models.GetAllPosts()
		res, err := json.Marshal(posts)

		if err != nil {
			log.Println("error get all posts: ", err)
			w.WriteHeader(http.StatusInternalServerError)
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
			log.Println("error create post: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]

	post, err := models.GetPostById(postId)

	if post == (models.Post{}) {
		w.WriteHeader(http.StatusNotFound)
		utils.SendErrorJson(w, ("Not found any record with id: " + postId))
		return
	}
	if err != nil {
		log.Println("error get post by id: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(post)

	if err != nil {
		log.Println("error parsing json: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
