package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/vmihailenco/msgpack/v5"
	"log/slog"
	"traffic/apps/events"
)

func createIngestorsWorker() gen.PoolWorkerBehavior {
	return &IngestorsWorker{}
}

type IngestorsWorker struct {
	gen.PoolWorker
}

func (w *IngestorsWorker) InitPoolWorker(process *gen.PoolWorkerProcess, args ...etf.Term) error {
	slog.Info("   started pool IngestorsWorker", slog.Any("process", process.Self()))
	return nil
}

func (w *IngestorsWorker) HandleWorkerCall(process *gen.PoolWorkerProcess, message etf.Term) etf.Term {
	slog.Debug("IngestorsWorker received Call request")
	return "pong"
}

func (w *IngestorsWorker) HandleWorkerCast(process *gen.PoolWorkerProcess, message etf.Term) {
	slog.Debug("IngestorsWorker received Cast message")
}

func (w *IngestorsWorker) HandleWorkerInfo(process *gen.PoolWorkerProcess, message etf.Term) {
	slog.Debug("IngestorsWorker received Info message")

	trafficEvt, ok := message.(events.TrafficEventMessage)
	if !ok {
		slog.Error("IngestorsWorker: type conversion from etf.Term to events.TrafficEventMessage failed", slog.Any("type", fmt.Sprintf("%T", message)))
		return
	}

	// Unmarshal from msgpack
	geoPosUpdate := events.GeoPos{}
	if err := msgpack.Unmarshal(trafficEvt.Payload, &geoPosUpdate); err != nil {
		slog.Error(err.Error())
		return
	}

	// Comment in to fake some time-consuming work.
	//time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))

	if err := process.Send("storage", events.VehiclePosition{
		ID:   geoPosUpdate.ID,
		Lat:  geoPosUpdate.Lat,
		Lon:  geoPosUpdate.Lon,
		Date: trafficEvt.Date,
	}); err != nil {
		slog.Error("IngestorsWorker: Send to storage", slog.Any("error", err))
	}
}
