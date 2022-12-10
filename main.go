package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
	docs "github.com/AlperRehaYAZGAN/temp-file-transfer-app/docs"
	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/handlers"
	plugins "github.com/AlperRehaYAZGAN/temp-file-transfer-app/plugins"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Path: temp-upload-service
// @Title Temp File Upload Service API
// @Description alya.temp-file.upload-service : microservice for temporary upload and retrieve file operations.
// @Version 1.0.0
// @Schemes http https
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	// server name
	APP_NAME = "alya.temp-file.upload-service"
	// server description
	APP_DESCRIPTION = "alya.temp-file.upload-service : microservice for temporary upload and retrieve file operations."
)

func main() {
	// parse application envs
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("INIT: Cannot get current working directory os.Getwd()")
	}
	config.ReadConfig(dir)
	config.Pwd = dir

	// init env
	env := config.C.App.Env
	port := config.C.App.Port
	// log env and port like "alya.temp-file.upload-service env: dev, port: 9097"
	log.Printf("INIT: %s env: %s, port: %s", APP_NAME, env, port)

	// create 3th party connections
	inAppCache := plugins.NewInAppCacheStore(time.Minute)

	// create application service
	uploadsvc := handlers.NewUploadService(
		inAppCache,
	)

	// check env and set gin mode
	gin.SetMode(gin.ReleaseMode)
	if env == "prod" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	uploadsvc.InitRouter(router)
	uploadsvc.InitUploadsOldFileCleaner()

	// check env and set swagger
	// if !(env == "prod" || env == "production") {
	docs.SwaggerInfo.BasePath = "/api/v1/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// }

	// start server
	port = os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	log.Println("INIT: Application " + APP_NAME + " started on port " + port)

	// fatal if error
	log.Fatal(router.Run(":" + port))
}
