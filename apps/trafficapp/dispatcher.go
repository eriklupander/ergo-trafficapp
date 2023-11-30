package trafficapp

import (
	"fmt"
	"log/slog"
	"traffic/apps/events"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func createDispatcher() gen.ServerBehavior {
	return &Dispatcher{}
}

type Dispatcher struct {
	gen.Server
}

// Init invoked on a start this process.
func (s *Dispatcher) Init(process *gen.ServerProcess, args ...etf.Term) error {
	slog.Info(fmt.Sprintf("Init process: %s with name %q and args %v", process.Self(), process.Name(), args))

	if err := process.RegisterEvent(events.TrafficEvent, events.TrafficEventMessage{}); err != nil {
		slog.Error("can't register event " + err.Error())
		return err
	}
	return nil
}

//
// Methods below are optional, so you can remove those that aren't be used
//

// HandleInfo invoked if this process received message sent with Process.Send(...).
func (s *Dispatcher) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("Dispatcher: HandleInfo")

	switch message.(type) {
	case events.TrafficEventMessage:
		if err := process.SendEventMessage(events.TrafficEvent, message); err != nil {
			slog.Error("Dispatcher: handle SendEventMessage(events.TrafficEvent) error", slog.Any("error", err))
		}
	}

	return gen.ServerStatusOK
}

// HandleCast invoked if this process received message sent with ServerProcess.Cast(...).
// Return ServerStatusStop to stop server with "normal" reason. Use ServerStatus(error)
// for the custom reason
func (s *Dispatcher) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("HandleCast [Dispatcher]")
	return gen.ServerStatusOK
}

// HandleCall invoked if this process got sync request using ServerProcess.Call(...)
func (s *Dispatcher) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	slog.Debug("HandleCall [Dispatcher]")
	return nil, gen.ServerStatusOK
}

// HandleDirect invoked on a direct request made with Process.Direct(...)
func (s *Dispatcher) HandleDirect(process *gen.ServerProcess, ref etf.Ref, message interface{}) (interface{}, gen.DirectStatus) {
	slog.Debug("HandleDirect [Dispatcher]")
	return nil, nil
}

// Terminate invoked on a termination process. ServerProcess.State is not locked during this callback.
func (s *Dispatcher) Terminate(process *gen.ServerProcess, reason string) {
	slog.Info("Terminated [Dispatcher]", slog.Any("process_id", process.Self()), slog.String("reason", reason))
}
