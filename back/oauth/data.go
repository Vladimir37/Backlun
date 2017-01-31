package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	nrand "math/rand"
)

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// DviUser structure for all data for the user
type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Descr    string `json:"description"`
}

type GhostDB struct {
	Users []User
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RndStr(n int) string {
	rnd_str := make([]rune, n)
	for i := range rnd_str {
		rnd_str[i] = letterRunes[nrand.Intn(len(letterRunes))]
	}
	return string(rnd_str)
}

// RandToken generates a random @length token.
func RandToken(length int) string {
	thisByte := make([]byte, length)
	rand.Read(thisByte)
	return base64.StdEncoding.EncodeToString(thisByte)
}

// SaveUser register a user so we know that we saw that user already.
func (db *GhostDB) SaveUser(user *User) {
	db.Users = append(db.Users, *user)
}

// LoadUser get data from a user.
func (db *GhostDB) LoadUser(Email string) (user User, err error) {
	for _, searchUser := range db.Users {
		if searchUser.Email == Email {
			return searchUser, nil
		}
	}
	return user, errors.New("user not found")
}

// PrintUsers print all users
func (ghostDB *GhostDB) PrintUsers() {
	fmt.Print(ghostDB)
}

// SetRnd set random data to user
func (user *User) SetRnd() {
	user.ID = RndStr(6)
	user.Email = RndStr(4) + "@" + RndStr(4) + ".com"
	user.Username = RndStr(4)
	user.Descr = RndStr(20)
}
