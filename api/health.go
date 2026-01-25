package api

import (
	"github.com/gin-gonic/gin"
	"zeusro.com/hermes/internal/core/logprovider"
	"zeusro.com/hermes/internal/core/webprovider"
	"zeusro.com/hermes/internal/service"
)

type IndexRoutes struct {
	logger      logprovider.Logger
	gin         webprovider.MyGinEngine
	health      service.HealthService
	hermes      service.TranslateService
	userHandler *UserHandler
	// m middleware.JWTMiddleware
}

func NewIndexRoutes(logger logprovider.Logger, gin webprovider.MyGinEngine,
	s service.HealthService, herms service.TranslateService, userHandler *UserHandler) IndexRoutes {
	return IndexRoutes{
		logger:      logger,
		gin:         gin,
		health:      s,
		hermes:      herms,
		userHandler: userHandler,
	}
}

func (r IndexRoutes) SetUp() {

	r.gin.Gin.GET("/index", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.gin.Gin.GET("/translate", func(c *gin.Context) {
		c.File("./static/translate.html")
	})
	r.gin.Gin.POST("/translate", r.hermes.Translate)

	index := r.gin.Gin.Group("/api")
	{
		//http://localhost:8080/api/health
		index.OPTIONS("health", r.health.Check)
		index.GET("health", r.health.Check)
		index.OPTIONS("healthz", r.health.Check)
		index.GET("healthz", r.health.Check)

		// User routes
		users := index.Group("/users")
		{
			users.POST("", r.userHandler.CreateUser)
			users.GET("", r.userHandler.ListUsers)
			users.GET(":id", r.userHandler.GetUser)
		}

		// Prometheus metrics endpoint
		index.GET("/metrics", func(c *gin.Context) {
			// Prometheus handler will be added via middleware
			c.Next()
		})
	}

}
