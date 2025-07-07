package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"projects/services/gateway/prometheus"
	"projects/services/gateway/receiver"
)

func InitRoute(engine *gin.Engine) {

	{
		groupV1 := engine.Group("/v1")
		groupV1.Use(prometheus.PrometheusMiddleware)
		groupV1.POST("/report", receiver.Receive)
		groupV1.GET("/test", receiver.Test)
	}

	// prometheus
	{
		engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

}
