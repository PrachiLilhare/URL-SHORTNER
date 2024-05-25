package handler

import (
	"net/http"
	"urlshortner/models"

	"github.com/go-chi/chi/v5"
)

func RedirectToShortGeneratedUrl(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key field is empty"))
		return
	}

	redirectUrl := FetchUrlShortnerUsingKey(key)
	if redirectUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is not exist with respect to given key"))
		return
	}
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}

func FetchUrlShortnerUsingKey(key string) string {
	if models.UrlMapper != nil {
		models.UrlMapper.Lock.Lock()
		defer models.UrlMapper.Lock.Unlock()

		return models.UrlMapper.KeyMapping[key]
	}
	return ""
}
