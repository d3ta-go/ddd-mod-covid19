package repository

import (
	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
	"github.com/d3ta-go/system/system/identity"
)

// ICurrentRepo represent CurrentRepo Interface
type ICurrentRepo interface {
	DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest, i identity.Identity) (*domSchema.DisplayCurrentDataByCountryResponse, error)
}
