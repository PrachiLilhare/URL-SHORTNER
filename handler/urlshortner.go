package handler

import (
	"fmt"
	"log"
	"net/http"
	"urlshortner/models"

	"github.com/lithammer/shortuuid/v4"
)

func CreatUrlShortnerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Generated Short URL "))
	r.ParseForm()
	url := r.Form.Get("URL")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url field is empty"))
		return
	}
	/*generate key*/
	key := shortuuid.New()
	InsertKeyAndUrlToMap(key, url)
	log.Println("url mapped successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("http://localhost:3007/short/%s", key)))
}

func InsertKeyAndUrlToMap(key string, url string) {
	if models.UrlMapper != nil {
		models.UrlMapper.Lock.Lock()
		defer models.UrlMapper.Lock.Unlock()
		models.UrlMapper.KeyMapping[key] = url
	}
}
