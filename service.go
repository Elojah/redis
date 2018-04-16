package redis

import (
	"github.com/go-redis/redis"
)

// Service embed a connected redis client.
type Service struct {
	*redis.Client
}

// Dial connects client to external redis service.
func (s *Service) Dial(cfg Config) error {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping().Result()
	return err
}
