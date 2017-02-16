package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	nrand "math/rand"
)

// DviUser structure for all data for the user
type User struct {
	ID       string `form:"id" binding:"required"`
	Username string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Descr    string `form:"description" binding:"required"`
}

type GhostDB struct {
	Users []User
}

// ========== addition methods

// random {{{
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
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
} // }}}

// ========== database methods

// SaveUser register a user so we know that we saw that user already.// {{{
func (db *GhostDB) SaveUser(user *User) {
	db.Users = append(db.Users, *user)
} // }}}

// LoadUser get data from a user.// {{{
func (db *GhostDB) LoadUser(Email string) (user User, err error) {
	for _, searchUser := range db.Users {
		if searchUser.Email == Email {
			return searchUser, nil
		}
	}
	return user, errors.New("user not found")
} // }}}

// PrintUsers print all users// {{{
func (ghostDB *GhostDB) PrintUsers() {
	fmt.Print(ghostDB)
} // }}}

// ========== user methods

// SetRnd set random data to user// {{{
func (user *User) SetRnd() {
	user.ID = rndStr(6)
	user.Email = rndStr(4) + "@" + rndStr(4) + ".com"
	user.Username = rndStr(4)
	user.Descr = rndStr(20)
} // }}}
