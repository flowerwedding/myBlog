/**
 * @Title  limiter
 * @description  接口限流控制
 * @Author  沈来
 * @Update  2020/8/7 20:15
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"myBlog/pkg/app"
	"myBlog/pkg/errcode"
	"myBlog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
