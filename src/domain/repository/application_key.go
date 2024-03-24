package repository

import (
	"context"

	"github.com/coma/coma/src/domain/entity"
)

//counterfeiter:generate . RepositoryApplicationKeyWriter
type RepositoryApplicationKeyWriter interface {
	CreateOrSaveApplicationKey(ctx context.Context, data entity.ApplicationKey) error
}

//counterfeiter:generate . RepositoryApplicationKeyReader
type RepositoryApplicationKeyReader interface {
	FindApplicationKey(ctx context.Context, filter entity.FilterApplicationKey) (entity.ApplicationKey, error)
}
