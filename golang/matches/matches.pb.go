// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: matches/matches.proto

package matches

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MatchedUserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float32 `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *MatchedUserLocation) Reset() {
	*x = MatchedUserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedUserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedUserLocation) ProtoMessage() {}

func (x *MatchedUserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedUserLocation.ProtoReflect.Descriptor instead.
func (*MatchedUserLocation) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{0}
}

func (x *MatchedUserLocation) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *MatchedUserLocation) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type MatchedUserAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Residential string `protobuf:"bytes,1,opt,name=residential,proto3" json:"residential,omitempty"`
	Town        string `protobuf:"bytes,2,opt,name=town,proto3" json:"town,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Region      string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Country     string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	CountryCode string `protobuf:"bytes,6,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
}

func (x *MatchedUserAddress) Reset() {
	*x = MatchedUserAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedUserAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedUserAddress) ProtoMessage() {}

func (x *MatchedUserAddress) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedUserAddress.ProtoReflect.Descriptor instead.
func (*MatchedUserAddress) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{1}
}

func (x *MatchedUserAddress) GetResidential() string {
	if x != nil {
		return x.Residential
	}
	return ""
}

func (x *MatchedUserAddress) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *MatchedUserAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *MatchedUserAddress) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *MatchedUserAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MatchedUserAddress) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type MatchedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dob           string               `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Sex           string               `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Bio           string               `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	LookingFor    string               `protobuf:"bytes,6,opt,name=lookingFor,proto3" json:"lookingFor,omitempty"`
	Interests     []string             `protobuf:"bytes,7,rep,name=interests,proto3" json:"interests,omitempty"`
	Location      *MatchedUserLocation `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	ProfileImages []string             `protobuf:"bytes,9,rep,name=profileImages,proto3" json:"profileImages,omitempty"`
	TelegramId    int64                `protobuf:"varint,10,opt,name=telegramId,proto3" json:"telegramId,omitempty"`
	Address       *MatchedUserAddress  `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MatchedUser) Reset() {
	*x = MatchedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedUser) ProtoMessage() {}

func (x *MatchedUser) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedUser.ProtoReflect.Descriptor instead.
func (*MatchedUser) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{2}
}

func (x *MatchedUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MatchedUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MatchedUser) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *MatchedUser) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *MatchedUser) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *MatchedUser) GetLookingFor() string {
	if x != nil {
		return x.LookingFor
	}
	return ""
}

func (x *MatchedUser) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *MatchedUser) GetLocation() *MatchedUserLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *MatchedUser) GetProfileImages() []string {
	if x != nil {
		return x.ProfileImages
	}
	return nil
}

func (x *MatchedUser) GetTelegramId() int64 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *MatchedUser) GetAddress() *MatchedUserAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type CreateMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []string `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateMatchRequest) Reset() {
	*x = CreateMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchRequest) ProtoMessage() {}

func (x *CreateMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchRequest.ProtoReflect.Descriptor instead.
func (*CreateMatchRequest) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMatchRequest) GetUser() []string {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User []string `protobuf:"bytes,2,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateMatchResponse) Reset() {
	*x = CreateMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchResponse) ProtoMessage() {}

func (x *CreateMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchResponse.ProtoReflect.Descriptor instead.
func (*CreateMatchResponse) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMatchResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateMatchResponse) GetUser() []string {
	if x != nil {
		return x.User
	}
	return nil
}

type GetMatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMatchesRequest) Reset() {
	*x = GetMatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchesRequest) ProtoMessage() {}

func (x *GetMatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchesRequest.ProtoReflect.Descriptor instead.
func (*GetMatchesRequest) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{5}
}

func (x *GetMatchesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetMatchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*MatchedUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetMatchesResponse) Reset() {
	*x = GetMatchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchesResponse) ProtoMessage() {}

func (x *GetMatchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchesResponse.ProtoReflect.Descriptor instead.
func (*GetMatchesResponse) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{6}
}

func (x *GetMatchesResponse) GetUsers() []*MatchedUser {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_matches_matches_proto protoreflect.FileDescriptor

var file_matches_matches_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x13, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x6f, 0x77, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xcc, 0x02,
	0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x6f, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x46, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xaf, 0x01, 0x0a, 0x07, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6a, 0x75, 0x6e, 0x77, 0x61,
	0x31, 0x2f, 0x6b, 0x6f, 0x6e, 0x65, 0x6b, 0x73, 0x68, 0x65, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matches_matches_proto_rawDescOnce sync.Once
	file_matches_matches_proto_rawDescData = file_matches_matches_proto_rawDesc
)

func file_matches_matches_proto_rawDescGZIP() []byte {
	file_matches_matches_proto_rawDescOnce.Do(func() {
		file_matches_matches_proto_rawDescData = protoimpl.X.CompressGZIP(file_matches_matches_proto_rawDescData)
	})
	return file_matches_matches_proto_rawDescData
}

var file_matches_matches_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_matches_matches_proto_goTypes = []interface{}{
	(*MatchedUserLocation)(nil), // 0: MatchedUserLocation
	(*MatchedUserAddress)(nil),  // 1: MatchedUserAddress
	(*MatchedUser)(nil),         // 2: MatchedUser
	(*CreateMatchRequest)(nil),  // 3: CreateMatchRequest
	(*CreateMatchResponse)(nil), // 4: CreateMatchResponse
	(*GetMatchesRequest)(nil),   // 5: GetMatchesRequest
	(*GetMatchesResponse)(nil),  // 6: GetMatchesResponse
}
var file_matches_matches_proto_depIdxs = []int32{
	0, // 0: MatchedUser.location:type_name -> MatchedUserLocation
	1, // 1: MatchedUser.address:type_name -> MatchedUserAddress
	2, // 2: GetMatchesResponse.users:type_name -> MatchedUser
	3, // 3: Matcher.CreateMatch:input_type -> CreateMatchRequest
	5, // 4: Matcher.GetMatches:input_type -> GetMatchesRequest
	4, // 5: Matcher.CreateMatch:output_type -> CreateMatchResponse
	6, // 6: Matcher.GetMatches:output_type -> GetMatchesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_matches_matches_proto_init() }
func file_matches_matches_proto_init() {
	if File_matches_matches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matches_matches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedUserLocation); i {
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
		file_matches_matches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedUserAddress); i {
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
		file_matches_matches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedUser); i {
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
		file_matches_matches_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchRequest); i {
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
		file_matches_matches_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchResponse); i {
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
		file_matches_matches_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchesRequest); i {
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
		file_matches_matches_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchesResponse); i {
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
			RawDescriptor: file_matches_matches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matches_matches_proto_goTypes,
		DependencyIndexes: file_matches_matches_proto_depIdxs,
		MessageInfos:      file_matches_matches_proto_msgTypes,
	}.Build()
	File_matches_matches_proto = out.File
	file_matches_matches_proto_rawDesc = nil
	file_matches_matches_proto_goTypes = nil
	file_matches_matches_proto_depIdxs = nil
}
