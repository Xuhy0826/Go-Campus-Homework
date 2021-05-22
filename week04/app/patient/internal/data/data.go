package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kwdms/app/patient/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPatientRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
