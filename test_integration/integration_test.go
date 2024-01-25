package testintegration

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/uala_challenge/app"
	adaptersMocks "github.com/FrancoLiberali/uala_challenge/app/mocks/adapters"
)

const (
	host     = "localhost:6379"
	password = "uala_challenge2024"
)

func TestMain(t *testing.T) {
	t.Setenv(app.CacheURLEnvVar, host)
	t.Setenv(app.CachePasswordEnvVar, password)

	followService, rdb, err := app.NewService()
	require.NoError(t, err)

	// mock time
	mockClock := adaptersMocks.NewIClock(t)

	var nowDecoded time.Time

	// marshal and unmarshal time to ensure its equal to value unmarshalled from cache
	nowEncoded, err := json.Marshal(time.Now())
	require.NoError(t, err)
	err = json.Unmarshal(nowEncoded, &nowDecoded)
	require.NoError(t, err)

	mockClock.On("Now").Return(nowDecoded).Maybe()

	followService.Clock = mockClock

	suite.Run(t, &IntTestSuite{rdb: rdb, followService: followService, now: nowDecoded})
}
