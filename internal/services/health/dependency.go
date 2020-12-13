package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckDependencyServices(c *gin.Context) {
	// service, err := services.DependencyServicesCheck()
	// if err != nil {
	// 	logger.ErrorDependencyService(service, err)
	// 	c.JSON(http.StatusOK, response.ResponseDependencyError())
	// 	return
	// }

	// c.JSON(http.StatusOK, response.ResponseServerOk())
	// logger.InfoAPIDenpendencyOk()
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	logger.Info("/dependency")
}
