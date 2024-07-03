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

	// departures, _ := client.GetArrivals(when, duration, stationId)

	// foo, _ := json.Marshal(&departures)
	// println(string(foo))

	journey, _ := client.GetJourney("MnwjVk4jMSNTVCMxNzE5ODYyOTkyI1BJIzAjWkkjMzA4NzMyI1RBIzAjREEjMzA3MjQjMVMjODAwMjU1MyMxVCMyMTM0I0xTIzgwMDAyNjEjTFQjMTA1NDAjUFUjODAjUlQjMSNDQSNJQ0UjWkUjMTUwNyNaQiNJQ0UgMTUwNyNQQyMwI0ZSIzgwMDI1NTMjRlQjMjEzNCNUTyM4MDAwMjYxI1RUIzEwNTQwIw")

	foo, _ := json.Marshal(&journey)
	println(string(foo))
}
