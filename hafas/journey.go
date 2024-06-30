package hafas

import (
	"encoding/json"
	"time"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type Journey struct {
	ID      string  `json:"id"`
	Product Product `json:"product"`
	Stops   []Stop  `json:"stops"`
}

type Stop struct {
	Station            Station   `json:"station"`
	DepartureScheduled time.Time `json:"departure_scheduled"`
	DepartureReal      time.Time `json:"departure_real"`
	ArrivalScheduled   time.Time `json:"arrival_scheduled"`
	ArrivalReal        time.Time `json:"arrival_real"`
}

func (c *HafasClient) GetJourney(journeyId string) (Journey, error) {
	journey := Journey{ID: journeyId}

	if c.base64urlMode {
		journeyId, _ = c.decodeBase64URL(journeyId)
	}

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

	journey.Product = Product{
		Name:   product.Name,
		NameS:  product.NameS,
		Number: product.Number,
		Class:  product.Cls,
	}

	for _, stop := range result.Journey.StopL {
		location := result.Common.LocL[stop.LocX]

		departureScheduled := c.parseTime(stop.DTimeS, result.Journey.TrainStartDate, time.Time{})
		departureReal := c.parseTime(stop.DTimeR, result.Journey.TrainStartDate, time.Time{})
		arrivalScheduled := c.parseTime(stop.ATimeS, result.Journey.TrainStartDate, time.Time{})
		arrivalReal := c.parseTime(stop.ATimeR, result.Journey.TrainStartDate, time.Time{})

		stationId := location.ExtId
		if c.base64urlMode {
			stationId = c.encodeBase64URL(stationId)
		}

		journey.Stops = append(journey.Stops, Stop{
			Station: Station{
				ID:        stationId,
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
