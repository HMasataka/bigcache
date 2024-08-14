package main

import (
	"context"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

func main() {
	ctx := context.Background()

	eviction := 10 * time.Minute

	cache, err := bigcache.New(ctx, bigcache.DefaultConfig(eviction))
	if err != nil {
		panic(err)
	}

	const key = "cache-key"

	if err := cache.Set(key, []byte("value")); err != nil {
		panic(err)
	}

	value, err := cache.Get(key)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(value))
}
