/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alex-guoba/gin-clean-template/global"
	"github.com/alex-guoba/gin-clean-template/internal/dao/seed"

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
	Run: func(cmd *cobra.Command, args []string) {
		if err := setupDBEngine(); err != nil {
			return
		}

		if err := seed.Seed(global.DBEngine, seedCount); err != nil {
			log.Error("seed db error. ", err)
			return
		}

		log.Info(seedCount, " articles tag seed success.")
	},
}

func init() {
	seedCmd.Flags().IntVar(&seedCount, "count", 1, "seed record count.")

	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
