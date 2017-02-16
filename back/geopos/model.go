package geopos

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// ========== addition methods

// random string {{{
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	rndStr := make([]rune, n)
	for i := range rndStr {
		rndStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rndStr)
} // }}}

// ========== data section

type TokenReq struct {
	Token string `form:"token" binding:"required"`
}

type DistanceReq struct {
	Distance float64 `form:"distance" binding:"required"`
}

// GeoPoint for example {lat: 1.011111, lng: 1.0000450}
type GeoPoint struct {
	Type        string     `form:"-"`
	Token       string     `form:"token" binding:"required"`
	Coordinates [2]float64 `form:"coordinates" binding:"required"`
}

// GeoState is map(array) of points
type GeoState struct {
	Location map[string]GeoPoint `json:"location"`
	sync.RWMutex
}

// ========== GeoState methods

// NewGeoState will return a new state {{{
func NewGeoState() *GeoState {
	return &GeoState{
		Location: make(map[string]GeoPoint),
	}
} // }}}

// Add new point with token {{{
func (geost *GeoState) Add(point *GeoPoint) {
	geost.Lock()
	defer geost.Unlock()
	geost.Location[point.Token] = *point
} // }}}

// SetRnd fill GeoState the n points {{{
func (geost *GeoState) SetRnd(num int) {
	geost.Lock()
	defer geost.Unlock()

	point := new(GeoPoint)
	for i := 0; i < num; i++ {
		point.SetRnd()
		geost.Location[point.Token] = *point
	}
} // }}}

// Clear state {{{
func (geost *GeoState) Clear() {
	geost.Lock()
	defer geost.Unlock()

	geost.Location = make(map[string]GeoPoint)
} // }}}

// Len return lenght state {{{
func (geost *GeoState) Len() int {
	return len(geost.Location)
} // }}}

// Print print poinsts to a dafault stream {{{
func (geost *GeoState) Print() {
	fmt.Print(geost)
} // }}}

// GetPoint new point with token// {{{
func (geost *GeoState) GetPoint(token string) (point GeoPoint, ok bool) {
	geost.Lock()
	defer geost.Unlock()
	point, ok = geost.Location[token]
	return point, ok
} // }}}

// ========== GeoPoint methods

// GetDistance set random data to a point// {{{
func (point *GeoPoint) GetDistance(toPoint *GeoPoint) (distance float64) {
	distance = math.Sqrt(
		math.Pow(point.Coordinates[0]-toPoint.Coordinates[0], 2) +
			math.Pow(point.Coordinates[1]-toPoint.Coordinates[1], 2))
	return distance
} // }}}

// NewGeoPoint will return a new point {{{
func NewGeoPoint() *GeoPoint {
	point := new(GeoPoint)
	point.SetRnd()
	return point
} // }}}

// SetRnd set random data to a point// {{{
func (point *GeoPoint) SetRnd() {
	point.Type = "Point"
	point.Token = rndStr(8)
	point.Coordinates[0] = (rand.Float64() * 5) + 5
	point.Coordinates[1] = (rand.Float64() * 5) + 5
} // }}}
