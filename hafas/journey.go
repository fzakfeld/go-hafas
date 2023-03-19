package hafas

import (
	"encoding/json"
	"time"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type Journey struct {
	JourneyId string
	Product   Product
	Stops     []Stop
}

type Stop struct {
	Station            Station
	DepartureScheduled time.Time
	DepartureReal      time.Time
	ArrivalScheduled   time.Time
	ArrivalReal        time.Time
}

func (c *hafasClient) GetJourney(journeyId string) (Journey, error) {
	journey := Journey{}

	request := hrequests.JourneyDetailsRequest{
		Jid:         journeyId,
		GetPolyline: false, // this is currently not supported.
	}
	data, err := c.makeRequest("JourneyDetails", request)

	if err != nil {
		return journey, err
	}

	var result hresponse.JourneyDetailsResult

	err = json.Unmarshal(data.Res, &result)

	if err != nil {
		return journey, err
	}

	product := result.Common.ProdL[result.Journey.ProdX]

	journey = Journey{
		JourneyId: result.Journey.Jid,
		Product: Product{
			Name:   product.Name,
			NameS:  product.NameS,
			Number: product.Number,
			Class:  product.Cls,
		},
	}

	for _, x := range result.Journey.StopL {
		location := result.Common.LocL[x.LocX]

		departureScheduled := c.parseTime(x.DTimeS, result.Journey.TrainStartDate, time.Time{})
		departureReal := c.parseTime(x.DTimeR, result.Journey.TrainStartDate, time.Time{})
		arrivalScheduled := c.parseTime(x.ATimeS, result.Journey.TrainStartDate, time.Time{})
		arrivalReal := c.parseTime(x.ATimeR, result.Journey.TrainStartDate, time.Time{})

		journey.Stops = append(journey.Stops, Stop{
			Station: Station{
				ID:        location.ExtId,
				Name:      location.Name,
				Latitude:  float32(location.Crd.Y) / 1000000,
				Longitude: float32(location.Crd.X) / 1000000,
				Floor:     location.Crd.Floor,
			},
			DepartureScheduled: departureScheduled,
			DepartureReal:      departureReal,
			ArrivalScheduled:   arrivalScheduled,
			ArrivalReal:        arrivalReal,
		})
	}

	return journey, nil
}
