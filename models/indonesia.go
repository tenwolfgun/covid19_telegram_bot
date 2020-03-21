package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Indonesia struct {
	Meninggal   int `json:"meninggal"`
	Sembuh      int `json:"sembuh"`
	Perawatan   int `json:"perawatan"`
	JumlahKasus int `json:"jumlahKasus"`
	PerKasus    struct {
		JSON  string `json:"json"`
		Csv   string `json:"csv"`
		Links string `json:"links"`
		Old   string `json:"old"`
	} `json:"perKasus"`
	PerProvinsi struct {
		JSON string `json:"json"`
		Csv  string `json:"csv"`
	} `json:"perProvinsi"`
	PerHari struct {
		JSON string `json:"json"`
		Csv  string `json:"csv"`
	} `json:"perHari"`
}

func (i *Indonesia) GetIndonesiaCovid19() (*Indonesia, error) {
	resp, err := http.Get("https://indonesia-covid-19.mathdro.id/api/")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &i)

	if err != nil {
		return nil, err
	}

	return i, nil
}
