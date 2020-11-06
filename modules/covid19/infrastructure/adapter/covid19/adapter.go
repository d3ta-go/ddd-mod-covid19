package adapter

import domSchema "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/domain/schema"

// ICovid19Adapter represent Covid19Adapter interface
type ICovid19Adapter interface {
	DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest) (*domSchema.TotalCountryProviderData, error)
}

// BaseCovid19Adapter represent BaseCovid19 Adapter
type BaseCovid19Adapter struct {
	info Covid19AdapterInfo
}

// SetInfo set Info
func (b *BaseCovid19Adapter) SetInfo(info Covid19AdapterInfo) {
	b.info = info
}

// GetInfo set Info
func (b *BaseCovid19Adapter) GetInfo() Covid19AdapterInfo {
	return b.info
}
