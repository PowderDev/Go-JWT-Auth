package apiHelpers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name string, value string) {
	domain := os.Getenv("DOMAIN")

	//one day
	maxAge := 86400

	c.SetCookie(name, value, maxAge, "/", domain, false, true)
}

func DeleteCookie(c *gin.Context, name string) {
	domain := os.Getenv("DOMAIN")
	c.SetCookie(name, "", 0, "/", domain, false, true)
}
