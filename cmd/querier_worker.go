package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/tidwall/buntdb"
	"log/slog"
	"strings"
	"time"
	"traffic/apps/events"
)

const maxItems = 10

func createQueriesWorker(db *buntdb.DB) gen.PoolWorkerBehavior {
	return &QueriesWorker{db: db}
}

type QueriesWorker struct {
	gen.PoolWorker
	db *buntdb.DB
}

func (w *QueriesWorker) InitPoolWorker(process *gen.PoolWorkerProcess, args ...etf.Term) error {
	slog.Info("   started pool queries worker", slog.Any("process_id", process.Self()))

	return nil
}

func (w *QueriesWorker) HandleWorkerCall(process *gen.PoolWorkerProcess, message etf.Term) etf.Term {
	slog.Debug("QueriesWorker received Call request", slog.Any("process_id", process.Self().String()))
	query, ok := message.(events.NearbyQueryEventMessage)

	elems := make([]events.VehiclePosition, 0)
	if !ok {
		slog.Error("QueriesWorker: invalid type passed to HandleWorkerCall", slog.Any("type", fmt.Sprintf("%T", message)))
		return elems
	}

	// Call DB
	if err := w.db.View(func(tx *buntdb.Tx) error {
		count := 0
		return tx.Nearby("fleet", buntdb.Point(query.Lat, query.Lon), func(key, val string, dist float64) bool {
			if count >= maxItems {
				return false
			}

			lat, lon, ok := events.CoordToLatLon(val)
			if !ok {
				return false
			}

			elems = append(elems, events.VehiclePosition{
				ID:   strings.Split(key, ":")[1],
				Lat:  lat,
				Lon:  lon,
				Date: time.Now().UnixMilli(),
			})
			count++
			return true
		})
	}); err != nil {
		slog.Error("QueriesWorker query error", slog.Any("error", err))
	}

	return elems
}

func (w *QueriesWorker) HandleWorkerCast(process *gen.PoolWorkerProcess, message etf.Term) {
	slog.Debug("QueriesWorker received Cast message")
}

func (w *QueriesWorker) HandleWorkerInfo(process *gen.PoolWorkerProcess, message etf.Term) {
	slog.Debug("QueriesWorker received Info message", slog.Any("process_id", process.Self().String()))

	query := message.(events.NearbyQueryEventMessage)

	// Call DB
	if err := w.db.View(func(tx *buntdb.Tx) error {
		return tx.Nearby("fleet", fmt.Sprintf("[%f %f]", query.Lat, query.Lon), func(key, val string, dist float64) bool {
			slog.Info("found item", slog.String("key", key), slog.String("val", val), slog.Float64("dist", dist))
			return true
		})
	}); err != nil {
		slog.Error("QueriesWorker query error", slog.Any("error", err))
	}
}
