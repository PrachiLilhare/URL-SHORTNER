package main

import (
	"net/http"
	handler "urlshortner/handler"
	models "urlshortner/models"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func init() {
	//initialize mapper which is global variable
	models.UrlMapper = &models.UrlStructure{
		KeyMapping: make(map[string]string),
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running..."))
	})

	r.Post("/short-url", handler.CreatUrlShortnerHandler)
	r.Get("/short/{key}", handler.RedirectToShortGeneratedUrl)
	http.ListenAndServe(":3007", r)
}
