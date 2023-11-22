package events

type VehiclePosition struct {
	ID   string
	Lon  float64
	Lat  float64
	Date int64
}

// GeoPos is the representation used by the clients, e.g. msgpack +> struct should use this struct.
type GeoPos struct {
	ID        string
	Lon       float64
	Lat       float64
	Emergency bool
}
