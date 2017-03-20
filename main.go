package main

import (
	"log"

	"github.com/amitm1/basic-go-service/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
