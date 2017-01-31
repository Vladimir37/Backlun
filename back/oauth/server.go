package oauth

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Server for you
type Server struct{}

var Port string = "8000"

func testData(cont *gin.Context) {
	cont.JSON(200, gin.H{"message: ": "test data"})
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
	router.Use(static.Serve("/", static.LocalFile("./front/oauth", true)))

	// login
	router.GET("/login", LoginHandler)
	router.GET("/auth", AuthHandler)

	// v1 group: here is API for authorized query
	authorized := router.Group("/v1")
	authorized.Use(AuthorizeRequest())
	{
		authorized.GET("/test", testData)
	}

	router.Run(":" + port)
}

func Start(args []string) {
	// Set port
	if len(args) == 4 {
		Port = args[3]
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + Port)
	fmt.Println("---------------")

	// Star server
	thisServer := new(Server)
	thisServer.NewEngine(Port)
}
