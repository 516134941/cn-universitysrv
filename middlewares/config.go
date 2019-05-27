package middlewares

import (
	"cn-universitysrv/config"

	"github.com/gin-gonic/gin"
)

// Config .
func Config(tomlConfig *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", tomlConfig)
		c.Next()
	}
}
