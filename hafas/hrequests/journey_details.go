package hrequests

type JourneyDetailsRequest struct {
	Jid         string `json:"jid"`
	GetPolyline bool   `json:"getPolyline"`
}
