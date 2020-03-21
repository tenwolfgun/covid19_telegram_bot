package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// CurrentCase **
type CurrentCase struct {
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
	DailySummary    string `json:"dailySummary"`
	DailyTimeSeries struct {
		Pattern string `json:"pattern"`
		Example string `json:"example"`
	} `json:"dailyTimeSeries"`
	Image         string `json:"image"`
	Source        string `json:"source"`
	Countries     string `json:"countries"`
	CountryDetail struct {
		Pattern string `json:"pattern"`
		Example string `json:"example"`
	} `json:"countryDetail"`
	LastUpdate time.Time `json:"lastUpdate"`
}

// GetCurrentCovid19 **
func (c *CurrentCase) GetCurrentCovid19() (*CurrentCase, error) {
	resp, err := http.Get("https://covid19.mathdro.id/api/")

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
