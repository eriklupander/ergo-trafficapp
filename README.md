# Project: "traffic"

Just-for-fun and make-believe app built using Ergo Framework.

## Start

`go run cmd/*.go -level info`

## Usage examples

curl the HTTP API:
`curl -k https://localhost:9090\?lon=43.5454\&lat=76.232 | jq .`

Use the tools/udpclient with 2 clients.
`go run main.go localhost:9092 2`