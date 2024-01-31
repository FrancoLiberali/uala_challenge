package handlers

import (
	"context"
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

	followedID, err := getUserID(event, userIDJSONKey)
	if err != nil {
		return nil, err
	}

	followerID, err := getUserID(event, followerIDJSONKey)
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

func getUserID(event *events.APIGatewayV2HTTPRequest, paramKey string) (uint, error) {
	userIDString, isPresent := event.PathParameters[paramKey]
	if !isPresent {
		return 0, badRequest(paramKey)
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return 0, badRequest(paramKey)
	}

	return uint(userID), nil
}
