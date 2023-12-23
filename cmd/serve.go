/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/alex-guoba/gin-clean-template/global"
	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/middleware/ratelimit"
	"github.com/alex-guoba/gin-clean-template/internal/routers"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/alex-guoba/gin-clean-template/pkg/signals"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"

	// for file migration source directory.
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-clean-template",
	Short: "A clean architecture template for Golang Gin services",

	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(global.Config.Server.RunMode)

		// init logger
		logger.SetupLogger(
			filepath.Join(global.Config.Log.LogSavePath, global.Config.Log.LogFileName),
			global.Config.Log.MaxSize, global.Config.Log.MaxBackups, global.Config.Log.Compress,
			global.Config.Log.Level)

		// init db
		if err := dbInit(&global.Config.Database); err != nil {
			log.Error("init db failed.", err)
			return
		}

		// start http server
		srv, err := startHttpServer()
		if err != nil {
			log.Error("start http server failed.", err)
			return
		}

		// graceful shutdown
		stopCh := signals.SetupSignalHandler()
		sd, _ := signals.NewShutdown(global.Config.App.ServerShutdownTimeout)
		sd.Graceful(stopCh, srv, global.DBEngine)
	},
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}

func startHttpServer() (*http.Server, error) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// global rate limit middleware
	if global.Config.Ratelimit.Enable {
		limiter := ratelimit.New(global.Config.Ratelimit.ConfigFile,
			global.Config.Ratelimit.CPULoadThresh, global.Config.Ratelimit.CPULoadStrategy)
		if limiter == nil {
			log.Error("init rate limit middleware failed, ignored")
		} else {
			r.Use(limiter)
		}
	}
	routers.SetRouters(r)

	// Timeout: https://adam-p.ca/blog/2022/01/golang-http-server-timeouts/
	srv := &http.Server{
		Addr:           ":" + global.Config.Server.HTTPPort,
		Handler:        r,
		ReadTimeout:    global.Config.Server.ReadTimeout,
		WriteTimeout:   global.Config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Info("Starting HTTP Server at :", global.Config.Server.HTTPPort)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("HTTP server expcetpion. ", err)
		}
	}()

	return srv, nil
}

func dbInit(dbconfig *setting.DatabaseSettingS) error {
	var err error
	global.DBEngine, err = dao.NewDBEngine(&global.Config.Database)
	if err != nil {
		return err
	}

	// Run db migration
	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbconfig.DBType,
		dbconfig.UserName,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.DBName,
		dbconfig.Charset,
		dbconfig.ParseTime,
	)
	migration, err := migrate.New(dbconfig.MigrationURL, dsn)
	if err != nil {
		log.Error("migration init error.", err)
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to run migrate up.", err)
		return err
	}

	log.Info("db migrated successfully.")
	return nil
}

func init() {
	// cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gin-clean-template.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	if err := setting.LoadConfig(&global.Config); err != nil {
		log.Fatal("loading config file failed.", err)
	}
}
