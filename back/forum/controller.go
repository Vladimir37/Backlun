package forum

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var request RegistrationReq
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

	userExist := false
	for _, user := range UserList {
		if user.Login == request.Login {
			userExist = true
			break
		}
	}

	if userExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "User with this login is exists",
			"body":    nil,
		})
		return
	}

	var emptyVoted map[int]time.Time
	newUser := UserStruct{
		ID:         UserNum,
		Login:      request.Login,
		Password:   request.Password,
		Text:       request.Text,
		Token:      GenerateToken(25),
		PostCount:  0,
		Reputation: 0,
		Voted:      emptyVoted,
	}
	UserNum++

	UserList = append(UserList, newUser)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    newUser,
	})
}

func Login(c *gin.Context) {
	var request LoginReq
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

	founded := false
	var targetIndex int
	for index, user := range UserList {
		if (user.Login == request.Login) && (user.Password == request.Password) {
			founded = true
			targetIndex = index
			break
		}
	}

	if !founded {
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect login or password",
			"body":    nil,
		})
		return
	} else {
		UserList[targetIndex].Token = GenerateToken(25)
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    UserList[targetIndex].Token,
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

	targetIndex := CheckTokenUtility(request.Token)

	if targetIndex == -1 {
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	} else {
		UserList[targetIndex].Token = GenerateToken(25)
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    nil,
		})
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

	userIndex := CheckTokenUtility(request.Token)

	if userIndex == -1 {
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
