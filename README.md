# TCPTUNNEL-CLIENT
This is a raw (without sh module) reverse tcp tunnel implementation

## USAGE
export TCPTUNNELHOST=192.168.0.100 && go run main.go
A service needs to be running on PORT 8080, such as a proxy
See: https://github.com/elazarl/goproxy

## SERVER-SIDE
https://github.com/tenenwurcel/tcptunnel-server

## WARNING
This project was made for fun and learning purposes and it's not intended to be a production ready app
