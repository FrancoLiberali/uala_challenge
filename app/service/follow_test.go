package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	repositoryMocks "github.com/FrancoLiberali/uala_challenge/app/mocks/repository"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

func TestFollowReturnsErrorIfTryToFollowYourself(t *testing.T) {
	followService := service.FollowService{}

	err := followService.Follow(1, 1)
	require.ErrorIs(t, err, service.ErrCantFollowYourself)
}

func TestFollowReturnsErrorIfRepositoryReturnsError(t *testing.T) {
	mockUserRepository := repositoryMocks.NewIUserRepository(t)

	followService := service.FollowService{
		UserRepository: mockUserRepository,
	}

	mockUserRepository.On("AddFollower", uint(2), uint(1)).Return(errors.New("an error"))

	err := followService.Follow(1, 2)
	require.ErrorIs(t, err, service.ErrFollow)
	require.ErrorContains(t, err, "from user 1 to user 2")
}

func TestFollowReturnsNilIfRepositoryWorks(t *testing.T) {
	mockUserRepository := repositoryMocks.NewIUserRepository(t)

	followService := service.FollowService{
		UserRepository: mockUserRepository,
	}

	mockUserRepository.On("AddFollower", uint(2), uint(1)).Return(nil)

	err := followService.Follow(1, 2)
	require.NoError(t, err)
}
