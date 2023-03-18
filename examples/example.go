package main

import (
	"encoding/json"

	"github.com/fzakfeld/go-hafas/hafas"
)

func main() {
	client := hafas.NewHafasClient(&hafas.Config{
		Url:  "hafas-endpoint/mgate.exe",
		Salt: "hafas-salt",
		Aid:  "hafas-auth-aid",
	})

	departures, _ := client.GetDepartures()

	foo, _ := json.Marshal(&departures)
	println(string(foo))
}
