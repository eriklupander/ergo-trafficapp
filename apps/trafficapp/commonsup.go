package trafficapp

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/tidwall/buntdb"
)

func createCommonSup(db *buntdb.DB) gen.SupervisorBehavior {
	return &CommonSup{db: db}
}

type CommonSup struct {
	gen.Supervisor
	db *buntdb.DB
}

func (sup *CommonSup) Init(args ...etf.Term) (gen.SupervisorSpec, error) {

	spec := gen.SupervisorSpec{
		Name: "commonsup",
		Children: []gen.SupervisorChildSpec{
			{
				Name:    "dispatcher",
				Child:   createDispatcher(),
				Options: gen.ProcessOptions{MailboxSize: 16384, DirectboxSize: 16384},
			},
			{
				Name:    "emergency",
				Child:   createEmergency(),
				Options: gen.ProcessOptions{MailboxSize: 16384, DirectboxSize: 16384},
			},
			{
				Name:    "storage",
				Child:   createStorage(sup.db),
				Options: gen.ProcessOptions{MailboxSize: 16384, DirectboxSize: 16384},
			},
		},
		Strategy: gen.SupervisorStrategy{
			Type:      gen.SupervisorStrategyRestForOne,
			Intensity: 2, // How big bursts of restarts you want to tolerate.
			Period:    5, // In seconds.
			Restart:   gen.SupervisorStrategyRestartTransient,
		},
	}
	return spec, nil
}
