package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"kwdms/pkg/utils/common"
	"time"
)

type Patient struct {
	PatientId   string
	PatientName string
	PactCode    string
	PactName    string
	BirthDay    time.Time
	SexCode     string
	IdenNo      string
	IdcardType  int32
	HomeAddress string
	PhoneNumber string
}

type PatientRepo interface {
	GetPatient(ctx context.Context, id string) (*Patient, error)
	ListPatient(ctx context.Context, ids []string) ([]*Patient, error)
}

type PatientUseCase struct {
	repo PatientRepo
	log  *log.Helper
}

//NewPatientUseCase ctor
func NewPatientUseCase(repo PatientRepo, logger log.Logger) *PatientUseCase {
	return &PatientUseCase{
		repo: repo,
		log:  log.NewHelper("usecase/patient", logger),
	}
}

//Get query patient entity
func (uc *PatientUseCase) Get(ctx context.Context, pid string) (*Patient, error) {
	pid, err := common.RepairPatientId(pid)
	if err != nil {
		return nil, errors.BadRequest("kwdms.patient.get", "invalid patient-id", err.Error())
	}
	return uc.repo.GetPatient(ctx, pid)
}

//List query patients entity
func (uc *PatientUseCase) List(ctx context.Context, pids []string) ([]*Patient, error) {
	return uc.repo.ListPatient(ctx, pids)
}
