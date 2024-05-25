package handler

import (
	"fmt"
	"log"
	"net/http"
	"urlshortner/models"

	"github.com/lithammer/shortuuid/v4"
)

/*
Technique 1 : CreateUrlShortnerHandler - using mapping
*/
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
	w.Write([]byte(fmt.Sprintf(models.ApiBaseUrl+models.Port+"/short/%s/m ", key)))
}

func InsertKeyAndUrlToMap(key string, url string) {
	if models.UrlMapper != nil {
		models.UrlMapper.Lock.Lock()
		defer models.UrlMapper.Lock.Unlock()
		models.UrlMapper.KeyMapping[key] = url
	}
}

/*
Technique 2 : CreateUrlShortnerHandler - using mapping
*/

func CreateUrlShortnerHandlerUsingRedis(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Generate Short URL - "))
	r.ParseForm()
	url := r.Form.Get("URL")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is empty"))
		return
	}

	//generate key using shortuuid
	key := shortuuid.New()
	InsertKeyAndUrlToRedis(key, url)
	log.Println("url mapped successfully.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(models.ApiBaseUrl+models.Port+"/short/%s/r ", key)))
}

func InsertKeyAndUrlToRedis(key, url string) {
	if models.RedisClient == nil {
		log.Println("Redis client is not initialized")
		return
	}
	err := models.RedisClient.Set(models.Ctx, key, url, 0).Err()
	if err != nil {
		log.Printf("Could not set key %s : %v", key, err)
		return
	}
}
