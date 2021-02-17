package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	// 记一些日志
	if err != nil {
		panic(err)
	}
	// 中间件
	const keyRequestId = "requestId"
	r.Use(func(c *gin.Context) {
		// 开始时间
		s := time.Now()
		c.Next()

		// path,response code,log latency
		logger.Info("incoming request",
			// 记录访问路径
			zap.String("path", c.Request.URL.Path),
			// 记录状态
			zap.Int("status", c.Writer.Status()),
			// 开始时间减去结束时间的结果
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		// 让每一个request都带上一个ID
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		// 获取设置好的requestid
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
