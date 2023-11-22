package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/vmihailenco/msgpack/v5"
	"log/slog"
	"math/rand"
	"traffic/apps/events"
)

func createIngestorsWorker() gen.PoolWorkerBehavior {
	return &IngestorsWorker{}
}

type IngestorsWorker struct {
	gen.PoolWorker
}

func (w *IngestorsWorker) InitPoolWorker(process *gen.PoolWorkerProcess, args ...etf.Term) error {
	fmt.Println("   started pool IngestorsWorker: ", process.Self())

	return nil
}

func (w *IngestorsWorker) HandleWorkerCall(process *gen.PoolWorkerProcess, message etf.Term) etf.Term {
	fmt.Printf("IngestorsWorker [%s] received Call request: %v\n", process.Self(), message)
	return "pong"
}

func (w *IngestorsWorker) HandleWorkerCast(process *gen.PoolWorkerProcess, message etf.Term) {
	fmt.Printf("IngestorsWorker [%s] received Cast message: %v\n", process.Self(), message)
}

func (w *IngestorsWorker) HandleWorkerInfo(process *gen.PoolWorkerProcess, message etf.Term) {
	slog.Debug("IngestorsWorker received Info message")

	trafficEvt, ok := message.(events.TrafficEventMessage)
	if !ok {
		slog.Error("type conversion from etf.Term to events.TrafficEventMessage failed")
		return
	}

	// Unmarshal from msgpack
	geoPosUpdate := events.GeoPos{}
	if err := msgpack.Unmarshal(trafficEvt.Payload, &geoPosUpdate); err != nil {
		slog.Error(err.Error())
		return
	}

	// Comment in to fake some time-consuming work.
	//time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))

	if err := process.Send("storage", events.VehiclePosition{
		ID:   geoPosUpdate.ID,
		Lon:  geoPosUpdate.Lon,
		Lat:  geoPosUpdate.Lat,
		Date: trafficEvt.Date,
	}); err != nil {
		slog.Error("IngestorsWorker: Send to storage", slog.Any("error", err))
	}
}

func rndGeo() float64 {
	return -90 + rand.Float64()*180.0
}
