package oauth

import (
	"Backlun/back/conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// configure structurs
type Server struct{}

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

type Config struct {
	Port    string
	Host    string
	KeyFile string
	Cred    Credentials
}

// configure vars
var config *Config
var msgState *conf.MsgState
var confTemp *oauth2.Config
var ghostDB *GhostDB

func (cred *Credentials) SetFromFile(keyf string) *conf.ApiError { // {{{
	file, err := ioutil.ReadFile(keyf)
	if err != nil {
		return conf.NewApiError(err)
	}

	err = json.Unmarshal(file, &cred)
	if err != nil {
		return conf.NewApiError(err)
	}
	return conf.NewApiError(err)
} // }}}

func (config *Config) SetDefault() { // {{{
	config.Host = "localhost"
	config.Port = "8000"
	config.KeyFile = "key.json"
	config.Cred.Cid = "295529031882-ap6njd8e8p0bmggmvkb7t0iflhcetjn1.apps.googleusercontent.com"
	config.Cred.Csecret = "ICiVhKO51UxbNfIQVR7WudxH"
} // }}}

// ========== homepage

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

// ========== middlewares

// AuthorizeRequest is used to authorize a request for a certain end-point group.// {{{
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
} // }}}

// CORSMiddleware middleware witch headers for any RESTful requests// {{{
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
} // }}}

func noRoute(c *gin.Context) { // {{{
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, msgState.Errors[http.StatusNotFound])
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
} // }}}

// ========== server

// NewEngine return the new gin server// {{{
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

	// Current path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// all frontend
	router.LoadHTMLGlob(dir + "/front/oauth/index.html")
	router.Static("/src", dir+"/front/oauth/static/")
	router.StaticFile("/favicon.ico", dir+"/favicon/favicon.ico")

	// set specification
	router.GET("/", serveHome)

	// set api/handlers
	api := router.Group("api")
	{
		api.GET("/login", LoginHandler)
		api.GET("/auth", AuthHandler)
		//  group: here is API for authorized query
		authorized := router.Group("/lock")
		authorized.Use(AuthorizeRequest())
		{
			authorized.GET("/test", lockTest)
		}
	}

	// no route, bad url
	router.NoRoute(noRoute)
	router.Run(":" + port)
} // }}}

func Start(args []string) {
	// configure
	config = &Config{}
	config.SetDefault()
	ghostDB = &GhostDB{}
	msgState = conf.NewMsgState()
	msgState.SetErrors()

	if len(args) > 3 { // set port
		config.Port = args[3]
	}
	if len(args) > 4 { // set host
		config.Host = args[4]
	}
	if len(args) > 5 { // set key
		config.KeyFile = args[5]
	}
	err := config.Cred.SetFromFile(config.KeyFile)
	if err != nil {
		fmt.Println(err)
	}
	// init oauth config
	//scope: https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	confTemp = &oauth2.Config{
		ClientID:     config.Cred.Cid,
		ClientSecret: config.Cred.Csecret,
		RedirectURL:  "http://" + config.Host + ":" + config.Port + "/api/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + config.Port)
	fmt.Println("Selected host: " + config.Host)
	fmt.Println("Selected key file: " + config.KeyFile)
	fmt.Println("---------------")

	// star server
	thisServer := new(Server)
	thisServer.NewEngine(config.Port)
}
