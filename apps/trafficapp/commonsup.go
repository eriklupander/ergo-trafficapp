package trafficapp

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/tidwall/buntdb"
)

const (
	DefaultMailboxSize       = 128
	DefaultDirectMailboxSize = 32
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
				Options: gen.ProcessOptions{MailboxSize: DefaultMailboxSize, DirectboxSize: DefaultDirectMailboxSize},
			},
			{
				Name:    "emergency",
				Child:   createEmergency(),
				Options: gen.ProcessOptions{MailboxSize: DefaultMailboxSize * 32, DirectboxSize: DefaultDirectMailboxSize},
			},
			{
				Name:    "storage",
				Child:   createStorage(sup.db),
				Options: gen.ProcessOptions{MailboxSize: DefaultMailboxSize * 256, DirectboxSize: DefaultDirectMailboxSize},
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
