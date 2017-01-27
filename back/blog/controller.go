package blog

import (
	"fmt"
	"math/rand"
	"time"

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
	//
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
