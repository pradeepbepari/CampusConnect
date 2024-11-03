package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/pradeepbepari/golang_microservices/pkg/rest/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func registerRoutes(r *gin.Engine, handular *controller.Controller) {
	router := r.Group("/api")
	router.Use(otelgin.Middleware("golang"))
	router.Use(handleOptionsRequests())
	router.POST("/", handular.Create)
}
func handleOptionsRequests() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Header("Access-Control-Allow-origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Authorization,Accept")
			ctx.Set("requestID", "12052002")
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
