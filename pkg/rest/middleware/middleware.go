package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pradeepbepari/golang_microservices/pkg/logger"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/interfaces"
	"github.com/pradeepbepari/golang_microservices/pkg/rest/models"
	"go.opentelemetry.io/otel"
)

type Controller struct {
	service interfaces.Service
	logger  *logger.Logger
}

func NewController(service interfaces.Service, logger *logger.Logger) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}
func (s *Controller) Create(c *gin.Context) {
	tracer := otel.Tracer("Middleware")
	ctx, span := tracer.Start(c.Request.Context(), "Create")
	defer span.End()
	var req models.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
