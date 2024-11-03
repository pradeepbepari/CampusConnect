package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pradeepbepari/golang_microservices/pkg/logger"
	controller "github.com/pradeepbepari/golang_microservices/pkg/rest/middleware"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/repository"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/service"
)

func InatiliazeCependencies(db *sql.DB, router *gin.Engine, logger *logger.Logger) {
	repository := repository.NewRepository(db, logger)
	serviceHandular := service.NewService(repository, logger)
	controllerHandular := controller.NewController(serviceHandular, logger)
	registerRoutes(router, controllerHandular)
}
