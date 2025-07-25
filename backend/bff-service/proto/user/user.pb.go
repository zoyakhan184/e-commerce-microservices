// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/user.proto

package userpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_proto_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_proto_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Dob           string                 `protobuf:"bytes,6,opt,name=dob,proto3" json:"dob,omitempty"`
	AvatarUrl     string                 `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Role          string                 `protobuf:"bytes,8,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // ✅ Add this
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_proto_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProfile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserProfile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserProfile) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *UserProfile) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserProfile) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Dob           string                 `protobuf:"bytes,5,opt,name=dob,proto3" json:"dob,omitempty"`
	AvatarUrl     string                 `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_proto_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *UpdateUserRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type GenericResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	mi := &file_proto_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressLine   string                 `protobuf:"bytes,5,opt,name=address_line,json=addressLine,proto3" json:"address_line,omitempty"`
	City          string                 `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Zip           string                 `protobuf:"bytes,8,opt,name=zip,proto3" json:"zip,omitempty"`
	Country       string                 `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	IsDefault     bool                   `protobuf:"varint,10,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	mi := &file_proto_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *AddressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddressRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddressRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddressRequest) GetAddressLine() string {
	if x != nil {
		return x.AddressLine
	}
	return ""
}

func (x *AddressRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AddressRequest) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *AddressRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AddressRequest) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

type AddressList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []*AddressRequest      `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressList) Reset() {
	*x = AddressList{}
	mi := &file_proto_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressList) ProtoMessage() {}

func (x *AddressList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressList.ProtoReflect.Descriptor instead.
func (*AddressList) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *AddressList) GetAddresses() []*AddressRequest {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type WishlistRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WishlistRequest) Reset() {
	*x = WishlistRequest{}
	mi := &file_proto_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistRequest) ProtoMessage() {}

func (x *WishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistRequest.ProtoReflect.Descriptor instead.
func (*WishlistRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{7}
}

func (x *WishlistRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WishlistRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type WishlistResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductIds    []string               `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WishlistResponse) Reset() {
	*x = WishlistResponse{}
	mi := &file_proto_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistResponse) ProtoMessage() {}

func (x *WishlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistResponse.ProtoReflect.Descriptor instead.
func (*WishlistResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{8}
}

func (x *WishlistResponse) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type ListProfilesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Profiles      []*UserProfile         `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProfilesResponse) Reset() {
	*x = ListProfilesResponse{}
	mi := &file_proto_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesResponse) ProtoMessage() {}

func (x *ListProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesResponse.ProtoReflect.Descriptor instead.
func (*ListProfilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{9}
}

func (x *ListProfilesResponse) GetProfiles() []*UserProfile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type AddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressLine   string                 `protobuf:"bytes,5,opt,name=address_line,json=addressLine,proto3" json:"address_line,omitempty"`
	City          string                 `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Zip           string                 `protobuf:"bytes,8,opt,name=zip,proto3" json:"zip,omitempty"`
	Country       string                 `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	IsDefault     bool                   `protobuf:"varint,10,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressResponse) Reset() {
	*x = AddressResponse{}
	mi := &file_proto_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResponse) ProtoMessage() {}

func (x *AddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResponse.ProtoReflect.Descriptor instead.
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{10}
}

func (x *AddressResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddressResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddressResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddressResponse) GetAddressLine() string {
	if x != nil {
		return x.AddressLine
	}
	return ""
}

func (x *AddressResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AddressResponse) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *AddressResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AddressResponse) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{11}
}

var File_proto_user_proto protoreflect.FileDescriptor

