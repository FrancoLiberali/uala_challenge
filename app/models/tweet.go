package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Tweet struct {
	UserID    uint      `json:"userId"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}

const MaxTweetLen = 280 // characters

var ErrTweetTooLong = errors.New("tweets must have less than 280 characters")

// NewTweet creates a tweet.
//
// Returns ErrTweetTooLong if content len is bigger than 280 characters
func NewTweet(userID uint, timestamp time.Time, content string) (*Tweet, error) {
	if len(content) > MaxTweetLen {
		return nil, ErrTweetTooLong
	}

	return &Tweet{
		UserID:    userID,
		Timestamp: timestamp,
		Content:   content,
	}, nil
}

// Implement encoding.BinaryMarshaler
func (tweet Tweet) MarshalBinary() (data []byte, err error) {
	return json.Marshal(tweet)
}
