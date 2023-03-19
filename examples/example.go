package main

import (
	"encoding/json"
	"os"

	"github.com/fzakfeld/go-hafas/hafas"
)

func main() {
	client := hafas.NewHafasClient(&hafas.Config{
		Url:  os.Getenv("HAFAS_URL"),
		Salt: os.Getenv("HAFAS_SALT"),
		Aid:  os.Getenv("HAFAS_AUTH_AID"),
	})

	// when := time.Now()
	// duration := 20
	// stationId := "8000068"

	// departures, _ := client.GetDepartures(when, duration, stationId)

	// foo, _ := json.Marshal(&departures)
	// println(string(foo))

	journey, _ := client.GetJourney("1|1348228|0|80|19032023")

	foo, _ := json.Marshal(&journey)
	println(string(foo))
}
