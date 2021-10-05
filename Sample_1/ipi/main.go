package main

import (
	"Sample_1/ipi/handlers"
	middleware1 "Sample_1/ipi/middleware"
	"Sample_1/ipi/repositories"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		connecting, err := repositories.NewDbmanager()
		if err != nil {
			log.Fatal("error conecting database", err)
			fmt.Println("err")
			return
		}

		r.Use(middleware.SetHeader("Content-Type", "application/json"))

		ath := middleware1.New(connecting)
		r.Use(ath.Authenticator())

		h, err := handlers.NewUser(connecting)
		if err != nil {
			panic(err)
		}
		r.Get("/api/me", h.GetInforUser)
		http.ListenAndServe(":8080", r)
	})
}
