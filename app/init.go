package app

import (
	"context"
	"errors"
	"os"

	"github.com/redis/go-redis/v9"

	"github.com/FrancoLiberali/uala_challenge/app/repository"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

//go:generate mockery --all --keeptree

const (
	CacheURLEnvVar      = "CACHE_URL"
	CachePasswordEnvVar = "CACHE_PASSWORD"
)

var ErrCacheNotConfigured = errors.New("cache env variables not configured")

// Creates a new service.FollowService that connects to cache from information in env vars.
//
// Returns ErrCacheNotConfigured if env vars are not configured or an error is the connection cannot be established
func NewFollowService() (*service.FollowService, *redis.Client, error) {
	// get cache information from env vars
	cacheURL := os.Getenv(CacheURLEnvVar)
	cachePassword := os.Getenv(CachePasswordEnvVar)

	if cacheURL == "" || cachePassword == "" {
		return nil, nil, ErrCacheNotConfigured
	}

	// create cache client
	rdb := redis.NewClient(&redis.Options{
		Addr:     cacheURL,
		Password: cachePassword,
	})

	// test client connection
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, nil, err
	}

	// return new service
	return &service.FollowService{
		UserRepository: repository.UserRepository{
			RDB: rdb,
		},
	}, rdb, nil
}
