package forum

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

type UserStruct struct {
	ID         int
	Login      string
	Password   string
	Text       string
	Token      string
	PostCount  int
	Reputation int
}

type PublicUserStruct struct {
	ID         int
	Login      string
	Text       string
	PostCount  int
	Reputation int
}

type CategoryStruct struct {
	ID   int
	Name string
}

type ThreadStruct struct {
	ID       int
	Category int
	Title    string
	PostNum  int
	Posts    []PostStruct
}

type PostStruct struct {
	ID     int
	Author int
	Text   string
}

type PublicThreadStruct struct {
	ID       int
	Category int
	Title    string
	PostNum  int
	Posts    []PublicPostStruct
}

type PublicPostStruct struct {
	ID     int
	Author PublicUserStruct
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

type IDReq struct {
	ID int `form:"id" binding:"required"`
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
	Thread   int    `form:"thread" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

type ReputationReq struct {
	Token  string `form:"token" binding:"required"`
	Target int    `form:"target" binding:"required"`
	Inc    bool   `form:"inc"`
}

// Current

var UserList []UserStruct
var CategoryList []CategoryStruct
var ThreadList []ThreadStruct

var UserNum int = 1
var CategoryNum int = 1
var ThreadNum int = 1

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

func GetUserUtility(id int) (error, PublicUserStruct) {
	var founded bool
	var targetIndex int
	for index, user := range UserList {
		if id == user.ID {
			founded = true
			targetIndex = index
			break
		}
	}

	if !founded {
		var emptyUser PublicUserStruct
		return errors.New("User not found"), emptyUser
	}

	targetUser := PublicUserStruct{
		ID:         UserList[targetIndex].ID,
		Login:      UserList[targetIndex].Login,
		Text:       UserList[targetIndex].Text,
		PostCount:  UserList[targetIndex].PostCount,
		Reputation: UserList[targetIndex].Reputation,
	}
	return nil, targetUser
}

func ThreadToPublicThread(thread ThreadStruct) (error, PublicThreadStruct) {
	var newPosts []PublicPostStruct
	for _, post := range thread.Posts {
		err, user := GetUserUtility(post.Author)
		if err != nil {
			var emptyThread PublicThreadStruct
			return errors.New("Author not found"), emptyThread
		}
		newPost := PublicPostStruct{
			ID:     post.ID,
			Author: user,
			Text:   post.Text,
		}
		newPosts = append(newPosts, newPost)
	}

	targetThread := PublicThreadStruct{
		ID:       thread.ID,
		Category: thread.Category,
		Title:    thread.Title,
		PostNum:  thread.PostNum,
		Posts:    newPosts,
	}

	return nil, targetThread
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

func GenerateCategory() {
	categoryNum := randomdata.Number(4, 12)
	for index := 0; index < categoryNum; index++ {
		newCategory := CategoryStruct{
			ID:   CategoryNum,
			Name: strings.Title(randomdata.Noun()),
		}
		CategoryNum++
		CategoryList = append(CategoryList, newCategory)
	}
}
