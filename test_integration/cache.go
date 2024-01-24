package testintegration

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func CleanCache(rdb *redis.Client) {
	rdb.FlushDB(context.Background())
}
