package routes

import (
	"go-crud/app/controllers"
	"go-crud/app/middleware"

	"github.com/go-chi/chi/v5"
)

func ApiRoutes() *chi.Mux {
	r := chi.NewRouter()

	// auth
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Get("/users", controllers.IndexUser)
		r.Get("/users/{id}", controllers.ShowUser)
		r.Put("/users/{id}", controllers.UpdateUser)
		r.Delete("/users/{id}", controllers.DeleteUser)
		r.Get("/users/pagination", controllers.PaginatedUsers)
		r.Post("/users/create/data", controllers.CreateUserBaru())
	})

	return r
}
