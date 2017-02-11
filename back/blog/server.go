package blog

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
	router.LoadHTMLGlob("front/blog/index.html")

	router.Static("/src", "./front/blog/static/")
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
	api := baseRouter.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/get", GetAuthData)
			auth.GET("/check", CheckToken)
			auth.POST("/login", Login)
			auth.POST("/logout", Logout)
		}

		get := api.Group("/get")
		{
			get.GET("/all", GetAllPosts)
			get.GET("/one", GetOnePost)
			get.GET("/search_tag", SearchTag)
			get.GET("/search_text", SearchText)
		}

		posts := api.Group("/posts")
		{
			posts.POST("/create", CreatePost)
			posts.POST("/edit", EditPost)
			posts.POST("/delete", DeletePost)
		}

		comments := api.Group("/comments")
		{
			comments.POST("/create", CreateComment)
			comments.POST("/delete", DeleteComment)
		}
	}
}
