package calendar

import "time"

type ShortEvent struct {
	ID          int
	Category    int
	Title       string
	Description string
	Time        time.Time
}

type LongEvent struct {
	ID          int
	Category    int
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type Category struct {
	ID    int
	Name  string
	Color string
}

// Expanded category

type ShortEventExpanded struct {
	ID          int
	Category    Category
	Title       string
	Description string
	Time        time.Time
}

type LongEventExpanded struct {
	ID          int
	Category    Category
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

// Requests

type NewCategoryReq struct {
	Name  string `form:"name" binding:"required"`
	Color string `form:"color" binding:"required"`
}

type EditCategoryReq struct {
	ID    int    `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Color string `form:"color" binding:"required"`
}

type NewShortEventReq struct {
	Category    int       `form:"category"`
	Title       string    `form:"title" binding:"required"`
	Description string    `form:"description" binding:"required"`
	Time        time.Time `form:"time" binding:"required"`
}

type EditShortEventReq struct {
	ID          int       `form:"id" binding:"required"`
	Category    int       `form:"category"`
	Title       string    `form:"title" binding:"required"`
	Description string    `form:"description" binding:"required"`
	Time        time.Time `form:"time" binding:"required"`
}

type NewLongEventReq struct {
	Title       string    `form:"title" binding:"required"`
	Category    int       `form:"category"`
	Description string    `form:"description" binding:"required"`
	StartTime   time.Time `form:"start_time" binding:"required"`
	EndTime     time.Time `form:"end_time" binding:"required"`
}

type EditLongEventReq struct {
	ID          int       `form:"id" binding:"required"`
	Category    int       `form:"category"`
	Title       string    `form:"title" binding:"required"`
	Description string    `form:"description" binding:"required"`
	StartTime   time.Time `form:"start_time" binding:"required"`
	EndTime     time.Time `form:"end_time" binding:"required"`
}

type IDReq struct {
	ID int `form:"id" binding:"required"`
}

type TimeReq struct {
	Time time.Time `form:"day" binding:"required"`
}

// Current

var CurrentShortEventID int = 1
var CurrentLongEventID int = 1
var CurrentCategoryID int = 1

var AllShortEvents []ShortEvent
var AllLongEvents []LongEvent

var AllCategories []Category

// Utlity

func GetShortCategory(event ShortEvent) ShortEventExpanded {
	founded := false
	targetIndex := 0
	emptyCategory := Category{}

	expandedEvent := ShortEventExpanded{
		ID:          event.ID,
		Category:    emptyCategory,
		Title:       event.Title,
		Description: event.Description,
		Time:        event.Time,
	}

	if event.Category == 0 {
		return expandedEvent
	}

	for index, category := range AllCategories {
		if category.ID == event.Category {
			founded = true
			targetIndex = index
			break
		}
	}

	if founded {
		expandedEvent.Category = AllCategories[targetIndex]
	}

	return expandedEvent
}

func GetLongCategory(event LongEvent) LongEventExpanded {
	founded := false
	targetIndex := 0
	emptyCategory := Category{}

	expandedEvent := LongEventExpanded{
		ID:          event.ID,
		Category:    emptyCategory,
		Title:       event.Title,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
	}

	if event.Category == 0 {
		return expandedEvent
	}

	for index, category := range AllCategories {
		if category.ID == event.Category {
			founded = true
			targetIndex = index
			break
		}
	}

	if founded {
		expandedEvent.Category = AllCategories[targetIndex]
	}

	return expandedEvent
}

func FindShortEvent(targetEvent int) (bool, int) {
	founded := false
	targetIndex := 0

	for index, event := range AllShortEvents {
		if event.ID == targetEvent {
			founded = true
			targetIndex = index
			break
		}
	}

	if founded {
		return true, targetIndex
	} else {
		return false, targetIndex
	}
}

func FindLongEvent(targetEvent int) (bool, int) {
	founded := false
	targetIndex := 0

	for index, event := range AllLongEvents {
		if event.ID == targetEvent {
			founded = true
			targetIndex = index
			break
		}
	}

	if founded {
		return true, targetIndex
	} else {
		return false, targetIndex
	}
}

func CheckCategoryExist(targetCategory int) bool {
	founded := false

	if targetCategory == 0 {
		return true
	}

	for _, category := range AllCategories {
		if category.ID == targetCategory {
			founded = true
			break
		}
	}

	return founded
}

func FindCategory(num int) (bool, int) {
	founded := false
	targetIndex := 0

	for index, category := range AllCategories {
		if category.ID == num {
			founded = true
			targetIndex = index
			break
		}
	}

	return founded, targetIndex
}

func GetStartDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetEndDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}
