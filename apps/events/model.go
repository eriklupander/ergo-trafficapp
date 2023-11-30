package events

import (
	"strconv"
	"strings"
)

type VehiclePosition struct {
	ID   string
	Lat  float64
	Lon  float64
	Date int64
}

// GeoPos is the representation used by the clients, e.g. msgpack +> struct should use this struct.
type GeoPos struct {
	ID        string
	Lat       float64
	Lon       float64
	Emergency bool
}

func CoordToLatLon(val string) (float64, float64, bool) {
	if len(val) == 0 {
		return 0, 0, false
	}
	// Ugly, parse coords string
	coords := strings.Split(val[1:len(val)-1], " ")
	if len(coords) < 2 {
		return 0, 0, false
	}
	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return 0, 0, false
	}
	lon, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return 0, 0, false
	}
	return lat, lon, true
}
