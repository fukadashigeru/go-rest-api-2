package inject

import (
	"go-rest-api/db"

	"gorm.io/gorm"
)

type Injector struct {
	db *gorm.DB
}

func NewInjector(db *gorm.DB) Injector {
	return Injector{db}
}

func (i *Injector) DB() *gorm.DB {
	return db.NewDB()
}
