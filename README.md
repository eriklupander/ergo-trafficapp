## Project: "traffic"

### Generated with
 - Types for network messaging: false
 - Enabled Cloud feature: false

### Supervision Tree

Applications
 - `TrafficApp{}` traffic/apps/trafficapp/trafficapp.go
   - `CommonSup{}` traffic/apps/trafficapp/commonsup.go
     - `Dispatcher{}` traffic/apps/trafficapp/dispatcher.go

Process list that is starting by node directly
 - `Storage{}` traffic/cmd/storage.go
 - `Web{}` traffic/cmd/web.go
 - `TCPReceiver{}` traffic/cmd/tcpreceiver.go
 - `UDPReceiver{}` traffic/cmd/udpreceiver.go


#### Used command
`ergo -init traffic -with-app TrafficApp -with-sup TrafficApp:CommonSup{type:rfo} -with-actor CommonSup:Dispatcher -with-actor Storage -with-tcp TCPReceiver{ssl:false,port:9091,handlers:5} -with-udp UDPReceiver{port:9092,handlers:4} -with-web Web{ssl:yes,port:9090,handlers:3}`

## Usage examples

curl the HTTP API:
`curl -k https://localhost:9090\?lon=43.5454\&lat=76.232 | jq .`

Use the tools/udpclient:
`go run main.go localhost:9092 2`