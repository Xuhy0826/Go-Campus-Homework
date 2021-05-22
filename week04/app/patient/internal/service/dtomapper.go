package service

import (
	v1 "kwdms/api/patient/v1"
	"kwdms/app/patient/internal/biz"
	"kwdms/pkg/utils/timeutil"
)

func ToPatientDto(p *biz.Patient) (dto v1.PatientInfo, err error) {
	dto = v1.PatientInfo{}
	if p == nil {
		return
	}
	dto.PatientId = p.PatientId
	dto.PatientName = p.PatientName
	dto.PactCode = p.PactCode
	dto.PactName = p.PactName
	dto.HomeAddress = p.HomeAddress
	dto.IdenNo = p.IdenNo
	dto.IdcardType = p.IdcardType
	dto.PhoneNumber = p.PhoneNumber
	dto.BirthDay = timeutil.ToDateString(p.BirthDay)
	dto.Age = timeutil.GetAge(p.BirthDay)
	if p.SexCode == "M" {
		dto.Gender = "男"
	} else if p.SexCode == "F" {
		dto.Gender = "女"
	} else {
		dto.Gender = "未知"
	}
	return
}
