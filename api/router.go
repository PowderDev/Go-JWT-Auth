package api

import (
	"encoding/json"

	apiExceptions "auth/api/exceptions"
	"auth/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(api *config.API) *gin.Engine {
	r := gin.New()
	setupDefaultMiddlewares(r)

	r.POST("/registration", api.Handlers.Registration)
	r.POST("/login", api.Handlers.Login)
	r.POST("/logout", api.Handlers.Logout)
	r.POST("/refresh", api.Handlers.Refresh)

	return r
}

func setupDefaultMiddlewares(r *gin.Engine) {
	r.Use(errorHandler())
	r.Use(cors.Default())
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.Last()
		if lastError == nil {
			return
		}

		var errData apiExceptions.HTTPError
		_ = json.Unmarshal([]byte(lastError.Error()), &errData)

		c.JSON(errData.StatusCode, gin.H{"error": errData})
	}
}
