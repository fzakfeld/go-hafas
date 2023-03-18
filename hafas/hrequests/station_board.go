package hrequests

type StationBoardRequest struct {
	Type     string          `json:"type"`
	Date     string          `json:"date"`
	Time     string          `json:"time"`
	StbLoc   Location        `json:"stbLoc"`
	JnyFltrL []JourneyFilter `json:"jnyFltrL"`
	Dur      int             `json:"dur"`
}
