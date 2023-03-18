package hafas

import (
	"encoding/json"
	"fzakfeld/go-hafas/hafas/hrequests"
	"fzakfeld/go-hafas/hafas/hresponse"
	"time"
)

type Departure struct {
	JourneyId          string
	Product            Product
	Direction          string
	Station            Station
	DepartureScheduled time.Time
	DepartureReal      time.Time
}

type Product struct {
	Name   string
	Number string
}

type Station struct {
	ID        string
	Name      string
	Latitude  float32
	Longitude float32
	Floor     int
}

func (c *hafasClient) GetDepartures() ([]Departure, error) {
	departures := []Departure{}

	request := hrequests.StationBoardRequest{
		Type: "DEP",
		Date: "20230317",
		Time: "170638",
		StbLoc: hrequests.Location{
			Type:   "S",
			LineId: "A=1@L=694884@",
		},
		JnyFltrL: []hrequests.JourneyFilter{
			{
				Type:  "PROD",
				Mode:  "INC",
				Value: "1023",
			},
		},
		Dur: 20,
	}

	data, err := c.makeRequest("StationBoard", request)
	var result hresponse.StationBoardResult

	if err != nil {
		return departures, err
	}

	err = json.Unmarshal(data.Res, &result)

	if err != nil {
		return departures, err
	}

	for _, x := range result.JnyL {
		product := result.Common.ProdL[x.ProdX]
		station := result.Common.LocL[x.StbStop.LocX]

		departureScheduled, err := c.parseTime(x.StbStop.DTimeS, x.TrainStartDate, time.Time{})
		departureReal := departureScheduled

		if err != nil {
			return departures, err
		}

		if x.StbStop.DTimeR != "" {
			departureReal, err = c.parseTime(x.StbStop.DTimeS, x.TrainStartDate, departureScheduled)
		} // @todo: what if real is < scheduled?

		if err != nil {
			return departures, err
		}

		departures = append(departures, Departure{
			JourneyId: x.Jid,
			Product: Product{
				Name:   product.Name,
				Number: product.Number,
			},
			Direction: x.DirTxt,
			Station: Station{
				ID:        station.ExtId,
				Name:      station.Name,
				Latitude:  float32(station.Crd.X) / 1000000,
				Longitude: float32(station.Crd.Y) / 1000000,
				Floor:     station.Crd.Floor,
			},
			DepartureScheduled: departureScheduled,
			DepartureReal:      departureReal, // @todo
		})
	}

	return departures, nil
}
