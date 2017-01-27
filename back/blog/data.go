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

// Current

var CurrentAuthData AuthData = AuthData{
	Login:    randomdata.FirstName(randomdata.Female),
	Password: GenerateToken(5),
}
var AuthToken string = GenerateToken(12)
var CurrentPostID int = 1
var PostList []PostStruct
