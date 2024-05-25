package models

import "sync"

type UrlStructure struct {
	KeyMapping map[string]string /* store key : url */
	Lock       sync.Mutex
}

// global variable to store data using map {key : url}
var UrlMapper *UrlStructure

// Base url and port on which server is running
var Port = "3000"
var ApiBaseUrl = "http://localhost:"

// two api end points
var (
	//using mapping
	ShortUrl         = "/short-url"
	RedirectShortUrl = "/short/{key}/{method}"

	//using redis
	ShortUrlRedis = "/short-rurl"
)
