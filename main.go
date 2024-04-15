package main

import (
	"app/adapter/mysql"
	"app/config"
	"app/delivery/httpserver"
	"app/repository/migrator"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("./config.yml")

	mysqlAdapter := mysql.New(cfg.Mysql)
	m := migrator.New(mysqlAdapter.Conn())
	m.Up()
	server := httpserver.New(cfg)
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
