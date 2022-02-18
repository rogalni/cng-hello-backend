package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupMetrics(r *gin.Engine) {
	r.GET("/metrics", getMetrics)
}

func getMetrics(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)

}
