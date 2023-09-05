package main

import (
	"duongGoGo/infra/caching"
	"duongGoGo/infra/database"
	"duongGoGo/infra/rabbitmq"
	"duongGoGo/middleware"
	"duongGoGo/modules"
	"duongGoGo/modules/auth"
	"duongGoGo/modules/user"
	"duongGoGo/tasks"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://3016983b82447dfcb2cb8aa81e82258d@o4505792933593088.ingest.sentry.io/4505792935362560",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	//Load the .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	modules.Init()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(sentrygin.New(sentrygin.Options{}))

	database.InitDB()
	database.MigrateDB()
	caching.Init()
	rabbitmq.Init()
	tasks.Init()

	v1 := r.Group("api/v1")
	{
		auth.Routes(v1)
		user.Routes(v1, middleware.JWTMiddleware(), middleware.CSRFMiddleware())
	}

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	err = r.Run(":" + port)
	if err != nil {
		return
	}

	rabbitmq.Close()
}
