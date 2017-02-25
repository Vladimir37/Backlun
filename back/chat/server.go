package chat

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

var msgState *conf.MsgState

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

func (server *Server) NewEngine(port string) {
	// router
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("front/chat/index.html")
	router.Static("/src", "./front/chat/static/")

	// specification page
	router.GET("/", serveHome)

	// start chat hub
	hub := newHub()
	go hub.run()

	// set api/handlers
	api := router.Group("/api")
	{
		api.GET("/ws", func(c *gin.Context) { serveWs(hub, c) })
	}

	// no route, bad url
	router.NoRoute(noRoute)

	// start server
	router.Run(":" + port)
}

func Start(args []string) {
	// configure
	config := Config{}
	config.SetDefault()
	msgState = conf.NewMsgState()
	msgState.SetErrors()

	if len(args) > 2 { // set port
		config.Port = args[2]
	}
	if len(args) > 2 { // set port
		config.Host = args[2]
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
