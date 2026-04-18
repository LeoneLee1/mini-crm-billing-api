package redis

import (
	"context"
	"fmt"
	"mini-crm-billing-api/source/config"
	"mini-crm-billing-api/source/pkg/logger"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config) (*redis.Client, error) {
	db, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		db = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Error().Err(err).Msg("Failed to connect to Redis")
		return nil, err
	}

	return client, nil
}
