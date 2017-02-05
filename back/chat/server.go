package chat

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server structure
type Server struct{}

var Port string = "8000"
var Host string = "localhost"

func serveHome(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.JSON(404, gin.H{"message": "Not found"})
	}
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "index.html", "")
}

func (server *Server) NewEngine(port string) {
	// router
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("front/chat/index.html")
	router.Static("/src", "./front/chat/static/")

	// start chat hub
	go hub.run()

	// set api/handlers
	router.GET("/", serveHome)
	router.GET("/ws", serveWs)

	// start server
	router.Run(":" + port)
}

func Start(args []string) {
	// set port
	if len(args) > 3 {
		Port = args[3]
	}
	// set host
	if len(args) > 4 {
		Host = args[4]
	}

	// info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + Port)
	fmt.Println("Selected host: " + Host)
	fmt.Println("---------------")

	// star server
	thisServer := new(Server)
	thisServer.NewEngine(Port)
}
