// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: internal/proto/config-backend.proto

package config_backend

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

type ErrorCode int32

const (
	ErrorCode_SUCCESS                   ErrorCode = 0
	ErrorCode_ERR_INVALID_SERVICE_ID    ErrorCode = 1
	ErrorCode_ERR_INVALID_SERVICE_KEY   ErrorCode = 2
	ErrorCode_ERR_INTERNAL_SERVER_ERROR ErrorCode = 999
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:   "SUCCESS",
		1:   "ERR_INVALID_SERVICE_ID",
		2:   "ERR_INVALID_SERVICE_KEY",
		999: "ERR_INTERNAL_SERVER_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"SUCCESS":                   0,
		"ERR_INVALID_SERVICE_ID":    1,
		"ERR_INVALID_SERVICE_KEY":   2,
		"ERR_INTERNAL_SERVER_ERROR": 999,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_config_backend_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_internal_proto_config_backend_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{0}
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode ErrorCode `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=config_backend.victor_leee.github.com.ErrorCode" json:"err_code,omitempty"`
	ErrMsg  string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetErrCode() ErrorCode {
	if x != nil {
		return x.ErrCode
	}
	return ErrorCode_SUCCESS
}

func (x *BaseResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId  string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceKey string `protobuf:"bytes,2,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	Key        string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *GetConfigRequest) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *GetConfigRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	KeyExist     bool          `protobuf:"varint,2,opt,name=key_exist,json=keyExist,proto3" json:"key_exist,omitempty"`
	Value        string        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{2}
}

func (x *GetConfigResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *GetConfigResponse) GetKeyExist() bool {
	if x != nil {
		return x.KeyExist
	}
	return false
}

func (x *GetConfigResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId  string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceKey string `protobuf:"bytes,2,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	Key        string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutConfigRequest) Reset() {
	*x = PutConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConfigRequest) ProtoMessage() {}

func (x *PutConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConfigRequest.ProtoReflect.Descriptor instead.
func (*PutConfigRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{3}
}

func (x *PutConfigRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *PutConfigRequest) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *PutConfigRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutConfigRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
}

func (x *PutConfigResponse) Reset() {
	*x = PutConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConfigResponse) ProtoMessage() {}

func (x *PutConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConfigResponse.ProtoReflect.Descriptor instead.
func (*PutConfigResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{4}
}

func (x *PutConfigResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type GetAllKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId  string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceKey string `protobuf:"bytes,2,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
}

func (x *GetAllKeysRequest) Reset() {
	*x = GetAllKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKeysRequest) ProtoMessage() {}

func (x *GetAllKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKeysRequest.ProtoReflect.Descriptor instead.
func (*GetAllKeysRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllKeysRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *GetAllKeysRequest) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

type GetAllKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	Keys         []string      `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *GetAllKeysResponse) Reset() {
	*x = GetAllKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_config_backend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKeysResponse) ProtoMessage() {}

func (x *GetAllKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_config_backend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKeysResponse.ProtoReflect.Descriptor instead.
func (*GetAllKeysResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_config_backend_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllKeysResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *GetAllKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_internal_proto_config_backend_proto protoreflect.FileDescriptor

var file_internal_proto_config_backend_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65,
	0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x74, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08,
	0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x22, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58,
	0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65,
	0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a, 0x0a, 0x10, 0x50,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6d, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x2a, 0x71, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xe7, 0x07, 0x32, 0xa2, 0x03, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x80, 0x01, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76,
	0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c,
	0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x50,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x83, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x69, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x6c, 0x65,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_config_backend_proto_rawDescOnce sync.Once
	file_internal_proto_config_backend_proto_rawDescData = file_internal_proto_config_backend_proto_rawDesc
)

func file_internal_proto_config_backend_proto_rawDescGZIP() []byte {
	file_internal_proto_config_backend_proto_rawDescOnce.Do(func() {
		file_internal_proto_config_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_config_backend_proto_rawDescData)
	})
	return file_internal_proto_config_backend_proto_rawDescData
}

var file_internal_proto_config_backend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_proto_config_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_proto_config_backend_proto_goTypes = []interface{}{
	(ErrorCode)(0),             // 0: config_backend.victor_leee.github.com.ErrorCode
	(*BaseResponse)(nil),       // 1: config_backend.victor_leee.github.com.BaseResponse
	(*GetConfigRequest)(nil),   // 2: config_backend.victor_leee.github.com.GetConfigRequest
	(*GetConfigResponse)(nil),  // 3: config_backend.victor_leee.github.com.GetConfigResponse
	(*PutConfigRequest)(nil),   // 4: config_backend.victor_leee.github.com.PutConfigRequest
	(*PutConfigResponse)(nil),  // 5: config_backend.victor_leee.github.com.PutConfigResponse
	(*GetAllKeysRequest)(nil),  // 6: config_backend.victor_leee.github.com.GetAllKeysRequest
	(*GetAllKeysResponse)(nil), // 7: config_backend.victor_leee.github.com.GetAllKeysResponse
}
var file_internal_proto_config_backend_proto_depIdxs = []int32{
	0, // 0: config_backend.victor_leee.github.com.BaseResponse.err_code:type_name -> config_backend.victor_leee.github.com.ErrorCode
	1, // 1: config_backend.victor_leee.github.com.GetConfigResponse.base_response:type_name -> config_backend.victor_leee.github.com.BaseResponse
	1, // 2: config_backend.victor_leee.github.com.PutConfigResponse.base_response:type_name -> config_backend.victor_leee.github.com.BaseResponse
	1, // 3: config_backend.victor_leee.github.com.GetAllKeysResponse.base_response:type_name -> config_backend.victor_leee.github.com.BaseResponse
	2, // 4: config_backend.victor_leee.github.com.ConfigBackendService.GetConfig:input_type -> config_backend.victor_leee.github.com.GetConfigRequest
	4, // 5: config_backend.victor_leee.github.com.ConfigBackendService.PutConfig:input_type -> config_backend.victor_leee.github.com.PutConfigRequest
	6, // 6: config_backend.victor_leee.github.com.ConfigBackendService.GetAllKeys:input_type -> config_backend.victor_leee.github.com.GetAllKeysRequest
	3, // 7: config_backend.victor_leee.github.com.ConfigBackendService.GetConfig:output_type -> config_backend.victor_leee.github.com.GetConfigResponse
	5, // 8: config_backend.victor_leee.github.com.ConfigBackendService.PutConfig:output_type -> config_backend.victor_leee.github.com.PutConfigResponse
	7, // 9: config_backend.victor_leee.github.com.ConfigBackendService.GetAllKeys:output_type -> config_backend.victor_leee.github.com.GetAllKeysResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_proto_config_backend_proto_init() }
func file_internal_proto_config_backend_proto_init() {
	if File_internal_proto_config_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_config_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_internal_proto_config_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_internal_proto_config_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
		file_internal_proto_config_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConfigRequest); i {
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
		file_internal_proto_config_backend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConfigResponse); i {
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
		file_internal_proto_config_backend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKeysRequest); i {
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
		file_internal_proto_config_backend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKeysResponse); i {
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
			RawDescriptor: file_internal_proto_config_backend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_config_backend_proto_goTypes,
		DependencyIndexes: file_internal_proto_config_backend_proto_depIdxs,
		EnumInfos:         file_internal_proto_config_backend_proto_enumTypes,
		MessageInfos:      file_internal_proto_config_backend_proto_msgTypes,
	}.Build()
	File_internal_proto_config_backend_proto = out.File
	file_internal_proto_config_backend_proto_rawDesc = nil
	file_internal_proto_config_backend_proto_goTypes = nil
	file_internal_proto_config_backend_proto_depIdxs = nil
}
