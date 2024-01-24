package testintegration

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/uala_challenge/app/models"
	"github.com/FrancoLiberali/uala_challenge/app/repository"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

type IntTestSuite struct {
	suite.Suite
	rdb           *redis.Client
	followService *service.TwitterService
	now           time.Time
}

func (ts *IntTestSuite) SetupTest() {
	CleanCache(ts.rdb)
}

func (ts *IntTestSuite) TestFollowCreatesSetIfIsTheFirstFollower() {
	err := ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(1), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1"}, followers)
}

func (ts *IntTestSuite) TestFollowAddsToSetIfAlreadyExists() {
	err := ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	err = ts.followService.Follow(3, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(2), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1", "3"}, followers)
}

func (ts *IntTestSuite) TestFollowNotAddIfAlreadyFollower() {
	err := ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	err = ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	followersLen, err := ts.rdb.SCard(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal(int64(1), followersLen)

	followers, err := ts.rdb.SMembers(context.Background(), repository.UserFollowersKey(2)).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1"}, followers)
}

func (ts *IntTestSuite) TestTweetCreatesTweet() {
	tweetID, err := ts.followService.Tweet(1, "aguante banfield")
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
