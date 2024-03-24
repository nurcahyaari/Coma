package repository

import (
	"context"
	"fmt"

	"github.com/coma/coma/infrastructure/database"
	"github.com/coma/coma/src/domain/entity"
	"github.com/coma/coma/src/domain/repository"
	"github.com/ostafen/clover"
)

type RepositoryRead struct {
	dbName string
	db     *database.Clover
}

func NewRepositoryReader(db *database.Clover, name string) repository.RepositoryAuthReader {
	db.DB.CreateCollection(name)
	return &RepositoryRead{
		db:     db,
		dbName: name,
	}
}

func (r *RepositoryRead) FindTokenById(ctx context.Context, id int64) (entity.Apikey, error) {
	var apiKey entity.Apikey
	doc, err := r.db.DB.Query("apikey").FindById(fmt.Sprintf("%d", id))
	if err != nil {
		return apiKey, err
	}

	err = doc.Unmarshal(&apiKey)
	if err != nil {
		return apiKey, err
	}

	return apiKey, nil
}

func (r *RepositoryRead) FindTokenByToken(ctx context.Context, token string) (entity.Apikey, error) {
	var apiKey entity.Apikey

	criteria := clover.Field("key").
		Eq(token)

	doc, err := r.db.DB.Query("apikey").
		Where(criteria).
		FindFirst()
	if err != nil {
		return apiKey, err
	}

	err = doc.Unmarshal(&apiKey)
	if err != nil {
		return apiKey, err
	}

	return apiKey, nil
}
