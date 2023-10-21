package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"url-shortener/api/handlers"
	cacheManagement "url-shortener/pkg/cache-management"
	configReader "url-shortener/pkg/config"
	logHandler "url-shortener/pkg/log"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	path := "E:/WS/WS-GO-URL-SHORTENER/url-shortener/configuration/config.yml"
	configReader.GetInstance(path)
	logHandler.SetupLogger()
	cacheManagement.NewCacheManagerFactory()()

	engine := gin.Default()
	handlers.NewShorterConfigRoutes(engine)

	// Create a CORS configuration with your desired settings.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	engine.Use(cors.New(corsConfig))

	logHandler.Logger().Info("start program")

	// Serve the Swagger documentation on /swagger URL
	// http://localhost:8080/swagger/index.html
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
