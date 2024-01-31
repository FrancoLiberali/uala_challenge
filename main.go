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

	c.String(http.StatusCreated, fmt.Sprintf("%d started to follow %d", followerID, followedID))
}

type TweetRequestBody struct {
	Content string `json:"content"`
}

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
