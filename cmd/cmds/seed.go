/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmds

import (
	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/dao/seed"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	seedCount int
)

// seedCmd represents the seed command.
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed blog service database",
	Run: func(_ *cobra.Command, _ []string) {
		if err := svrInit(); err != nil {
			log.Error("init server failed.", err)
			return
		}

		engine, err := dao.NewDBEngine(&setting.Conf.Database)
		if err != nil {
			log.Error("init db error. ", err)
			return
		}

		if err = seed.Seed(engine, seedCount); err != nil {
			log.Error("seed db error. ", err)
			return
		}

		log.Info(seedCount, " articles tag seed success.")
	},
}

func init() {
	seedCmd.Flags().IntVar(&seedCount, "count", 1, "seed record count.")

	rootCmd.AddCommand(seedCmd)
}
