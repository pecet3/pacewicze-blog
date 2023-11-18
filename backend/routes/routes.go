package router

import (
	controllers "backend/controllers"
	"backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":30111"

	mux.HandleFunc("/api/post", controllers.Post).Methods("GET")
	mux.HandleFunc("/api/post", utils.ValidateJWT(controllers.Post)).Methods("POST")

	mux.HandleFunc("/api/post/{id}", controllers.PostById).Methods("GET")
	mux.HandleFunc("/api/post/{id}", utils.ValidateJWT(controllers.PostById)).Methods("PUT")

	mux.HandleFunc("/api/user/register", controllers.Register).Methods("POST")
	mux.HandleFunc("/api/user/login", controllers.Login).Methods("POST")

	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
