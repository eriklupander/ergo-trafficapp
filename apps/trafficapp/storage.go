package trafficapp

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"github.com/vmihailenco/msgpack/v5"
	"log/slog"
	"traffic/apps/events"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func createStorage(db *buntdb.DB) gen.ServerBehavior {
	if err := db.CreateSpatialIndex("fleet", "fleet:*:pos", buntdb.IndexRect); err != nil {
		slog.Error(err.Error())
	}
	if err := db.CreateIndex("vehicle", "vehicle:*", buntdb.IndexString); err != nil {
		slog.Error(err.Error())
	}
	return &Storage{db: db}
}

type Storage struct {
	gen.Server
	db *buntdb.DB
}

// Init invoked on a start this process.
func (s *Storage) Init(process *gen.ServerProcess, args ...etf.Term) error {
	slog.Info(fmt.Sprintf("Init process: %s with name %q and args %v", process.Self(), process.Name(), args))
	return nil
}

//
// Methods below are optional, so you can remove those that aren't be used
//

// HandleInfo invoked if this process received message sent with Process.Send(...).
func (s *Storage) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("[Storage] HandleInfo")

	trafficEvt, ok := message.(events.VehiclePosition)
	if !ok {
		slog.Error("Storage: Type conversion from etf.Term to events.VehiclePosition failed")
		return gen.ServerStatusIgnore
	}

	var err error
	// Store position on fleet index
	if err := s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set("fleet:"+trafficEvt.ID+":pos", buntdb.Point(trafficEvt.Lat, trafficEvt.Lon), nil)
		return err
	}); err != nil {
		slog.Error(err.Error())
		return gen.ServerStatusIgnore
	}
	slog.Debug("stored position event in DB", slog.Float64("lat", trafficEvt.Lat), slog.Float64("lon", trafficEvt.Lon))

	// Marshal the PositionUpdate into msgpack and dump into the DB.
	str, err := msgpack.Marshal(trafficEvt)
	if err != nil {
		return gen.ServerStatusIgnore
	}

	if err := s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("vehicle:"+trafficEvt.ID, string(str), nil)
		return err
	}); err != nil {
		slog.Error(err.Error())
		return gen.ServerStatusIgnore
	}

	return gen.ServerStatusOK
}

// HandleCast invoked if this process received message sent with ServerProcess.Cast(...).
// Return ServerStatusStop to stop server with "normal" reason. Use ServerStatus(error)
// for the custom reason
func (s *Storage) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("HandleCast [Storage]")
	return gen.ServerStatusOK
}

// HandleCall invoked if this process got synchronous request using ServerProcess.Call(...)
func (s *Storage) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	slog.Info("HandleCall", slog.Any("from", from.Pid))
	return nil, gen.ServerStatusOK
}

// HandleDirect invoked on a direct request made with Process.Direct(...)
func (s *Storage) HandleDirect(process *gen.ServerProcess, ref etf.Ref, message interface{}) (interface{}, gen.DirectStatus) {
	slog.Info("HandleDirect")
	return nil, nil
}

// Terminate invoked on a termination process. ServerProcess.State is not locked during this callback.
func (s *Storage) Terminate(process *gen.ServerProcess, reason string) {
	slog.Info("Terminated: %s with reason %s", slog.String("reason", reason))
}
