package service

import (
	"context"

	"github.com/pradeepbepari/golang_microservices/pkg/logger"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/interfaces"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/models"
)

type services struct {
	repository interfaces.Repository
	logger     *logger.Logger
}

func NewService(repository interfaces.Repository, logger *logger.Logger) *services {
	return &services{
		repository: repository,
		logger:     logger,
	}
}
func (r *services) Create(ctx context.Context, u models.User) (*models.User, error) {
	return r.repository.Create(ctx, u)
}
