package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/tidwall/buntdb"
	"log/slog"
)

type Queries struct {
	gen.Pool
	db *buntdb.DB
}

func createQueries(db *buntdb.DB) gen.PoolBehavior {
	return &Queries{db: db}
}

func (p *Queries) InitPool(process *gen.PoolProcess, args ...etf.Term) (gen.PoolOptions, error) {
	slog.Info(fmt.Sprintf("Init pool process: %s with name %q and args %v", process.Self(), process.Name(), args))

	opts := gen.PoolOptions{
		Worker:     createQueriesWorker(p.db),
		NumWorkers: 3,
	}

	return opts, nil
}
