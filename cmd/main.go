package main

import (
	"log"
	"github.com/it59com/ws-go-proto/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
