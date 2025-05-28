package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ratelimiter "github.com/philippe-berto/rate-limiter"
	"github.com/philippe-berto/rate-limiter/database/redis"
)

func main() {
	ctx := context.Background()
	cfg := LoadConfig()

	fmt.Println("Config:", cfg)

	db := redis.New(ctx, cfg.RedisConfig)

	rl := ratelimiter.New(ctx, cfg.RateLimiterConfig, db)

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(rl.Middleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on :8080")
}
