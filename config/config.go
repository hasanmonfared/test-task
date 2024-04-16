package config

import (
	"app/adapter/mysql"
	"app/adapter/redis"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type HTTPServer struct {
	Port int `koanf:"port"`
}
type WalletGrpcServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	Application      Application      `koanf:"application"`
	HTTPServer       HTTPServer       `koanf:"http_server"`
	WalletGrpcServer WalletGrpcServer `koanf:"wallet_grpc_server"`
	Mysql            mysql.Config     `koanf:"mysql"`
	Redis            redis.Config     `koanf:"redis"`
}
