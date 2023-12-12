package ratelimit

import (
	"log"
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/system"

	"github.com/gin-gonic/gin"
)

func initConfig(cfg string) {
	// config file format: https://github.com/alibaba/sentinel-golang/wiki/%E5%90%AF%E5%8A%A8%E9%85%8D%E7%BD%AE
	if len(cfg) == 0 {
		if err := sentinel.InitDefault(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := sentinel.InitWithConfigFile(cfg); err != nil {
			log.Fatal(err)
		}
	}

	// init rules. see https://github.com/alibaba/sentinel-golang/wiki/%E6%B5%81%E9%87%8F%E6%8E%A7%E5%88%B6
	// we just used system rules here by default.
	// how system adaptive rules works: https://sentinelguard.io/zh-cn/docs/system-adaptive-protection.html
	_, err := system.LoadRules([]*system.Rule{
		{
			ID:           "DefaultSystemRule",
			MetricType:   system.CpuUsage,
			TriggerCount: 0.8, // 80%
			Strategy:     system.BBR,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

// See: https://pkg.go.dev/github.com/alibaba/sentinel-golang/pkg/adapters/gin
func New(cfg string) gin.HandlerFunc {
	// init default
	initConfig(cfg)

	return func(c *gin.Context) {
		// resourceName := c.Request.Method + ":" + c.FullPath()
		resourceName := c.GetHeader("X-Real-IP")

		// TODO: adjust default config
		entry, err := sentinel.Entry(
			resourceName,
			sentinel.WithResourceType(base.ResTypeWeb),
			sentinel.WithTrafficType(base.Inbound),
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, map[string]interface{}{
				"err":  "too many request; the quota used up",
				"code": 10222,
			})
			return
		}

		defer entry.Exit()
		c.Next()
	}

}
