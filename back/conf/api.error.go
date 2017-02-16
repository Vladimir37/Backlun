package conf

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "os"
)

// ApiError structure for processing errors
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Href    string `json:"href"`
}

// NewApiError
func NewApiError(err error) *ApiError {
	if err == nil {
		return ErrNoError
	} else {
		return &ApiError{http.StatusInternalServerError, err.Error(), ""}
	}
}

func (err *ApiError) Error() string {
	return err.Message
}

func (err *ApiError) PrintError() {
	fmt.Printf("\ncode: %d, %s", err.Code, err.Message)
	if err.Href != "" {
		fmt.Print(err.Href)
	}
}

var (
	ErrNoError = &ApiError{0, "No error.", ""}
	ErrJSON    = &ApiError{http.StatusInternalServerError, "error of json.marshal.", ""}
)

// type Header map[string]interface{}

type Header gin.H

type MsgState struct {
	Errors   map[int]gin.H
	Messages map[int]gin.H
}

// NewMsgState will return a new state
func NewMsgState() *MsgState {
	return &MsgState{
		Errors:   make(map[int]gin.H),
		Messages: make(map[int]gin.H),
	}
}

// SetRnd fill GeoState the n points
func (errst *MsgState) SetErrors() {
	// good and info answer
	errst.Messages[http.StatusOK] = gin.H{"status": 0, "message": "correct data", "body": nil}
	// bad and error answer
	errst.Errors[http.StatusBadRequest] = gin.H{"status": 1, "message": "Bad request, incorrect data", "body": nil}
	errst.Errors[http.StatusInternalServerError] = gin.H{"status": 5, "message": "Internal error, server can't return correct data", "body": nil}
	errst.Errors[http.StatusMethodNotAllowed] = gin.H{"status": 1, "message": "Method not allowed", "body": nil}
	errst.Errors[http.StatusNotFound] = gin.H{"status": 3, "message": "Request not found", "body": nil}
	errst.Errors[http.StatusRequestTimeout] = gin.H{"status": 10, "message": "url req not found", "body": nil}
	errst.Errors[http.StatusNotImplemented] = gin.H{"status": 10, "message": "req not implemented", "body": nil}
	errst.Errors[http.StatusNoContent] = gin.H{"status": 3, "message": "empty content", "body": nil}
}

func GiveResponse(some interface{}) (response *gin.H) {
	return &gin.H{"status": 0, "message": "Success", "body": some}
}
