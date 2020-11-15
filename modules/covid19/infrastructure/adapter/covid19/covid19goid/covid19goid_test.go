package covid19goid

import (
	"encoding/json"
	"fmt"
	"testing"

	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"
	infC19Adp "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/adapter/covid19"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	return c, v, nil
}

func newCovid19goidAdapter(t *testing.T) (infC19Adp.ICovid19Adapter, *handler.Handler, error) {
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
	viperTest.AddConfigPath("../../../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

	adp, _, err := NewCovid19goidAdapter(h)
	if err != nil {
		return nil, nil, err
	}

	return adp, h, nil
}

func TestCovid19goidAdapter_DisplayCurrentDataByCountry(t *testing.T) {
	adp, h, err := newCovid19goidAdapter(t)
	if err != nil {
		t.Errorf("newCovid19goidAdapter: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.covid19.covid19.infra-layer.adapter.covid19goid")

	req := domSchema.DisplayCurrentDataByCountryRequest{}
	req.CountryCode = testData["country-code"]
	req.Providers = append(req.Providers, &domSchema.Provider{Code: testData["provider-code"]})

	res, err := adp.DisplayCurrentDataByCountry(&req)
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
