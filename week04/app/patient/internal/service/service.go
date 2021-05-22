package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kwdms/api/patient/v1"
	"kwdms/app/patient/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPatientService)

type PatientService struct {
	v1.UnimplementedPatientServer

	pc  *biz.PatientUseCase
	log *log.Helper
}

func NewPatientService(pc *biz.PatientUseCase, logger log.Logger) *PatientService {
	return &PatientService{
		pc:  pc,
		log: log.NewHelper("service/catalog", logger),
	}
}
