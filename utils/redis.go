package utils

import (
    "context"
    "github.com/redis/go-redis/v9"
    "time"
)

type RedisClient struct {
    client *redis.Client
}

func NewRedisClient(addr string, password string, db int) *RedisClient {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })
    return &RedisClient{client: client}
}

func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
    ctx := context.Background()
    return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
    ctx := context.Background()
    return r.client.Get(ctx, key).Result()
}