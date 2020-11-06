package repository

import (
	"encoding/json"
	"testing"

	domRepo "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/repository"
	domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newRepoIdent(t *testing.T) (domRepo.ICurrentRepo, identity.Identity, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, identity.Identity{}, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, identity.Identity{}, err
	}

	h.SetDefaultConfig(c)

	r, err := NewCurrentRepo(h)
	if err != nil {
		return nil, identity.Identity{}, err
	}

	i := newIdentity(h, t)

	return r, i, nil
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
	repo, i, err := newRepoIdent(t)
	if err != nil {
		t.Errorf("REPO: [%#v]", err.Error())
	}

	var req domSchema.DisplayCurrentDataByCountryRequest

	err = json.Unmarshal([]byte(`{ "countryCode": "ID", "providers": [  {"code": "_ALL_" } ] }`), &req)
	// err = json.Unmarshal([]byte(`{ "countryCode": "ID", "providers": [  {"code": "WHO" }, {"code": "WHO" } ] }`), &req)
	if err != nil {
		t.Errorf("Request: [%s]", err.Error())
	}

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
