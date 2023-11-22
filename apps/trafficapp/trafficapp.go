package trafficapp

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log/slog"
	"traffic/apps/trafficapp/web"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func CreateTrafficApp(db *buntdb.DB) gen.ApplicationBehavior {
	return &TrafficApp{db: db}
}

type TrafficApp struct {
	gen.Application
	db *buntdb.DB
}

func (app *TrafficApp) Load(args ...etf.Term) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        "trafficapp",
		Description: "Traffic and vehicle position tracker",
		Version:     "v.1.0",
		Children: []gen.ApplicationChildSpec{
			{
				Name:  "commonsup",
				Child: createCommonSup(app.db),
			},
			{
				Name:  "web",
				Child: web.CreateWeb(),
			},
		},
	}, nil
}

func (app *TrafficApp) Start(process gen.Process, args ...etf.Term) {
	slog.Info(fmt.Sprintf("Application TrafficApp started with Pid %s", process.Self()))
}
