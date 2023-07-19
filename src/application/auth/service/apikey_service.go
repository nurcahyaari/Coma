package service

import (
	"context"
	"fmt"

	"github.com/coma/coma/src/application/auth/dto"
	"github.com/coma/coma/src/domains/entity"
	"github.com/coma/coma/src/domains/repository"
	"github.com/coma/coma/src/domains/service"
	"github.com/rs/zerolog/log"
)

type ApiKeyService struct {
	repositoryReader repository.RepositoryAuthReader
	repositoryWriter repository.RepositoryAuthWriter
}

type ApiKeyServiceOption func(s *ApiKeyService)

func SetApiKeyRepository(repositoryReader repository.RepositoryAuthReader, repositoryWriter repository.RepositoryAuthWriter) ApiKeyServiceOption {
	return func(s *ApiKeyService) {
		s.repositoryWriter = repositoryWriter
		s.repositoryReader = repositoryReader
	}
}

func NewApiKey(opts ...ApiKeyServiceOption) service.ApiKeyServicer {
	svc := &ApiKeyService{}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *ApiKeyService) ValidateToken(ctx context.Context, request dto.RequestValidateToken) (dto.ResponseValidateKey, error) {
	var (
		resp   = dto.ResponseValidateKey{}
		apikey entity.Apikey
		err    error
	)

	apikey, err = s.repositoryReader.FindTokenByToken(ctx, request.Token)
	if err != nil {
		log.Error().Err(err).Msg("[ApiKeyService][ValidateToken] err on FindTokenById")
		return resp, err
	}

	if apikey.Id == 0 {
		log.Error().Err(err).Msg("[ApiKeyService][ValidateToken] err token is not found")
		return resp, fmt.Errorf("error: token is not found")
	}

	return dto.ResponseValidateKey{
		Valid: true,
	}, nil
}

func (s *ApiKeyService) CreateApplicationKey(ctx context.Context) error {
	return nil
}
