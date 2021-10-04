package main

import (
	"Sample_1/ipi/handlers"
	middleware1 "Sample_1/ipi/middleware"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		//r.Use(middleware.SetHeader("Content-Type", "application/json"))
		ath := middleware1.New()
		r.Use(ath.Authenticator())
		h := handlers.NewUser()
		r.Get("/api/me", h.GetInforUser)
		http.ListenAndServe(":8080", r)
	})
}
