package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/C0C045/Prueba-Tecnica/db"
	"github.com/C0C045/Prueba-Tecnica/routes"
)

func main() {
    db.DataBaseConnection()

    r := mux.NewRouter()
    r.HandleFunc("/", routes.HomeHandler)
    r.HandleFunc("/users", routes.GetUsers).Methods("GET")
    r.HandleFunc("/users", routes.PostUser).Methods("POST")
    r.HandleFunc("/details", routes.GetDetail).Methods("GET")
    r.HandleFunc("/details/{id}", routes.PostDetail).Methods("POST")
    r.HandleFunc("/tutorials", routes.GetTutorials).Methods("GET")
    r.HandleFunc("/tutorials/{id}", routes.PostTutorial).Methods("POST")
    r.HandleFunc("/tutorials/{id}", routes.DeleteTutorial).Methods("DELETE")
    r.HandleFunc("/tutorials/{id}", routes.UpdateTutorial).Methods("PUT")

    http.ListenAndServe(":3000", r)
}
