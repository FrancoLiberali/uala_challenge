package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/FrancoLiberali/uala_challenge/app/models"
)

type IRepository interface {
	// AddFollower adds newFollowerID to the list of followers of userID
	AddFollower(userID, newFollowerID uint) error
	// GetFollowers returns the list of followers of user userID
	GetFollowers(userID uint) ([]uint, error)
	// CreateTweet creates a tweet returning its id
	CreateTweet(tweet models.Tweet) (uuid.UUID, error)
	// AddTweetToTimeline adds a tweetID to the user's timeline, maintaining a maximin length of 40 tweets
	AddTweetToTimeline(tweetID uuid.UUID, userID uint) error
	// GetTimeline returns the list of tweets ids in a user timeline
	GetTimeline(userID uint) ([]uuid.UUID, error)
}

type Repository struct {
	RDB *redis.Client
}

const (
	Forever             = 0 // Forever indicates an infinite TTL
	TweetTTL            = 2 * time.Hour
	MaxTweetsInTimeline = 40
)

// AddFollower adds newFollowerID to the list of followers of userID
func (repository Repository) AddFollower(userID, newFollowerID uint) error {
	followersKey := UserFollowersKey(userID)

	return repository.RDB.SAdd(context.Background(), followersKey, newFollowerID).Err()
}

// GetFollowers returns the list of followers of user userID
func (repository Repository) GetFollowers(userID uint) ([]uint, error) {
	followersKey := UserFollowersKey(userID)

	idsString, err := repository.RDB.SMembers(context.Background(), followersKey).Result()
	if err != nil {
		return nil, err
	}

	ids := make([]uint, 0, len(idsString))

	for _, idString := range idsString {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return nil, err
		}

		ids = append(ids, uint(id))
	}

	return ids, nil
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

// AddTweetToTimeline adds a tweetID to the user's timeline, maintaining a maximin length of 40 tweets
func (repository Repository) AddTweetToTimeline(tweetID uuid.UUID, userID uint) error {
	timelineKey := TimelineKey(userID)

	timelineLen, err := repository.RDB.LLen(context.Background(), timelineKey).Result()
	if err != nil {
		return err
	}

	if timelineLen == MaxTweetsInTimeline {
		err := repository.RDB.RPop(context.Background(), timelineKey).Err()
		if err != nil {
			return err
		}
	}

	return repository.RDB.LPush(context.Background(), timelineKey, tweetID.String()).Err()
}

// GetTimeline returns the list of tweets ids in a user timeline
func (repository Repository) GetTimeline(userID uint) ([]uuid.UUID, error) {
	timelineKey := TimelineKey(userID)

	idsString, err := repository.RDB.LRange(context.Background(), timelineKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, 0, len(idsString))

	for _, idString := range idsString {
		id, err := uuid.Parse(idString)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// TimelineKey returns the key that stores an user's timeline
func TimelineKey(userID uint) string {
	return fmt.Sprintf("tl-%d", userID)
}
