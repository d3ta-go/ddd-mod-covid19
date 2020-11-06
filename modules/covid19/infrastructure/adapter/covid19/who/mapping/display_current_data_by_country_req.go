package mapping

import (
	con19type "github.com/d3ta-go/connector-covid19/connector/covid19/who/types"
	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
)

// MapDisplayCurrentDataByCountryReq mapping DisplayCurrentDataByCountryReq
func MapDisplayCurrentDataByCountryReq(req *domSchema.DisplayCurrentDataByCountryRequest) (*con19type.GetCountryRequest, error) {

	reqCon := new(con19type.GetCountryRequest)
	reqCon.CountryCode = req.CountryCode

	return reqCon, nil
}
