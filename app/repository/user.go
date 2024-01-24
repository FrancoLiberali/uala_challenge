package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type IUserRepository interface {
	AddFollower(userID, newFollowerID uint) error
}

type UserRepository struct {
	RDB *redis.Client
}

const Forever = 0

func (repository UserRepository) AddFollower(userID, newFollowerID uint) error {
	followersKey := UserFollowersKey(userID)

	_, err := repository.RDB.RPush(context.Background(), followersKey, newFollowerID).Result()
	if errors.Is(err, redis.Nil) {
		return repository.RDB.Set(context.Background(), followersKey, []uint{newFollowerID}, Forever).Err()
	} else if err != nil {
		return err
	}

	return nil
}

func UserFollowersKey(userID uint) string {
	return fmt.Sprintf("%d-followers", userID)
}
