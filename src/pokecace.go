package src

import (
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data 	  []byte
}

func NewCache() {
	
}
