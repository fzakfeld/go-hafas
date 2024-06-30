package main

import (
	"encoding/json"
	"os"

	"github.com/fzakfeld/go-hafas/hafas"
)

func main() {
	client := hafas.NewHafasClient(&hafas.Config{
		Url:           os.Getenv("HAFAS_URL"),
		Salt:          os.Getenv("HAFAS_SALT"),
		Aid:           os.Getenv("HAFAS_AUTH_AID"),
		Base64urlMode: true,
	})

	// when := time.Now()
	// duration := 20
	// stationId := "ODAwMDA2OA"

	// departures, _ := client.GetDepartures(when, duration, stationId)

	// foo, _ := json.Marshal(&departures)
	// println(string(foo))

	journey, _ := client.GetJourney("MnwjVk4jMSNTVCMxNzE5NDI4MzQxI1BJIzAjWkkjOTcwMDEzI1RBIzEjREEjMzAwNjI0IzFTIzExNjE0OSMxVCMyMzMzI0xTIzEyNDUzNyNMVCMxMDAwMyNQVSM4MCNSVCMxI0NBI1NUUiNaRSM5I1pCI1NUUiAgICA5I1BDIzgjRlIjMTE2MTQ5I0ZUIzIzMzMjVE8jMTI0NTM3I1RUIzEwMDAzIw")

	foo, _ := json.Marshal(&journey)
	println(string(foo))
}
