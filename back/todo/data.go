package todo

type ToDoStruct struct {
	ID       int    `form:"id"`
	Title    string `form:"title" binding:"required"`
	Text     string `form:"text" binding:"required"`
	Category int    `form:"category"`
	Status   bool   `form:"status"`
}

type CategoryStruct struct {
	ID    int    `form:"id"`
	Name  string `form:"name" binding:"required"`
	Color string `form:"color" binding:"required"`
}

// Requests

type IDRequest struct {
	ID int `form:"id" binding:"required"`
}

type ToDoStructEdit struct {
	ID       int    `form:"id" binding:"required"`
	Title    string `form:"title"`
	Text     string `form:"text"`
	Category int    `form:"category"`
	Status   bool   `form:"status"`
}

type CategoryStructEdit struct {
	ID    int    `form:"id" binding:"required"`
	Name  string `form:"name"`
	Color string `form:"color"`
}

// Current

var TasksList []ToDoStruct
var CategoriesList []CategoryStruct
var CurrentTaskID int = 1
var CurrentCategoryID int = 1
