package main

import (
	"log"

	"github.com/alex-guoba/gin-clean-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
