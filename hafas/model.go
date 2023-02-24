package hafas

type RequestBodyStationBoard struct {
	Lang    string                            `json:"lang"`
	SvcReqL [1]RequestBodyStationBoardSvcReqL `json:"svcReqL"`
	Client  RequestBodyClient                 `json:"client"`
	Ext     string                            `json:"ext"`
	Ver     string                            `json:"ver"`
	Auth    RequestBodyAuth                   `json:"auth"`
}

type RequestBodyStationBoardSvcReqL struct {
	Meth string                     `json:"meth"`
	Req  RequestBodyStationBoardReq `json:"req"`
	Cfg  RequestBodyCfg             `json:"cfg"`
}

//
type RequestBodyClient struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	V    int    `json:"v"`
	Name string `json:"name"`
}

type RequestBodyAuth struct {
	Type string `json:"type"`
	Aid  string `json:"aid"`
}

type RequestBodyCfg struct {
	RtMode string `json:"rtMode"`
}

type RequestBodyJnyFltrL struct {
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

// Specific to StationBoard
type RequestBodyStationBoardReq struct {
	Type     string                          `json:"type"` // DEP or ARR
	Date     string                          `json:"date"`
	Time     string                          `json:"time"`
	StbLoc   RequestBodyStationBoardLocation `json:"stbLoc"`
	JnyFltrL []RequestBodyJnyFltrL           `json:"jnyFltrL"`
	Dur      int                             `json:"dur"`
}

// Specific to StationBoard
type RequestBodyStationBoardLocation struct {
	Type string `json:"type"` // POI = P, Station = S, Address = A
	Lid  string `json:"lid"`  // A=1 = Station?, L = Station ID
}
