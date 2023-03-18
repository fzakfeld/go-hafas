package hrequests

type RequestBase struct {
	Lang    string     `json:"lang"`
	SvcReqL [1]Request `json:"svcReqL"`
	Ext     string     `json:"ext"`
	Ver     string     `json:"ver"`
	Client  Client     `json:"client"`
	Auth    Auth       `json:"auth"`
}

type Request struct {
	Meth string      `json:"meth"`
	Req  interface{} `json:"req"`
	Cfg  Config      `json:"cfg"`
}

type Config struct {
	RtMode string `json:"rtMode"`
}

type Client struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	V    int    `json:"v"`
	Name string `json:"name"`
}

type Auth struct {
	Type string `json:"type"`
	AID  string `json:"aid"`
}

type Location struct {
	Type   string `json:"type"`
	LineId string `json:"lid"`
}

type JourneyFilter struct {
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	Value string `json:"value"`
}
