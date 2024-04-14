package main

import (
	"app/adapter/mysql"
	"app/config"
	"app/delivery/httpserver"
	"app/repository/mysql/mysqluser"
	"app/service/userservice"
	"app/validator/uservalidator"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("./config.yml")
	time.Sleep(5 * time.Second)
	//mgr := migrator.New(cfg.Mysql)
	//mgr.Up()

	//userSvc, userValidator := setupServices(cfg)
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

func setupServices(cfg config.Config) (userservice.Service, uservalidator.Validator) {
	// MYSQL
	mysqlAdapter := mysql.New(cfg.Mysql)
	mysqlUser := mysqluser.New(mysqlAdapter)
	uV := uservalidator.New(&mysqlUser)
	userSvc := userservice.New(&mysqlUser)
	return userSvc, uV
}
