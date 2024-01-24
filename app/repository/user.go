package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type IUserRepository interface {
	AddFollower(userID, newFollowerID uint) error
}

type UserRepository struct {
	RDB *redis.Client
}

const Forever = 0 // Forever indicates an infinite TTL

// AddFollower adds newFollowerID to the list of followers of userID
func (repository UserRepository) AddFollower(userID, newFollowerID uint) error {
	followersKey := UserFollowersKey(userID)

	return repository.RDB.SAdd(context.Background(), followersKey, newFollowerID).Err()
}

// UserFollowersKey returns the key that stores the list of followers of userID in the cache
func UserFollowersKey(userID uint) string {
	return fmt.Sprintf("%d-followers", userID)
}
