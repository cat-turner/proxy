package proxy

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client     redis.Client
	KeyTimeout time.Duration
}

func NewRedisClient(keyTimeout *time.Duration) RedisClient {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if keyTimeout != nil {
		return RedisClient{
			Client:     *rc,
			KeyTimeout: *keyTimeout,
		}
	}

	return RedisClient{
		Client: *rc,
	}
}

// Put ...
func (rc RedisClient) Put(key string, value string) error {
	var ctx = context.Background()
	err := rc.Client.Set(ctx, key, value, rc.KeyTimeout).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (rc RedisClient) Get(key string) (*string, error) {
	var ctx = context.Background()
	val, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil

}
