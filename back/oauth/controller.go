package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"dvij.geoloc/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var cred Credentials
var confTemp *oauth2.Config

func getLoginURL(state string) string { // {{{
	return confTemp.AuthCodeURL(state)
} // }}}

func init() { // {{{
	// file, err := ioutil.ReadFile("keys/clientid.google.json")
	// if err != nil {
	// log.Printf("File error: %v\n", err)
	// os.Exit(1)
	// }
	// json.Unmarshal(file, &cred)
	cred.Cid = "295529031882-ap6njd8e8p0bmggmvkb7t0iflhcetjn1.apps.googleusercontent.com"
	cred.Csecret = "ICiVhKO51UxbNfIQVR7WudxH"

	confTemp = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://localhost" + Port + "/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
} // }}}

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

	// seen := false
	// db := GhostDB{}
	// if _, mongoErr := db.LoadUser(user.Email); mongoErr == nil {
	// seen = true
	// } else {
	// err = db.SaveUser(&user)
	// if err != nil {
	// log.Println(err)
	// cont.JSON(http.StatusBadRequest, gin.H{"message": "Error while saving user. Please try again."})
	// return
	// }
	// }

	cont.JSON(http.StatusOK, gin.H{"email": user.Email, "seen": true})
} // }}}

// LoginHandler handles the login procedure.
func LoginHandler(cont *gin.Context) { // {{{

	// session
	state := models.RandToken(32)
	session := sessions.Default(cont)
	session.Set("state", state)
	session.Save()

	// response
	link := getLoginURL(state)
	cont.JSON(http.StatusOK, gin.H{
		"auth_url":     confTemp.Endpoint.AuthURL,
		"client_id":    confTemp.ClientID,
		"redirect_uri": confTemp.RedirectURL,
		"scope":        strings.Join(confTemp.Scopes, " "),
		"state":        state,
		"link":         link,
	})
} // }}}

// FieldHandler is a rudementary handler for logged in users.
func FieldHandler(cont *gin.Context) { // {{{
	session := sessions.Default(cont)
	userID := session.Get("user-id")
	cont.JSON(http.StatusOK, gin.H{"user": userID})
} // }}}
