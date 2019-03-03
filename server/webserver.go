package server

import (
	"log"
	"net/http"
)

//StartWebServer - starts server
func StartWebServer(port string) error {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
		return err
	}
	return nil
}
