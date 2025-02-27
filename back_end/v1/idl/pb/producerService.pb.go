// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: producerService.proto

package pb

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

type ProducerModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" form:"id"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @gotags: json:"name" form:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	// @gotags: json:"short_name" form:"short_name"
	ShortName string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name" form:"short_name"`
	// @gotags: json:"address" form:"address"
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address" form:"address"`
	// @gotags: json:"tel" form:"tel"
	Tel string `protobuf:"bytes,5,opt,name=tel,proto3" json:"tel" form:"tel"`
	// @gotags: json:"email" form:"email"
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email" form:"email"`
	// @gotags: json:"contact_name" form:"contact_name"
	ContactName string `protobuf:"bytes,7,opt,name=contact_name,json=contactName,proto3" json:"contact_name" form:"contact_name"`
	// @gotags: json:"contact_tel" form:"contact_tel"
	ContactTel string `protobuf:"bytes,8,opt,name=contact_tel,json=contactTel,proto3" json:"contact_tel" form:"contact_tel"`
	// @gotags: json:"note" form:"note"
	Note string `protobuf:"bytes,9,opt,name=note,proto3" json:"note" form:"note"`
}

func (x *ProducerModel) Reset() {
	*x = ProducerModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProducerModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProducerModel) ProtoMessage() {}

func (x *ProducerModel) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProducerModel.ProtoReflect.Descriptor instead.
func (*ProducerModel) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{0}
}

func (x *ProducerModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProducerModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProducerModel) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *ProducerModel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProducerModel) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *ProducerModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProducerModel) GetContactName() string {
	if x != nil {
		return x.ContactName
	}
	return ""
}

func (x *ProducerModel) GetContactTel() string {
	if x != nil {
		return x.ContactTel
	}
	return ""
}

func (x *ProducerModel) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type InsertProducerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"producer_list" form:"producer_list"
	ProducerList []*ProducerModel `protobuf:"bytes,1,rep,name=producer_list,json=producerList,proto3" json:"producer_list" form:"producer_list"`
}

func (x *InsertProducerRequest) Reset() {
	*x = InsertProducerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertProducerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertProducerRequest) ProtoMessage() {}

func (x *InsertProducerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertProducerRequest.ProtoReflect.Descriptor instead.
func (*InsertProducerRequest) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{1}
}

func (x *InsertProducerRequest) GetProducerList() []*ProducerModel {
	if x != nil {
		return x.ProducerList
	}
	return nil
}

type InsertProducerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"success_list" form:"success_list"
	SuccessList []*ProducerModel `protobuf:"bytes,1,rep,name=success_list,json=successList,proto3" json:"success_list" form:"success_list"`
	// @gotags: json:"fail_list" form:"fail_list"
	FailList []*ProducerModel `protobuf:"bytes,2,rep,name=fail_list,json=failList,proto3" json:"fail_list" form:"fail_list"`
	// @gotags: json:"code" form:"code"
	Code int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code" form:"code"`
}

func (x *InsertProducerResponse) Reset() {
	*x = InsertProducerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertProducerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertProducerResponse) ProtoMessage() {}

func (x *InsertProducerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertProducerResponse.ProtoReflect.Descriptor instead.
func (*InsertProducerResponse) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{2}
}

func (x *InsertProducerResponse) GetSuccessList() []*ProducerModel {
	if x != nil {
		return x.SuccessList
	}
	return nil
}

func (x *InsertProducerResponse) GetFailList() []*ProducerModel {
	if x != nil {
		return x.FailList
	}
	return nil
}

func (x *InsertProducerResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type UpdateProducerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"producer" form:"producer"
	Producer *ProducerModel `protobuf:"bytes,1,opt,name=producer,proto3" json:"producer" form:"producer"`
}

func (x *UpdateProducerRequest) Reset() {
	*x = UpdateProducerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProducerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProducerRequest) ProtoMessage() {}

func (x *UpdateProducerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProducerRequest.ProtoReflect.Descriptor instead.
func (*UpdateProducerRequest) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProducerRequest) GetProducer() *ProducerModel {
	if x != nil {
		return x.Producer
	}
	return nil
}

type UpdateProducerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"success" form:"success"
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success" form:"success"`
	// @gotags: json:"code" form:"code"
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code" form:"code"`
}

func (x *UpdateProducerResponse) Reset() {
	*x = UpdateProducerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProducerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProducerResponse) ProtoMessage() {}

func (x *UpdateProducerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProducerResponse.ProtoReflect.Descriptor instead.
func (*UpdateProducerResponse) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProducerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateProducerResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SelectProducerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SelectProducerRequest) Reset() {
	*x = SelectProducerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectProducerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectProducerRequest) ProtoMessage() {}

func (x *SelectProducerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectProducerRequest.ProtoReflect.Descriptor instead.
func (*SelectProducerRequest) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{5}
}

type SelectProducerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"producer_list" form:"producer_list"
	ProducerList []*ProducerModel `protobuf:"bytes,1,rep,name=producer_list,json=producerList,proto3" json:"producer_list" form:"producer_list"`
	// @gotags: json:"code" form:"code"
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code" form:"code"`
}

func (x *SelectProducerResponse) Reset() {
	*x = SelectProducerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectProducerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectProducerResponse) ProtoMessage() {}

