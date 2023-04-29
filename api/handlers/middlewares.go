package api

import (
	"strings"

	apiExceptions "auth/api/exceptions"
	"auth/api/services"
	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct {
	Services *api.Services
}

func (h *MiddlewareHandler) auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.Error(apiExceptions.Error401())
		return
	}

	accessToken := strings.Split(authHeader, " ")[1]

	if accessToken == "" {
		c.Error(apiExceptions.Error401())
		return
	}

	payload, err := h.Services.TokenService.Validate(accessToken, api.AccessTokenType)

	if err != nil {
		c.Error(err)
		return
	}

	c.Set("userID", payload.Sub)
}
