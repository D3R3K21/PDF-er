package main

import (
	"log"
	"pdf-er/server"
)

func main() {

	err := server.StartWebServer("80")
	if err != nil {
		log.Fatalf("Cannot Start Server: %s|n", err)
	}
	// var name = "Derek Rose"
	// var company = "Integrate Inc"
	// var email = "drose@integrate.com"
	// var ip = "127.0.0.1"
	// var date = "March 1 2019"
	// // sig := newSignature(name, company, email, ip, date)
	// // err = savePDF(sig)
	// if err != nil {

	// }
}
