package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pradeepbepari/golang_microservices/pkg/logger"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/models"
)

type repository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewRepository(db *sql.DB, logger *logger.Logger) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
func (r *repository) Create(ctx context.Context, u models.User) (*models.User, error) {
	u.ID = uuid.New()
	u.Name = "pradeep bepari"
	return &u, nil
}
