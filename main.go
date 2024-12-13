package main

import (
	"log"

	"github.com/alex-guoba/gin-clean-template/cmd"
	_ "github.com/alex-guoba/gin-clean-template/server/docs"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
