package hresponse

type JourneyDetailsResult struct {
	Common  Common  `json:"common"`
	Journey Journey `json:"journey"`
	FpB     string  `json:"fpB"`
	FpE     string  `json:"fpE"`
}
