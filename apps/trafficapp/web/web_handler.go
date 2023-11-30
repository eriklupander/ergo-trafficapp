package web

import (
	"encoding/json"
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"traffic/apps/events"
)

func createWebHandler() gen.WebHandlerBehavior {
	return &WebHandler{}
}

type WebHandler struct {
	gen.WebHandler
}

//
// Mandatory callbacks
//

// HandleRequest invokes on a HTTP-request
func (r *WebHandler) HandleRequest(process *gen.WebHandlerProcess, request gen.WebMessageRequest) gen.WebHandlerStatus {
	slog.Debug("[Web handler] Received request", slog.String("method", request.Request.Method))
	st := time.Now()

	if request.Request.URL == nil {
		request.Response.WriteHeader(http.StatusBadRequest)
		return gen.WebHandlerStatusDone
	}

	lat, err := strconv.ParseFloat(request.Request.URL.Query().Get("lat"), 64)
	if err != nil {
		request.Response.WriteHeader(http.StatusBadRequest)
		return gen.WebHandlerStatusDone
	}

	lon, err := strconv.ParseFloat(request.Request.URL.Query().Get("lon"), 64)
	if err != nil {
		request.Response.WriteHeader(http.StatusBadRequest)
		return gen.WebHandlerStatusDone
	}

	elems, err := process.Call("queries_workers", events.NearbyQueryEventMessage{Lat: lat, Lon: lon})
	if err != nil {
		slog.Error("Call storage error", slog.Any("error", err))
		request.Response.WriteHeader(http.StatusNotFound)
		return gen.WebHandlerStatusDone
	}

	data, err := json.Marshal(elems)
	if err != nil {
		request.Response.WriteHeader(http.StatusInternalServerError)
		return gen.ServerStatusIgnore
	}
	request.Response.Header().Set("Content-Type", "application/json")
	request.Response.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	request.Response.WriteHeader(http.StatusOK)
	_, _ = request.Response.Write(data)

	slog.Info("WebHandler - served query", slog.Duration("duration", time.Since(st)), slog.Int("vehicle_count", len(elems.([]events.VehiclePosition))), slog.Float64("lon", lon), slog.Float64("lat", lat))
	return gen.WebHandlerStatusDone
}

//
// Optional gen.Server's callbacks
//

// HandleWebHandlerCall this callback is invoked on ServerProcess.Call(...).
func (r *WebHandler) HandleWebHandlerCall(process *gen.WebHandlerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	return nil, gen.ServerStatusOK
}

// HandleWebHandlerCast this callback is invoked on ServerProcess.Cast(...).
func (r *WebHandler) HandleWebHandlerCast(process *gen.WebHandlerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleWebHandlerInfo this callback is invoked on Process.Send(...).
func (r *WebHandler) HandleWebHandlerInfo(process *gen.WebHandlerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleWebHandlerTerminate this callback is invoked on the process termiation, providing the reason of termination
// along with the counter of handled requests
func (r *WebHandler) HandleWebHandlerTerminate(process *gen.WebHandlerProcess, reason string, count int64) {

}
