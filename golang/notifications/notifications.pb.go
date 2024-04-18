// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: notifications/notifications.proto

package notifications

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

type NotificationUserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float32 `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *NotificationUserLocation) Reset() {
	*x = NotificationUserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationUserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationUserLocation) ProtoMessage() {}

func (x *NotificationUserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationUserLocation.ProtoReflect.Descriptor instead.
func (*NotificationUserLocation) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationUserLocation) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *NotificationUserLocation) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type NotificationUserAddress struct {
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

func (x *NotificationUserAddress) Reset() {
	*x = NotificationUserAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationUserAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationUserAddress) ProtoMessage() {}

func (x *NotificationUserAddress) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationUserAddress.ProtoReflect.Descriptor instead.
func (*NotificationUserAddress) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationUserAddress) GetResidential() string {
	if x != nil {
		return x.Residential
	}
	return ""
}

func (x *NotificationUserAddress) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *NotificationUserAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *NotificationUserAddress) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *NotificationUserAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NotificationUserAddress) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type NotificationUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dob           string                    `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Sex           string                    `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Bio           string                    `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	LookingFor    string                    `protobuf:"bytes,6,opt,name=lookingFor,proto3" json:"lookingFor,omitempty"`
	Interests     []string                  `protobuf:"bytes,7,rep,name=interests,proto3" json:"interests,omitempty"`
	Location      *NotificationUserLocation `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	ProfileImages []string                  `protobuf:"bytes,9,rep,name=profileImages,proto3" json:"profileImages,omitempty"`
	TelegramId    int64                     `protobuf:"varint,10,opt,name=telegramId,proto3" json:"telegramId,omitempty"`
	Address       *NotificationUserAddress  `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NotificationUser) Reset() {
	*x = NotificationUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationUser) ProtoMessage() {}

func (x *NotificationUser) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationUser.ProtoReflect.Descriptor instead.
func (*NotificationUser) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationUser) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *NotificationUser) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *NotificationUser) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *NotificationUser) GetLookingFor() string {
	if x != nil {
		return x.LookingFor
	}
	return ""
}

func (x *NotificationUser) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *NotificationUser) GetLocation() *NotificationUserLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *NotificationUser) GetProfileImages() []string {
	if x != nil {
		return x.ProfileImages
	}
	return nil
}

func (x *NotificationUser) GetTelegramId() int64 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *NotificationUser) GetAddress() *NotificationUserAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type SendNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string            `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string            `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Title string            `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body  string            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	User  *NotificationUser `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SendNotificationRequest) Reset() {
	*x = SendNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationRequest) ProtoMessage() {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{3}
}

func (x *SendNotificationRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendNotificationRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendNotificationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotificationRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SendNotificationRequest) GetUser() *NotificationUser {
	if x != nil {
		return x.User
	}
	return nil
}

type SendNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *SendNotificationResponse) Reset() {
	*x = SendNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationResponse) ProtoMessage() {}

func (x *SendNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationResponse.ProtoReflect.Descriptor instead.
func (*SendNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{4}
}

func (x *SendNotificationResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendNotificationResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type ReceiveNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To   string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *ReceiveNotificationRequest) Reset() {
	*x = ReceiveNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveNotificationRequest) ProtoMessage() {}

func (x *ReceiveNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveNotificationRequest.ProtoReflect.Descriptor instead.
func (*ReceiveNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{5}
}

func (x *ReceiveNotificationRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ReceiveNotificationRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type ReceiveNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string            `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string            `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Title string            `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body  string            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	User  *NotificationUser `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ReceiveNotificationResponse) Reset() {
	*x = ReceiveNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveNotificationResponse) ProtoMessage() {}

func (x *ReceiveNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveNotificationResponse.ProtoReflect.Descriptor instead.
func (*ReceiveNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{6}
}

func (x *ReceiveNotificationResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ReceiveNotificationResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ReceiveNotificationResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReceiveNotificationResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ReceiveNotificationResponse) GetUser() *NotificationUser {
	if x != nil {
		return x.User
	}
	return nil
}

var File_notifications_notifications_proto protoreflect.FileDescriptor

var file_notifications_notifications_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x18, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x17, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0xdb, 0x02, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x3e, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x1a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x22, 0x92, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xf8, 0x01, 0x0a, 0x0d, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6e, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22,
	0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x77, 0x0a, 0x13, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6a, 0x75, 0x6e, 0x77, 0x61, 0x31, 0x2f, 0x6b, 0x6f, 0x6e, 0x65, 0x6b,
	0x73, 0x68, 0x65, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_notifications_proto_rawDescOnce sync.Once
	file_notifications_notifications_proto_rawDescData = file_notifications_notifications_proto_rawDesc
)

func file_notifications_notifications_proto_rawDescGZIP() []byte {
	file_notifications_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_notifications_proto_rawDescData)
	})
	return file_notifications_notifications_proto_rawDescData
}

var file_notifications_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_notifications_notifications_proto_goTypes = []interface{}{
	(*NotificationUserLocation)(nil),    // 0: NotificationUserLocation
	(*NotificationUserAddress)(nil),     // 1: NotificationUserAddress
	(*NotificationUser)(nil),            // 2: NotificationUser
	(*SendNotificationRequest)(nil),     // 3: SendNotificationRequest
	(*SendNotificationResponse)(nil),    // 4: SendNotificationResponse
	(*ReceiveNotificationRequest)(nil),  // 5: ReceiveNotificationRequest
	(*ReceiveNotificationResponse)(nil), // 6: ReceiveNotificationResponse
}
var file_notifications_notifications_proto_depIdxs = []int32{
	0, // 0: NotificationUser.location:type_name -> NotificationUserLocation
	1, // 1: NotificationUser.address:type_name -> NotificationUserAddress
	2, // 2: SendNotificationRequest.user:type_name -> NotificationUser
	2, // 3: ReceiveNotificationResponse.user:type_name -> NotificationUser
	3, // 4: Notifications.SendNotification:input_type -> SendNotificationRequest
	5, // 5: Notifications.ReceiveNotification:input_type -> ReceiveNotificationRequest
	4, // 6: Notifications.SendNotification:output_type -> SendNotificationResponse
	6, // 7: Notifications.ReceiveNotification:output_type -> ReceiveNotificationResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_notifications_notifications_proto_init() }
func file_notifications_notifications_proto_init() {
	if File_notifications_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationUserLocation); i {
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
		file_notifications_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationUserAddress); i {
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
		file_notifications_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationUser); i {
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
		file_notifications_notifications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationRequest); i {
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
		file_notifications_notifications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationResponse); i {
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
		file_notifications_notifications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveNotificationRequest); i {
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
		file_notifications_notifications_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveNotificationResponse); i {
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
			RawDescriptor: file_notifications_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifications_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_notifications_proto_depIdxs,
		MessageInfos:      file_notifications_notifications_proto_msgTypes,
	}.Build()
	File_notifications_notifications_proto = out.File
	file_notifications_notifications_proto_rawDesc = nil
	file_notifications_notifications_proto_goTypes = nil
	file_notifications_notifications_proto_depIdxs = nil
}