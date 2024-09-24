package main

import (
	"acide/src"
	"embed"
	"fmt"
)

//go:embed public
var staticFiles embed.FS

func main() {
	server := src.NewServer(&staticFiles)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
