package proxy

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// RedisClient is used in this package for the external cache
type RedisClient struct {
	Client     redis.Client
	KeyTimeout time.Duration
}

// NewRedisClient creates new redis client
func NewRedisClient(keyTimeout *time.Duration, redisUrl string) RedisClient {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	if keyTimeout != nil {
		return RedisClient{
			Client:     *client,
			KeyTimeout: *keyTimeout,
		}
	}

	return RedisClient{
		Client: *client,
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
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &val, nil

}
