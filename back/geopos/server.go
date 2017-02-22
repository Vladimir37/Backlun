package geopos

import (
	"Backlun/back/conf"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Server structure
type Server struct{}

type Config struct {
	Port string
	Host string
}

// declare states
var msgState *conf.MsgState
var geoState *GeoState
var checkPoint *GeoPoint

func (config *Config) SetDefault() { // {{{
	config.Port = "8000"
	config.Host = "localhost"
} // }}}

func serveHome(c *gin.Context) { // {{{
	if c.Request.URL.Path != "/" {
		c.JSON(http.StatusNotFound, msgState.Errors[http.StatusNotFound])
	}

	if c.Request.Method != "GET" {
		c.JSON(http.StatusMethodNotAllowed, msgState.Errors[http.StatusMethodNotAllowed])
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	c.HTML(http.StatusOK, "index.html", "")
} // }}}

func noRoute(c *gin.Context) { // {{{
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, msgState.Errors[http.StatusNotFound])
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
} // }}}

// CORSMiddleware middleware witch headers for any RESTful requests {{{
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

// NewEngine will return gin.server {{{
func (server *Server) NewEngine(port string) {
	router := gin.Default()

	// router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// all frontend
	router.LoadHTMLGlob("front/geopos/index.html")
	router.Static("/src", "./front/geopos/static/")

	// set api/handlers
	router.GET("/", serveHome)

	api := router.Group("api")
	{
		// add headers middleware
		// api.Use(CORSMiddleware())

		// points
		points := api.Group("points")
		{
			points.GET("/get", GetPoints)
			points.GET("/get_from_token", GetPointFromToken)
			points.POST("/post", PostPoint)
		}
		random_point := api.Group("random_point")
		{
			random_point.GET("/get", GetRndPoint)
			random_point.POST("/post", PostRndPoint)
		}
		check_point := api.Group("check_point")
		{
			check_point.PUT("/put_distance", PutDistance)
			check_point.GET("/get", GetCheckPoint)
			check_point.POST("/post", PostCheckPoint)
		}
	}

	// no route, bad url
	router.NoRoute(noRoute)
	// start server
	router.Run(":" + port)
} // }}}

// Start will start new server
func Start(args []string) {
	// config
	config := Config{}
	config.SetDefault()
	msgState = conf.NewMsgState()
	msgState.SetErrors()
	geoState = NewGeoState()
	checkPoint = NewGeoPoint()

	if len(args) > 3 { // set port
		config.Port = args[3]
	}
	if len(args) > 4 { // set host
		config.Host = args[4]
	}

	// info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + config.Port)
	fmt.Println("Selected host: " + config.Host)
	fmt.Println("---------------")

	// star server
	thisServer := new(Server)
	thisServer.NewEngine(config.Port)
}
