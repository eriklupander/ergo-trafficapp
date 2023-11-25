package trafficapp

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"log/slog"
	"time"
	"traffic/apps/events"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func createEmergency() gen.ServerBehavior {
	return &Emergency{}
}

type Emergency struct {
	gen.Server
}

// Init invoked on a start this process.
func (s *Emergency) Init(process *gen.ServerProcess, args ...etf.Term) error {
	slog.Info(fmt.Sprintf("Init process: %s with name %q and args %v", process.Self(), process.Name(), args))
	if err := process.MonitorEvent(events.TrafficEvent); err != nil {
		slog.Error("can't monitor event", slog.Any("error", err))
	}
	return nil
}

//
// Methods below are optional, so you can remove those that aren't be used
//

// HandleInfo invoked if this process received message sent with Process.Send(...).
func (s *Emergency) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("[Emergency] HandleInfo")

	trafficEvt, ok := message.(events.TrafficEventMessage)
	if !ok {
		slog.Error("Emergency: type conversion from etf.Term to events.TrafficEventMessage failed", slog.Any("type", fmt.Sprintf("%T", message)))
		return gen.ServerStatusIgnore
	}

	// Unmarshal from msgpack
	update := events.GeoPos{}
	if err := msgpack.Unmarshal(trafficEvt.Payload, &update); err != nil {
		slog.Error(err.Error())
		return gen.ServerStatusIgnore
	}

	if update.Emergency {
		slog.Info("Emergency actor: got emergency event. Sending notification to Emergency services.",
			slog.String("vehicle_id", update.ID),
			slog.Float64("lon", update.Lon),
			slog.Float64("lat", update.Lat),
			slog.Int("queue_len", process.Info().MessageQueueLen),
		)
		time.Sleep(time.Millisecond * 10) // fake a slow external call here.
	}

	return gen.ServerStatusOK
}
