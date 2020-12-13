package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping responds `pong` for health check
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	logger.Info("/ping")
}
