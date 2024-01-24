package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FrancoLiberali/uala_challenge/app"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

var followService *service.FollowService

func init() {
	var err error

	followService, _, err = app.NewFollowService()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := gin.Default()

	r.POST("user/:followedID/follower/:followerID", follow)

	log.Fatalln(r.Run())
}

func follow(c *gin.Context) {
	followedID, err := strconv.Atoi(c.Param("followedID"))
	if err != nil {
		returnError(c, err)
		return
	}

	followerID, err := strconv.Atoi(c.Param("followerID"))
	if err != nil {
		returnError(c, err)
		return
	}

	err = followService.Follow(uint(followerID), uint(followedID))
	if err != nil {
		returnError(c, err)
		return
	}

	c.String(http.StatusCreated, "OK")
}

func returnError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err.Error())
}
