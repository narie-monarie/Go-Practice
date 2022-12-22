package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

cors := cors.New(cors.Options{
  AllowedOrigins:   []string{"*"},
  AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
  AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Toke\
n"},
  AllowCredentials: true,
  MaxAge:           300, // Maximum value not ignored by any of major browsers
  })
r.Use(cors.Handler)

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!\n", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		println(err)
	}
}
