package main

import (
	"log"
	"net/http"

	"goWithTestsAcceptance/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
