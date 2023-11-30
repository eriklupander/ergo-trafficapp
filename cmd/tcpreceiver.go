package main

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"log/slog"

	"crypto/tls"
	"github.com/ergo-services/ergo/lib"
)

func createTCPReceiver() gen.TCPBehavior {
	return &TCPReceiver{}
}

type TCPReceiver struct {
	gen.TCP
}

//
// Mandatory callbacks
//

// InitTCP
func (ts *TCPReceiver) InitTCP(process *gen.TCPProcess, args ...etf.Term) (gen.TCPOptions, error) {
	var options gen.TCPOptions

	options.Handler = createTCPReceiverHandler()
	options.Port = 9091
	options.NumHandlers = 5

	// Enable SSL with self-signed certificate
	slog.Info("Enabled SSL for TCP server with self-signed certificate. You may check it with command below:")
	slog.Info(fmt.Sprintf("   $ openssl s_client -connect %s:%d\n", options.Host, options.Port))
	cert, _ := lib.GenerateSelfSignedCert("TCPReceiver Service")
	certUpdater := lib.CreateCertUpdater(cert)
	tlsConfig := &tls.Config{
		GetCertificate: certUpdater.GetCertificateFunc(),
	}
	options.TLS = tlsConfig
	return options, nil
}

//
// Optional gen.Server's callbacks
//

// HandleTCPCall this callback is invoked on ServerProcess.Call(...).
func (us *TCPReceiver) HandleTCPCall(process *gen.TCPProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	return nil, gen.ServerStatusOK
}

// HandleTCPCast this callback is invoked on ServerProcess.Cast(...).
func (us *TCPReceiver) HandleTCPCast(process *gen.TCPProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleTCPInfo this callback is invoked on Process.Send(...).
func (us *TCPReceiver) HandleTCPInfo(process *gen.TCPProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleTCPTerminate this callback invoked on a process termination
func (us *TCPReceiver) HandleTCPTerminate(process *gen.TCPProcess, reason string) {

}
