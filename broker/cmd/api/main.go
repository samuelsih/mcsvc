package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Starting broker on %s", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}