func (x *SelectProducerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectProducerResponse.ProtoReflect.Descriptor instead.
func (*SelectProducerResponse) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{6}
}

func (x *SelectProducerResponse) GetProducerList() []*ProducerModel {
	if x != nil {
		return x.ProducerList
	}
	return nil
}

func (x *SelectProducerResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DeleteProducerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"producer" form:"producer"
	Producer *ProducerModel `protobuf:"bytes,1,opt,name=producer,proto3" json:"producer" form:"producer"`
}

func (x *DeleteProducerRequest) Reset() {
	*x = DeleteProducerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProducerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProducerRequest) ProtoMessage() {}

func (x *DeleteProducerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProducerRequest.ProtoReflect.Descriptor instead.
func (*DeleteProducerRequest) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProducerRequest) GetProducer() *ProducerModel {
	if x != nil {
		return x.Producer
	}
	return nil
}

type DeleteProducerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"success" form:"success"
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success" form:"success"`
	// @gotags: json:"code" form:"code"
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code" form:"code"`
}

func (x *DeleteProducerResponse) Reset() {
	*x = DeleteProducerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producerService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProducerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProducerResponse) ProtoMessage() {}

func (x *DeleteProducerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producerService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProducerResponse.ProtoReflect.Descriptor instead.
func (*DeleteProducerResponse) Descriptor() ([]byte, []int) {
	return file_producerService_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteProducerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteProducerResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_producerService_proto protoreflect.FileDescriptor

var file_producerService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x22, 0x55, 0x0a, 0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x17,
	0x0a, 0x15, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x16, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x4c, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x22, 0x46, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xe5, 0x02, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a,
	0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_producerService_proto_rawDescOnce sync.Once
	file_producerService_proto_rawDescData = file_producerService_proto_rawDesc
)

func file_producerService_proto_rawDescGZIP() []byte {
	file_producerService_proto_rawDescOnce.Do(func() {
		file_producerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_producerService_proto_rawDescData)
	})
	return file_producerService_proto_rawDescData
}

var file_producerService_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_producerService_proto_goTypes = []any{
	(*ProducerModel)(nil),          // 0: services.ProducerModel
	(*InsertProducerRequest)(nil),  // 1: services.InsertProducerRequest
	(*InsertProducerResponse)(nil), // 2: services.InsertProducerResponse
	(*UpdateProducerRequest)(nil),  // 3: services.UpdateProducerRequest
	(*UpdateProducerResponse)(nil), // 4: services.UpdateProducerResponse
	(*SelectProducerRequest)(nil),  // 5: services.SelectProducerRequest
	(*SelectProducerResponse)(nil), // 6: services.SelectProducerResponse
	(*DeleteProducerRequest)(nil),  // 7: services.DeleteProducerRequest
	(*DeleteProducerResponse)(nil), // 8: services.DeleteProducerResponse
}
var file_producerService_proto_depIdxs = []int32{
	0,  // 0: services.InsertProducerRequest.producer_list:type_name -> services.ProducerModel
	0,  // 1: services.InsertProducerResponse.success_list:type_name -> services.ProducerModel
	0,  // 2: services.InsertProducerResponse.fail_list:type_name -> services.ProducerModel
	0,  // 3: services.UpdateProducerRequest.producer:type_name -> services.ProducerModel
	0,  // 4: services.SelectProducerResponse.producer_list:type_name -> services.ProducerModel
	0,  // 5: services.DeleteProducerRequest.producer:type_name -> services.ProducerModel
	1,  // 6: services.ProducerService.InsertProducer:input_type -> services.InsertProducerRequest
	3,  // 7: services.ProducerService.UpdateProducer:input_type -> services.UpdateProducerRequest
	5,  // 8: services.ProducerService.SelectProducer:input_type -> services.SelectProducerRequest
	7,  // 9: services.ProducerService.DeleteProducer:input_type -> services.DeleteProducerRequest
	2,  // 10: services.ProducerService.InsertProducer:output_type -> services.InsertProducerResponse
	4,  // 11: services.ProducerService.UpdateProducer:output_type -> services.UpdateProducerResponse
	6,  // 12: services.ProducerService.SelectProducer:output_type -> services.SelectProducerResponse
	8,  // 13: services.ProducerService.DeleteProducer:output_type -> services.DeleteProducerResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_producerService_proto_init() }
func file_producerService_proto_init() {
	if File_producerService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_producerService_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProducerModel); i {
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
		file_producerService_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InsertProducerRequest); i {
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
		file_producerService_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InsertProducerResponse); i {
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
		file_producerService_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProducerRequest); i {
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
		file_producerService_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProducerResponse); i {
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
		file_producerService_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SelectProducerRequest); i {
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
		file_producerService_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SelectProducerResponse); i {
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
		file_producerService_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProducerRequest); i {
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
		file_producerService_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProducerResponse); i {
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
			RawDescriptor: file_producerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_producerService_proto_goTypes,
		DependencyIndexes: file_producerService_proto_depIdxs,
		MessageInfos:      file_producerService_proto_msgTypes,
	}.Build()
	File_producerService_proto = out.File
	file_producerService_proto_rawDesc = nil
	file_producerService_proto_goTypes = nil
	file_producerService_proto_depIdxs = nil
}
