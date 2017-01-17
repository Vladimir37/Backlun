package todo

type ToDo struct {
	id       int
	title    string
	text     string
	category int
	status   bool
}

type Category struct {
	id    int
	name  string
	color string
}

var TasksList []ToDo
var CategoriesList []Category
