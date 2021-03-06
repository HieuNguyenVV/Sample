package main

import (
	"Sample_1/helper"
	"Sample_1/ipi/handlers"
	middleware1 "Sample_1/ipi/middleware"
	"Sample_1/ipi/psql"
	"Sample_1/ipi/repositories"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		log.Printf("error bind config file. Err: %s", err)
		return
	}
	r := chi.NewRouter()
	dbmanager, err := psql.NewDbmanager()
	if err != nil {
		log.Fatalf("error conecting database . Err: %s", err)
	}
	r.Group(func(r chi.Router) {
		connecting := repositories.NewUserRepository(dbmanager)

		r.Use(middleware.SetHeader("Content-Type", "application/json"))

		ath := middleware1.New(connecting)
		r.Use(ath.Authenticator())

		h := handlers.NewUser(connecting)
		r.Get("/api/me", h.GetInforUser)
		http.ListenAndServe(":8080", r)
	})
}
