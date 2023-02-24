package main

import "fzakfeld/hafas/hafas"

func main() {
	client := hafas.NewHafasClient(hafas.HafasClientOptions{
		Endpoint:      "https://reiseauskunft.bahn.de/bin/mgate.exe",
		Salt:          "6264493855566A34304B356676787766",
		Aid:           "-",
		ClientType:    "AND",
		ClientId:      "DB",
		ClientVersion: 21120000,
		ClientName:    "DB Navigator",
	})

	client.GetDepartures("693125")
}
