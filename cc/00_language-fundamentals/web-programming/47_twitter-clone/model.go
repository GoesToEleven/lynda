package main

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn      bool
	LoginFail     bool
	Tweets        []Tweet
	ViewingUser   string
	FollowingUser bool
	Following     []F
}

type Tweet struct {
	Msg      string
	Time     time.Time
	UserName string
}

type F struct {
	Following string
	Follower  string
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
