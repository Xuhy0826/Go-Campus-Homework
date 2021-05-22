package service

import (
	"context"

	v1 "kwdms/api/patient/v1"
)

//GetPatientInfo Get patient basic information by patient_id
func (s *PatientService) GetPatientInfo(ctx context.Context, req *v1.GetPatientInfoRequest) (*v1.GetPatientInfoReply, error) {
	patient, err := s.pc.Get(ctx, req.GetPatientId())
	if err != nil {
		return nil, err
	}
	patientInfo, _ := ToPatientDto(patient)
	return &v1.GetPatientInfoReply{
		Patient: &patientInfo,
	}, err
}

//ListPatientInfo Get patient basic info collection by given patient id slice
func (s *PatientService) ListPatientInfo(ctx context.Context, req *v1.ListPatientInfoRequest) (*v1.ListPatientInfoReply, error) {
	return &v1.ListPatientInfoReply{}, nil
}
