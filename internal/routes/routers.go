package routes

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/configs"
	"github.com/DropKit/DropKit-Adapter/internal/services/health"
	"github.com/DropKit/DropKit-Adapter/pkg/logging"
	"github.com/gin-gonic/gin"
)

var logger logging.Logger

func init() {
	logger, _ = logging.NewLogger(configs.Configs.Adapter.LoggerLevel, "routes")
}

// SetupRouter Create a new router object
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	corsMiddleware := func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "User-Agent, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", http.MethodPost)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
	r.Use(corsMiddleware)

	healthRoute := r.Group("/health")
	{
		healthRoute.GET("/dependency", health.CheckDependencyServices)
		healthRoute.GET("/ping", health.Ping)
	}

	// userRoute := r.Group("/user")
	// {
	// 	userRoute.GET("/create", user.GenerateRandomAccount)
	// }

	// // Add routing group for DB operations
	// dbRoute := r.Group("/db")
	// {
	// 	dbRoute.POST("/create", db.HandleDBCreation)
	// 	dbRoute.POST("/insert", db.HandleDBInsertion)
	// 	dbRoute.POST("/select", db.HandleDBSelection)
	// 	dbRoute.POST("/update", db.HandleDBUpdate)
	// 	dbRoute.POST("/delete", db.HandleDBDeletion)
	// }

	// grantRoute := r.Group("/permission/grant/table")
	// {
	// 	grantRoute.POST("/owner", permission.GrantTableOwner)
	// 	grantRoute.POST("/maintainer", permission.GrantTableMaintainer)
	// 	grantRoute.POST("/viewer", permission.GrantTableViewer)
	// }

	// revokeRoute := r.Group("/permission/revoke/table")
	// {
	// 	revokeRoute.POST("/owner", permission.RevokeTableOwner)
	// 	revokeRoute.POST("/maintainer", permission.RevokeTableMaintainer)
	// 	revokeRoute.POST("/viewer", permission.RevokeTableViewer)
	// }

	// verifyRoute := r.Group("/permission/verify/table")
	// {
	// 	verifyRoute.POST("/owner", permission.VerifyTableOwner)
	// 	verifyRoute.POST("/maintainer", permission.VerifyTableMaintainer)
	// 	verifyRoute.POST("/viewer", permission.TableViewer)
	// }

	// paymentRoute := r.Group("/payment")
	// {
	// 	paymentRoute.POST("/mint", payment.MintToken)
	// 	paymentRoute.POST("/burn", payment.BurnToken)
	// 	paymentRoute.POST("/transfer", payment.TransferToken)
	// 	paymentRoute.POST("/balance", payment.GetAccountBalance)
	// }

	// roleRoute := r.Group("/role")
	// {
	// 	roleRoute.POST("/create", role.CreateColumnRole)
	// }
	return r
}
