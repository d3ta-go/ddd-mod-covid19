package repository

import (
	"encoding/json"
	"fmt"
	"testing"

	domRepo "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/repository"
	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
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

func newRepoIdent(t *testing.T) (domRepo.ICurrentRepo, identity.Identity, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, identity.Identity{}, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, identity.Identity{}, nil, err
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

	r, err := NewCurrentRepo(h)
	if err != nil {
		return nil, identity.Identity{}, nil, err
	}

	i := newIdentity(h, t)

	return r, i, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}

	return i
}

func TestRepoDisplayCurrentDataByCountry(t *testing.T) {
	repo, i, h, err := newRepoIdent(t)
	if err != nil {
		t.Errorf("REPO: [%#v]", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.covid19.covid19.infra-layer.repo.display-current-data-by-country")

	req := domSchema.DisplayCurrentDataByCountryRequest{}
	req.CountryCode = testData["country-code"]
	req.Providers = append(req.Providers, &domSchema.Provider{Code: testData["provider-code"]})

	res, err := repo.DisplayCurrentDataByCountry(&req, i)
	if err != nil {
		t.Errorf("Request: [%s]", err.Error())
	}

	if res == nil {
		t.Fail()
	}

	if res != nil {
		resJSON, err := json.Marshal(res)
		if err != nil {
			t.Errorf("json.Marshal: [%s]", err.Error())
		}
		t.Logf("Resp: %s", string(resJSON))
	}
}
