package forum

import (
	"math/rand"
	"time"
)

type UserStruct struct {
	ID         int
	Login      string
	Password   string
	Text       string
	Token      string
	PostCount  int
	Reputation int
	Voted      map[int]time.Time
}

type CategoryStruct struct {
	ID   int
	Name string
}

type ThreadStruct struct {
	ID       int
	Category int
	Title    string
	Posts    []PostStruct
}

type PostStruct struct {
	ID     int
	Thread int
	Author int
	Text   string
}

// Requests

type RegistrationReq struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
	Text     string `form:"text"`
}

type LoginReq struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type TokenReq struct {
	Token string `form:"token" binding:"required"`
}

type CategoryReq struct {
	Category int `form:"category" binding:"required"`
}

type GetThreadReq struct {
	Thread int `form:"thread" binding:"required"`
}

type ThreadReq struct {
	Token    string `form:"token" binding:"required"`
	Category int    `form:"category" binding:"required"`
	Title    string `form:"title" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

type PostReq struct {
	Token    string `form:"token" binding:"required"`
	Category int    `form:"category" binding:"required"`
	Thread   string `form:"thread" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

// Current

var UserList []UserStruct
var CategoryList []CategoryStruct
var ThreadList []ThreadStruct
var PostList []PostStruct

var UserNum int = 0
var CategoryNum int = 0
var ThreadNum int = 0
var PostNum int = 0

// Utility

func CheckTokenUtility(token string) int {
	targetIndex := -1
	for index, user := range UserList {
		if user.Token == token {
			targetIndex = index
			break
		}
	}
	return targetIndex
}

func GenerateToken(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
