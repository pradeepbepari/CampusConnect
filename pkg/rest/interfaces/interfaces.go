package interfaces

import (
	"context"

	"github.com/pradeepbepari/golang_microservices/pkg/rest/models"
)

type Service interface {
	Create(context.Context, models.User) (*models.User, error)
}
type Repository interface {
	Create(context.Context, models.User) (*models.User, error)
}
