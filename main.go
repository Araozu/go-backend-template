package main

import (
	"acide/src"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := src.NewServer()

	log.Print("starting HTTP server")
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Print("server closed")
	} else {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
