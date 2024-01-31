package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const (
	userIDJSONKey  = "userID"
	contentJSONKey = "content"
)

func HandleTweet(_ context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if event == nil {
		return nil, ErrNilEvent
	}

	userID, err := getUserID(event, userIDJSONKey)
	if err != nil {
		return nil, err
	}

	body := map[string]string{}

	err = json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return nil, err
	}

	content, isPresent := body[contentJSONKey]
	if !isPresent {
		return nil, badRequest(contentJSONKey)
	}

	tweetID, err := twService.Tweet(uint(userID), content)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       "Created tweet " + tweetID.String(),
	}, nil
}
