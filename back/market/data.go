package market

import (
	"math/rand"
	"time"

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

type OrderStruct struct {
	ID       int
	User     int
	Products []LotStruct
	Price    int
	Paid     bool
}

type UserStruct struct {
	ID       int
	Login    string
	Password string
	Address  string
	Money    int
	Backet   []LotStruct
	Token    string
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

// func checkTokenUtility(token string) bool {
// 	return token == AuthToken
// }

func GenerateCategories() {
	categoryNum := randomdata.Number(5, 12)
	for index := 0; index < categoryNum; index++ {
		newCategory := CategoryStruct{
			ID:   CategoryNum,
			Name: randomdata.Noun(),
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
				Name:         randomdata.Adjective() + " " + randomdata.Noun(),
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
