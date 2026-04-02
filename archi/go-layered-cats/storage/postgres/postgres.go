package postgres

import (
	"go-layered-cats/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (*storage.Store, error) {
	// Implémenter la logique de connexion à la base de données PostgreSQL
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	var storeCat PostgresStoreCat
	storeCat.db = db
	var storeDog PostgresStoreDog
	storeDog.db = db

	return &storage.Store{
		Cat: &storeCat,
		Dog: &storeDog,
	}, nil
}
