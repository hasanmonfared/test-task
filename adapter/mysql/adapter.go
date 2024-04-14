package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Username string `koanf="username"`
	Password string `koanf="password"`
	Port     int    `koanf="port"`
	Host     string `koanf="host"`
	DBName   string `koanf="dbname"`
}
type Adapter struct {
	config Config
	db     *sql.DB
}

func (m *Adapter) Conn() *sql.DB {
	return m.db
}
func New(config Config) Adapter {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db:%v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return Adapter{config: config, db: db}
}
