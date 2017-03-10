package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// AuthHandler handles authentication of a user and initiates a session.
func AuthHandler(cont *gin.Context) { // {{{
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(cont)
	retrievedState := session.Get("state")
	queryState := cont.Request.URL.Query().Get("state")
	fmt.Print("retrievedState:\n")
	fmt.Print(retrievedState)
	fmt.Print("\nqueryState:\n")
	fmt.Print(queryState)
	fmt.Print("\n")

	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		cont.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session state."})
		return
	}

	code := cont.Request.URL.Query().Get("code")
	tok, err := confTemp.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		cont.JSON(http.StatusBadRequest, gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := confTemp.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		cont.AbortWithStatus(http.StatusBadRequest)
		return
	}

	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	user := User{}
	if err = json.Unmarshal(data, &user); err != nil {
		log.Println(err)
		cont.JSON(http.StatusBadRequest, gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}

	session.Set("user-id", user.Email)
	err = session.Save()
	if err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	fmt.Printf("\nuser: %v\n", user)
	log.Println(err)

	seen := false
	// db := GhostDB{}
	if _, mongoErr := ghostDB.LoadUser(user.Email); mongoErr == nil {
		seen = true
	} else {
		ghostDB.SaveUser(&user)
	}

	cont.JSON(http.StatusOK, gin.H{"email": user.Email, "seen": seen})
} // }}}

// LoginHandler handles the login procedure.
func LoginHandler(cont *gin.Context) { // {{{

	// session
	state := RandToken(32)
	session := sessions.Default(cont)
	session.Set("state", state)
	session.Save()

	scopes := strings.Join(confTemp.Scopes, " ")
	linkStr := string(confTemp.Endpoint.AuthURL +
		"?client_id=" + confTemp.ClientID +
		"&redirect_uri=" + confTemp.RedirectURL +
		"&response_type=code&scope=" + scopes +
		"&state=" + state)

	cont.JSON(http.StatusOK, gin.H{"link": linkStr})
} // }}}

// FieldHandler is a rudementary handler for logged in users.
func FieldHandler(c *gin.Context) { // {{{
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.JSON(http.StatusOK, gin.H{"user": userID})
} // }}}

func lockTest(c *gin.Context) {
	c.JSON(http.StatusOK, msgState.Messages[http.StatusOK])
}
