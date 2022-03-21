package main

import (
	"log"

	"github.com/bariis/distributed-log-service/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
