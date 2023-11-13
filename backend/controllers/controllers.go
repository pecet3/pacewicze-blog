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
	method := r.Method
	vars := mux.Vars(r)
	postId := vars["id"]

	if method == "GET" {
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

	if method == "PUT" {
		vars := mux.Vars(r)
		id := vars["id"]
		postRequest := &models.Post{}
		utils.ParseJsonBody(r, postRequest)

		post, err := models.GetPostById(id)
		if err != nil {
			log.Println("error update post: ", err)
		}

		if postRequest.Title != "" {
			post.Title = postRequest.UserId
		}

		if postRequest.Content != "" {
			post.Content = postRequest.Content
		}
		if postRequest.ImageUrl != "" {
			post.ImageUrl = postRequest.ImageUrl
		}

		_, err = post.EditAPost()
		if err != nil {
			log.Println("error update post: ", err)
		}

		w.WriteHeader(http.StatusOK)
	}

}
