package calendar

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	var ShortEventsList map[time.Time][]ShortEventExpanded
	var LongEventsList map[time.Time][]LongEventExpanded

	// Short categories
	for date, _ := range AllShortEvents {
		var emptyPointEvents []ShortEventExpanded
		ShortEventsList[date] = emptyPointEvents

		for _, targetShortEvent := range AllShortEvents[date] {
			ShortEventsList[date] = append(ShortEventsList[date], GetShortCategory(targetShortEvent))
		}
	}

	// Long categories
	for date, _ := range LongEventsList {
		var emptyPointEvents []LongEventExpanded
		LongEventsList[date] = emptyPointEvents

		for _, targetLongEvent := range AllLongEvents[date] {
			LongEventsList[date] = append(LongEventsList[date], GetLongCategory(targetLongEvent))
		}
	}

	fullList := map[string]interface{}{
		"short": ShortEventsList,
		"long":  LongEventsList,
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    fullList,
	})
}

func GetShortEvents(c *gin.Context) {
	var ShortEventsList map[time.Time][]ShortEventExpanded

	for date, _ := range AllShortEvents {
		var emptyPointEvents []ShortEventExpanded
		ShortEventsList[date] = emptyPointEvents

		for _, targetShortEvent := range AllShortEvents[date] {
			ShortEventsList[date] = append(ShortEventsList[date], GetShortCategory(targetShortEvent))
		}
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    ShortEventsList,
	})
}

func GetLongEvents(c *gin.Context) {
	var LongEventsList map[time.Time][]LongEventExpanded

	for date, _ := range LongEventsList {
		var emptyPointEvents []LongEventExpanded
		LongEventsList[date] = emptyPointEvents

		for _, targetLongEvent := range AllLongEvents[date] {
			LongEventsList[date] = append(LongEventsList[date], GetLongCategory(targetLongEvent))
		}
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    LongEventsList,
	})
}
