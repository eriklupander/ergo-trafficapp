package web

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"

	"crypto/tls"
	"github.com/ergo-services/ergo/lib"
)

func CreateWeb() gen.WebBehavior {
	return &Web{}
}

type Web struct {
	gen.Web
}

//
// Mandatory callbacks
//

// InitWeb invoked on starting Web server
func (w *Web) InitWeb(process *gen.WebProcess, args ...etf.Term) (gen.WebOptions, error) {
	slog.Info(fmt.Sprintf("Init process: %s with name %q and args %v", process.Self(), process.Name(), args))
	var options gen.WebOptions
	options.Port = 9090
	options.Host = "localhost"

	// enable TLS with self-signed certificate
	cert, _ := lib.GenerateSelfSignedCert("Web Service")
	certUpdater := lib.CreateCertUpdater(cert)
	tlsConfig := &tls.Config{
		GetCertificate: certUpdater.GetCertificateFunc(),
	}
	options.TLS = tlsConfig

	mux := http.NewServeMux()
	handlerOptions := gen.WebHandlerOptions{
		NumHandlers:    3,
		IdleTimeout:    10,
		RequestTimeout: 20,
	}
	webRoot := process.StartWebHandler(createWebHandler(), handlerOptions)
	mux.Handle("/", webRoot)
	options.Handler = mux
	slog.Info("Start Web server", slog.String("addr", fmt.Sprintf("https://%s:%d/", options.Host, options.Port)))
	return options, nil
}

//
// Optional gen.Server's callbacks
//

// HandleWebCall this callback is invoked on ServerProcess.Call(...).
func (w *Web) HandleWebCall(process *gen.WebProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	slog.Debug("handle web call")
	return nil, gen.ServerStatusOK
}

// HandleWebCast this callback is invoked on ServerProcess.Cast(...).
func (w *Web) HandleWebCast(process *gen.WebProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("handle web caast")
	return gen.ServerStatusOK
}

// HandleWebInfo this callback is invoked on Process.Send(...).
func (w *Web) HandleWebInfo(process *gen.WebProcess, message etf.Term) gen.ServerStatus {
	slog.Debug("handle web info")
	return gen.ServerStatusOK
}
