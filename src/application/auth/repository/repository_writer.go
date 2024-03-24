package repository

import (
	"github.com/coma/coma/infrastructure/database"
	"github.com/coma/coma/src/domain/repository"
)

type RepositoryWrite struct {
	dbName string
	db     *database.Clover
}

func NewRepositoryWriter(db *database.Clover, name string) repository.RepositoryAuthWriter {
	db.DB.CreateCollection(name)
	return &RepositoryWrite{
		db:     db,
		dbName: name,
	}
}
