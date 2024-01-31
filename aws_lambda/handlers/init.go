package handlers

import (
	"log"

	"github.com/FrancoLiberali/uala_challenge/app"
	"github.com/FrancoLiberali/uala_challenge/app/service"
)

var twService *service.TwitterService

func init() {
	var err error

	twService, _, err = app.NewService()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
