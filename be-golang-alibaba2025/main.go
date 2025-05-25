package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	cfg "github.com/ZihxS/be-alibabacloud-genai-2025/config"
	"github.com/ZihxS/be-alibabacloud-genai-2025/constant"
	"github.com/ZihxS/be-alibabacloud-genai-2025/handler"
	"github.com/ZihxS/be-alibabacloud-genai-2025/helper"
	"github.com/ZihxS/be-alibabacloud-genai-2025/middleware"
	"github.com/ZihxS/be-alibabacloud-genai-2025/product"
	"github.com/ZihxS/be-alibabacloud-genai-2025/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	isDocker := os.Getenv("IS_DOCKER") == "true"

envLocation := ".env" // default
if isDocker {
    log.Println("Running in Docker. Loading default .env in container...")
    envLocation = ".env" // let docker handle file location via mounting
} else {
    log.Println("Running locally. Using project root .env...")
    _, b, _, _ := runtime.Caller(0)
    projectRootPath := filepath.Join(filepath.Dir(b), "")
    envLocation = filepath.Join(projectRootPath, ".env")
}

zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
if isDocker {
    zerolog.SetGlobalLevel(zerolog.InfoLevel)
} else {
    zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

if err := godotenv.Load(envLocation); err != nil {
    log.Fatal("Error loading .env file: ", err.Error())
}
	constant.InitDBConstant()
	constant.InitAuthConstant()
	constant.InitRedisConstant()
	constant.InitOSSConstant()

	rateLimitStorage := cfg.InitLimiterStorage(isDocker)
	app := gin.New()
	app.Use(gin.Recovery())

	app.Use(ratelimit.RateLimiter(rateLimitStorage, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, info ratelimit.Info) {
			response := helper.BasicAPIResponse(http.StatusTooManyRequests)
			ctx.JSON(http.StatusTooManyRequests, response)
		},
		KeyFunc: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Host", "Origin", "Content-Length", "Content-Type", "Authorization", "User-Agent", "X-Forwarded-For", "Accept-Encoding", "Connection", "X-API-Key", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Encoding", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.Use(logger.SetLogger(logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
		return l.Output(gin.DefaultWriter).With().Logger()
	})))

	// initial redis
	redisClient := cfg.InitRedis()

	// initial database
	db := cfg.InitDB(isDocker, redisClient)

	productRepository := product.NewRepository(db)
	userRepository := user.NewRepository(db)

	productService := product.NewService(productRepository)
	userService := user.NewService(userRepository)

	aiHandler := handler.NewAIHandler(redisClient)
	productHandler := handler.NewProductHandler(redisClient, productService)
	userHandler := handler.NewUserHandler(redisClient, userService)

	_ = aiHandler

	// for activate release mode
	if isDocker {
		gin.SetMode(gin.ReleaseMode)
	}

	app.SetTrustedProxies(nil)
	app.Static("/images", "./images")
	app.Use(gzip.Gzip(gzip.BestCompression))
	app.RedirectTrailingSlash = false

	api := app.Group("/api/v1", middleware.Auth())
	{
		product := api.Group("/products")
		{
			product.GET("", productHandler.GetAll)
			product.GET("/:id", productHandler.GetByID)
			product.POST("", productHandler.Add)
			product.PATCH("/:id", productHandler.Update)
		}

		order := api.Group("/orders")
		{
			order.GET("", nil)
			order.GET("/:id", nil)
			order.POST("", nil)
			order.PATCH("", nil)
		}

		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetByID)
		}

		ai := api.Group("/ai")
		{
			ai.GET("", nil)
		}
	}

	_ = api

	// handle invalid method
	app.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, helper.BasicAPIResponse(http.StatusMethodNotAllowed))
	})

	// handle invalid path or invalid endpoint
	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, helper.BasicAPIResponse(http.StatusNotFound))
	})

	// run http server
	app.Run(os.Getenv("APP_RUN_ON"))
}
