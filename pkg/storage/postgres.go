package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDb struct {
	db *gorm.DB
}

func (p *postgresDb) Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *postgresDb) DB() *gorm.DB {
	return p.db
}

func NewStorage(dsn string) (Storage, error) {
	db := &postgresDb{}

	if err := db.Connect(dsn); err != nil {
		return nil, err
	}

	return db, nil
}
