package repository

import (
	"context"

	"github.com/coma/coma/src/domain/entity"
)

//counterfeiter:generate . RepositoryUserReader
type RepositoryUserReader interface {
	FindUser(ctx context.Context, filter entity.FilterUser) (entity.User, error)
	FindUsers(ctx context.Context, filter entity.FilterUser) (entity.Users, error)
}

//counterfeiter:generate . RepositoryUserWriter
type RepositoryUserWriter interface {
	SaveUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, filter entity.FilterUser) error
	UpdateUser(ctx context.Context, user entity.User) error
}

//counterfeiter:generate . RepositoryUserApplicationScopeReader
type RepositoryUserApplicationScopeReader interface {
	FindUserApplicationScope(ctx context.Context, filter entity.FilterUserApplicationScope) (entity.UserApplicationScope, bool, error)
	FindUserApplicationsScope(ctx context.Context, filter entity.FilterUserApplicationScope) (entity.UserApplicationsScope, error)
}

//counterfeiter:generate . RepositoryUserApplicationScopeWriter
type RepositoryUserApplicationScopeWriter interface {
	SaveUserApplicationScope(ctx context.Context, userApplicationScope entity.UserApplicationScope) error
	UpdateUserApplicationScope(ctx context.Context, userApplicationScope entity.UserApplicationScope) error
	RevokeUserApplicationScope(ctx context.Context, filter entity.FilterUserApplicationScope) error
}
