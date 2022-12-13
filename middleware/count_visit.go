package middleware

import (
	"fmt"
	"strings"

	"github.com/beslow/goblog/db"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func CountVisit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipAddr := c.ClientIP()
		if !strings.HasPrefix(c.Request.RequestURI, "/public") && ipAddr != "" {
			func() {
				key := fmt.Sprintf("site_visit#ip#%s", ipAddr)

				conn := db.RedisPool.Get()
				defer conn.Close()

				reply, err := conn.Do("get", key)
				if err != nil {
					return
				}

				if reply == nil {
					conn.Do("setex", key, 24*3600, 1)

					keySiteVisitCount := "site_visit_count"
					currentVisitCount, err := conn.Do("get", keySiteVisitCount)
					if err != nil {
						return
					}

					var val int64
					if currentVisitCount == nil {
						val = 1
					} else {
						val, err = redis.Int64(currentVisitCount, err)
						if err != nil {
							return
						}

						val += 1
					}

					conn.Do("set", keySiteVisitCount, val)
				}
			}()
		}

		c.Next()
	}
}
