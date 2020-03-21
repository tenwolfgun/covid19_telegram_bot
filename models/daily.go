package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Daily struct {
	ReportDate       int64       `json:"reportDate"`
	MainlandChina    int         `json:"mainlandChina"`
	OtherLocations   int         `json:"otherLocations"`
	TotalConfirmed   int         `json:"totalConfirmed"`
	TotalRecovered   interface{} `json:"totalRecovered"`
	ReportDateString string      `json:"reportDateString"`
	DeltaConfirmed   int         `json:"deltaConfirmed"`
	DeltaRecovered   interface{} `json:"deltaRecovered"`
	Objectid         int         `json:"objectid"`
}

// GetDailyCovid19 **
func (*Daily) GetDailyCovid19() ([]Daily, error) {
	resp, err := http.Get("https://covid19.mathdro.id/api/daily")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	daily := []Daily{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &daily)

	if err != nil {
		return nil, err
	}

	return daily, nil
}
