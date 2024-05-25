package models

import "sync"

type UrlStructure struct {
	KeyMapping map[string]string /* store key : url */
	Lock       sync.Mutex
}

// global variable to store data using map {key : url}
var UrlMapper *UrlStructure
