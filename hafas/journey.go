package hafas

import (
	"encoding/json"
	"fzakfeld/go-hafas/hafas/hrequests"
	"fzakfeld/go-hafas/hafas/hresponse"
)

type Journey struct {
	JourneyId string
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

	journey.JourneyId = result.Journey.Jid

	return journey, nil
}
