package forum

import (
	"fmt"

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

	newUser := UserStruct{
		ID:         UserNum,
		Login:      request.Login,
		Password:   request.Password,
		Text:       request.Text,
		Token:      GenerateToken(25),
		PostCount:  0,
		Reputation: 0,
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

func GetAllCategories(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CategoryList,
	})
}

func GetAllThreads(c *gin.Context) {
	var request CategoryReq
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

	for _, category := range CategoryList {
		if category.ID == request.Category {
			founded = true
			break
		}
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  11,
			"message": "Category not found",
			"body":    nil,
		})
		return
	}

	var allThreads []ThreadStruct

	for _, thread := range ThreadList {
		if thread.Category == request.Category {
			allThreads = append(allThreads, thread)
		}
	}

	var publicThreadsList []PublicThreadStruct
	for _, thread := range allThreads {
		err, publicThread := ThreadToPublicThread(thread)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  8,
				"message": "Author not found",
				"body":    nil,
			})
			return
		} else {
			publicThreadsList = append(publicThreadsList, publicThread)
		}
	}

	c.JSON(200, gin.H{
		"status":  1,
		"message": "Success",
		"body":    publicThreadsList,
	})
}

func GetThread(c *gin.Context) {
	var request GetThreadReq
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

	for index, thread := range ThreadList {
		if thread.ID == request.Thread {
			founded = true
			targetIndex = index
		}
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  7,
			"message": "Thread not found",
			"body":    nil,
		})
		return
	}

	err, publicThread := ThreadToPublicThread(ThreadList[targetIndex])

	if err != nil {
		c.JSON(500, gin.H{
			"status":  8,
			"message": "Author not found",
			"body":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    publicThread,
	})
}

func GetUser(c *gin.Context) {
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
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    UserList[userIndex],
	})
}

func GetTargetUser(c *gin.Context) {
	var request IDReq
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

	var founded bool
	var targetUser PublicUserStruct

	for _, user := range UserList {
		if user.ID == request.ID {
			founded = true
			targetUser.ID = user.ID
			targetUser.Login = user.Login
			targetUser.Text = user.Text
			targetUser.PostCount = user.PostCount
			targetUser.Reputation = user.Reputation
		}
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "User not found",
			"body":    nil,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    targetUser,
		})
	}
}

func CreateThread(c *gin.Context) {
	var request ThreadReq
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
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	foundedCategory := false
	for _, category := range CategoryList {
		if category.ID == request.Category {
			foundedCategory = true
		}
	}

	if !foundedCategory {
		c.JSON(400, gin.H{
			"status":  4,
			"message": "Category not found",
			"body":    nil,
		})
		return
	}

	firstPost := PostStruct{
		ID:     0,
		Author: UserList[userIndex].ID,
		Text:   request.Text,
	}

	var firstPostArr []PostStruct
	firstPostArr = append(firstPostArr, firstPost)

	newThread := ThreadStruct{
		ID:       ThreadNum,
		Category: request.Category,
		Title:    request.Title,
		PostNum:  1,
		Posts:    firstPostArr,
	}

	ThreadList = append(ThreadList, newThread)
	UserList[userIndex].PostCount++

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}

func SendPost(c *gin.Context) {
	var request PostReq
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
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	var targetIndex int
	var founded bool

	for index, thread := range ThreadList {
		if (thread.Category == request.Category) && (thread.ID == request.Thread) {
			founded = true
			targetIndex = index
			break
		}
	}

	if !founded {
		c.JSON(403, gin.H{
			"status":  5,
			"message": "Category or thread not found",
			"body":    nil,
		})
		return
	}

	newPost := PostStruct{
		ID:     ThreadList[targetIndex].PostNum,
		Author: UserList[userIndex].ID,
		Text:   request.Text,
	}
	ThreadList[targetIndex].PostNum++
	ThreadList[targetIndex].Posts = append(ThreadList[targetIndex].Posts, newPost)
	UserList[userIndex].PostCount++

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}

func ChangeReputation(c *gin.Context) {
	var request ReputationReq
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
		c.JSON(403, gin.H{
			"status":  11,
			"message": "Incorrect token",
			"body":    nil,
		})
		return
	}

	var targetIndex int
	var founded bool

	for index, user := range UserList {
		if user.ID == request.Target {
			founded = true
			targetIndex = index
		}
	}

	if !founded {
		c.JSON(403, gin.H{
			"status":  6,
			"message": "User not found",
			"body":    nil,
		})
		return
	}

	if UserList[targetIndex].ID == UserList[userIndex].ID {
		c.JSON(403, gin.H{
			"status":  9,
			"message": "You can not change a reputation for yourself",
			"body":    nil,
		})
		return
	}

	if request.Inc {
		UserList[targetIndex].Reputation++
	} else {
		UserList[targetIndex].Reputation--
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}
