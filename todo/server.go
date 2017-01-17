package todo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Start(args []string) {
	// Selecting port
	port := "8000"
	if len(args) >= 4 {
		port = args[3]
	}

	// Selecting language
	lang := "en"
	if len(args) >= 5 {
		lang = args[4]
	}
	if (lang != "en") && (lang != "ru") {
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Incorrect language - must be \"en\" or \"ru\"")
		fmt.Println("---------------")
		return
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + port)
	fmt.Println("Selected language: " + lang)
	fmt.Println("---------------")

	// Creating router
	router := gin.Default()
	router.LoadHTMLGlob("front/todo/index.html")

	router.Static("/src", "./front/todo/static/")
	getApiRouter(router)

	router.NoRoute(func(c *gin.Context) {
		path := strings.Split(c.Request.URL.Path, "/")
		if (path[1] != "") && (path[1] == "api") {
			c.JSON(404, gin.H{
				"status": 10,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", "")
		}
	})

	// Starting
	router.Run(":" + port)
}

func getApiRouter(baseRouter *gin.Engine) {
	//
}
