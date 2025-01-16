package main

import (
	"myblogapi/internal/db"
	"myblogapi/internal/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.Init()
	defer db.Close()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", routes.AuthRoutes)
	http.ListenAndServe(":5000", r)
}
