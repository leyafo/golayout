package api

import (
	"golayout/internal/api/user"
	"golayout/pkg/httpctrl"
	"golayout/pkg/logger"
	"net/http"

	"golayout/internal/api/helper"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(middleware.RequestID)
	r.Use(logger.NewStructuredLogger())
	r.Use(middleware.Recoverer)
	r.Use(httpctrl.SetResponseHeader) //set reqID in header

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(helper.GetJWTAuth()))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Route("/v1", func(r chi.Router) {
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {

		r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		})

		r.Post("/login", user.Login)
	})

	return r
}
