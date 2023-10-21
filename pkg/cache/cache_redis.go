package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// RedisCacheManager is a Redis-based cache implementation.
type RedisCacheManager struct {
	client *redis.Client
}

func (client *RedisCacheManager) Set(key string, value interface{}, expire int) error {
	ctx := context.Background()
	err := client.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisCacheManager) SetString(key, value string, expire int) error {
	ctx := context.Background()
	err := client.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisCacheManager) Sets(values map[string]string, expire int) error {
	ctx := context.Background()
	pipe := client.client.TxPipeline()
	for key, value := range values {
		pipe.Set(ctx, key, value, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("Transaction failed:", err)
		err := pipe.Discard()
		if err != nil {
			return err
		}
		fmt.Println("Transaction rolled back.")
	}
	return nil
}

func (client *RedisCacheManager) Get(key string) (*string, error) {
	ctx := context.Background()
	value, err := client.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &value, nil
}

func (client *RedisCacheManager) Delete(key string) error {
	ctx := context.Background()
	err := client.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
