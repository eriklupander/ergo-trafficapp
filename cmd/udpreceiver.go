package main

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func createUDPReceiver() gen.UDPBehavior {
	return &UDPReceiver{}
}

type UDPReceiver struct {
	gen.UDP
}

//
// Mandatory callbacks
//

// InitUDP invoked on starting UDP server
func (us *UDPReceiver) InitUDP(process *gen.UDPProcess, args ...etf.Term) (gen.UDPOptions, error) {
	var options gen.UDPOptions

	options.Handler = createUDPReceiverHandler()
	options.Port = 9092
	options.NumHandlers = 4
	options.ExtraHandlers = true
	options.QueueLength = 4096

	return options, nil
}

//
// Optional gen.Server's callbacks
//

// HandleUDPCall this callback is invoked on ServerProcess.Call(...).
func (us *UDPReceiver) HandleUDPCall(process *gen.UDPProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	return nil, gen.ServerStatusOK
}

// HandleUDPCast this callback is invoked on ServerProcess.Cast(...).
func (us *UDPReceiver) HandleUDPCast(process *gen.UDPProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleUDPInfo this callback is invoked on Process.Send(...).
func (us *UDPReceiver) HandleUDPInfo(process *gen.UDPProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleUDPTerminate this callback invoked on a process termination
func (us *UDPReceiver) HandleUDPTerminate(process *gen.UDPProcess, reason string) {

}
