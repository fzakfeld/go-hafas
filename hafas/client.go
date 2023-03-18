package hafas

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fzakfeld/go-hafas/hafas/hrequests"
	"fzakfeld/go-hafas/hafas/hresponse"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type hafasClient struct {
	url      string
	salt     string
	aid      string
	language string
}

// type HafasConfig struct {
// 	url      string
// 	salt     string
// 	aid      string
// 	language string
// }

type Config struct {
	Url  string
	Salt string
	Aid  string
}

func NewHafasClient(config *Config) *hafasClient {
	return &hafasClient{
		url:      config.Url,
		salt:     config.Salt,
		aid:      config.Aid,
		language: "en",
	}
}

func (c *hafasClient) makeRequest(method string, payload interface{}) (hresponse.HafasResult, error) {
	request := hrequests.RequestBase{
		Lang: c.language,
		SvcReqL: [1]hrequests.Request{
			{
				Meth: method,
				Req:  payload,
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

	requestData, _ := json.Marshal(request)

	url := c.url + "?checksum=" + c.createChecksum(requestData)

	resp, _ := http.Post(url, "application/json", bytes.NewReader(requestData))
	body, _ := ioutil.ReadAll(resp.Body)

	var response hresponse.HafasResponse
	var result hresponse.HafasResult
	err := json.Unmarshal([]byte(body), &response)

	if err != nil {
		return result, err
	}

	if response.Err != "OK" {
		return result, errors.New(strings.ToLower(response.Err))
	}

	if len(response.SvcResL) < 1 {
		return result, errors.New("received no payload")
	}

	result = response.SvcResL[0]

	return result, nil
}

func (c *hafasClient) createChecksum(requestData []byte) string {
	salt, _ := hex.DecodeString(c.salt)
	hash := md5.Sum([]byte(string(requestData) + string(salt)))
	return hex.EncodeToString(hash[:])
}

func (c *hafasClient) parseTime(timestamp string, day string, startDate time.Time) (time.Time, error) {
	layout := "20060102150405"

	ts, err := time.Parse(layout, day+timestamp)

	if startDate.After(ts) {
		// this is neccessary for overnight journeys.
		ts = ts.AddDate(0, 0, 1)
	}

	return ts, err
}
