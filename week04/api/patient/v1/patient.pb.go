// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: patient.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PatientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId   string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	PatientName string `protobuf:"bytes,2,opt,name=patient_name,json=patientName,proto3" json:"patient_name,omitempty"`
	PactCode    string `protobuf:"bytes,3,opt,name=pact_code,json=pactCode,proto3" json:"pact_code,omitempty"`
	PactName    string `protobuf:"bytes,4,opt,name=pact_name,json=pactName,proto3" json:"pact_name,omitempty"`
	BirthDay    string `protobuf:"bytes,5,opt,name=birth_day,json=birthDay,proto3" json:"birth_day,omitempty"`
	Age         int32  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	Gender      string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	IdenNo      string `protobuf:"bytes,8,opt,name=iden_no,json=idenNo,proto3" json:"iden_no,omitempty"`
	IdcardType  int32  `protobuf:"varint,9,opt,name=idcard_type,json=idcardType,proto3" json:"idcard_type,omitempty"`
	HomeAddress string `protobuf:"bytes,10,opt,name=home_address,json=homeAddress,proto3" json:"home_address,omitempty"`
	PhoneNumber string `protobuf:"bytes,11,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *PatientInfo) Reset() {
	*x = PatientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientInfo) ProtoMessage() {}

func (x *PatientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientInfo.ProtoReflect.Descriptor instead.
func (*PatientInfo) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{0}
}

func (x *PatientInfo) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *PatientInfo) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *PatientInfo) GetPactCode() string {
	if x != nil {
		return x.PactCode
	}
	return ""
}

func (x *PatientInfo) GetPactName() string {
	if x != nil {
		return x.PactName
	}
	return ""
}

func (x *PatientInfo) GetBirthDay() string {
	if x != nil {
		return x.BirthDay
	}
	return ""
}

func (x *PatientInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PatientInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientInfo) GetIdenNo() string {
	if x != nil {
		return x.IdenNo
	}
	return ""
}

func (x *PatientInfo) GetIdcardType() int32 {
	if x != nil {
		return x.IdcardType
	}
	return 0
}

func (x *PatientInfo) GetHomeAddress() string {
	if x != nil {
		return x.HomeAddress
	}
	return ""
}

func (x *PatientInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetPatientInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
}

func (x *GetPatientInfoRequest) Reset() {
	*x = GetPatientInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientInfoRequest) ProtoMessage() {}

func (x *GetPatientInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPatientInfoRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{1}
}

func (x *GetPatientInfoRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

type GetPatientInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *PatientInfo `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *GetPatientInfoReply) Reset() {
	*x = GetPatientInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientInfoReply) ProtoMessage() {}

func (x *GetPatientInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientInfoReply.ProtoReflect.Descriptor instead.
func (*GetPatientInfoReply) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{2}
}

func (x *GetPatientInfoReply) GetPatient() *PatientInfo {
	if x != nil {
		return x.Patient
	}
	return nil
}

type ListPatientInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientIds []string `protobuf:"bytes,1,rep,name=patient_ids,json=patientIds,proto3" json:"patient_ids,omitempty"`
}

func (x *ListPatientInfoRequest) Reset() {
	*x = ListPatientInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientInfoRequest) ProtoMessage() {}

func (x *ListPatientInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientInfoRequest.ProtoReflect.Descriptor instead.
func (*ListPatientInfoRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{3}
}

func (x *ListPatientInfoRequest) GetPatientIds() []string {
	if x != nil {
		return x.PatientIds
	}
	return nil
}

type ListPatientInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient []*PatientInfo `protobuf:"bytes,1,rep,name=patient,proto3" json:"patient,omitempty"`
}

func (x *ListPatientInfoReply) Reset() {
	*x = ListPatientInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientInfoReply) ProtoMessage() {}

func (x *ListPatientInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientInfoReply.ProtoReflect.Descriptor instead.
func (*ListPatientInfoReply) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{4}
}

func (x *ListPatientInfoReply) GetPatient() []*PatientInfo {
	if x != nil {
		return x.Patient
	}
	return nil
}

var File_patient_proto protoreflect.FileDescriptor

var file_patient_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22,
	0xd0, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x6e, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x65, 0x6e, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x64, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x64, 0x63, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5c,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2b, 0x0a,
	0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x17, 0x6b, 0x77, 0x64, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_patient_proto_rawDescOnce sync.Once
	file_patient_proto_rawDescData = file_patient_proto_rawDesc
)

func file_patient_proto_rawDescGZIP() []byte {
	file_patient_proto_rawDescOnce.Do(func() {
		file_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_patient_proto_rawDescData)
	})
	return file_patient_proto_rawDescData
}

var file_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_patient_proto_goTypes = []interface{}{
	(*PatientInfo)(nil),            // 0: api.patient.v1.PatientInfo
	(*GetPatientInfoRequest)(nil),  // 1: api.patient.v1.GetPatientInfoRequest
	(*GetPatientInfoReply)(nil),    // 2: api.patient.v1.GetPatientInfoReply
	(*ListPatientInfoRequest)(nil), // 3: api.patient.v1.ListPatientInfoRequest
	(*ListPatientInfoReply)(nil),   // 4: api.patient.v1.ListPatientInfoReply
}
var file_patient_proto_depIdxs = []int32{
	0, // 0: api.patient.v1.GetPatientInfoReply.patient:type_name -> api.patient.v1.PatientInfo
	0, // 1: api.patient.v1.ListPatientInfoReply.patient:type_name -> api.patient.v1.PatientInfo
	1, // 2: api.patient.v1.Patient.GetPatientInfo:input_type -> api.patient.v1.GetPatientInfoRequest
	3, // 3: api.patient.v1.Patient.ListPatientInfo:input_type -> api.patient.v1.ListPatientInfoRequest
	2, // 4: api.patient.v1.Patient.GetPatientInfo:output_type -> api.patient.v1.GetPatientInfoReply
	4, // 5: api.patient.v1.Patient.ListPatientInfo:output_type -> api.patient.v1.ListPatientInfoReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_patient_proto_init() }
func file_patient_proto_init() {
	if File_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatientInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatientInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_patient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_patient_proto_goTypes,
		DependencyIndexes: file_patient_proto_depIdxs,
		MessageInfos:      file_patient_proto_msgTypes,
	}.Build()
	File_patient_proto = out.File
	file_patient_proto_rawDesc = nil
	file_patient_proto_goTypes = nil
	file_patient_proto_depIdxs = nil
}