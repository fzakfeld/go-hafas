package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/fzakfeld/go-hafas/hafas"
)

func main() {
	client := hafas.NewHafasClient(&hafas.Config{
		Url:  os.Getenv("HAFAS_URL"),
		Salt: os.Getenv("HAFAS_SALT"),
		Aid:  os.Getenv("HAFAS_AUTH_AID"),
	})

	when := time.Now()
	duration := 20
	stationId := "8000068"

	departures, _ := client.GetDepartures(when, duration, stationId)

	foo, _ := json.Marshal(&departures)
	println(string(foo))
}
