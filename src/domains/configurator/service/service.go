package service

import (
	"context"
	"errors"

	"github.com/coma/coma/src/domains/configurator/dto"
	"github.com/coma/coma/src/domains/configurator/model"
	"github.com/coma/coma/src/domains/configurator/repository"
	selfextsvc "github.com/coma/coma/src/external/self/service"
	"github.com/rs/zerolog/log"
)

type Servicer interface {
	GetClientConfiguration(ctx context.Context, req dto.RequestGetConfiguration) (dto.ResponseGetClientConfiguration, error)
	GetConfiguration(ctx context.Context, req dto.RequestGetConfiguration) (dto.ResponseGetConfigurations, error)
	SetConfiguration(ctx context.Context, req dto.RequestSetConfiguration) error
	UpdateConfiguration(ctx context.Context, req dto.RequestUpdateConfiguration) error
	DeleteConfiguration(ctx context.Context, req dto.RequestDeleteConfiguration) error
}

type Service struct {
	selfExtSvc selfextsvc.WSServicer
	readerRepo repository.RepositoryReader
	writerRepo repository.RepositoryWriter
}

type ServiceOption func(svc *Service)

func SetExternalService(selfExtService selfextsvc.WSServicer) ServiceOption {
	return func(svc *Service) {
		svc.selfExtSvc = selfExtService
	}
}

func SetRepository(reader repository.RepositoryReader, writer repository.RepositoryWriter) ServiceOption {
	return func(svc *Service) {
		svc.readerRepo = reader
		svc.writerRepo = writer
	}
}

func New(opts ...ServiceOption) Servicer {
	svc := &Service{}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *Service) GetClientConfiguration(ctx context.Context, req dto.RequestGetConfiguration) (dto.ResponseGetClientConfiguration, error) {
	var (
		response dto.ResponseGetClientConfiguration
		err      error
	)

	configurations, err := s.readerRepo.FindClientConfiguration(ctx, model.FilterConfiguration{
		ClientKey: req.XClientKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetConfiguration] error FindClientConfiguration")
		return response, err
	}

	response, err = dto.NewResponseGetClientConfiguration[model.Configurations](configurations)
	if err != nil {
		log.Error().Err(err).Msg("[GetConfiguration] error NewResponseGetConfiguration")
		return response, err
	}

	return response, nil
}

func (s *Service) GetConfiguration(ctx context.Context, req dto.RequestGetConfiguration) (dto.ResponseGetConfigurations, error) {
	var (
		response dto.ResponseGetConfigurations
		err      error
	)

	configurations, err := s.readerRepo.FindClientConfiguration(ctx, model.FilterConfiguration{
		ClientKey: req.XClientKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetConfiguration] error FindClientConfiguration")
		return response, err
	}

	response = dto.NewResponseGetConfigurations(configurations)

	return response, nil
}

func (s *Service) SetConfiguration(ctx context.Context, req dto.RequestSetConfiguration) error {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("[SetConfiguration] error validate dto")
		return err
	}

	clientConfigurations, err := s.readerRepo.FindClientConfiguration(ctx, model.FilterConfiguration{
		ClientKey: req.XClientKey,
		Field:     req.Field,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("field", req.Field).
			Msg("[SetConfiguration] error on search configuration")
		return err
	}
	if clientConfigurations.Exists() {
		log.Error().
			Err(err).
			Str("field", req.Field).
			Msg("[SetConfiguration] error duplicate field name")
		return errors.New("err: duplicate field name")
	}

	_, err = s.writerRepo.SetConfiguration(ctx, req.Configuration())
	if err != nil {
		log.Error().Err(err).Msg("[SetConfiguration] error SetConfiguration")
		return err
	}

	return nil
}

func (s *Service) UpdateConfiguration(ctx context.Context, req dto.RequestUpdateConfiguration) error {
	clientConfigurations, err := s.readerRepo.FindClientConfiguration(ctx, model.FilterConfiguration{
		ClientKey: req.XClientKey,
		Id:        req.Id,
		Field:     req.Field,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("field", req.Field).
			Msg("[UpdateConfiguration] error on search configuration")
		return err
	}
	if !clientConfigurations.Exists() {
		log.Error().
			Err(err).
			Str("field", req.Field).
			Msg("[UpdateConfiguration] error configuration is empty")
		return errors.New("err: configuration is empty")
	}

	var (
		configuration        = req.Configuration()
		configurations       = model.Configurations{configuration}
		mapConfigurationById = configurations.MapConfigurationById()
	)

	clientConfigurations.Update(mapConfigurationById)

	for _, configuration := range clientConfigurations {
		err = s.writerRepo.UpdateConfiguration(ctx, configuration)
		if err != nil {
			log.Error().
				Err(err).
				Str("field", req.Field).
				Msg("[UpdateConfiguration] error on update configuration")
			return err
		}
	}

	return nil
}

func (s *Service) DeleteConfiguration(ctx context.Context, req dto.RequestDeleteConfiguration) error {
	err := s.writerRepo.DeleteConfiguration(ctx, req.FilterConfiguration())
	if err != nil {
		log.Error().Err(err).Msg("[DeleteConfiguration] error when deleting configuration")
		return err
	}

	return nil
}
