package market

import (
	"math/rand"
	"time"

	"errors"

	"strings"

	"github.com/Pallinder/go-randomdata"
)

type ProductStruct struct {
	ID           int
	Name         string
	Manufacturer string
	Price        int
	Description  string
	Count        int
	Category     int
}

type CategoryStruct struct {
	ID   int
	Name string
}

type LotStruct struct {
	Product int
	Count   int
}

type BacketLotStruct struct {
	Product ProductStruct
	Count   int
}

type OrderStruct struct {
	ID       int
	User     int
	Date     time.Time
	Products []LotStruct
	Price    int
	Paid     bool
}

type UserStruct struct {
	ID       int
	Login    string
	Password string
	FullName string
	Address  string
	Money    int
	Backet   []LotStruct
	Token    string
}

// Requests

type RegistrationReq struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
	FullName string `form:"name" binding:"required"`
	Address  string `form:"address" binding:"required"`
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

type AddCreditReq struct {
	Token  string `form:"token" binding:"required"`
	Credit int    `form:"credit" binding:"required"`
}

type ProductReq struct {
	Token   string `form:"token" binding:"required"`
	Product int    `form:"product" binding:"required"`
	Count   int    `form:"count" binding:"required"`
}

// Current

var UserList []UserStruct
var OrderList []OrderStruct
var CategoryList []CategoryStruct
var ProductList []ProductStruct

// Numbers

var CategoryNum int = 1
var ProductNum int = 1
var OrderNum int = 1
var UserNum int = 1

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

func GetProductUtility(id int) (ProductStruct, error) {
	founded := false
	var targetProduct ProductStruct

	for index, product := range ProductList {
		if product.ID == id {
			founded = true
			targetProduct = ProductList[index]
			break
		}
	}

	if founded {
		return targetProduct, nil
	} else {
		return targetProduct, errors.New("Not found")
	}
}

func GenerateCategories() {
	categoryNum := randomdata.Number(5, 12)
	for index := 0; index < categoryNum; index++ {
		newCategory := CategoryStruct{
			ID:   CategoryNum,
			Name: strings.Title(randomdata.Noun()),
		}
		CategoryNum++
		CategoryList = append(CategoryList, newCategory)
	}
	GenerateProducts()
}

func GenerateProducts() {
	for _, category := range CategoryList {
		productsNum := randomdata.Number(3, 18)
		for i := 0; i < productsNum; i++ {
			newProduct := ProductStruct{
				ID:           ProductNum,
				Name:         strings.Title(randomdata.Adjective()) + " " + randomdata.Noun(),
				Manufacturer: randomdata.Country(randomdata.FullCountry),
				Price:        randomdata.Number(2, 160),
				Description:  randomdata.Paragraph(),
				Count:        randomdata.Number(50, 200),
				Category:     category.ID,
			}
			ProductNum++
			ProductList = append(ProductList, newProduct)
		}
	}
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
