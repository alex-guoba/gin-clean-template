/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/alex-guoba/gin-clean-template/global"
	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/middleware/ratelimit"
	"github.com/alex-guoba/gin-clean-template/internal/routers"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gin-clean-template",
		Short: "A clean architecture template for Golang Gin services",

		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(global.ServerSetting.RunMode)

			r := gin.New()

			r.Use(gin.Logger())
			r.Use(gin.Recovery())

			// global rate limit middleware
			if global.RatelimitSetting.Enable {
				limiter := ratelimit.New(global.RatelimitSetting.ConfigFile,
					global.RatelimitSetting.CPULoadThresh, global.RatelimitSetting.CPULoadStrategy)
				if limiter == nil {
					log.Error("init rate limit middleware failed")
					return
				}
				r.Use(limiter)
			}

			routers.SetRouters(r)

			// use http server
			s := &http.Server{
				Addr:           ":" + global.ServerSetting.HTTPPort,
				Handler:        r,
				ReadTimeout:    global.ServerSetting.ReadTimeout,
				WriteTimeout:   global.ServerSetting.WriteTimeout,
				MaxHeaderBytes: 1 << 20,
			}

			log.Info("server started at " + s.Addr)

			_ = s.ListenAndServe()
		},
	}
)

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// setup log file with configuration.
	logger.SetupLogger(
		filepath.Join(global.LogSetting.LogSavePath, global.LogSetting.LogFileName),
		global.LogSetting.MaxSize, global.LogSetting.MaxBackups, global.LogSetting.Compress,
		global.LogSetting.Level)

	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gin-clean-template.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// parsed by section
	if err := setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("Log", &global.LogSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("Ratelimit", &global.RatelimitSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = dao.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
