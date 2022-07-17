package api

import (
	"go-auth/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://example.com/?clientId=test&clientSecret=secret", nil)

	handler := RestAPIHandler{service.MockService{}}

	handler.Token(c)

	assert.Equal(t, 200, w.Code) // or what value you need it to be
	assert.Equal(t, "{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0\"}", w.Body.String())
}

func TestTokenMissingClientId(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://example.com/?clientSecret=secret", nil)

	handler := RestAPIHandler{service.MockService{}}

	handler.Token(c)

	assert.Equal(t, 401, w.Code) // or what value you need it to be
	assert.Equal(t, "Access Denied", w.Body.String())
}

func TestTokenMissingClientSecret(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://example.com/?clientId=test", nil)

	handler := RestAPIHandler{service.MockService{}}

	handler.Token(c)

	assert.Equal(t, 401, w.Code) // or what value you need it to be
	assert.Equal(t, "Access Denied", w.Body.String())
}
