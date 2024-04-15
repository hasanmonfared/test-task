package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db     *gorm.DB
}

func (m *Adapter) Conn() *gorm.DB {
	return m.db
}
func New(config Config) Adapter {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN: fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Username, config.Password, config.Host, config.Port, config.DBName),
	}), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Errorf("mysql connect err: %v", err)
	}
	return Adapter{config: config, db: db}
}
