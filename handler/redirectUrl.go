package handler

import (
	"log"
	"net/http"
	"urlshortner/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

func RedirectToShortGeneratedUrl(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key field is empty"))
		return
	}
	method := chi.URLParam(r, "method")
	redirectUrl := ""
	switch method {
	//using mapping
	case "m":
		redirectUrl = FetchUrlShortnerUsingKey(key)
	//using redis
	case "r":
		redirectUrl = FetchUrlShortnerUsingKeyInRedis(key)
	default:
		log.Printf("key does not exist %s ", key)
	}

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

func FetchUrlShortnerUsingKeyInRedis(key string) string {
	url, err := models.RedisClient.Get(models.Ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		return ""
	}
	return url
}
