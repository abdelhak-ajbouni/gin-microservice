package routers

import (
	"github.com/abdelhak-ajbouni/gin-microservice/pkg/app"
	"github.com/gin-gonic/gin"
)

type Health struct {
	Healthy bool `json:"healthy"`
}

func GetHealth(c *gin.Context) {
	response := app.NewResponse(c)
	health := Health{
		Healthy: true,
	}
	response.Success("healthy", health)
}
