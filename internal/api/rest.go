package api

import (
	"go-auth/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestAPIHandler struct {
	AuthService service.AuthenticateService
}

func (ra RestAPIHandler) Token(c *gin.Context) {
	log.Println("Request for token")

	clientId := c.Query("clientId")
	clientSecret := c.Query("clientSecret")

	if len(clientId) == 0 {
		log.Println("Client Id is either missing or invalid")
		c.String(http.StatusUnauthorized, "Access Denied")
		return
	}

	if len(clientSecret) == 0 {
		log.Println("Client Secret is either missing or invalid")
		c.String(http.StatusUnauthorized, "Access Denied")
		return
	}

	token, err := ra.AuthService.GenerateJwtToken(clientId, clientSecret)

	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusUnauthorized, "Access Denied")
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ra RestAPIHandler) WellKnown(c *gin.Context) {
	output := ra.AuthService.WellKnown()

	c.JSON(http.StatusOK, output)
}

func (ra RestAPIHandler) Jwks(c *gin.Context) {
	output, err := ra.AuthService.Jwks()

	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusUnauthorized, "Access Denied")
	}

	c.JSON(http.StatusOK, output)
}
