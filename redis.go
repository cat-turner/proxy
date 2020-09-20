package proxy

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client     redis.Client
	KeyTimeout time.Duration
}

func NewRedisClient(keyTimeout time.Duration) RedisClient {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// set duration by some con
	return RedisClient{
		Client: *rc, KeyTimeout: keyTimeout,
	}
}

func (rc RedisClient) Put(key string, value string) error {
	var ctx = context.Background()
	err := rc.Client.Set(ctx, key, value, rc.KeyTimeout).Err()
	if err != nil {
		return err
	}
	return nil
}
