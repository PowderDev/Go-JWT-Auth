package api

import (
	"net/http"

	apiExceptions "auth/api/exceptions"
	apiHelpers "auth/api/helpers"
	services "auth/api/services"
	"auth/api/services/dtos"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Services *services.Services
}

func (h *AuthHandler) Registration(c *gin.Context) {
	var authInput dtos.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.Error(apiExceptions.ErrorBodyMismatch())
		return
	}

	user, err := h.Services.AuthService.Authenticate(&authInput)

	if err != nil {
		c.Error(err)
		return
	}

	jwtPayload := &dtos.JWTPayload{Sub: user.ID}
	tokens, tokensError := h.Services.TokenService.GetNewTokens(jwtPayload)

	if tokensError != nil {
		c.Error(tokensError)
		return
	}

	apiHelpers.SetCookie(c, "refresh_token", tokens.RefreshToken)

	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var authInput dtos.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.Error(apiExceptions.ErrorBodyMismatch())
		return
	}

	user, err := h.Services.AuthService.Login(&authInput)

	if err != nil {
		c.Error(err)
		return
	}

	jwtPayload := &dtos.JWTPayload{Sub: user.ID}
	tokens, tokensError := h.Services.TokenService.GetNewTokens(jwtPayload)

	if tokensError != nil {
		c.Error(tokensError)
		return
	}

	apiHelpers.SetCookie(c, "refresh_token", tokens.RefreshToken)

	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		return
	}

	dbErr := h.Services.TokenService.Delete(refreshToken)

	if dbErr != nil {
		c.Error(err)
		return
	}

	apiHelpers.DeleteCookie(c, "refresh_token")
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.Error(apiExceptions.Error401())
		return
	}

	tokens, refreshErr := h.Services.TokenService.Refresh(refreshToken)

	if refreshErr != nil {
		c.Error(refreshErr)
		return
	}

	c.JSON(http.StatusOK, tokens)
}
