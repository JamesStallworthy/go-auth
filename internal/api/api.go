package api

import (
	"go-auth/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type APIHandler interface {
	Token(c *gin.Context)
	WellKnown(c *gin.Context)
}

func CreateRestAPIHandler(authService service.AuthenticateService, port string) APIHandler {
	handler := RestAPIHandler{AuthService: authService}

	r := gin.Default()
	r.POST("/token", handler.Token)

	r.GET("/.well-known/openid-configuration", handler.WellKnown)

	log.Fatal(r.Run(":" + port))

	return handler
}
