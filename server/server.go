package server

import (
	"context"
	"net/http"

	"github.com/alex-guoba/gin-clean-template/internal/middleware/ratelimit"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/alex-guoba/gin-clean-template/server/api"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	gswag "github.com/swaggo/gin-swagger"  // gin-swagger middleware

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//	@title			Gin-Clean-Template
//	@version		1.0
//	@description	Clean Architecture template for Golang Gin services
//	@contact.name	gelco
//	@contact.url	https://github.com/alex-guoba/gin-clean-template
//	@license.name	MIT License
//	@license.url	https://github.com/alex-guoba/gin-clean-template/blob/main/LICENSE
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http https

type Server struct {
	Router *gin.Engine
	Svr    *http.Server
	Config *setting.Configuration
	DB     *gorm.DB
}

func NewServer(cfg *setting.Configuration, db *gorm.DB) *Server {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	srv := &http.Server{
		Addr:           ":" + cfg.Server.HTTPPort,
		Handler:        r,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s := &Server{
		Config: cfg,
		Svr:    srv,
		Router: r,
		DB:     db,
	}

	// rate limit middleware
	if cfg.Ratelimit.Enable {
		limiter := ratelimit.New(cfg.Ratelimit.ConfigFile,
			cfg.Ratelimit.CPULoadThresh, cfg.Ratelimit.CPULoadStrategy)
		if limiter == nil {
			log.Error("init rate limit middleware failed, ignored")
		} else {
			r.Use(limiter)
		}
	}
	api.SetRouters(r, cfg, db)

	// swagger middleware
	r.GET("/swagger/*any", gswag.WrapHandler(swaggerfiles.Handler))

	return s
}

func (s *Server) Start() error {
	// Timeout: https://adam-p.ca/blog/2022/01/golang-http-server-timeouts/
	go func() {
		log.Info("Starting HTTP Server at :", s.Config.Server.HTTPPort)
		if err := s.Svr.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("HTTP server expcetpion. ", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	// TODO: add code
	return s.Svr.Shutdown(ctx)
}
