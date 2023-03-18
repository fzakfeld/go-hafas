package hafas

import (
	"encoding/json"
	"fzakfeld/go-hafas/hafas/hrequests"
	"fzakfeld/go-hafas/hafas/hresponse"
)

type Departure struct {
	JourneyId string
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
		departures = append(departures, Departure{
			JourneyId: x.Jid,
		})
	}

	return departures, nil
}
