package http

import (
	"github.com/coma/coma/container"
	service "github.com/coma/coma/src/domain/service"
	"github.com/go-chi/chi/v5"
)

type HttpHandle struct {
	authSvc             service.AuthServicer
	configurationSvc    service.ApplicationConfigurationServicer
	applicationStageSvc service.ApplicationStageServicer
	applicationSvc      service.ApplicationServicer
	applicationKeySvc   service.ApplicationKeyServicer
	userSvc             service.UserServicer
}

func (h HttpHandle) Router(r *chi.Mux) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Route("/oauth", func(r chi.Router) {
				r.Post("/login", h.OauthLogin)
			})
		})

		r.Route("/applications", func(r chi.Router) {
			r.Get("/", h.FindApplications)
			r.Post("/", h.CreateApplication)
			r.Delete("/{applicationId}", h.DeleteApplications)
		})

		r.Route("/stages", func(r chi.Router) {
			r.Get("/", h.FindApplicationStages)
			r.Post("/", h.CreateApplicationStages)
			r.Delete("/{stageName}", h.DeleteApplicationStages)
		})

		r.Route("/keys", func(r chi.Router) {
			r.Get("/", h.FindApplicationKey)
			r.Post("/", h.CreateOrUpdateApplicationKey)
		})

		r.Route("/configuration", func(r chi.Router) {
			r.Use(h.MiddlewareCheckIsClientKeyExists)
			r.Get("/", h.GetConfiguration)
			r.Post("/", h.SetConfiguration)
			r.Put("/", h.UpdateConfiguration)
			r.Post("/upsert", h.UpsertConfiguration)
			r.Delete("/{id}", h.DeleteConfiguration)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.FindUsers)
			r.Get("/{id}", h.FindUser)
			r.Post("/", h.CreateUser)
			r.Delete("/{id}", h.DeleteUser)
			r.Put("/{id}", h.UpdateUser)
			r.Patch("/password/{id}", h.UpdateUserPassword)
		})
	})
}

func NewHttpHandler(c container.Service) *HttpHandle {
	httpHandle := &HttpHandle{
		authSvc:             c.AuthServicer,
		configurationSvc:    c.ApplicationConfigurationServicer,
		applicationStageSvc: c.ApplicationStageServicer,
		applicationSvc:      c.ApplicationServicer,
		applicationKeySvc:   c.ApplicationKeyServicer,
		userSvc:             c.UserServicer,
	}
	return httpHandle
}
