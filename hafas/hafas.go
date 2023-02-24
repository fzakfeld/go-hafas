package hafas

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type HafasClient interface {
	GetDepartures(stationId string)
}

type hafasClient struct {
	Endpoint      string
	Salt          string
	Aid           string
	ClientType    string
	ClientId      string
	ClientVersion int
	ClientName    string
	httpClient    http.Client
}

type HafasClientOptions struct {
	Endpoint      string
	Salt          string
	Aid           string
	ClientType    string
	ClientId      string
	ClientVersion int
	ClientName    string
}

func NewHafasClient(options HafasClientOptions) HafasClient {
	return &hafasClient{
		Endpoint:      options.Endpoint,
		Salt:          options.Salt,
		Aid:           options.Aid,
		ClientType:    options.ClientType,
		ClientId:      options.ClientId,
		ClientVersion: options.ClientVersion,
		ClientName:    options.ClientName,
		httpClient:    http.Client{},
	}
}

func (c *hafasClient) GetDepartures(stationId string) {
	requestBody := RequestBodyStationBoard{
		Lang: "en",
		SvcReqL: [1]RequestBodyStationBoardSvcReqL{
			{
				Meth: "StationBoard",
				Req: RequestBodyStationBoardReq{
					Type: "DEP",
					Date: "20230223",
					Time: "221907",
					StbLoc: RequestBodyStationBoardLocation{
						Type: "S",
						Lid:  "A=1@L=" + stationId + "@",
					},
					JnyFltrL: []RequestBodyJnyFltrL{},
					Dur:      10,
				},
				Cfg: RequestBodyCfg{
					RtMode: "",
				},
			},
		},
		Ext:    "DB.R21.12.a",
		Ver:    "1.34",
		Client: c.getClient(),
		Auth:   c.getAuth(),
	}

	rb, err := json.Marshal(requestBody)
	check(err)

	bodyReader := bytes.NewReader(rb)

	requestUrl := c.Endpoint + "?checksum=" + c.generateChecksum(string(rb), c.Salt)
	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	check(err)

	resp, err := c.httpClient.Do(req)
	check(err)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(b))
}

func (c *hafasClient) generateChecksum(requestBody string, salt string) string {
	saltStr, err := hex.DecodeString(salt)
	if err != nil {
		panic(err)
	}
	cs := md5.Sum([]byte(requestBody + string(saltStr)))
	return hex.EncodeToString(cs[:])
}

func (c *hafasClient) getAuth() RequestBodyAuth {
	return RequestBodyAuth{
		Type: "AID",
		Aid:  c.Aid,
	}
}

func (c *hafasClient) getClient() RequestBodyClient {
	return RequestBodyClient{
		Type: c.ClientType,
		Id:   c.ClientId,
		V:    c.ClientVersion,
		Name: c.ClientName,
	}
}
