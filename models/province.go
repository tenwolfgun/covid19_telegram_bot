package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Province struct {
	Data []struct {
		Fid       int    `json:"fid"`
		KodeProvi int    `json:"kodeProvi"`
		Provinsi  string `json:"provinsi"`
		KasusPosi int    `json:"kasusPosi"`
		KasusSemb int    `json:"kasusSemb"`
		KasusMeni int    `json:"kasusMeni"`
	} `json:"data"`
}

func (p *Province) GetCovid19ByProvince() (*Province, error) {
	resp, err := http.Get("https://indonesia-covid-19.mathdro.id/api/provinsi")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
