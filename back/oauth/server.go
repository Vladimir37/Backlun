package oauth

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Server structure
type Server struct{}

var Host string = "localhost"
var Port string = "8000"
var KeyFile string = "key.json"

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

func lockData(cont *gin.Context) {
	cont.JSON(200, gin.H{"message: ": "here is locked data, you have got auth data!"})
}

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(thisContext *gin.Context) {
		session := sessions.Default(thisContext)
		v := session.Get("user-id")
		if v == nil {
			thisContext.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			thisContext.Abort()
		}
		thisContext.Next()
	}
}

// CORSMiddleware middleware witch headers for any RESTful requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// NewEngine return the new gin server
func (server *Server) NewEngine(port string) {
	router := gin.Default()

	// support sessions
	store := sessions.NewCookieStore([]byte(RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	// router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("goquestsession", store))
	// add headers middleware
	router.Use(CORSMiddleware())

	// all frontend
	router.LoadHTMLGlob("front/oauth/index.html")
	router.Static("/src", "./front/oauth/static/")

	// set api/handlers
	router.GET("/", serveHome)
	router.GET("/login", LoginHandler)
	router.GET("/auth", AuthHandler)

	// v1 group: here is API for authorized query
	authorized := router.Group("/v1")
	authorized.Use(AuthorizeRequest())
	{
		authorized.GET("/test", lockData)
	}

	router.Run(":" + port)
}

func Start(args []string) {
	if len(args) > 2 { // set port
		Port = args[2]
	} else if len(args) > 3 { // set host
		Host = args[3]
	} else if len(args) > 4 { // set key
		KeyFile = args[4]
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
