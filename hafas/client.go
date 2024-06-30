package hafas

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/fzakfeld/go-hafas/hafas/hrequests"
	"github.com/fzakfeld/go-hafas/hafas/hresponse"
)

type HafasClient struct {
	url           string
	salt          string
	aid           string
	language      string
	userAgent     string
	base64urlMode bool
}

type Config struct {
	Url           string
	Salt          string
	Aid           string
	Base64urlMode bool
}

func NewHafasClient(config *Config) *HafasClient {
	return &HafasClient{
		url:           config.Url,
		salt:          config.Salt,
		aid:           config.Aid,
		language:      "en",
		userAgent:     "go-hafas",
		base64urlMode: config.Base64urlMode,
	}
}

func (c *HafasClient) encodeBase64URL(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}

func (c *HafasClient) decodeBase64URL(input string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(input)
	return string(data), err
}

func (c *HafasClient) makeRequest(method string, request interface{}) (hresponse.HafasResult, error) {
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response hresponse.HafasResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	if response.Err != "OK" {
		panic(response.Err)
	}

	if len(response.SvcResL) != 1 {
		panic(errors.New(""))
	}

	result := response.SvcResL[0]

	return result, nil
}
