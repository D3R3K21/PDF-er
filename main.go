package main

import (
	"integrate-pdf-service/server"
	"log"
)

func main() {

	err := server.StartWebServer("80")
	if err != nil {
		log.Fatalf("Cannot Start Server: %s|n", err)
	}
}
