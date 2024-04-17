package main

import (
	"app/adapter/mysql"
	"app/adapter/redis"
	"app/adapter/walletadapter"
	"app/config"
	"app/delivery/httpserver"
	"app/repository/migrator"
	"app/service/discountservice"
	"app/validator/discountvalidator"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("./config.yml")

	mysqlAdapter := mysql.New(cfg.Mysql)
	m := migrator.New(mysqlAdapter.Conn())
	m.Up()
	discountSvc, discountValidator := setupServices(cfg)
	server := httpserver.New(cfg, discountSvc, discountValidator)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefulShutdownTimeout)
	defer cancel()
	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)

	}
	time.Sleep(cfg.Application.GracefulShutdownTimeout)
	<-ctxWithTimeout.Done()
}
func setupServices(cfg config.Config) (discountservice.Service, discountvalidator.Validator) {
	// MYSQL
	mysqlAdapter := mysql.New(cfg.Mysql)
	redisAdapter := redis.New(cfg.Redis)
	walletGrpc := fmt.Sprintf(":%d", cfg.WalletGrpcServer.Port)
	walletAdapter := walletadapter.New(walletGrpc)

	diSvc := discountservice.New(mysqlAdapter, redisAdapter, walletAdapter)
	diVal := discountvalidator.New(mysqlAdapter)
	return diSvc, diVal
}
