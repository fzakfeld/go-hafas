package hafas

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type hafasClient struct {
	url       string
	salt      string
	aid       string
	language  string
	userAgent string
}

type Config struct {
	Url  string
	Salt string
	Aid  string
}

func NewHafasClient(config *Config) *hafasClient {
	return &hafasClient{
		url:       config.Url,
		salt:      config.Salt,
		aid:       config.Aid,
		language:  "en",
		userAgent: "go-hafas",
	}
}

func (c *hafasClient) makeRequest(method string, request interface{}) (hresponse.HafasResult, error) {
	payload := hrequests.RequestBase{
		Lang: c.language,
		SvcReqL: [1]hrequests.Request{
			{
				Meth: method,
				Req:  request,
				Cfg: hrequests.Config{
					RtMode: "REALTIME",
				},
			},
		},
		Ext: "DB.R21.12.a",
		Ver: "1.34",
		Client: hrequests.Client{
			Type: "AND",
			ID:   "DB",
			V:    21120000,
			Name: "DB Navigator",
		},
		Auth: hrequests.Auth{
			Type: "AID",
			AID:  c.aid,
		},
	}

	payloadJson, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	url := c.url + "?checksum=" + c.createChecksum(payloadJson)

	resp, err := http.Post(url, "application/json", bytes.NewReader(payloadJson))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response hresponse.HafasResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	if response.Err != "OK" {
		panic(err)
	}

	if len(response.SvcResL) != 1 {
		panic(err)
	}

	result := response.SvcResL[0]

	return result, nil
}
