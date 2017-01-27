package blog

import "time"

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
	name string
	date time.Time
	text string
}

var CurrentPostID int = 1
