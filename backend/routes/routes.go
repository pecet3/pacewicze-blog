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

	mux.HandleFunc("/api/post", controllers.Post).Methods("GET", "POST")
	mux.HandleFunc("/api/post/{id}", controllers.PostById).Methods("GET", "PUT")
	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
