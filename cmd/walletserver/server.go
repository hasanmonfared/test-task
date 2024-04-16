package main

import (
	"app/adapter/mysql"
	"app/config"
	"app/delivery/grpcserver/walletgrpcserver"
	"app/service/walletservice"
)

func main() {
	cfg := config.Load("config.yml")
	mysqlAdapter := mysql.New(cfg.Mysql)
	svc := walletservice.New(mysqlAdapter.Conn())
	server := walletgrpcserver.New(cfg, svc)
	server.Start()
}
