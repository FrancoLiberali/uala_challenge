package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

const (
	followerIDJSONKey = "followerID"
)

func HandleFollow(_ context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if event == nil {
		return nil, ErrNilEvent
	}

	body := map[string]string{}

	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return nil, err
	}

	followedID, err := getUserID(body, userIDJSONKey)
	if err != nil {
		return nil, err
	}

	followerID, err := getUserID(body, followerIDJSONKey)
	if err != nil {
		return nil, err
	}

	err = twService.Follow(followerID, followedID)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("User %d started to follow %d", followerID, followedID),
	}, nil
}

func getUserID(body map[string]string, jsonKey string) (uint, error) {
	userIDString, isPresent := body[jsonKey]
	if !isPresent {
		return 0, badRequest(jsonKey)
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return 0, badRequest(jsonKey)
	}

	return uint(userID), nil
}
