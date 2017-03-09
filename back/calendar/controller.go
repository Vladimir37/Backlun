package calendar

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	// Short categories
	var ShortEventList []ShortEventExpanded

	for _, targetShortEvent := range AllShortEvents {
		ShortEventList = append(ShortEventList, GetShortCategory(targetShortEvent))
	}

	// Long categories
	var LongEventList []LongEventExpanded

	for _, targetLongEvent := range AllLongEvents {
		LongEventList = append(LongEventList, GetLongCategory(targetLongEvent))
	}

	fullList := map[string]interface{}{
		"short": ShortEventList,
		"long":  LongEventList,
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    fullList,
	})
}

func GetShortEvents(c *gin.Context) {
	var ShortEventList []ShortEventExpanded

	for _, targetShortEvent := range AllShortEvents {
		ShortEventList = append(ShortEventList, GetShortCategory(targetShortEvent))
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    ShortEventList,
	})
}

func GetLongEvents(c *gin.Context) {
	var LongEventList []LongEventExpanded

	for _, targetLongEvent := range AllLongEvents {
		LongEventList = append(LongEventList, GetLongCategory(targetLongEvent))
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    LongEventList,
	})
}

func GetCategoryEvents(c *gin.Context) {
	var request IDReq
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

	var ShortEventList []ShortEventExpanded
	var LongEventList []LongEventExpanded

	// Short events
	for _, event := range AllShortEvents {
		if event.Category == request.ID {
			targetEvent := GetShortCategory(event)
			ShortEventList = append(ShortEventList, targetEvent)
		}
	}

	// Long events
	for _, event := range AllLongEvents {
		if event.Category == request.ID {
			targetEvent := GetLongCategory(event)
			LongEventList = append(LongEventList, targetEvent)
		}
	}

	fullList := map[string]interface{}{
		"short": ShortEventList,
		"long":  LongEventList,
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    fullList,
	})
}

func GetDayData(c *gin.Context) {
	//
}

func CreateShortEvent(c *gin.Context) {
	var request NewShortEventReq
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

	categoryExist := CheckCategoryExist(request.Category)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	newShortEvent := ShortEvent{
		ID:          CurrentShortEventID,
		Category:    request.Category,
		Title:       request.Title,
		Description: request.Description,
		Time:        request.Time,
	}

	AllShortEvents = append(AllShortEvents, newShortEvent)

	CurrentShortEventID++

	expandedNewEvent := GetShortCategory(newShortEvent)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    expandedNewEvent,
	})
}

func EditShortEvent(c *gin.Context) {
	var request EditShortEventReq
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

	categoryExist := CheckCategoryExist(request.Category)

	founded, index := FindShortEvent(request.ID)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllShortEvents[index].Category = request.Category
	AllShortEvents[index].Description = request.Description
	AllShortEvents[index].Time = request.Time
	AllShortEvents[index].Title = request.Title

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    AllShortEvents[index],
	})
}

func DeleteShortEvent(c *gin.Context) {
	var request IDReq
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

	founded, index := FindShortEvent(request.ID)

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllShortEvents = append(AllShortEvents[:index], AllShortEvents[index+1:]...)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}

func CreateLongEvent(c *gin.Context) {
	var request NewLongEventReq
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

	categoryExist := CheckCategoryExist(request.Category)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	newLongEvent := LongEvent{
		ID:          CurrentShortEventID,
		Category:    request.Category,
		Title:       request.Title,
		Description: request.Description,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
	}

	AllLongEvents = append(AllLongEvents, newLongEvent)

	CurrentLongEventID++

	expandedNewEvent := GetLongCategory(newLongEvent)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    expandedNewEvent,
	})
}

func EditLongEvent(c *gin.Context) {
	var request EditLongEventReq
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

	categoryExist := CheckCategoryExist(request.Category)

	founded, index := FindLongEvent(request.ID)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllLongEvents[index].Category = request.Category
	AllLongEvents[index].Description = request.Description
	AllLongEvents[index].StartTime = request.StartTime
	AllLongEvents[index].EndTime = request.EndTime
	AllLongEvents[index].Title = request.Title

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    AllLongEvents[index],
	})
}

func DeleteLongEvent(c *gin.Context) {
	var request IDReq
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

	founded, index := FindLongEvent(request.ID)

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllLongEvents = append(AllLongEvents[:index], AllLongEvents[index+1:]...)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}

func CreateCategory(c *gin.Context) {
	var request NewCategoryReq
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

	newCategory := Category{
		ID:    CurrentCategoryID,
		Name:  request.Name,
		Color: request.Color,
	}

	AllCategories = append(AllCategories, newCategory)

	CurrentCategoryID++

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    newCategory,
	})
}

func EditCategory(c *gin.Context) {
	var request EditCategoryReq
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

	exists, index := FindCategory(request.ID)

	if !exists {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	AllCategories[index].Name = request.Name
	AllCategories[index].Color = request.Color

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    AllCategories[index],
	})
}

func DeleteCategory(c *gin.Context) {
	var request IDReq
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

	for _, event := range AllShortEvents {
		if event.Category == request.ID {
			founded = true
			break
		}
	}

	for _, event := range AllLongEvents {
		if event.Category == request.ID {
			founded = true
			break
		}
	}

	if founded {
		c.JSON(400, gin.H{
			"status":  4,
			"message": "There are events with this category - category can not be deleted",
			"body":    nil,
		})
		return
	}

	exists, index := FindCategory(request.ID)

	if !exists {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	AllCategories = append(AllCategories[:index], AllCategories[index+1:]...)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    nil,
	})
}
