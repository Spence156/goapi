package handlers

import (
	"github.com/Spence156/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes) // Strips trailing / from the end of the URL to prevent 404 errors

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization) // Enforces Authentication using the internal/middleware/authorization.go

		router.Get("/coins", GetCoinBalance) // Produces a API endpoint which calls the internal/handlers/GetCoinBalance
	})
}
