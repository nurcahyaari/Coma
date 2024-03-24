package repository

import (
	"context"

	"github.com/coma/coma/src/domain/entity"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RepositoryAuthReader
type RepositoryAuthReader interface {
	FindTokenById(ctx context.Context, id int64) (entity.Apikey, error)
	FindTokenByToken(ctx context.Context, token string) (entity.Apikey, error)
}

//counterfeiter:generate . RepositoryAuthWriter
type RepositoryAuthWriter interface{}

//counterfeiter:generate . AuthRepositorier
type AuthRepositorier interface {
	NewRepositoryReader() RepositoryAuthReader
	NewRepositoryWriter() RepositoryAuthWriter
	NewRepositoryUserAuthReader() RepositoryUserAuthReader
	NewRepositoryUserAuthWriter() RepositoryUserAuthWriter
}

//counterfeiter:generate . RepositoryOAuthReader
type RepositoryOAuthReader interface{}

//counterfeiter:generate . RepositoryOAuthWriter
type RepositoryOAuthWriter interface{}

//counterfeiter:generate . OauthRepositorier
type OauthRepositorier interface {
	NewRepositoryOAuthReader() RepositoryOAuthReader
	NewRepositoryOAuthWriter() RepositoryOAuthWriter
}

//counterfeiter:generate . RepositoryUserAuthReader
type RepositoryUserAuthReader interface {
	FindTokenBy(ctx context.Context, filter entity.FilterUserAuth) (*entity.UserAuth, error)
}

//counterfeiter:generate . RepositoryUserAuthWriter
type RepositoryUserAuthWriter interface {
	CreateUserToken(ctx context.Context, userAuth entity.UserAuth) error
}
