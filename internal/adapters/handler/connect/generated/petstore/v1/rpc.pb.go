// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: petstore/v1/rpc.proto

package petstorev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PetType int32

const (
	PetType_PET_TYPE_UNSPECIFIED PetType = 0
	PetType_PET_TYPE_DOG         PetType = 1
	PetType_PET_TYPE_CAT         PetType = 2
	PetType_PET_TYPE_RABBIT      PetType = 3
)

// Enum value maps for PetType.
var (
	PetType_name = map[int32]string{
		0: "PET_TYPE_UNSPECIFIED",
		1: "PET_TYPE_DOG",
		2: "PET_TYPE_CAT",
		3: "PET_TYPE_RABBIT",
	}
	PetType_value = map[string]int32{
		"PET_TYPE_UNSPECIFIED": 0,
		"PET_TYPE_DOG":         1,
		"PET_TYPE_CAT":         2,
		"PET_TYPE_RABBIT":      3,
	}
)

func (x PetType) Enum() *PetType {
	p := new(PetType)
	*p = x
	return p
}

func (x PetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PetType) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_v1_rpc_proto_enumTypes[0].Descriptor()
}

func (PetType) Type() protoreflect.EnumType {
	return &file_petstore_v1_rpc_proto_enumTypes[0]
}

func (x PetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PetType.Descriptor instead.
func (PetType) EnumDescriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{0}
}

type PetStatus int32

const (
	PetStatus_PET_STATUS_UNSPECIFIED PetStatus = 0
	PetStatus_PET_STATUS_AVAILABLE   PetStatus = 1
	PetStatus_PET_STATUS_PENDING     PetStatus = 2
	PetStatus_PET_STATUS_SOLD        PetStatus = 3
)

// Enum value maps for PetStatus.
var (
	PetStatus_name = map[int32]string{
		0: "PET_STATUS_UNSPECIFIED",
		1: "PET_STATUS_AVAILABLE",
		2: "PET_STATUS_PENDING",
		3: "PET_STATUS_SOLD",
	}
	PetStatus_value = map[string]int32{
		"PET_STATUS_UNSPECIFIED": 0,
		"PET_STATUS_AVAILABLE":   1,
		"PET_STATUS_PENDING":     2,
		"PET_STATUS_SOLD":        3,
	}
)

func (x PetStatus) Enum() *PetStatus {
	p := new(PetStatus)
	*p = x
	return p
}

func (x PetStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PetStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_v1_rpc_proto_enumTypes[1].Descriptor()
}

func (PetStatus) Type() protoreflect.EnumType {
	return &file_petstore_v1_rpc_proto_enumTypes[1]
}

func (x PetStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PetStatus.Descriptor instead.
func (PetStatus) EnumDescriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{1}
}

type UpdatePetStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pet           *Pet                   `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePetStatusResponse) Reset() {
	*x = UpdatePetStatusResponse{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetStatusResponse) ProtoMessage() {}

func (x *UpdatePetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdatePetStatusResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePetStatusResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type CreatePetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pet           *Pet                   `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePetResponse) Reset() {
	*x = CreatePetResponse{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetResponse) ProtoMessage() {}

func (x *CreatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetResponse.ProtoReflect.Descriptor instead.
func (*CreatePetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type GetPetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pet           *Pet                   `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPetResponse) Reset() {
	*x = GetPetResponse{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetResponse) ProtoMessage() {}

func (x *GetPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetResponse.ProtoReflect.Descriptor instead.
func (*GetPetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type Pet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          PetType                `protobuf:"varint,3,opt,name=type,proto3,enum=petstore.v1.PetType" json:"type,omitempty"`
	Status        PetStatus              `protobuf:"varint,4,opt,name=status,proto3,enum=petstore.v1.PetStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pet) Reset() {
	*x = Pet{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetType() PetType {
	if x != nil {
		return x.Type
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

func (x *Pet) GetStatus() PetStatus {
	if x != nil {
		return x.Status
	}
	return PetStatus_PET_STATUS_UNSPECIFIED
}

func (x *Pet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Pet) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreatePetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          PetType                `protobuf:"varint,2,opt,name=type,proto3,enum=petstore.v1.PetType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePetRequest) Reset() {
	*x = CreatePetRequest{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetRequest) ProtoMessage() {}

func (x *CreatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetRequest.ProtoReflect.Descriptor instead.
func (*CreatePetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePetRequest) GetType() PetType {
	if x != nil {
		return x.Type
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

type GetPetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPetRequest) Reset() {
	*x = GetPetRequest{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetRequest) ProtoMessage() {}

func (x *GetPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetRequest.ProtoReflect.Descriptor instead.
func (*GetPetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetPetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPetsRequest) Reset() {
	*x = ListPetsRequest{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsRequest) ProtoMessage() {}

func (x *ListPetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsRequest.ProtoReflect.Descriptor instead.
func (*ListPetsRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{6}
}

type ListPetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pets          []*Pet                 `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPetsResponse) Reset() {
	*x = ListPetsResponse{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsResponse) ProtoMessage() {}

func (x *ListPetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsResponse.ProtoReflect.Descriptor instead.
func (*ListPetsResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *ListPetsResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type UpdatePetStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        PetStatus              `protobuf:"varint,2,opt,name=status,proto3,enum=petstore.v1.PetStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePetStatusRequest) Reset() {
	*x = UpdatePetStatusRequest{}
	mi := &file_petstore_v1_rpc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetStatusRequest) ProtoMessage() {}

func (x *UpdatePetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_rpc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetStatusRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePetStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePetStatusRequest) GetStatus() PetStatus {
	if x != nil {
		return x.Status
	}
	return PetStatus_PET_STATUS_UNSPECIFIED
}

var File_petstore_v1_rpc_proto protoreflect.FileDescriptor

var file_petstore_v1_rpc_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52,
	0x03, 0x70, 0x65, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x34, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03,
	0x70, 0x65, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x50, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x22,
	0x58, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x5c, 0x0a, 0x07, 0x50, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x47, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x54,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x41, 0x42, 0x42, 0x49, 0x54, 0x10, 0x03, 0x2a, 0x6e, 0x0a, 0x09, 0x50, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x45,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x03, 0x32, 0xc7, 0x02, 0x0a, 0x0f, 0x50, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xd1, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61,
	0x72, 0x6f, 0x6e, 0x42, 0x6f, 0x6e, 0x65, 0x74, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2d, 0x70, 0x65,
	0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x0b, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0b, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x50,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_petstore_v1_rpc_proto_rawDescOnce sync.Once
	file_petstore_v1_rpc_proto_rawDescData []byte
)

func file_petstore_v1_rpc_proto_rawDescGZIP() []byte {
	file_petstore_v1_rpc_proto_rawDescOnce.Do(func() {
		file_petstore_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_petstore_v1_rpc_proto_rawDesc), len(file_petstore_v1_rpc_proto_rawDesc)))
	})
	return file_petstore_v1_rpc_proto_rawDescData
}

var file_petstore_v1_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_petstore_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_petstore_v1_rpc_proto_goTypes = []any{
	(PetType)(0),                    // 0: petstore.v1.PetType
	(PetStatus)(0),                  // 1: petstore.v1.PetStatus
	(*UpdatePetStatusResponse)(nil), // 2: petstore.v1.UpdatePetStatusResponse
	(*CreatePetResponse)(nil),       // 3: petstore.v1.CreatePetResponse
	(*GetPetResponse)(nil),          // 4: petstore.v1.GetPetResponse
	(*Pet)(nil),                     // 5: petstore.v1.Pet
	(*CreatePetRequest)(nil),        // 6: petstore.v1.CreatePetRequest
	(*GetPetRequest)(nil),           // 7: petstore.v1.GetPetRequest
	(*ListPetsRequest)(nil),         // 8: petstore.v1.ListPetsRequest
	(*ListPetsResponse)(nil),        // 9: petstore.v1.ListPetsResponse
	(*UpdatePetStatusRequest)(nil),  // 10: petstore.v1.UpdatePetStatusRequest
	(*timestamppb.Timestamp)(nil),   // 11: google.protobuf.Timestamp
}
var file_petstore_v1_rpc_proto_depIdxs = []int32{
	5,  // 0: petstore.v1.UpdatePetStatusResponse.pet:type_name -> petstore.v1.Pet
	5,  // 1: petstore.v1.CreatePetResponse.pet:type_name -> petstore.v1.Pet
	5,  // 2: petstore.v1.GetPetResponse.pet:type_name -> petstore.v1.Pet
	0,  // 3: petstore.v1.Pet.type:type_name -> petstore.v1.PetType
	1,  // 4: petstore.v1.Pet.status:type_name -> petstore.v1.PetStatus
	11, // 5: petstore.v1.Pet.created_at:type_name -> google.protobuf.Timestamp
	11, // 6: petstore.v1.Pet.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 7: petstore.v1.CreatePetRequest.type:type_name -> petstore.v1.PetType
	5,  // 8: petstore.v1.ListPetsResponse.pets:type_name -> petstore.v1.Pet
	1,  // 9: petstore.v1.UpdatePetStatusRequest.status:type_name -> petstore.v1.PetStatus
	6,  // 10: petstore.v1.PetStoreService.CreatePet:input_type -> petstore.v1.CreatePetRequest
	7,  // 11: petstore.v1.PetStoreService.GetPet:input_type -> petstore.v1.GetPetRequest
	8,  // 12: petstore.v1.PetStoreService.ListPets:input_type -> petstore.v1.ListPetsRequest
	10, // 13: petstore.v1.PetStoreService.UpdatePetStatus:input_type -> petstore.v1.UpdatePetStatusRequest
	3,  // 14: petstore.v1.PetStoreService.CreatePet:output_type -> petstore.v1.CreatePetResponse
	4,  // 15: petstore.v1.PetStoreService.GetPet:output_type -> petstore.v1.GetPetResponse
	9,  // 16: petstore.v1.PetStoreService.ListPets:output_type -> petstore.v1.ListPetsResponse
	2,  // 17: petstore.v1.PetStoreService.UpdatePetStatus:output_type -> petstore.v1.UpdatePetStatusResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_petstore_v1_rpc_proto_init() }
func file_petstore_v1_rpc_proto_init() {
	if File_petstore_v1_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_petstore_v1_rpc_proto_rawDesc), len(file_petstore_v1_rpc_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_v1_rpc_proto_goTypes,
		DependencyIndexes: file_petstore_v1_rpc_proto_depIdxs,
		EnumInfos:         file_petstore_v1_rpc_proto_enumTypes,
		MessageInfos:      file_petstore_v1_rpc_proto_msgTypes,
	}.Build()
	File_petstore_v1_rpc_proto = out.File
	file_petstore_v1_rpc_proto_goTypes = nil
	file_petstore_v1_rpc_proto_depIdxs = nil
}
