package router

import (
	controllers "backend/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":5000"

	mux.HandleFunc("/post", controllers.GetAllPosts).Methods("GET")
	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}