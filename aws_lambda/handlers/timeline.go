package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleTimeline(_ context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if event == nil {
		return nil, ErrNilEvent
	}

	body := map[string]string{}

	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return nil, err
	}

	userID, err := getUserID(body, userIDJSONKey)
	if err != nil {
		return nil, err
	}

	timeline, err := twService.GetTimeline(userID)
	if err != nil {
		return nil, err
	}

	timelineString, err := json.Marshal(timeline)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(timelineString),
	}, nil
}
