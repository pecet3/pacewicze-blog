package router

import (
	controllers "backend/controllers"
	"backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}

func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":5000"

	mux.HandleFunc("/test", utils.ValidateJWT(Home)).Methods("GET")

	mux.HandleFunc("/api/post", controllers.Post).Methods("GET", "POST")
	mux.HandleFunc("/api/post/{id}", controllers.PostById).Methods("GET", "PUT")
	mux.HandleFunc("/api/user/register", controllers.Register).Methods("POST")
	mux.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
