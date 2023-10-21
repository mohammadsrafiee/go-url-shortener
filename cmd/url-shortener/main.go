package main

import (
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"url-shortener/api/handlers"
	cacheManagement "url-shortener/pkg/cache"
	configReader "url-shortener/pkg/config-reader"
	logHandler "url-shortener/pkg/log"

	"github.com/gin-gonic/gin"
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
	path := "C:/Users/rafiee/Desktop/WS-GO/src/go-url-shortener/configuration/config.yml"
	configReader.GetInstance(path)
	logHandler.SetupLogger()
	cacheManagement.NewCacheManagerFactory()()

	engine := gin.Default()
	defer func() {
		err := engine.Run(":9191")
		if err != nil {
			return
		}
	}()
	handlers.NewShorterConfigRoutes(engine)
	configSwagger(engine)
	configCORS(engine)
}

func configCORS(engine *gin.Engine) {
	// Create a CORS configuration with your desired settings.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	engine.Use(cors.New(corsConfig))
}

func configSwagger(engine *gin.Engine) {
	//// Serve the Swagger documentation on /swagger URL
	//// http://localhost:8080/swagger/index.html
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		//ginSwagger.URL("http://localhost:9191/docs/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(1),
		ginSwagger.DeepLinking(true)))
}
