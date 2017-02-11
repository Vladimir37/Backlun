package geopos

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Server structure
type Server struct{}

type Config struct {
	Port string
	Host string
}

func (config *Config) SetDefault() { // {{{
	config.Port = "8000"
	config.Host = "localhost"
} // }}}

// home whitch specification {{{
func serveHome(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.JSON(404, gin.H{"message": "Not found"})
	}
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "index.html", "")
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

	// add headers middleware
	router.Use(CORSMiddleware())

	// all frontend
	router.LoadHTMLGlob("front/geopos/index.html")
	router.Static("/src", "./front/geopos/static/")

	// set api/handlers
	router.GET("/", serveHome)

	router.GET("/getPoints", GetPoints)
	router.GET("/getPointOnToken", GetPointOnToken)

	router.POST("/postPoint", PostPoint)

	router.PUT("/putDistance", PutDistance)
	router.GET("/getCheckPoint", GetCheckPoint)
	router.POST("/postCheckPoint", PostCheckPoint)

	router.GET("/getRndPoint", GetRndPoint)
	router.POST("/postRndPoint", PostRndPoint)

	// start server
	router.Run(":" + port)
} // }}}

// Start will start new server {{{
func Start(args []string) {
	config := Config{}
	config.SetDefault()

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
} // }}}
