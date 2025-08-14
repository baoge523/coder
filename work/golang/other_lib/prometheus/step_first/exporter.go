package step_first

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func exporter() {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "ping")
	})

	engine.POST("/metric", gin.WrapF(promhttp.Handler().ServeHTTP))

}
