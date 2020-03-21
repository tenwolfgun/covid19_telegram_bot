package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Country struct {
	Confirmed struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"confirmed"`
	Recovered struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"recovered"`
	Deaths struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"deaths"`
	LastUpdate time.Time `json:"lastUpdate"`
}

func (c *Country) GetIndonesiaCovid19() (*Country, error) {
	resp, err := http.Get("https://covid19.mathdro.id/api/countries/ID")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
