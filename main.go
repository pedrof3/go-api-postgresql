package main

import (
	"api-postgresql/configs"
	"api-postgresql/handlers"
	"fmt"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /", handlers.Create)
	mux.HandleFunc("PUT /{id}", handlers.Update)
	mux.HandleFunc("DELETE /{id}", handlers.Delete)
	mux.HandleFunc("GET /", handlers.List)
	mux.HandleFunc("GET /{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), mux)
}
