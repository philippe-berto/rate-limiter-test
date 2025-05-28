package main

import (
	"github.com/caarlos0/env/v10"
	rl "github.com/philippe-berto/rate-limiter"
	r "github.com/philippe-berto/rate-limiter/database/redis"
)

type (
	BaseConfigRL struct {
		RateLimiterConfig rl.RateLimiterConfig `json:"rate_limiter_config"`
		RedisConfig       r.RedisConfig        `json:"redis_config"`
	}
)

func LoadConfig() BaseConfigRL {
	cfg := BaseConfigRL{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return cfg

}
