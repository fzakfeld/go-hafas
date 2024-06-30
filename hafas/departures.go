package hafas

import (
	"encoding/json"
	"time"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type Departure struct {
	ID                 string    `json:"id"`
	Direction          string    `json:"direction"`
	Product            Product   `json:"product"`
	Station            Station   `json:"station"`
	DepartureScheduled time.Time `json:"departure_scheduled"`
	DepartureReal      time.Time `json:"departure_real"`
}

type Product struct {
	Name     string   `json:"name"`
	NameS    string   `json:"-"`
	Number   string   `json:"number"`
	Operator Operator `json:"operator"`
	Class    int      `json:"-"`
}

type Operator struct {
	Name string `json:"name"`
}

type Station struct { // @todo rename to location?
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Floor     int     `json:"floor"`
}

func (c *HafasClient) GetDepartures(when time.Time, duration int, stationId string) ([]Departure, error) {
	departures := []Departure{}

	if c.base64urlMode {
		stationId, _ = c.decodeBase64URL(stationId)
	}

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

	for _, journey := range result.JnyL {
		product := result.Common.ProdL[journey.ProdX]
		station := result.Common.LocL[journey.StbStop.LocX]
		operator := hresponse.Operator{}
		// if true { // @todo default value for OpX that is not 0
		// 	operator = result.Common.OpL[product.OprX]
		// }

		departureScheduled := c.parseTime(journey.StbStop.DTimeS, journey.TrainStartDate, time.Time{})
		departureReal := c.parseTime(journey.StbStop.DTimeS, journey.TrainStartDate, time.Time{})
		latitude := float32(station.Crd.X) / 1000000
		longitude := float32(station.Crd.Y) / 1000000

		journeyId := journey.Jid
		if c.base64urlMode {
			journeyId = c.encodeBase64URL(journeyId)
			station.ExtId = c.encodeBase64URL(station.ExtId)
		}

		departures = append(departures, Departure{
			ID: journeyId,
			Product: Product{
				Name:   product.Name,
				NameS:  product.NameS,
				Number: product.Number,
				Operator: Operator{
					Name: operator.Name,
				},
				Class: 0,
			},
			Direction: journey.DirTxt,
			Station: Station{
				ID:        station.ExtId, // stationId
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
