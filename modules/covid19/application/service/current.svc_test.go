package service

import (
	"encoding/json"
	"testing"

	"github.com/d3ta-go/ddd-mod-covid19/modules/covid19/application/dto"
	"github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newCurrentSvc(t *testing.T) (*CurrentSvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}
	if err := initialize.OpenAllCacheConnection(h); err != nil {
		return nil, nil, err
	}

	r, err := NewCurrentSvc(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	if err := i.SetCasbinEnforcer("../../../../conf/casbin/casbin_rbac_rest_model.conf"); err != nil {
		t.Errorf("SetCasbinEnforcer: %s", err.Error())
	}

	i.Claims.Username = "test.d3tago"
	i.Claims.AuthorityID = "group:default"

	i.RequestInfo.RequestObject = "/api/v1/covid19/current/by-country"
	i.RequestInfo.RequestAction = "POST"

	return i
}

func TestCurrentSvc_DisplayCurrentDataByCountry(t *testing.T) {
	svc, h, err := newCurrentSvc(t)
	if err != nil {
		t.Errorf("newCurrentSvc: %s", err.Error())
		return
	}

	req := dto.DisplayCurrentDataByCountryReqDTO{}
	req.CountryCode = "ID"
	req.Providers = append(req.Providers, &schema.Provider{Code: "_ALL_"})

	i := newIdentity(h, t)

	resp, err := svc.DisplayCurrentDataByCountry(&req, i)
	if err != nil {
		t.Errorf("DisplayCurrentDataByCountry: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}
