package events

import "github.com/ergo-services/ergo/gen"

const NearbyQueryEvent gen.Event = "nearby_query"

type NearbyQueryEventMessage struct {
	Lon float64
	Lat float64
}
