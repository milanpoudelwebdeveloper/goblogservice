package routes

import (
	"myblogapi/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
	r.Post("/signup", controllers.SignUp)
	r.Post("/login", controllers.LoginUser)
}
