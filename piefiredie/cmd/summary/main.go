package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"piefiredie/internal/bootstrap"
	"piefiredie/internal/route"
)

func main() {
	bootstrap.LoadEnv()

	gin.SetMode(bootstrap.Env.GinMode)
	r := gin.Default()
	r.Use(otelgin.Middleware(bootstrap.Env.AppName))

	r.GET("/health", route.Health)

	routerBeef := r.Group("/beef")
	routerBeef.GET("/summary", route.BeefStockSummary)

	if err := r.Run(bootstrap.Env.ServerAddress); err != nil {
		log.Fatalf("failed to start the server: %e", err)
	}
}
