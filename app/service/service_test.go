package service_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/FrancoLiberali/uala_challenge/app/adapters"
	adaptersMocks "github.com/FrancoLiberali/uala_challenge/app/mocks/adapters"
	repositoryMocks "github.com/FrancoLiberali/uala_challenge/app/mocks/repository"
	"github.com/FrancoLiberali/uala_challenge/app/models"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

func TestFollowReturnsErrorIfTryToFollowYourself(t *testing.T) {
	followService := service.TwitterService{}

	err := followService.Follow(1, 1)
	require.ErrorIs(t, err, service.ErrCantFollowYourself)
}

func TestFollowReturnsErrorIfRepositoryReturnsError(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)

	followService := service.TwitterService{
		Repository: mockRepository,
	}

	mockRepository.On("AddFollower", uint(2), uint(1)).Return(errors.New("an error"))

	err := followService.Follow(1, 2)
	require.ErrorIs(t, err, service.ErrFollow)
	require.ErrorContains(t, err, "from user 1 to user 2")
}

func TestFollowReturnsNilIfRepositoryWorks(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)

	followService := service.TwitterService{
		Repository: mockRepository,
	}

	mockRepository.On("AddFollower", uint(2), uint(1)).Return(nil)

	err := followService.Follow(1, 2)
	require.NoError(t, err)
}

func TestTweetReturnsErrorIfContentLenIfBiggerThan280Characters(t *testing.T) {
	followService := service.TwitterService{
		Clock: adapters.Clock{},
	}

	_, err := followService.Tweet(0, strings.Repeat("aguante banfield", 20))
	require.ErrorIs(t, err, models.ErrTweetTooLong)
}

func TestTweetReturnsErrorIfRepositoryCreateTweetReturnsError(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)
	mockClock := adaptersMocks.NewIClock(t)

	followService := service.TwitterService{
		Repository: mockRepository,
		Clock:      mockClock,
	}

	now := time.Now()

	mockClock.On("Now").Return(now)
	mockRepository.On("CreateTweet", models.Tweet{UserID: 1, Timestamp: now, Content: "aguante banfield"}).Return(uuid.Nil, errors.New("an error"))

	_, err := followService.Tweet(1, "aguante banfield")
	require.ErrorIs(t, err, service.ErrTweet)
	require.ErrorContains(t, err, "from user 1")
}

func TestTweetReturnsErrorIfRepositoryGetFollowersReturnsError(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)
	mockClock := adaptersMocks.NewIClock(t)

	followService := service.TwitterService{
		Repository: mockRepository,
		Clock:      mockClock,
	}

	now := time.Now()
	mockClock.On("Now").Return(now)

	id := uuid.New()
	mockRepository.On("CreateTweet", models.Tweet{UserID: 1, Timestamp: now, Content: "aguante banfield"}).Return(id, nil)
	mockRepository.On("GetFollowers", uint(1)).Return(nil, errors.New("an error"))

	_, err := followService.Tweet(1, "aguante banfield")
	require.ErrorIs(t, err, service.ErrTweet)
	require.ErrorContains(t, err, "from user 1")
}

func TestTweetDoesNotReturnErrorIfRepositoryAddToTimelineReturnsError(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)
	mockClock := adaptersMocks.NewIClock(t)

	followService := service.TwitterService{
		Repository: mockRepository,
		Clock:      mockClock,
	}

	now := time.Now()
	mockClock.On("Now").Return(now)

	id := uuid.New()
	mockRepository.On("CreateTweet", models.Tweet{UserID: 1, Timestamp: now, Content: "aguante banfield"}).Return(id, nil)
	mockRepository.On("GetFollowers", uint(1)).Return([]uint{2, 3}, nil)
	mockRepository.On("AddTweetToTimeline", id, uint(2)).Return(errors.New("an error"))
	mockRepository.On("AddTweetToTimeline", id, uint(3)).Return(nil)

	tweetID, err := followService.Tweet(1, "aguante banfield")
	require.NoError(t, err)
	assert.Equal(t, id, tweetID)
}

func TestTweetReturnsTweetIDIfNoFollowers(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)
	mockClock := adaptersMocks.NewIClock(t)

	followService := service.TwitterService{
		Repository: mockRepository,
		Clock:      mockClock,
	}

	now := time.Now()
	mockClock.On("Now").Return(now)

	id := uuid.New()
	mockRepository.On("CreateTweet", models.Tweet{UserID: 1, Timestamp: now, Content: "aguante banfield"}).Return(id, nil)
	mockRepository.On("GetFollowers", uint(1)).Return([]uint{}, nil)

	tweetID, err := followService.Tweet(1, "aguante banfield")
	require.NoError(t, err)
	assert.Equal(t, id, tweetID)
}

func TestTweetAddsTweetToTimelineIfFollowers(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)
	mockClock := adaptersMocks.NewIClock(t)

	followService := service.TwitterService{
		Repository: mockRepository,
		Clock:      mockClock,
	}

	now := time.Now()
	mockClock.On("Now").Return(now)

	id := uuid.New()
	mockRepository.On("CreateTweet", models.Tweet{UserID: 1, Timestamp: now, Content: "aguante banfield"}).Return(id, nil)
	mockRepository.On("GetFollowers", uint(1)).Return([]uint{2, 3}, nil)
	mockRepository.On("AddTweetToTimeline", id, uint(2)).Return(nil)
	mockRepository.On("AddTweetToTimeline", id, uint(3)).Return(nil)

	tweetID, err := followService.Tweet(1, "aguante banfield")
	require.NoError(t, err)
	assert.Equal(t, id, tweetID)
}

func TestGetTimelineReturnsErrorIfRepositoryReturnsError(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)

	twService := service.TwitterService{
		Repository: mockRepository,
	}

	mockRepository.On("GetTimeline", uint(1)).Return(nil, errors.New("an error"))

	_, err := twService.GetTimeline(1)
	require.ErrorIs(t, err, service.ErrTimeline)
	require.ErrorContains(t, err, "of user 1")
}

func TestGetTimelineReturnsAListOfTweets(t *testing.T) {
	mockRepository := repositoryMocks.NewIRepository(t)

	twService := service.TwitterService{
		Repository: mockRepository,
	}

	id := uuid.New()
	tweets := []models.Tweet{{UserID: 1, Content: "hola"}}

	mockRepository.On("GetTimeline", uint(1)).Return([]uuid.UUID{id}, nil)
	mockRepository.On("GetTweets", []uuid.UUID{id}).Return(tweets, nil)

	tweetsReturned, err := twService.GetTimeline(1)
	require.NoError(t, err)
	assert.Equal(t, tweets, tweetsReturned)
}
