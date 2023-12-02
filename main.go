package main

import (
	"golang-proto-example/person"
	"log"

	"google.golang.org/protobuf/proto"
)

const (
	name             = "Alex"
	age              = 31
	youtubeFollowers = 1000
	twitterFollowers = 100
)

func main() {
	user := person.Person{
		Name: name,
		Age:  age,
		SocialFollowers: &person.SocialFollowers{
			Youtube: youtubeFollowers,
			Twitter: twitterFollowers,
		},
	}

	data, err := proto.Marshal(&user)
	if err != nil {
		log.Fatal("Marshaling error: ", err)
	}

	readUser := &person.Person{}
	err = proto.Unmarshal(data, readUser)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	log.Println(readUser.GetName())
	log.Println(readUser.GetAge())
	log.Println(readUser.GetSocialFollowers().GetYoutube())
	log.Println(readUser.GetSocialFollowers().GetTwitter())
}
