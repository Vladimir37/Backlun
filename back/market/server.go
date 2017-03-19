package market

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ========== middleware

// CORSMiddleware middleware headers for any RESTful requests {{{
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
} // }}}

func Start(args []string) {
	// Selecting port
	port := "8000"
	if len(args) == 3 {
		port = args[2]
	}

	// Current path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + port)
	fmt.Println("---------------")

	// Creating router
	router := gin.Default()
	router.LoadHTMLGlob(dir + "/front/market/index.html")

	router.Static("/src", dir+"/front/market/static/")
	router.StaticFile("/favicon.ico", dir+"/favicon/favicon.ico")

	// add headers middleware
	router.Use(CORSMiddleware())

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
			c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
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
