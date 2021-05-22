package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kwdms/app/patient/internal/biz"
)

func NewPatientRepo(data *Data, logger log.Logger) biz.PatientRepo {
	return &patientRepo{
		data: data,
		log:  log.NewHelper("data/beer", logger),
	}
}

type patientRepo struct {
	data *Data
	log  *log.Helper
}

func (p patientRepo) GetPatient(ctx context.Context, id string) (*biz.Patient, error) {
	panic("implement me")
}

func (p patientRepo) ListPatient(ctx context.Context, ids []string) ([]*biz.Patient, error) {
	panic("implement me")
}
