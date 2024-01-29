package testintegration

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/uala_challenge/app/models"
	"github.com/FrancoLiberali/uala_challenge/app/repository"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

type IntTestSuite struct {
	suite.Suite
	rdb     *redis.Client
	service *service.TwitterService
	now     time.Time
}

func (ts *IntTestSuite) SetupTest() {
	CleanCache(ts.rdb)
}

func (ts *IntTestSuite) TestFollowCreatesSetIfIsTheFirstFollower() {
	err := ts.service.Follow(1, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(1), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1"}, followers)
}

func (ts *IntTestSuite) TestFollowAddsToSetIfAlreadyExists() {
	err := ts.service.Follow(1, 2)
	ts.Require().NoError(err)

	err = ts.service.Follow(3, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(2), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1", "3"}, followers)
}

func (ts *IntTestSuite) TestFollowNotAddIfAlreadyFollower() {
	err := ts.service.Follow(1, 2)
	ts.Require().NoError(err)

	err = ts.service.Follow(1, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(1), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1"}, followers)
}

func (ts *IntTestSuite) TestTweetCreatesTweetNoFollowers() {
	tweetID, err := ts.service.Tweet(1, "aguante banfield")
	ts.Require().NoError(err)

	tweet, err := ts.rdb.Get(context.Background(), repository.TweetKey(tweetID)).Result()
	ts.Require().NoError(err)

	var tweetStruct models.Tweet
	err = json.Unmarshal([]byte(tweet), &tweetStruct)
	ts.Require().NoError(err)

	ts.Equal(models.Tweet{
		UserID:    1,
		Timestamp: ts.now,
		Content:   "aguante banfield",
	}, tweetStruct)
}

func (ts *IntTestSuite) TestTweetCreatesTweetWithFollowersAddsToTimelines() {
	err := ts.service.Follow(2, 1)
	ts.Require().NoError(err)

	err = ts.service.Follow(3, 1)
	ts.Require().NoError(err)

	tweetID, err := ts.service.Tweet(1, "aguante banfield")
	ts.Require().NoError(err)

	timeline2, err := ts.service.Repository.GetTimeline(2)
	ts.Require().NoError(err)
	ts.Equal([]uuid.UUID{tweetID}, timeline2)

	timeline3, err := ts.service.Repository.GetTimeline(3)
	ts.Require().NoError(err)
	ts.Equal([]uuid.UUID{tweetID}, timeline3)
}

func (ts *IntTestSuite) TestMultipleTweetWithFollowersAddsToTimelines() {
	err := ts.service.Follow(2, 1)
	ts.Require().NoError(err)

	err = ts.service.Follow(3, 1)
	ts.Require().NoError(err)

	err = ts.service.Follow(3, 2)
	ts.Require().NoError(err)

	tweet1ID, err := ts.service.Tweet(1, "aguante banfield")
	ts.Require().NoError(err)

	tweet2ID, err := ts.service.Tweet(2, "aguante banfield")
	ts.Require().NoError(err)

	timeline2, err := ts.service.Repository.GetTimeline(2)
	ts.Require().NoError(err)
	ts.Equal([]uuid.UUID{tweet1ID}, timeline2)

	timeline3, err := ts.service.Repository.GetTimeline(3)
	ts.Require().NoError(err)
	ts.Equal([]uuid.UUID{tweet2ID, tweet1ID}, timeline3)
}

func (ts *IntTestSuite) TestTweetDeletesFromTimelineIfMaxReached() {
	err := ts.service.Follow(2, 1)
	ts.Require().NoError(err)

	tweetID, err := ts.service.Tweet(1, "aguante banfield")
	ts.Require().NoError(err)

	for i := 0; i <= repository.MaxTweetsInTimeline; i++ {
		_, err = ts.service.Tweet(1, "aguante banfield")
		ts.Require().NoError(err)
	}

	timeline2, err := ts.service.Repository.GetTimeline(2)
	ts.Require().NoError(err)
	ts.Len(timeline2, repository.MaxTweetsInTimeline)
	ts.NotContains(timeline2, tweetID)
}
