package dto

import (
	"encoding/json"

	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
)

// DisplayCurrentDataByCountryReqDTO represent DisplayCurrentDataByCountryReq DTO
type DisplayCurrentDataByCountryReqDTO struct {
	CountryCode string                `json:"countryCode"`
	Providers   []*domSchema.Provider `json:"providers"`
}

// DisplayCurrentDataByCountryResDTO represent DisplayCurrentDataByCountryRes DTO
type DisplayCurrentDataByCountryResDTO struct {
	Query interface{}                             `json:"query"`
	Data  *domSchema.TotalCountryProviderDataList `json:"data"`
}

// ToJSON covert to JSON
func (r *DisplayCurrentDataByCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
