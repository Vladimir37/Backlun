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
	if len(args) == 4 {
		port = args[3]
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + port)
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
		tasks := api.Group("/tasks")
		{
			tasks.GET("/get_all", GetAllTasks)
			tasks.POST("/add", AddNewTask)
			tasks.POST("/edit", EditTask)
			tasks.POST("/delete", DeleteTask)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/get_all", GetAllCategory)
			categories.POST("/add", AddNewCatergory)
			categories.POST("/edit", EditCategory)
			categories.POST("/delete", DeleteCategory)
		}
	}
}
