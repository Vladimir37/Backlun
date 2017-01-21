package todo

import "github.com/gin-gonic/gin"
import "fmt"

func GetAllTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    TasksList,
	})
}

func AddNewTask(c *gin.Context) {
	var request ToDoStruct
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if request.Category != 0 {
		checkCategory := categoryExist(request.Category)
		if !checkCategory {
			c.JSON(400, gin.H{
				"status":  3,
				"message": "Category not found",
				"body":    nil,
			})
			return
		}
	}

	request.Status = true
	request.ID = CurrentTaskID
	CurrentTaskID++

	TasksList = append(TasksList, request)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    request,
	})
}

func EditTask(c *gin.Context) {
	var request ToDoStructEdit
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if request.Category != 0 {
		checkCategory := categoryExist(request.Category)
		if !checkCategory {
			c.JSON(400, gin.H{
				"status":  3,
				"message": "Category not found",
				"body":    nil,
			})
			return
		}
	}

	founded := false
	var targetIndex int

	for index, task := range TasksList {
		if task.ID == request.ID {
			targetIndex = index
			TasksList[index].Title = request.Title
			TasksList[index].Text = request.Text
			TasksList[index].Category = request.Category
			TasksList[index].Status = request.Status
			founded = true
		}
	}

	if founded {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    TasksList[targetIndex],
		})
	} else {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Task not found",
			"body":    nil,
		})
	}
}

func DeleteTask(c *gin.Context) {
	var request IDRequest
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	founded := false

	for index, task := range TasksList {
		if task.ID == request.ID {
			TasksList = append(TasksList[:index], TasksList[index+1:]...)
			founded = true
		}
	}

	if founded {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    nil,
		})
	} else {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Task not found",
			"body":    nil,
		})
	}
}

func GetAllCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CategoriesList,
	})
}

func AddNewCatergory(c *gin.Context) {
	var request CategoryStruct
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	request.ID = CurrentCategoryID
	CurrentCategoryID++

	CategoriesList = append(CategoriesList, request)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    request,
	})
}

func EditCategory(c *gin.Context) {
	var request CategoryStructEdit
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	founded := false
	var targetIndex int

	for index, category := range CategoriesList {
		if category.ID == request.ID {
			targetIndex = index
			CategoriesList[index].Name = request.Name
			CategoriesList[index].Color = request.Color
			founded = true
		}
	}

	if founded {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    CategoriesList[targetIndex],
		})
	} else {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Task not found",
			"body":    nil,
		})
	}
}

func DeleteCategory(c *gin.Context) {
	var request IDRequest
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	categoryTask := false

	for _, task := range TasksList {
		if task.Category == request.ID {
			categoryTask = true
		}
	}

	if categoryTask {
		c.JSON(400, gin.H{
			"status":  4,
			"message": "There are tasks with this category - category can not be deleted",
			"body":    nil,
		})
		return
	}

	founded := false

	for index, category := range CategoriesList {
		if category.ID == request.ID {
			CategoriesList = append(CategoriesList[:index], CategoriesList[index+1:]...)
			founded = true
		}
	}

	if founded {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    nil,
		})
	} else {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Task not found",
			"body":    nil,
		})
	}
}

func categoryExist(num int) bool {
	founded := false

	for _, category := range CategoriesList {
		if category.ID == num {
			founded = true
		}
	}

	return founded
}
