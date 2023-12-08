package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"time"
)

var LocalCache = cache.New(5*time.Minute, 10*time.Minute)

// RememberCache creates a cache with a given expiration time and a function to fetch data in case of a cache miss.
func RememberCache(key string, expiration time.Duration, fetchFunc func() (interface{}, error)) (interface{}, error) {
	// Try to get the value from the cache
	if value, found := LocalCache.Get(key); found {
		log.Print(key, "Cache hit!")
		return value, nil
	}

	// Cache miss, fetch the data
	log.Print(key, "Cache miss! Fetching data...")
	data, err := fetchFunc()
	if err != nil {
		return nil, err
	}

	// Store the fetched data in the cache
	LocalCache.Set(key, data, expiration)

	return data, nil
}
