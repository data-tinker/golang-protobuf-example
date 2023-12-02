package models

import (
	"golang-proto-example/userpb"
	"testing"

	"google.golang.org/protobuf/proto"
)

const (
	name             = "Alex"
	age              = 31
	youtubeFollowers = 1000
	twitterFollowers = 100
)

func TestCreateUser(t *testing.T) {
	expected := &userpb.User{
		Name: name,
		Age:  age,
		SocialFollowers: &userpb.SocialFollowers{
			Youtube: youtubeFollowers,
			Twitter: twitterFollowers,
		},
	}

	result := CreateUser(name, age, youtubeFollowers, twitterFollowers)

	if !proto.Equal(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestMarshalUnmarshalUser(t *testing.T) {
	user := CreateUser(name, age, youtubeFollowers, twitterFollowers)

	data, err := MarshalUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	result, err := UnmarshalUser(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !proto.Equal(result, user) {
		t.Errorf("Expected %+v, got %+v", user, result)
	}
}
