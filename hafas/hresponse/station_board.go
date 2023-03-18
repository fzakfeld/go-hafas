package hresponse

type StationBoardResult struct {
	Type   string    `json:"type"`
	Common Common    `json:"common"`
	JnyL   []Journey `json:"jnyL"`
}
