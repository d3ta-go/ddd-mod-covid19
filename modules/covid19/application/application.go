package application

import (
	appSvc "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/application/service"
	"github.com/d3ta-go/system/system/handler"
)

// NewCovid19App new Covid19App
func NewCovid19App(h *handler.Handler) (*Covid19App, error) {
	var err error

	app := new(Covid19App)
	app.handler = h

	if app.CurrentSvc, err = appSvc.NewCurrentSvc(h); err != nil {
		return nil, err
	}

	return app, nil
}

// Covid19App represent Covid19App
type Covid19App struct {
	handler    *handler.Handler
	CurrentSvc *appSvc.CurrentSvc
}
