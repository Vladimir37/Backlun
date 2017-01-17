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

	request.Status = true
	request.ID = CurrentTaskID
	CurrentTaskID++

	TasksList = append(TasksList, request)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
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

	founded := false

	for index, task := range TasksList {
		if task.ID == request.ID {
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
