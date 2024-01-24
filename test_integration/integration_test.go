package testintegration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/uala_challenge/app"
)

const (
	host     = "localhost:6379"
	password = "uala_challenge2024"
)

func TestMain(t *testing.T) {
	t.Setenv(app.CacheURLEnvVar, host)
	t.Setenv(app.CachePasswordEnvVar, password)

	followService, rdb, err := app.NewFollowService()
	require.NoError(t, err)

	suite.Run(t, &IntTestSuite{rdb: rdb, followService: followService})
}
