package repository

import (
	"github.com/coma/coma/infrastructure/database"
	"github.com/coma/coma/src/domain/repository"
)

type Repository struct {
	db *database.Clover
}

func New(db *database.Clover) repository.AuthRepositorier {
	return &Repository{
		db: db,
	}
}

func (r Repository) NewRepositoryReader() repository.RepositoryAuthReader {
	return NewRepositoryReader(r.db, "apikey")
}

func (r Repository) NewRepositoryWriter() repository.RepositoryAuthWriter {
	return NewRepositoryWriter(r.db, "apikey")
}

func (r Repository) NewRepositoryUserAuthReader() repository.RepositoryUserAuthReader {
	return NewRepositoryUserReader(r.db, "user_auth")
}

func (r Repository) NewRepositoryUserAuthWriter() repository.RepositoryUserAuthWriter {
	return NewRepositoryUserWriter(r.db, "user_auth")
}
