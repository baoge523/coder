package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func buildRedisClient(ctx context.Context, options *redis.UniversalOptions) (redis.UniversalClient, error) {

	client := redis.NewUniversalClient(options)

	// 看是否可以ping通
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
