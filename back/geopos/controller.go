package geopos

import (
	"github.com/Vladimir37/Backlun/back/conf"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPoints get all points
func GetPoints(c *gin.Context) { // {{{
	c.JSON(http.StatusOK, conf.GiveResponse(geoState.Location))
} // }}}

func PostPoint(c *gin.Context) {
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, msgState.Errors[http.StatusBadRequest])
		return
	}

	geoState.Add(&request)

	c.JSON(http.StatusOK, conf.GiveResponse(request))
}

func GetRndPoint(c *gin.Context) { // {{{
	var request GeoPoint
	request.SetRnd()

	c.JSON(http.StatusOK, conf.GiveResponse(request))
} // }}}

func PostRndPoint(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, msgState.Errors[http.StatusBadRequest])
		return
	}

	geoState.Add(&request)

	c.JSON(http.StatusOK, conf.GiveResponse(request))
} // }}}

func GetPointFromToken(c *gin.Context) { // {{{

	token := c.Request.URL.Query().Get("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, msgState.Errors[http.StatusBadRequest])
		return
	}

	if point, ok := geoState.GetPoint(token); ok {
		c.JSON(http.StatusOK, conf.GiveResponse(point))
	} else {
		c.JSON(http.StatusNotFound, msgState.Errors[http.StatusNotFound])
	}
} // }}}

func PutDistance(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, msgState.Errors[http.StatusBadRequest])
		return
	}

	geoState.Add(&request)
	distance := checkPoint.GetDistance(&request)

	c.JSON(http.StatusOK, conf.GiveResponse(distance))
} // }}}

func GetCheckPoint(c *gin.Context) { // {{{
	c.JSON(http.StatusOK, conf.GiveResponse(checkPoint))
} // }}}

func PostCheckPoint(c *gin.Context) { // {{{
	var request GeoPoint

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, msgState.Errors[http.StatusBadRequest])
		return
	}

	checkPoint = &request

	c.JSON(http.StatusOK, conf.GiveResponse(checkPoint))
} // }}}
