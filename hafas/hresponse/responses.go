package hresponse

import "encoding/json"

type HafasResponse struct {
	Err     string        `json:"err"`
	SvcResL []HafasResult `json:"svcResL"`
}

type HafasResult struct {
	Meth string          `json:"meth"`
	Res  json.RawMessage `json:"res"`
}

type Common struct {
	LocL  []Location `json:"locL"`
	ProdL []Product  `json:"prodL"`
	OpL   []Operator `json:"opL"`
}

type Operator struct {
	Name string `json:"name"`
}

type Location struct {
	Lid   string `json:"lid"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	ExtId string `json:"extId"`
	Crd   struct {
		X     int `json:"x"`
		Y     int `json:"y"`
		Floor int `json:"floor"`
	} `json:"crd"`
}

type Product struct {
	Name   string `json:"name"`
	NameS  string `json:"nameS"`
	Number string `json:"number"`
	Cls    int    `json:"cls"`
	OprX   int    `json:"oprX"`
}

type Journey struct {
	Jid            string   `json:"jid"`
	Date           string   `json:"date"`
	ProdX          int      `json:"prodX"`
	DirTxt         string   `json:"dirTxt"`
	StbStop        Stop     `json:"stbStop"`
	StopL          []Stop   `json:"stopL"`
	Pos            Position `json:"pos"`
	TrainStartDate string   `json:"trainStartDate"`
}

type Stop struct {
	LocX      int    `json:"locX"`
	Idx       int    `json:"idx"`
	DProdX    int    `json:"dProdX"`
	DTimeS    string `json:"dTimeS"`
	DTimeR    string `json:"dTimeR"`
	DProgType string `json:"dProgType"`
	DTZOffset int    `json:"dTZOffset"`
}

type Position struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Floor int `json:"floor"`
}
