// Ualá challenge Twitter API
//
// This the swagger documentation for the API of the Ualá challenge developed by Franco Liberali.
// Check our project here: https://github.com/FrancoLiberali/uala_challenge
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0
//	Contact: Franco Liberali <franco.liberali@gmail.com> https://github.com/FrancoLiberali
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- text/plain
//	- application/json
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FrancoLiberali/uala_challenge/app"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

//go:generate swagger generate spec -o ./swagger.json

var twitterService *service.TwitterService

func init() {
	var err error

	twitterService, _, err = app.NewService()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := gin.Default()

	r.POST("/user/:userID/tweet", tweet)
	r.POST("/user/:userID/follower/:followerID", follow)
	r.GET("/user/:userID/timeline", timeline)

	log.Fatalln(r.Run())
}

// swagger:operation POST /user/{userID}/follower/{followerID} Follow
// Start following an user
// ---
// produces:
// - text/plain
// parameters:
//   - name: userID
//     in: path
//     description: The id of the user to be followed
//     required: true
//     type: uint
//   - name: followerID
//     in: path
//     description: The id of the user that starts to follow userID
//     required: true
//     type: uint
//
// responses:
//
//	'200':
//	    description: 'Correctly followed'
func follow(c *gin.Context) {
	followedID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		returnError(c, err)
		return
	}

	followerID, err := strconv.Atoi(c.Param("followerID"))
	if err != nil {
		returnError(c, err)
		return
	}

	err = twitterService.Follow(uint(followerID), uint(followedID))
	if err != nil {
		returnError(c, err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("%d started to follow %d", followerID, followedID))
}

type TweetRequestBody struct {
	Content string `json:"content"`
}

// swagger:operation POST /user/{userID}/tweet Tweet
// Creates a tweet.
// ---
// produces:
// - text/plain
// parameters:
//   - name: userID
//     in: path
//     description: The id of the author of the tweet
//     required: true
//     type: uint
//   - name: content
//     in: body
//     description: Content of the tweet
//     required: true
//     type: string
//
// responses:
//
//	'201':
//	    description: 'Tweet created'
func tweet(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		returnError(c, err)
		return
	}

	var requestBody TweetRequestBody

	if err = c.BindJSON(&requestBody); err != nil {
		returnError(c, err)
		return
	}

	tweetID, err := twitterService.Tweet(uint(userID), requestBody.Content)
	if err != nil {
		returnError(c, err)
		return
	}

	c.String(http.StatusCreated, fmt.Sprintf("%d tweet %s created", userID, tweetID))
}

// swagger:operation GET /user/{userID}/timeline Timeline
// Get the user's timelines
// ---
// produces:
// - application/json
// parameters:
//   - name: userID
//     in: path
//     description: The id of the user that is owner of the timeline
//     required: true
//     type: uint
//
// responses:
//
//	'200':
//	    description: 'Timeline obtained'
func timeline(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		returnError(c, err)
		return
	}

	timeline, err := twitterService.GetTimeline(uint(userID))
	if err != nil {
		returnError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, timeline)
}

func returnError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err.Error())
}
