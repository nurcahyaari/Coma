package repository

import (
	"context"

	"github.com/coma/coma/infrastructure/database"
	"github.com/coma/coma/src/domain/entity"
	"github.com/coma/coma/src/domain/repository"
	"github.com/ostafen/clover"
)

type RepositoryUserWrite struct {
	dbName string
	db     *database.Clover
}

func NewRepositoryUserWriter(db *database.Clover, name string) repository.RepositoryUserAuthWriter {
	db.DB.CreateCollection(name)
	return &RepositoryUserWrite{
		db:     db,
		dbName: name,
	}
}

func (r *RepositoryUserWrite) CreateUserToken(ctx context.Context, userAuth entity.UserAuth) error {
	dataMap, err := userAuth.MapStringInterface()
	if err != nil {
		return err
	}

	doc := clover.NewDocument()
	doc.SetAll(dataMap)

	_, err = r.db.DB.InsertOne(r.dbName, doc)
	if err != nil {
		return err
	}

	return nil
}
