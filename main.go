package main

import (
	"golang-proto-example/models"
	"log"
)

const (
	name             = "Alex"
	age              = 31
	youtubeFollowers = 1000
	twitterFollowers = 100
)

func main() {
	user := models.CreateUser(name, age, youtubeFollowers, twitterFollowers)

	data, err := models.MarshalUser(user)
	if err != nil {
		log.Fatal("Marshaling error: ", err)
	}

	readUser, err := models.UnmarshalUser(data)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	log.Println(readUser.GetName())
	log.Println(readUser.GetAge())
	log.Println(readUser.GetSocialFollowers().GetYoutube())
	log.Println(readUser.GetSocialFollowers().GetTwitter())
}
