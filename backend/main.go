package main

import (
	"backend/models"
	router "backend/routes"
)

func main() {
	models.Config()
	router.SetupAndRun()
}
