package api

import (
	"fmt"
	"go-auth/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type APIHandler interface {
	Token(c *gin.Context)
	WellKnown(c *gin.Context)
	Jwks(c *gin.Context)
}

func CreateRestAPIHandler(authService service.AuthenticateService, port int) APIHandler {
	handler := RestAPIHandler{AuthService: authService}

	r := gin.Default()
	r.POST("/token", handler.Token)

	r.GET("/.well-known/openid-configuration", handler.WellKnown)

	r.GET("/oauth/jwks", handler.Jwks)

	log.Fatal(r.Run(fmt.Sprintf(": %v", port)))

	return handler
}
