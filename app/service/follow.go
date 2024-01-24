package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/FrancoLiberali/uala_challenge/app/repository"
)

type FollowService struct {
	UserRepository repository.IUserRepository
}

var (
	ErrCantFollowYourself = errors.New("error follow yourself not allowed")
	ErrFollow             = errors.New("error adding follow")
)

func (service FollowService) Follow(followerID, followedID uint) error {
	if followedID == followerID {
		return ErrCantFollowYourself
	}

	err := service.UserRepository.AddFollower(followedID, followerID)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("%w from user %d to user %d", ErrFollow, followerID, followedID)
	}

	return nil
}
