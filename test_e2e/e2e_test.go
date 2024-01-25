package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	messages "github.com/cucumber/messages/go/v21"
	"github.com/elliotchance/pie/v2"
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
	sc.Step(`^user (\d+) tweets "([^"]*)"$`, userTweets)
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

	return assertResponseCreated(resp)
}

// userTweets adds a tweet to userID
func userTweets(userID int, content string) error {
	requestBodyMap := map[string]string{
		"content": content,
	}

	requestBody, err := json.Marshal(requestBodyMap)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("http://localhost:8080/user/%d/tweet", userID),
		"application/json", bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return assertResponseCreated(resp)
}
