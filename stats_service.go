package main

import (
	"log"
	"net/http"
	"./cache"
	route "./routehandlers"
)

func init() {
	cache.InitializeWithStats()
}

func main() {
	log.Fatal(http.ListenAndServe(":15000", route.NewRedisAppHandler()))
}
