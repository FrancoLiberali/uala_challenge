package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/FrancoLiberali/uala_challenge/app/adapters"
	"github.com/FrancoLiberali/uala_challenge/app/models"
	"github.com/FrancoLiberali/uala_challenge/app/repository"
)

type TwitterService struct {
	Repository repository.IRepository
	Clock      adapters.IClock
}

var (
	ErrCantFollowYourself = errors.New("follow yourself not allowed")
	ErrFollow             = errors.New("error adding follow")
	ErrTweet              = errors.New("error tweeting")
)

// Follow makes user followerID to follow user followedID
//
// Returns ErrCantFollowYourself if followerID and followedID are equal or
// ErrFollow if an error occurred during the execution
func (service TwitterService) Follow(followerID, followedID uint) error {
	if followedID == followerID {
		return ErrCantFollowYourself
	}

	err := service.Repository.AddFollower(followedID, followerID)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("%w from user %d to user %d", ErrFollow, followerID, followedID)
	}

	return nil
}

// Tweet creates a tweet of userID
//
// Returns ErrTweetTooLong if content len is bigger than 280 characters or
// ErrTweet if an error occurred during the execution
func (service TwitterService) Tweet(userID uint, content string) (uuid.UUID, error) {
	tweet, err := models.NewTweet(userID, service.Clock.Now(), content)
	if err != nil {
		return uuid.Nil, err
	}

	tweetID, err := service.Repository.CreateTweet(*tweet)
	if err != nil {
		log.Println(err.Error())
		return uuid.Nil, fmt.Errorf("%w from user %d", ErrTweet, userID)
	}

	followers, err := service.Repository.GetFollowers(userID)
	if err != nil {
		log.Println(err.Error())
		return uuid.Nil, fmt.Errorf("%w from user %d", ErrTweet, userID)
	}

	for _, follower := range followers {
		err = service.Repository.AddTweetToTimeline(tweetID, follower)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return tweetID, nil
}
