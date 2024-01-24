package testintegration

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/uala_challenge/app/repository"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

type IntTestSuite struct {
	suite.Suite
	rdb           *redis.Client
	followService *service.FollowService
}

func (ts *IntTestSuite) SetupTest() {
	CleanCache(ts.rdb)
}

func (ts *IntTestSuite) TestFollowCreatesListIfIsTheFirstFollower() {
	err := ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	followers, err := ts.rdb.LRange(context.Background(), repository.UserFollowersKey(2), 0, -1).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1"}, followers)
}

func (ts *IntTestSuite) TestFollowAddsToListIfAlreadyExists() {
	err := ts.followService.Follow(1, 2)
	ts.Require().NoError(err)

	err = ts.followService.Follow(3, 2)
	ts.Require().NoError(err)

	followers, err := ts.rdb.LRange(context.Background(), repository.UserFollowersKey(2), 0, -1).Result()
	ts.Require().NoError(err)
	ts.Equal([]string{"1", "3"}, followers)
}
