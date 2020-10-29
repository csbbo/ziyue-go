package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		costTime := time.Since(startTime)
		url := c.Request.URL.String()
		fmt.Printf("URL: %s cost %v\n", url, costTime)
	}
}
