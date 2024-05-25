package main

import (
	"net/http"
	handler "urlshortner/handler"
	models "urlshortner/models"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func init() {
	//Technique 1 : initialize mapper which is global variable.
	models.UrlMapper = &models.UrlStructure{
		KeyMapping: make(map[string]string),
	}

	//Technique 2 : using redis we can store and fetch short url.
	models.InitRedis()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running..."))
	})

	/*
		Technique 1 : two api end point is initialize
		POST - generate url and GET - redirect to the url
	*/
	r.Post(models.ShortUrl, handler.CreatUrlShortnerHandler)

	/*
		Technique 2 : two api end point is initialize using redis
		POST - generate url and GET - redirect to the url
	*/
	r.Post(models.ShortUrlRedis, handler.CreateUrlShortnerHandlerUsingRedis)
	/*
		GET - to redirect to url
	*/
	r.Get(models.RedirectShortUrl, handler.RedirectToShortGeneratedUrl)

	http.ListenAndServe(":"+models.Port, r)
}
