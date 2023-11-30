package main

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"log/slog"
	"time"
	"traffic/apps/events"
)

func createUDPReceiverHandler() gen.UDPHandlerBehavior {
	return &UDPReceiverHandler{}
}

type UDPReceiverHandler struct {
	gen.UDPHandler
}

//
// Mandatory callbacks
//

// HandlePacket invokes on receiving UDP datagram
func (uh *UDPReceiverHandler) HandlePacket(process *gen.UDPHandlerProcess, data []byte, packet gen.UDPPacket) {
	slog.Debug("[UDP handler] got message", slog.Int("queue_len", process.Info().MessageQueueLen), slog.String("process_id", process.Self().String()))

	err := process.Send("dispatcher", events.TrafficEventMessage{Date: time.Now().UnixMilli(), Payload: data})
	if err != nil {
		slog.Error("cannot send event message", slog.Any("error", err))
	}
	// If you want to send a reply message, use packet.Socket.Write(reply) for that.
	_, err = packet.Socket.Write([]byte(`\n`))
	if err != nil {
		slog.Error("writing response", slog.Any("error", err))
	}
}

//
// Optional callbacks
//

// HandleTimeout invokes on socket reading timeout, which is default 3 seconds
func (uh *UDPReceiverHandler) HandleTimeout(process *gen.UDPHandlerProcess) {
	slog.Debug("[UDP handler] timeout")
}

//
// Optional gen.Server's callbacks
//

// HandleUDPHandlerCall this callback is invoked on ServerProcess.Call(...).
func (r *UDPReceiverHandler) HandleUDPHandlerCall(process *gen.UDPHandlerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	slog.Info("[UDP handler] HandleUDPHandlerCall")
	return nil, gen.ServerStatusOK
}

// HandleUDPHandlerCast this callback is invoked on ServerProcess.Cast(...).
func (r *UDPReceiverHandler) HandleUDPHandlerCast(process *gen.UDPHandlerProcess, message etf.Term) gen.ServerStatus {
	slog.Info("[UDP handler] HandleUDPHandlerCast")
	return gen.ServerStatusOK
}

// HandleUDPHandlerInfo this callback is invoked on Process.Send(...).
func (r *UDPReceiverHandler) HandleUDPHandlerInfo(process *gen.UDPHandlerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleUDPHandlerTerminate this callback is invoked on the process termiation
func (r *UDPReceiverHandler) HandleUDPHandlerTerminate(process *gen.UDPHandlerProcess, reason string) {

}
