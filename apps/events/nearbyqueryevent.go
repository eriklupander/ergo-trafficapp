package events

import "github.com/ergo-services/ergo/gen"

const NearbyQueryEvent gen.Event = "nearby_query"

type NearbyQueryEventMessage struct {
	Lat float64
	Lon float64
}
