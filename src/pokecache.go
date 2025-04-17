package src

import (
	"time"
	"fmt"
)

type cacheEntry struct {
	createdAt time.Time
	data 	  []byte
}

func NewCache(input string) string {
	return fmt.Sprintf("%s", input)
}
