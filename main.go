package main

import (
	"log"
	_ "github.com/alex-guoba/gin-clean-template/server/docs"
	"github.com/alex-guoba/gin-clean-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
