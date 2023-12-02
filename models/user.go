package models

import (
	"golang-protobuf-example/userpb"

	"google.golang.org/protobuf/proto"
)

func CreateUser(name string, age int32, youtubeFollowers int32, twitterFollowers int32) *userpb.User {
	return &userpb.User{
		Name: name,
		Age:  age,
		SocialFollowers: &userpb.SocialFollowers{
			Youtube: youtubeFollowers,
			Twitter: twitterFollowers,
		},
	}
}

func MarshalUser(user *userpb.User) ([]byte, error) {
	return proto.Marshal(user)
}

func UnmarshalUser(data []byte) (*userpb.User, error) {
	readUser := &userpb.User{}
	err := proto.Unmarshal(data, readUser)
	return readUser, err
}
