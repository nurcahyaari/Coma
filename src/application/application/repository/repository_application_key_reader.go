package repository

import (
	"context"

	"github.com/coma/coma/infrastructure/database"
	"github.com/coma/coma/src/domain/entity"
	"github.com/coma/coma/src/domain/repository"
)

type RepositoryApplicationKeyRead struct {
	dbName string
	db     *database.Clover
}

func NewRepositoryApplicationKeyReader(db *database.Clover, name string) repository.RepositoryApplicationKeyReader {
	db.DB.CreateCollection(name)
	return &RepositoryApplicationKeyRead{
		db:     db,
		dbName: name,
	}
}

func (r *RepositoryApplicationKeyRead) FindApplicationKey(ctx context.Context, filter entity.FilterApplicationKey) (entity.ApplicationKey, error) {
	var applicationKey entity.ApplicationKey

	doc, err := r.db.DB.
		Query(r.dbName).
		Where(filter.Filter()).
		FindFirst()
	if err != nil {
		return applicationKey, err
	}
	if doc == nil {
		return applicationKey, nil
	}

	err = doc.Unmarshal(&applicationKey)
	if err != nil {
		return applicationKey, err
	}

	return applicationKey, nil
}
