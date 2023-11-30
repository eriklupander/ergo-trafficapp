package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"log/slog"
	"traffic/apps/events"
	"traffic/apps/trafficapp"
)

type Ingestors struct {
	gen.Pool
}

func createIngestors() gen.PoolBehavior {
	return &Ingestors{}
}

func (p *Ingestors) InitPool(process *gen.PoolProcess, args ...etf.Term) (gen.PoolOptions, error) {
	slog.Info(fmt.Sprintf("Init pool process: %s with name %q and args %v", process.Self(), process.Name(), args))

	opts := gen.PoolOptions{
		Worker:        createIngestorsWorker(),
		NumWorkers:    5,
		WorkerOptions: gen.ProcessOptions{MailboxSize: trafficapp.DefaultMailboxSize * 8},
	}
	if err := process.MonitorEvent(events.TrafficEvent); err != nil {
		slog.Error("can't monitor event", slog.Any("error", err))
	}
	return opts, nil
}
