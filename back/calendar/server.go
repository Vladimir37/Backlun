package calendar

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
	router.LoadHTMLGlob("front/calendar/index.html")

	router.Static("/src", "./front/calendar/static/")
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
		get := api.Group("/get")
		{
			get.GET("/all", GetAllEvents)
			get.GET("/short", GetShortEvents)
			get.GET("/long", GetLongEvents)
			get.GET("/category_events", GetCategoryEvents)
			get.GET("/all_categories", GetAllCategories)
			get.GET("/day", GetDayData)
		}

		short := api.Group("/short")
		{
			short.POST("/create", CreateShortEvent)
			short.POST("/edit", EditShortEvent)
			short.POST("/delete", DeleteShortEvent)
		}

		long := api.Group("/long")
		{
			long.POST("/create", CreateLongEvent)
			long.POST("/edit", EditLongEvent)
			long.POST("/delete", DeleteLongEvent)
		}

		categories := api.Group("/categories")
		{
			categories.POST("/create", CreateCategory)
			categories.POST("/edit", EditCategory)
			categories.POST("/delete", DeleteCategory)
		}
	}
}
