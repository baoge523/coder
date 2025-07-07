package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func init() {
	prometheus.Register(requestCounter)
	prometheus.Register(requestDurationTime)
}

var (
	requestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_total",
		Help: "all request",
	}, []string{"status", "method", "path"})

	/**
	request_duration_time_bucket{method="POST",path="/v1/test",le="5"} 0
	request_duration_time_bucket{method="POST",path="/v1/test",le="10"} 0
	request_duration_time_bucket{method="POST",path="/v1/test",le="50"} 0
	request_duration_time_bucket{method="POST",path="/v1/test",le="100"} 1
	request_duration_time_bucket{method="POST",path="/v1/test",le="200"} 1
	request_duration_time_bucket{method="POST",path="/v1/test",le="500"} 1
	request_duration_time_bucket{method="POST",path="/v1/test",le="800"} 1
	request_duration_time_bucket{method="POST",path="/v1/test",le="1000"} 1
	request_duration_time_bucket{method="POST",path="/v1/test",le="+Inf"} 1
	request_duration_time_sum{method="POST",path="/v1/test"} 98
	request_duration_time_count{method="POST",path="/v1/test"} 1
	*/
	// prometheus 中的histogram 是累积直方图
	requestDurationTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_time",
		Help:    "all request duration time",
		Buckets: []float64{5, 10, 50, 100, 200, 500, 800, 1000},
	}, []string{"method", "path"})
)

func PrometheusMiddleware(ctx *gin.Context) {

	start := time.Now()
	ctx.Next() // 执行目标handle
	cost := time.Since(start)
	f := float64(cost / time.Millisecond)
	fmt.Println(f)
	requestDurationTime.WithLabelValues(ctx.Request.Method, ctx.Request.RequestURI).Observe(f)
	status := ctx.Writer.Status()
	requestCounter.WithLabelValues(strconv.Itoa(status), ctx.Request.Method, ctx.Request.RequestURI).Inc()
}
