# UDP traffic simulator

Usage:

go run main.go <remote UDP address> <number of generators>

`go run main.go localhost:9092 400` will connect to UDP server listening on localhost:9092 with 400 "traffic simulators" each sending one `msgpack` UDP packet per second with a random lon/lat for a per-worker generated UUID.