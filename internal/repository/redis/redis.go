package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"event-tracking/config"
	"event-tracking/pkg/logger"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, cfg *config.Config) (CacheHelper, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addrs[0],
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.GLogger.Errorf("connect redis error %s", err)
		return &redisHelper{}, err
	}

	logger.GLogger.Info("connect redis successfully")
	return &redisHelper{
		client: client,
	}, nil
}

func ConnectRedis(ctx context.Context, cfg config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addrs[0],
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	logger.GLogger.Info("connect redis successfully")
	return client, nil
}

type CacheHelper interface {
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key string, value string, time time.Duration) error
	Del(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SAdd(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SisMember(ctx context.Context, key string, value interface{}) (bool, error)
	SetExpire(ctx context.Context, key string, expiration time.Duration) error
	SMembers(ctx context.Context, key string) ([]string, error)
}

type redisHelper struct {
	client *redis.Client
}

func (h *redisHelper) GetString(ctx context.Context, key string) (outValue string, err error) {
	if h.client == nil {
		return outValue, redis.Nil
	}
	outValue, err = h.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return outValue, nil
}

func (h *redisHelper) SetString(ctx context.Context, key string, inputValue string, time time.Duration) (err error) {
	return nil
}

func (h *redisHelper) Del(ctx context.Context, key string) (err error) {
	if h.client == nil {
		return redis.Nil
	}
	_, err = h.client.Del(ctx, key).Result()
	fmt.Println("delete redis err:", err)
	return err
}
func (h *redisHelper) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (isSuccess bool, err error) {
	if h.client == nil {
		return false, redis.Nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	isSuccess, err = h.client.SetNX(ctx, key, string(data), expiration).Result()
	if err != nil {
		return false, err
	}
	return isSuccess, nil
}

func (h *redisHelper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if h.client == nil {
		return redis.Nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = h.client.Set(ctx, key, string(data), expiration).Result()
	return err
}

func (h *redisHelper) SAdd(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if h.client == nil {
		return redis.Nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = h.client.SAdd(ctx, key, string(data)).Result()

	if err == nil {
		h.client.Expire(ctx, key, expiration)
	}

	return err
}

func (h *redisHelper) SisMember(ctx context.Context, key string, value interface{}) (bool, error) {
	if h.client == nil {
		return false, redis.Nil
	}

	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	return h.client.SIsMember(ctx, key, string(data)).Result()
}

func (h *redisHelper) SetExpire(ctx context.Context, key string, expiration time.Duration) error {
	if h.client == nil {
		return redis.Nil
	}
	_, err := h.client.Expire(ctx, key, expiration).Result()
	return err
}

func (h *redisHelper) SMembers(ctx context.Context, key string) ([]string, error) {
	if h.client == nil {
		return []string{}, redis.Nil
	}
	return h.client.SMembers(ctx, key).Result()
}
