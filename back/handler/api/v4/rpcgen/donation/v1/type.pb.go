// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: donation/v1/type.proto

package donationv1

import (
	sharedpb "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/sharedpb"
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

type PaymentType int32

const (
	PaymentType_PAYMENT_TYPE_UNSPECIFIED  PaymentType = 0
	PaymentType_PAYMENT_TYPE_ONE_TIME     PaymentType = 1
	PaymentType_PAYMENT_TYPE_SUBSCRIPTION PaymentType = 2
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "PAYMENT_TYPE_UNSPECIFIED",
		1: "PAYMENT_TYPE_ONE_TIME",
		2: "PAYMENT_TYPE_SUBSCRIPTION",
	}
	PaymentType_value = map[string]int32{
		"PAYMENT_TYPE_UNSPECIFIED":  0,
		"PAYMENT_TYPE_ONE_TIME":     1,
		"PAYMENT_TYPE_SUBSCRIPTION": 2,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_donation_v1_type_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_donation_v1_type_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{0}
}

type PaymentStatus int32

const (
	PaymentStatus_PAYMENT_STATUS_UNSPECIFIED PaymentStatus = 0
	PaymentStatus_PAYMENT_STATUS_PENDING     PaymentStatus = 1
	PaymentStatus_PAYMENT_STATUS_CANCELED    PaymentStatus = 2
	PaymentStatus_PAYMENT_STATUS_SUCCEEDED   PaymentStatus = 3
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PAYMENT_STATUS_UNSPECIFIED",
		1: "PAYMENT_STATUS_PENDING",
		2: "PAYMENT_STATUS_CANCELED",
		3: "PAYMENT_STATUS_SUCCEEDED",
	}
	PaymentStatus_value = map[string]int32{
		"PAYMENT_STATUS_UNSPECIFIED": 0,
		"PAYMENT_STATUS_PENDING":     1,
		"PAYMENT_STATUS_CANCELED":    2,
		"PAYMENT_STATUS_SUCCEEDED":   3,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_donation_v1_type_proto_enumTypes[1].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_donation_v1_type_proto_enumTypes[1]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{1}
}

type PaymentUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      *sharedpb.UUID `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DisplayName *string        `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	Link        *string        `protobuf:"bytes,4,opt,name=link,proto3,oneof" json:"link,omitempty"`
}

func (x *PaymentUser) Reset() {
	*x = PaymentUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_donation_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentUser) ProtoMessage() {}

func (x *PaymentUser) ProtoReflect() protoreflect.Message {
	mi := &file_donation_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentUser.ProtoReflect.Descriptor instead.
func (*PaymentUser) Descriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentUser) GetUserId() *sharedpb.UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *PaymentUser) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *PaymentUser) GetLink() string {
	if x != nil && x.Link != nil {
		return *x.Link
	}
	return ""
}

type PaymentHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      PaymentType               `protobuf:"varint,2,opt,name=type,proto3,enum=donation.v1.PaymentType" json:"type,omitempty"`
	Status    PaymentStatus             `protobuf:"varint,3,opt,name=status,proto3,enum=donation.v1.PaymentStatus" json:"status,omitempty"`
	Amount    int32                     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt *sharedpb.RFC3339DateTime `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PaymentHistory) Reset() {
	*x = PaymentHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_donation_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentHistory) ProtoMessage() {}

func (x *PaymentHistory) ProtoReflect() protoreflect.Message {
	mi := &file_donation_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentHistory.ProtoReflect.Descriptor instead.
func (*PaymentHistory) Descriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentHistory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentHistory) GetType() PaymentType {
	if x != nil {
		return x.Type
	}
	return PaymentType_PAYMENT_TYPE_UNSPECIFIED
}

func (x *PaymentHistory) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
}

func (x *PaymentHistory) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentHistory) GetCreatedAt() *sharedpb.RFC3339DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SubscriptionPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SubscriptionPlan) Reset() {
	*x = SubscriptionPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_donation_v1_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionPlan) ProtoMessage() {}

func (x *SubscriptionPlan) ProtoReflect() protoreflect.Message {
	mi := &file_donation_v1_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionPlan.ProtoReflect.Descriptor instead.
func (*SubscriptionPlan) Descriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionPlan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscriptionPlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionPlan) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Plan      *SubscriptionPlan         `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	IsActive  bool                      `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt *sharedpb.RFC3339DateTime `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_donation_v1_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_donation_v1_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{3}
}

