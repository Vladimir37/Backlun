package geopos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPoints get all points
func GetPoints(c *gin.Context) { // {{{
	if len(geostate.Location) > 0 {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    geostate.Location,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"status":  3,
			"message": "Geostate is empty",
			"body":    nil,
		})
	}
} // }}}

func PostPoint(c *gin.Context) { // {{{
	var point GeoPoint
	err := c.BindJSON(&point)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	geostate.Add(&point)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    point,
	})
} // }}}

func GetRndPoint(c *gin.Context) { // {{{
	var request GeoPoint
	request.SetRnd()

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    request,
	})
} // }}}

func PostRndPoint(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	geostate.Add(&request)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    request,
	})
} // }}}

func GetPointOnToken(c *gin.Context) { // {{{

	token := c.Request.URL.Query().Get("token")
	if token == "" {
		c.JSON(403, gin.H{
			"status":  1,
			"message": "Incorrect request params",
			"body":    nil,
		})
		return
	}
	fmt.Printf("\n## get point: %s\n", token)

	if point, ok := geostate.GetPoint(token); ok {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "Success",
			"body":    point,
		})
	} else {
		c.JSON(403, gin.H{
			"status":  13,
			"message": "point didn't found",
			"body":    nil,
		})
	}
} // }}}

func PutDistance(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect request",
			"body":    nil,
		})
		return
	}

	geostate.Add(&request)
	distance := checkPoint.GetDistance(&request)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    distance,
	})
} // }}}

func GetCheckPoint(c *gin.Context) { // {{{
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    checkPoint,
	})
} // }}}

func PostCheckPoint(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	checkPoint = &request

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    checkPoint,
	})
} // }}}
