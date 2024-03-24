package repository

import (
	"context"

	"github.com/coma/coma/src/domain/entity"
)

//counterfeiter:generate . RepositoryApplicationWriter
type RepositoryApplicationWriter interface {
	CreateApplication(ctx context.Context, data entity.Application) error
	DeleteApplication(ctx context.Context, filter entity.FilterApplication) error
}

//counterfeiter:generate . RepositoryApplicationReader
type RepositoryApplicationReader interface {
	FindApplication(ctx context.Context, filter entity.FilterApplication) (entity.Application, bool, error)
	FindApplications(ctx context.Context, filter entity.FilterApplication) (entity.Applications, error)
}