func (x *Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription) GetPlan() *SubscriptionPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *Subscription) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Subscription) GetCreatedAt() *sharedpb.RFC3339DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Contributor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string  `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Link        *string `protobuf:"bytes,2,opt,name=link,proto3,oneof" json:"link,omitempty"`
}

func (x *Contributor) Reset() {
	*x = Contributor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_donation_v1_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contributor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contributor) ProtoMessage() {}

func (x *Contributor) ProtoReflect() protoreflect.Message {
	mi := &file_donation_v1_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contributor.ProtoReflect.Descriptor instead.
func (*Contributor) Descriptor() ([]byte, []int) {
	return file_donation_v1_type_proto_rawDescGZIP(), []int{4}
}

func (x *Contributor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Contributor) GetLink() string {
	if x != nil && x.Link != nil {
		return *x.Link
	}
	return ""
}

var File_donation_v1_type_proto protoreflect.FileDescriptor

var file_donation_v1_type_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x6f,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x6f,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x4e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xa6, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x31, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46,
	0x43, 0x33, 0x33, 0x33, 0x39, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x52, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2a, 0x65, 0x0a, 0x0b,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50,
	0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x02, 0x2a, 0x86, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1c,
	0x0a, 0x18, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x03, 0x42, 0x4e, 0x5a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d,
	0x74, 0x65, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x2f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x34, 0x2f, 0x72,
	0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_donation_v1_type_proto_rawDescOnce sync.Once
	file_donation_v1_type_proto_rawDescData = file_donation_v1_type_proto_rawDesc
)

func file_donation_v1_type_proto_rawDescGZIP() []byte {
	file_donation_v1_type_proto_rawDescOnce.Do(func() {
		file_donation_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_donation_v1_type_proto_rawDescData)
	})
	return file_donation_v1_type_proto_rawDescData
}

var file_donation_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_donation_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_donation_v1_type_proto_goTypes = []any{
	(PaymentType)(0),                 // 0: donation.v1.PaymentType
	(PaymentStatus)(0),               // 1: donation.v1.PaymentStatus
	(*PaymentUser)(nil),              // 2: donation.v1.PaymentUser
	(*PaymentHistory)(nil),           // 3: donation.v1.PaymentHistory
	(*SubscriptionPlan)(nil),         // 4: donation.v1.SubscriptionPlan
	(*Subscription)(nil),             // 5: donation.v1.Subscription
	(*Contributor)(nil),              // 6: donation.v1.Contributor
	(*sharedpb.UUID)(nil),            // 7: shared.UUID
	(*sharedpb.RFC3339DateTime)(nil), // 8: shared.RFC3339DateTime
}
var file_donation_v1_type_proto_depIdxs = []int32{
	7, // 0: donation.v1.PaymentUser.user_id:type_name -> shared.UUID
	0, // 1: donation.v1.PaymentHistory.type:type_name -> donation.v1.PaymentType
	1, // 2: donation.v1.PaymentHistory.status:type_name -> donation.v1.PaymentStatus
	8, // 3: donation.v1.PaymentHistory.created_at:type_name -> shared.RFC3339DateTime
	4, // 4: donation.v1.Subscription.plan:type_name -> donation.v1.SubscriptionPlan
	8, // 5: donation.v1.Subscription.created_at:type_name -> shared.RFC3339DateTime
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_donation_v1_type_proto_init() }
func file_donation_v1_type_proto_init() {
	if File_donation_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_donation_v1_type_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentUser); i {
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
		file_donation_v1_type_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentHistory); i {
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
		file_donation_v1_type_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SubscriptionPlan); i {
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
		file_donation_v1_type_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Subscription); i {
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
		file_donation_v1_type_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Contributor); i {
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
	file_donation_v1_type_proto_msgTypes[0].OneofWrappers = []any{}
	file_donation_v1_type_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_donation_v1_type_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_donation_v1_type_proto_goTypes,
		DependencyIndexes: file_donation_v1_type_proto_depIdxs,
		EnumInfos:         file_donation_v1_type_proto_enumTypes,
		MessageInfos:      file_donation_v1_type_proto_msgTypes,
	}.Build()
	File_donation_v1_type_proto = out.File
	file_donation_v1_type_proto_rawDesc = nil
	file_donation_v1_type_proto_goTypes = nil
	file_donation_v1_type_proto_depIdxs = nil
}