const file_proto_user_proto_rawDesc = "" +
	"\n" +
	"\x10proto/user.proto\x12\x04user\")\n" +
	"\x0eGetUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"&\n" +
	"\vUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"\xeb\x01\n" +
	"\vUserProfile\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x12\x10\n" +
	"\x03dob\x18\x06 \x01(\tR\x03dob\x12\x1d\n" +
	"\n" +
	"avatar_url\x18\a \x01(\tR\tavatarUrl\x12\x12\n" +
	"\x04role\x18\b \x01(\tR\x04role\x12\x1d\n" +
	"\n" +
	"created_at\x18\t \x01(\tR\tcreatedAt\"\xa8\x01\n" +
	"\x11UpdateUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\x12\x10\n" +
	"\x03dob\x18\x05 \x01(\tR\x03dob\x12\x1d\n" +
	"\n" +
	"avatar_url\x18\x06 \x01(\tR\tavatarUrl\"+\n" +
	"\x0fGenericResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\xfb\x01\n" +
	"\x0eAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12!\n" +
	"\faddress_line\x18\x05 \x01(\tR\vaddressLine\x12\x12\n" +
	"\x04city\x18\x06 \x01(\tR\x04city\x12\x14\n" +
	"\x05state\x18\a \x01(\tR\x05state\x12\x10\n" +
	"\x03zip\x18\b \x01(\tR\x03zip\x12\x18\n" +
	"\acountry\x18\t \x01(\tR\acountry\x12\x1d\n" +
	"\n" +
	"is_default\x18\n" +
	" \x01(\bR\tisDefault\"A\n" +
	"\vAddressList\x122\n" +
	"\taddresses\x18\x01 \x03(\v2\x14.user.AddressRequestR\taddresses\"I\n" +
	"\x0fWishlistRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x02 \x01(\tR\tproductId\"3\n" +
	"\x10WishlistResponse\x12\x1f\n" +
	"\vproduct_ids\x18\x01 \x03(\tR\n" +
	"productIds\"E\n" +
	"\x14ListProfilesResponse\x12-\n" +
	"\bprofiles\x18\x01 \x03(\v2\x11.user.UserProfileR\bprofiles\"\xfc\x01\n" +
	"\x0fAddressResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12!\n" +
	"\faddress_line\x18\x05 \x01(\tR\vaddressLine\x12\x12\n" +
	"\x04city\x18\x06 \x01(\tR\x04city\x12\x14\n" +
	"\x05state\x18\a \x01(\tR\x05state\x12\x10\n" +
	"\x03zip\x18\b \x01(\tR\x03zip\x12\x18\n" +
	"\acountry\x18\t \x01(\tR\acountry\x12\x1d\n" +
	"\n" +
	"is_default\x18\n" +
	" \x01(\bR\tisDefault\"\a\n" +
	"\x05Empty2\xd9\x04\n" +
	"\vUserService\x126\n" +
	"\n" +
	"CreateUser\x12\x11.user.UserProfile\x1a\x15.user.GenericResponse\x122\n" +
	"\aGetUser\x12\x14.user.GetUserRequest\x1a\x11.user.UserProfile\x126\n" +
	"\n" +
	"UpdateUser\x12\x11.user.UserProfile\x1a\x15.user.GenericResponse\x129\n" +
	"\n" +
	"AddAddress\x12\x14.user.AddressRequest\x1a\x15.user.AddressResponse\x12<\n" +
	"\rUpdateAddress\x12\x14.user.AddressRequest\x1a\x15.user.GenericResponse\x124\n" +
	"\fGetAddresses\x12\x11.user.UserRequest\x1a\x11.user.AddressList\x12=\n" +
	"\rAddToWishlist\x12\x15.user.WishlistRequest\x1a\x15.user.GenericResponse\x12B\n" +
	"\x12RemoveFromWishlist\x12\x15.user.WishlistRequest\x1a\x15.user.GenericResponse\x128\n" +
	"\vGetWishlist\x12\x11.user.UserRequest\x1a\x16.user.WishlistResponse\x12:\n" +
	"\x0fListAllProfiles\x12\v.user.Empty\x1a\x1a.user.ListProfilesResponseB\x1bZ\x19user-service/proto;userpbb\x06proto3"

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData []byte
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)))
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_user_proto_goTypes = []any{
	(*GetUserRequest)(nil),       // 0: user.GetUserRequest
	(*UserRequest)(nil),          // 1: user.UserRequest
	(*UserProfile)(nil),          // 2: user.UserProfile
	(*UpdateUserRequest)(nil),    // 3: user.UpdateUserRequest
	(*GenericResponse)(nil),      // 4: user.GenericResponse
	(*AddressRequest)(nil),       // 5: user.AddressRequest
	(*AddressList)(nil),          // 6: user.AddressList
	(*WishlistRequest)(nil),      // 7: user.WishlistRequest
	(*WishlistResponse)(nil),     // 8: user.WishlistResponse
	(*ListProfilesResponse)(nil), // 9: user.ListProfilesResponse
	(*AddressResponse)(nil),      // 10: user.AddressResponse
	(*Empty)(nil),                // 11: user.Empty
}
var file_proto_user_proto_depIdxs = []int32{
	5,  // 0: user.AddressList.addresses:type_name -> user.AddressRequest
	2,  // 1: user.ListProfilesResponse.profiles:type_name -> user.UserProfile
	2,  // 2: user.UserService.CreateUser:input_type -> user.UserProfile
	0,  // 3: user.UserService.GetUser:input_type -> user.GetUserRequest
	2,  // 4: user.UserService.UpdateUser:input_type -> user.UserProfile
	5,  // 5: user.UserService.AddAddress:input_type -> user.AddressRequest
	5,  // 6: user.UserService.UpdateAddress:input_type -> user.AddressRequest
	1,  // 7: user.UserService.GetAddresses:input_type -> user.UserRequest
	7,  // 8: user.UserService.AddToWishlist:input_type -> user.WishlistRequest
	7,  // 9: user.UserService.RemoveFromWishlist:input_type -> user.WishlistRequest
	1,  // 10: user.UserService.GetWishlist:input_type -> user.UserRequest
	11, // 11: user.UserService.ListAllProfiles:input_type -> user.Empty
	4,  // 12: user.UserService.CreateUser:output_type -> user.GenericResponse
	2,  // 13: user.UserService.GetUser:output_type -> user.UserProfile
	4,  // 14: user.UserService.UpdateUser:output_type -> user.GenericResponse
	10, // 15: user.UserService.AddAddress:output_type -> user.AddressResponse
	4,  // 16: user.UserService.UpdateAddress:output_type -> user.GenericResponse
	6,  // 17: user.UserService.GetAddresses:output_type -> user.AddressList
	4,  // 18: user.UserService.AddToWishlist:output_type -> user.GenericResponse
	4,  // 19: user.UserService.RemoveFromWishlist:output_type -> user.GenericResponse
	8,  // 20: user.UserService.GetWishlist:output_type -> user.WishlistResponse
	9,  // 21: user.UserService.ListAllProfiles:output_type -> user.ListProfilesResponse
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
