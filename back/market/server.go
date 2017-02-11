package market

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Start(args []string) {
	// Selecting port
	port := "8000"
	if len(args) == 3 {
		port = args[2]
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + port)
	fmt.Println("---------------")

	// Creating router
	router := gin.Default()
	router.LoadHTMLGlob("front/market/index.html")

	router.Static("/src", "./front/market/static/")
	getApiRouter(router)

	router.NoRoute(func(c *gin.Context) {
		path := strings.Split(c.Request.URL.Path, "/")
		if (path[1] != "") && (path[1] == "api") {
			c.JSON(404, gin.H{
				"status":  10,
				"message": "Not found",
				"body":    nil,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", "")
		}
	})

	// Starting
	router.Run(":" + port)
}

func getApiRouter(baseRouter *gin.Engine) {
	GenerateCategories()

	api := baseRouter.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/check", CheckToken)
			auth.POST("/registration", Registration)
			auth.POST("/login", Login)
			auth.POST("/logout", Logout)
		}

		get := api.Group("/get")
		{
			get.GET("/products", GetAllProducts)
			get.GET("/product", GetOneProduct)
			get.GET("/categories", GetAllCategories)
			get.GET("/orders", GetAllOrders)
			get.GET("/backet", GetBasket)
			get.GET("/user", GetUser)
		}

		market := api.Group("/market")
		{
			market.POST("/credits", AddCredit)
			market.POST("/products", ProductInBacket)
			market.POST("/pay", PayOrder)
		}
	}
}
