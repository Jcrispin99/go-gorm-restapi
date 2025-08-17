package main

import (
	"net/http"

	"github.com/faztweb/go-gorm-restapi/db"
	"github.com/faztweb/go-gorm-restapi/models"
	"github.com/faztweb/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.User{}, models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler).Methods("GET")

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", routes.PutUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
