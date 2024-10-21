package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestRedisCache(t *testing.T) {

	opts := &redis.UniversalOptions{
		Addrs: []string{"127.0.0.1:6379"},
	}
	client, err := buildRedisClient(context.Background(), opts)
	if err != nil {
		t.Error("create redis client error")
	}
	defer client.Close()
	t.Log("create redis client success")

}
