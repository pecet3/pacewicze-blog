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

func Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseJsonBody(r, user)

	if user.Name == "" && user.Password == "" && user.Email == "" {
		utils.SendErrorJson(w, "incorrect request data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userDb, err := models.GetUserByEmail(user.Email)

	if userDb != (models.User{}) {
		utils.SendErrorJson(w, "there is an user with this email")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = user.CreateAnUser()
	if err != nil {
		utils.SendErrorJson(w, "cannot create a new user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		utils.SendErrorJson(w, "error during parse json, but user has been added")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseJsonBody(r, user)

	if user.Name == "" && user.Password == "" && user.Email == "" {
		utils.SendErrorJson(w, "incorrect request data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userDb, err := models.GetUserByEmail(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SendErrorJson(w, "error with server or db")
		return
	}

	if userDb == (models.User{}) {
		w.WriteHeader(http.StatusUnauthorized)
		utils.SendErrorJson(w, ("there is not any user with this email"))
		return
	}

	isPasswordCorrect, err := utils.VerifyPassword(user.Password, userDb.Salt)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.SendErrorJson(w, ("incorrect password"))
		return
	}

	if isPasswordCorrect == true {
		res, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	return
}
