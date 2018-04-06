package main

import (
	"log"
	"net/http"
	route "github.com/rapidclock/align-bot-stats-cache/routehandlers"
	"github.com/rapidclock/align-bot-stats-cache/cache"
)

func init() {
	cache.InitializeWithStats()
}

func main() {
	log.Fatal(http.ListenAndServe(":15000", route.NewRedisAppHandler()))
}
