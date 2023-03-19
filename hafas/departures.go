package hafas

import (
	"encoding/json"
	"time"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type Departure struct {
	JourneyId          string
	Direction          string
	Product            Product
	Station            Station
	DepartureScheduled time.Time
	DepartureReal      time.Time
}

type Product struct {
	Name     string
	NameS    string
	Number   string
	Operator Operator
	Class    int
}

type Operator struct {
	Name string
}

type Station struct {
	ID        string
	Name      string
	Latitude  float32
	Longitude float32
	Floor     int
}

func (c *hafasClient) GetDepartures(when time.Time, duration int, stationId string) ([]Departure, error) {
	departures := []Departure{}

	departuresDate := when.Format("20060102")
	departuresTime := when.Format("150405")

	request := hrequests.StationBoardRequest{
		Type: "DEP",
		Date: departuresDate,
		Time: departuresTime,
		StbLoc: hrequests.Location{
			Type:   "S",
			LineId: "A=1@L=" + stationId + "@",
		},
		JnyFltrL: []hrequests.JourneyFilter{ // @todo what does this do?
			{
				Type:  "PROD",
				Mode:  "INC",
				Value: "1023",
			},
		},
		Dur: duration,
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
		operator := hresponse.Operator{}
		// if true { // @todo default value for OpX that is not 0
		// 	operator = result.Common.OpL[product.OprX]
		// }

		departureScheduled := c.parseTime(x.StbStop.DTimeS, x.TrainStartDate, time.Time{})
		departureReal := c.parseTime(x.StbStop.DTimeS, x.TrainStartDate, time.Time{})
		latitude := float32(station.Crd.X) / 1000000
		longitude := float32(station.Crd.Y) / 1000000

		departures = append(departures, Departure{
			JourneyId: x.Jid,
			Product: Product{
				Name:   product.Name,
				NameS:  product.NameS,
				Number: product.Number,
				Operator: Operator{
					Name: operator.Name,
				},
				Class: 0,
			},
			Direction: x.DirTxt,
			Station: Station{
				ID:        station.ExtId,
				Name:      station.Name,
				Latitude:  latitude,
				Longitude: longitude,
				Floor:     station.Crd.Floor,
			},
			DepartureScheduled: departureScheduled,
			DepartureReal:      departureReal,
		})
	}

	return departures, nil
}
