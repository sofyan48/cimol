package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sofyan48/rll-daemon-new/docs/swagger/docs"
	v1 "github.com/sofyan48/rll-daemon-new/src/controller/v1"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	// Declare Swaggers Docs
	// routers.Use(static.Serve("/docs", static.LocalFile("./docs/swagger/docs/swagger.json", false)))
	serverHost := os.Getenv("SWAGGER_SERVER_ADDRESS")
	url := ginSwagger.URL(serverHost + "/swagger/doc.json") // The url pointing to API definition
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// End Declare Swagger

	version1 := &v1.V1RouterLoader{}
	version1.V1Router(routers)
}
