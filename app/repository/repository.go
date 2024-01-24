package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/FrancoLiberali/uala_challenge/app/models"
)

type IRepository interface {
	// AddFollower adds newFollowerID to the list of followers of userID
	AddFollower(userID, newFollowerID uint) error
	// CreateTweet creates a tweet returning its id
	CreateTweet(tweet models.Tweet) (uuid.UUID, error)
}

type Repository struct {
	RDB *redis.Client
}

const (
	Forever  = 0 // Forever indicates an infinite TTL
	TweetTTL = 2 * time.Hour
)

// AddFollower adds newFollowerID to the list of followers of userID
func (repository Repository) AddFollower(userID, newFollowerID uint) error {
	followersKey := UserFollowersKey(userID)

	return repository.RDB.SAdd(context.Background(), followersKey, newFollowerID).Err()
}

// UserFollowersKey returns the key that stores the list of followers of userID in the cache
func UserFollowersKey(userID uint) string {
	return fmt.Sprintf("%d-followers", userID)
}

// CreateTweet creates a tweet returning its id
func (repository Repository) CreateTweet(tweet models.Tweet) (uuid.UUID, error) {
	tweetID := uuid.New()
	tweetKey := TweetKey(tweetID)

	return tweetID, repository.RDB.Set(context.Background(), tweetKey, tweet, TweetTTL).Err()
}

// TweetKey returns the key that stores a tweet by id
func TweetKey(tweetID uuid.UUID) string {
	return fmt.Sprintf("tweet-%s", tweetID)
}
