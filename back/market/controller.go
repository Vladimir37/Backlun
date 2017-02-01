package market

import (
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    ProductList,
	})
}

func GetOneProduct(c *gin.Context) {
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

	product, err := GetProductUtility(request.ID)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  10,
			"message": "Product not found",
			"body":    nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    product,
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

	var emptyBacket []LotStruct
	newUser := UserStruct{
		ID:       UserNum,
		Login:    request.Login,
		Password: request.Password,
		FullName: request.FullName,
		Address:  request.Address,
		Money:    0,
		Backet:   emptyBacket,
		Token:    GenerateToken(25),
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

	founded := false
	var targetIndex int
	for index, user := range UserList {
		if user.Token == request.Token {
			founded = true
			targetIndex = index
			break
		}
	}

	if !founded {
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

func AddCredit(c *gin.Context) {
	var request AddCreditReq
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

	if request.Credit < 1 {
		c.JSON(400, gin.H{
			"status":  4,
			"message": "Incorrect credit",
			"body":    nil,
		})
		return
	}

	UserList[userIndex].Money += request.Credit

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
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
	} else {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    UserList[userIndex],
		})
		return
	}
}

func GetAllOrders(c *gin.Context) {
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

	var allOrders []OrderStruct

	for _, order := range OrderList {
		if order.User == UserList[userIndex].ID {
			allOrders = append(allOrders, order)
		}
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    allOrders,
	})
}

func PayOrder(c *gin.Context) {
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

	if len(UserList[userIndex].Backet) == 0 {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Backet is empty",
			"body":    nil,
		})
		return
	}

	sum := 0
	for _, productElem := range UserList[userIndex].Backet {
		product, err := GetProductUtility(productElem.Product)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  10,
				"message": "Product in backet not found",
				"body":    nil,
			})
			return
		}
		sum += productElem.Count * product.Price
		if productElem.Count > product.Count {
			c.JSON(400, gin.H{
				"status":  4,
				"message": "Not enough products in stock",
				"body":    nil,
			})
			return
		}
	}

	newOrder := OrderStruct{
		ID:       OrderNum,
		User:     UserList[userIndex].ID,
		Date:     time.Now(),
		Products: UserList[userIndex].Backet,
		Price:    sum,
		Paid:     false,
	}

	if sum > UserList[userIndex].Money {
		c.JSON(400, gin.H{
			"status":  5,
			"message": "Not enough money",
			"body":    nil,
		})
		return
	}

	newOrder.Paid = true
	OrderNum++
	OrderList = append(OrderList, newOrder)
	UserList[userIndex].Money -= sum

	c.JSON(400, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}
