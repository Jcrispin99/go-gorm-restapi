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

	http.ListenAndServe(":3000", r)
}
