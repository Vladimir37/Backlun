package blog

import (
	"fmt"
	"math/rand"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    PostList,
	})
}

func GetAuthData(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CurrentAuthData,
	})
}

func Login(c *gin.Context) {
	var request AuthDataReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if (request.Login != CurrentAuthData.Login) || (request.Password != CurrentAuthData.Password) {
		c.JSON(403, gin.H{
			"status":  2,
			"message": "Incorrect login or password",
			"body":    nil,
		})
		return
	} else {
		AuthToken = GenerateToken(12)
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    AuthToken,
		})
		return
	}
}

func Logout(c *gin.Context) {
	var request TokenReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if request.Token != AuthToken {
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	} else {
		AuthToken = GenerateToken(12)
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    nil,
		})
		return
	}
}

func CheckToken(c *gin.Context) {
	var request TokenReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if request.Token != AuthToken {
		c.JSON(200, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Token is correct",
			"body":    nil,
		})
		return
	}
}

func CreatePost(c *gin.Context) {
	var request PostStructReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if !checkTokenUtility(request.Token) {
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	var tags []string
	if len(request.Tags) > 0 {
		tags = strings.Split(request.Tags, " ")
	}

	var emptyComments []CommentStruct

	newPost := PostStruct{
		ID:               CurrentPostID,
		Title:            request.Title,
		Date:             time.Now(),
		Tags:             tags,
		Text:             request.Text,
		CurrentCommentID: 1,
		Comments:         emptyComments,
	}
	CurrentPostID++

	PostList = append(PostList, newPost)
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    newPost,
	})
}

func EditPost(c *gin.Context) {
	var request PostStructEditReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if !checkTokenUtility(request.Token) {
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	var founded bool = false
	var targetIndex int

	for index, post := range PostList {
		if post.ID == request.ID {
			founded = true
			targetIndex = index
		}
	}

	if founded {
		var tags []string
		if len(request.Tags) > 0 {
			tags = strings.Split(request.Tags, " ")
		}
		PostList[targetIndex].Title = request.Title
		PostList[targetIndex].Tags = tags
		PostList[targetIndex].Text = request.Text
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    PostList[targetIndex],
		})
	} else {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Post not found",
			"body":    nil,
		})
	}
}

// Utility

func checkTokenUtility(token string) bool {
	return token == AuthToken
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
