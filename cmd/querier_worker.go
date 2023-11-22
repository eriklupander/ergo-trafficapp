package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/tidwall/buntdb"
	"log/slog"
	"strconv"
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
	fmt.Println("   started pool queries worker: ", process.Self())

	return nil
}

func (w *QueriesWorker) HandleWorkerCall(process *gen.PoolWorkerProcess, message etf.Term) etf.Term {
	slog.Debug("QueriesWorker received Call request")
	query := message.(events.NearbyQueryEventMessage)

	elems := make([]events.VehiclePosition, 0)

	// Call DB
	if err := w.db.View(func(tx *buntdb.Tx) error {
		count := 0
		return tx.Nearby("fleet", buntdb.Point(query.Lon, query.Lat), func(key, val string, dist float64) bool {
			if count >= maxItems {
				return false
			}

			lon, lat, ok := coordToLonLat(val)
			if !ok {
				return false
			}

			elems = append(elems, events.VehiclePosition{
				ID:   strings.Split(key, ":")[1],
				Lon:  lon,
				Lat:  lat,
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

func coordToLonLat(val string) (float64, float64, bool) {
	if len(val) == 0 {
		return 0, 0, false
	}
	// Ugly, parse coords string
	coords := strings.Split(val[1:len(val)-1], " ")
	if len(coords) < 2 {
		return 0, 0, false
	}
	lon, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return 0, 0, false
	}
	lat, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return 0, 0, false
	}
	return lon, lat, true
}

func (w *QueriesWorker) HandleWorkerCast(process *gen.PoolWorkerProcess, message etf.Term) {
	fmt.Printf("QueriesWorker [%s] received Cast message: %v\n", process.Self(), message)
}

func (w *QueriesWorker) HandleWorkerInfo(process *gen.PoolWorkerProcess, message etf.Term) {
	fmt.Printf("QueriesWorker [%s] received Info message: %v\n", process.Self(), message)

	query := message.(events.NearbyQueryEventMessage)
	// Call DB
	if err := w.db.View(func(tx *buntdb.Tx) error {
		return tx.Nearby("fleet", fmt.Sprintf("[%f %f]", query.Lon, query.Lat), func(key, val string, dist float64) bool {
			slog.Info("found item", slog.String("key", key), slog.String("val", val), slog.Float64("dist", dist))
			return true
		})
	}); err != nil {
		slog.Error("QueriesWorker query error", slog.Any("error", err))
	}
}
