package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDb struct {
	db *gorm.DB
}

func (p *PostgresDb) Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *PostgresDb) DB() *gorm.DB {
	return p.db
}

func NewPostgresDb(dsn string) (*PostgresDb, error) {
	db := &PostgresDb{}

	if err := db.Connect(dsn); err != nil {
		return nil, err
	}

	return db, nil
}
