package main

import (
	"backend/models"
	router "backend/routes"
	"fmt"
)

func main() {
	models.Config()
	post := &models.Post{}

	post, _ = post.CreateAPost()
	r, _ := models.GetAllPosts()
	fmt.Println(r)
	router.SetupAndRun()
}
