// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.2
// source: managerService.proto

package manager

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

type Sex int32

const (
	Sex_Sex_unknown Sex = 0
	Sex_Sex_man     Sex = 1
	Sex_Sex_woman   Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "Sex_unknown",
		1: "Sex_man",
		2: "Sex_woman",
	}
	Sex_value = map[string]int32{
		"Sex_unknown": 0,
		"Sex_man":     1,
		"Sex_woman":   2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_managerService_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_managerService_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{0}
}

type RequestManagerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid   string   `protobuf:"bytes,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Manager *Manager `protobuf:"bytes,3,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *RequestManagerData) Reset() {
	*x = RequestManagerData{}
	mi := &file_managerService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestManagerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestManagerData) ProtoMessage() {}

func (x *RequestManagerData) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestManagerData.ProtoReflect.Descriptor instead.
func (*RequestManagerData) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{0}
}

func (x *RequestManagerData) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *RequestManagerData) GetManager() *Manager {
	if x != nil {
		return x.Manager
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_managerService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseManagerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Managers  []*Manager `protobuf:"bytes,2,rep,name=managers,proto3" json:"managers,omitempty"`
}

func (x *ResponseManagerData) Reset() {
	*x = ResponseManagerData{}
	mi := &file_managerService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseManagerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseManagerData) ProtoMessage() {}

func (x *ResponseManagerData) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseManagerData.ProtoReflect.Descriptor instead.
func (*ResponseManagerData) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseManagerData) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ResponseManagerData) GetManagers() []*Manager {
	if x != nil {
		return x.Managers
	}
	return nil
}

type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    []string          `protobuf:"bytes,2,rep,name=email,proto3" json:"email,omitempty"`
	PhoneNum string            `protobuf:"bytes,3,opt,name=phoneNum,proto3" json:"phoneNum,omitempty"`
	Age     int32             `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Sex     Sex               `protobuf:"varint,5,opt,name=sex,proto3,enum=proto.Sex" json:"sex,omitempty"`
	Hobbies map[string]string `protobuf:"bytes,6,rep,name=hobbies,proto3" json:"hobbies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Helper:
	//
	//	*Manager_ManHelper
	//	*Manager_WomanHelper
	Helper isManager_Helper `protobuf_oneof:"helper"`
	Job    *Job             `protobuf:"bytes,9,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	mi := &file_managerService_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{3}
}

func (x *Manager) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manager) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *Manager) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *Manager) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Manager) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_Sex_unknown
}

func (x *Manager) GetHobbies() map[string]string {
	if x != nil {
		return x.Hobbies
	}
	return nil
}

func (m *Manager) GetHelper() isManager_Helper {
	if m != nil {
		return m.Helper
	}
	return nil
}

func (x *Manager) GetManHelper() string {
	if x, ok := x.GetHelper().(*Manager_ManHelper); ok {
		return x.ManHelper
	}
	return ""
}

func (x *Manager) GetWomanHelper() string {
	if x, ok := x.GetHelper().(*Manager_WomanHelper); ok {
		return x.WomanHelper
	}
	return ""
}

func (x *Manager) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type isManager_Helper interface {
	isManager_Helper()
}

type Manager_ManHelper struct {
	ManHelper string `protobuf:"bytes,7,opt,name=man_helper,json=manHelper,proto3,oneof"`
}

type Manager_WomanHelper struct {
	WomanHelper string `protobuf:"bytes,8,opt,name=woman_helper,json=womanHelper,proto3,oneof"`
}

func (*Manager_ManHelper) isManager_Helper() {}

func (*Manager_WomanHelper) isManager_Helper() {}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName,omitempty"`
	Money   int32  `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	Info    string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_managerService_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{4}
}

func (x *Job) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *Job) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Job) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_managerService_proto protoreflect.FileDescriptor

var file_managerService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x5f, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x22, 0xe0, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x78, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x6d,
	0x61, 0x6e, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0c,
	0x77, 0x6f, 0x6d, 0x61, 0x6e, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x77, 0x6f, 0x6d, 0x61, 0x6e, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x1a,
	0x3a, 0x0a, 0x0c, 0x48, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x0a, 0x07,
	0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a,
	0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x2a, 0x32, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x78, 0x5f, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x78, 0x5f,
	0x6d, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x65, 0x78, 0x5f, 0x77, 0x6f, 0x6d,
	0x61, 0x6e, 0x10, 0x02, 0x32, 0x98, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42,
	0x18, 0x5a, 0x16, 0x61, 0x6e, 0x64, 0x79, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_managerService_proto_rawDescOnce sync.Once
	file_managerService_proto_rawDescData = file_managerService_proto_rawDesc
)

func file_managerService_proto_rawDescGZIP() []byte {
	file_managerService_proto_rawDescOnce.Do(func() {
		file_managerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_managerService_proto_rawDescData)
	})
	return file_managerService_proto_rawDescData
}

var file_managerService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_managerService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_managerService_proto_goTypes = []any{
	(Sex)(0),                    // 0: proto.Sex
	(*RequestManagerData)(nil),  // 1: proto.RequestManagerData
	(*Response)(nil),            // 2: proto.Response
	(*ResponseManagerData)(nil), // 3: proto.ResponseManagerData
	(*Manager)(nil),             // 4: proto.Manager
	(*Job)(nil),                 // 5: proto.Job
	nil,                         // 6: proto.Manager.HobbiesEntry
}
var file_managerService_proto_depIdxs = []int32{
	4, // 0: proto.RequestManagerData.manager:type_name -> proto.Manager
	4, // 1: proto.ResponseManagerData.managers:type_name -> proto.Manager
	0, // 2: proto.Manager.sex:type_name -> proto.Sex
	6, // 3: proto.Manager.hobbies:type_name -> proto.Manager.HobbiesEntry
	5, // 4: proto.Manager.job:type_name -> proto.Job
	1, // 5: proto.ManagerService.createManager:input_type -> proto.RequestManagerData
	1, // 6: proto.ManagerService.listManagers:input_type -> proto.RequestManagerData
	2, // 7: proto.ManagerService.createManager:output_type -> proto.Response
	3, // 8: proto.ManagerService.listManagers:output_type -> proto.ResponseManagerData
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_managerService_proto_init() }
func file_managerService_proto_init() {
	if File_managerService_proto != nil {
		return
	}
	file_managerService_proto_msgTypes[3].OneofWrappers = []any{
		(*Manager_ManHelper)(nil),
		(*Manager_WomanHelper)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managerService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managerService_proto_goTypes,
		DependencyIndexes: file_managerService_proto_depIdxs,
		EnumInfos:         file_managerService_proto_enumTypes,
		MessageInfos:      file_managerService_proto_msgTypes,
	}.Build()
	File_managerService_proto = out.File
	file_managerService_proto_rawDesc = nil
	file_managerService_proto_goTypes = nil
	file_managerService_proto_depIdxs = nil
}