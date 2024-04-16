package main

import (
	"app/config"
	"app/delivery/grpcserver/walletgrpcserver"
)

func main() {
	cfg := config.Load("config.yml")
	server := walletgrpcserver.New(cfg)
	server.Start()
}
