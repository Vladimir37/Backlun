package blog

import (
	"time"

	"github.com/Pallinder/go-randomdata"
)

type AuthData struct {
	Login    string
	Password string
}

type PostStruct struct {
	ID               int
	Title            string
	Date             time.Time
	Tags             []string
	Text             string
	CurrentCommentID int
	Comments         []CommentStruct
}

type CommentStruct struct {
	ID   int
	Name string
	Date time.Time
	Text string
}

// Requests

type AuthDataReq struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type TokenReq struct {
	Token string `form:"token" binding:"required"`
}

type TagReq struct {
	Tag string `form:"tag" binding:"required"`
}

type IDReq struct {
	ID int `form:"id" binding:"required"`
}

type IDTokenReq struct {
	ID    int    `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
}

type CommentTokenReq struct {
	Post    int    `form:"post" binding:"required"`
	Comment int    `form:"comment" binding:"required"`
	Token   string `form:"token" binding:"required"`
}

type PostStructReq struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
	Tags  string `form:"tags"`
	Text  string `form:"text" binding:"required"`
}

type PostStructEditReq struct {
	ID    int    `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
	Tags  string `form:"tags"`
	Text  string `form:"text" binding:"required"`
}

type CommentStructReq struct {
	Post int    `form:"post" binding:"required"`
	Name string `form:"name"`
	Text string `form:"text" binding:"required"`
}

// Current

var CurrentAuthData AuthData = AuthData{
	Login:    randomdata.FirstName(randomdata.Female),
	Password: GenerateToken(5),
}
var AuthToken string = GenerateToken(12)
var CurrentPostID int = 1
var PostList []PostStruct
