package sqlite

import (
	"go-layered-cats/storage"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbName string) (*storage.Store, error) {
	// Implémenter la logique de connexion à la base de données SQLite
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
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

	var store storage.Store
	store.Cat = &SqliteStoreCat{db: db}
	store.Dog = &SqliteStoreDog{db: db}

	return &store, nil
}
