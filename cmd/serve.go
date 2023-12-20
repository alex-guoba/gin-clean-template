/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		gin.SetMode(global.ServerSetting.RunMode)

		if err := dbInit(); err != nil {
			log.Error("init db failed.", err)
			return
		}

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

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// setup log file with configuration.
	logger.SetupLogger(
		filepath.Join(global.LogSetting.LogSavePath, global.LogSetting.LogFileName),
		global.LogSetting.MaxSize, global.LogSetting.MaxBackups, global.LogSetting.Compress,
		global.LogSetting.Level)

	return rootCmd.Execute()
}

func dbInit() error {
	if err := setupDBEngine(); err != nil {
		return err
	}

	// Run db migration
	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		global.DatabaseSetting.DBType,
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.DBName,
		global.DatabaseSetting.Charset,
		global.DatabaseSetting.ParseTime,
	)
	migration, err := migrate.New(global.DatabaseSetting.MigrationURL, dsn)
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

	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
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
