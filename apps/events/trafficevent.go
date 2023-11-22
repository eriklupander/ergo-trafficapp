package events

import "github.com/ergo-services/ergo/gen"

const TrafficEvent gen.Event = "traffic"

type TrafficEventMessage struct {
	Date    int64
	Payload []byte
}
