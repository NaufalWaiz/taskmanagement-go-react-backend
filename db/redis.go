package db

import (
    "context"
    "log"
    "project/config"

    "github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() error {
    Rdb = redis.NewClient(&redis.Options{
        Addr:     config.GetEnv("REDIS_ADDR", "localhost:6379"),
        Password: config.GetEnv("REDIS_PASSWORD", ""),
        DB:       0,
    })

    ctx := context.Background()
    if _, err := Rdb.Ping(ctx).Result(); err != nil {
        return err
    }

    log.Println("Successfully connected to Redis!")
    return nil
}