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
	port := ":5000"

	mux.HandleFunc("/post", controllers.Post).Methods("GET")
	mux.HandleFunc("/post", utils.ValidateJWT(controllers.Post)).Methods("POST")

	mux.HandleFunc("/post/{id}", controllers.PostById).Methods("GET")
	mux.HandleFunc("/post/{id}", utils.ValidateJWT(controllers.PostById)).Methods("PUT")

	mux.HandleFunc("/user/register", controllers.Register).Methods("POST")
	mux.HandleFunc("/user/login", controllers.Login).Methods("POST")

	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
