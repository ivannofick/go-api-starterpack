package routes

import (
	"go-crud/app/controllers"
	"go-crud/app/middleware"

	"github.com/go-chi/chi/v5"
)

func ApiRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		// r.Use(middleware.Recover) // if any

		r.Route("/users", func(r chi.Router) {
			r.Get("/pagination", controllers.PaginatedUsers)
			r.Get("/", controllers.IndexUser)
			r.Post("/", controllers.CreateUserBaru())
			r.Get("/{id}", controllers.ShowUser)
			r.Put("/{id}", controllers.UpdateUser)
			r.Delete("/{id}", controllers.DeleteUser)
		})
	})

	return r
}
