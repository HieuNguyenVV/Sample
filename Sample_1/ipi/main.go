package main

import (
	"Sample_1/ipi/handlers"
	"Sample_1/ipi/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		ath := middleware.New()
		r.Use(ath.Authenticator())
		h := handlers.NewUser()
		r.Get("/ipi/me", h.GetInforUser)
		http.ListenAndServe(":8080", r)
	})
}
