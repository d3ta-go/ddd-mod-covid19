package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/d3ta-go/ddd-mod-covid19/modules/covid19/application/dto"
	"github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	return c, v, nil
}

func newCurrentSvc(t *testing.T) (*CurrentSvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.covid19.covid19.app-layer.service.display-current-data-by-country")

	req := dto.DisplayCurrentDataByCountryReqDTO{}
	req.CountryCode = testData["country-code"]
	req.Providers = append(req.Providers, &schema.Provider{Code: testData["provider-code"]})

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
