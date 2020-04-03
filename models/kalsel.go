package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// KalSel **
type KalSel struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	CallCenter           string    `json:"call_center"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	CovPositiveCount     int       `json:"cov_positive_count"`
	CovNegativeCount     int       `json:"cov_negative_count"`
	CovRecoveredCount    int       `json:"cov_recovered_count"`
	CovDiedCount         int       `json:"cov_died_count"`
	CovOdpCount          int       `json:"cov_odp_count"`
	CovPdpCount          int       `json:"cov_pdp_count"`
	Slug                 string    `json:"slug"`
	DeletedAt            time.Time `json:"deleted_at"`
	CovOdpProcessedCount int       `json:"cov_odp_processed_count"`
	CovPdpProcessedCount int       `json:"cov_pdp_processed_count"`
	Hotline              string    `json:"hotline"`
	Code                 string    `json:"code"`
}

// GetKalSelCovid19 **
func (*KalSel) GetKalSelCovid19() ([]KalSel, error) {
	resp, err := http.Get("http://corona.kalselprov.go.id/cov_map")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	kalsel := []KalSel{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &kalsel)

	if err != nil {
		return nil, err
	}

	return kalsel, nil
}
