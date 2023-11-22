package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"github.com/lmittmann/tint"
	"github.com/tidwall/buntdb"
	"log/slog"
	"os"
	"time"
	"traffic/apps/trafficapp"

	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

var (
	OptionNodeName   string
	OptionNodeCookie string

	OptionLevel string
)

func init() {
	// generate random value for cookie
	buff := make([]byte, 12)
	_, _ = rand.Read(buff)
	randomCookie := hex.EncodeToString(buff)

	flag.StringVar(&OptionNodeName, "name", "traffic@localhost", "node name")
	flag.StringVar(&OptionNodeCookie, "cookie", randomCookie, "a secret cookie for interaction within the cluster")
	flag.StringVar(&OptionLevel, "level", "info", "log level")
}

func main() {
	var options node.Options
	var process gen.Process

	flag.Parse()

	lvl := slog.LevelInfo
	switch OptionLevel {
	case "debug":
		fallthrough
	case "DEBUG":
		lvl = slog.LevelDebug
	}

	//slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      lvl, //slog.LevelInfo,
			TimeFormat: time.Kitchen,
			AddSource:  true,
		}),
	))

	// Set up DB
	db, err := buntdb.Open("data.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create applications that must be started
	apps := []gen.ApplicationBehavior{
		trafficapp.CreateTrafficApp(db),
	}
	options.Applications = apps

	// Starting node
	ctx := context.Background()
	trafficNode, err := ergo.StartNodeWithContext(ctx, OptionNodeName, OptionNodeCookie, options)
	if err != nil {
		panic(err)
	}
	slog.Info("Node is started", slog.String("node", trafficNode.Name()))

	// starting process TCPReceiver
	process, err = trafficNode.Spawn("tcpreceiver", gen.ProcessOptions{}, createTCPReceiver())
	if err != nil {
		panic(err)
	}
	slog.Info("  process is started", slog.String("process_name", process.Name()), slog.Any("pid", process.Self()))

	// starting process UDPReceiver
	udpProcess, err := trafficNode.Spawn("udpreceiver", gen.ProcessOptions{MailboxSize: 4096, DirectboxSize: 4096}, createUDPReceiver())
	if err != nil {
		panic(err)
	}
	slog.Info("  process is started", slog.String("process_name", udpProcess.Name()), slog.Any("pid", udpProcess.Self()))

	// starting process IngestorsWorkers
	process, err = trafficNode.Spawn("ingestors_workers", gen.ProcessOptions{}, createIngestors())
	if err != nil {
		panic(err)
	}
	slog.Info("  process is started", slog.String("process_name", process.Name()), slog.Any("pid", process.Self()))

	// starting process QueriesWorkers
	process, err = trafficNode.Spawn("queries_workers", gen.ProcessOptions{}, createQueries(db))
	if err != nil {
		panic(err)
	}
	slog.Info("  process is started", slog.String("process_name", process.Name()), slog.Any("pid", process.Self()))

	// Dump stats to stdout every 30 seconds
	//go func() {
	//	ticker := time.NewTicker(time.Second * 10)
	//OUTER:
	//	for {
	//		select {
	//		case <-ticker.C:
	//			slog.Info(fmt.Sprintf("Traffic node stats: %+v\n", trafficNode.Stats()))
	//			slog.Info("storage message box sizes", slog.Int("mailbox", storageProcess.Info().MessageQueueLen))
	//			slog.Info("UDP message box sizes", slog.Int("mailbox", udpProcess.Info().MessageQueueLen))
	//		case <-ctx.Done():
	//			break OUTER
	//		}
	//	}
	//
	//}()

	trafficNode.Wait()
}
