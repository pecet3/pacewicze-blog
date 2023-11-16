package main

import (
	"backend/models"
	router "backend/routes"
	"backend/utils"
	"fmt"
)

func main() {
	models.Config()
	token, _ := utils.CreateJWT()
	fmt.Println(token)
	router.SetupAndRun()
}
