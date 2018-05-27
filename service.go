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
	s.Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if _, err := s.Client.Ping().Result(); err != nil {
		return err
	}
	if cfg.Output != "" {
		return s.Client.Process(redis.NewStringCmd("OUTPUT", cfg.Output))
	}

	return nil
}
