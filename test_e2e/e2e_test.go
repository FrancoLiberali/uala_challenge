package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	messages "github.com/cucumber/messages/go/v21"
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

const executorID = 1

func init() {
	opts := godog.Options{Output: colors.Colored(os.Stdout)}
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^I follow users$`, iFollowUsers)
}

// Takes a list of users to follow and starts to follow them
func iFollowUsers(users *godog.Table) error {
	userIDs := pie.Map(users.Rows, func(row *messages.PickleTableRow) string {
		return row.Cells[0].Value
	})

	for _, userID := range userIDs {
		err := follow(userID)
		if err != nil {
			return err
		}
	}

	return nil
}

// makes executor user to follow userID
func follow(userID string) error {
	resp, err := http.Post(
		fmt.Sprintf("http://localhost:8080/user/%s/follower/%d", userID, executorID),
		"", nil,
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = assertExpectedAndActual(assert.Equal, http.StatusCreated, resp.StatusCode)
	if err != nil {
		return err
	}

	return nil
}
