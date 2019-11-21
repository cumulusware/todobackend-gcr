package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"todobackend-gcr/data/firestoreds"
	"todobackend-gcr/resources/todos"
)

func createRoutes(ds *firestoreds.Store) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	todos.AddRoutes(r, ds.Todos, "/api/todos")
	return r
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "OPTIONS", "POST", "DELETE", "PUT", "PATCH"},
		AllowedHeaders:   []string{"accept", "content-type"},
		Debug:            true,
	})
}
