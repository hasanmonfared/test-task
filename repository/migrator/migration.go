package migrator

import (
	"app/model"
	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func New(db *gorm.DB) Migrator {
	return Migrator{db: db}
}
func (m Migrator) Up() {
	m.db.Migrator().AutoMigrate(
		&model.User{},
	)
}
