// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: blog.proto

package blog

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Permissions int32

const (
	Permissions_NONE           Permissions = 0
	Permissions_CREATE_ARTICLE Permissions = 1
	Permissions_READ_ARTICLE   Permissions = 2
	Permissions_UPDATE_ARTICLE Permissions = 3
	Permissions_DELETE_ARTICLE Permissions = 4
)

// Enum value maps for Permissions.
var (
	Permissions_name = map[int32]string{
		0: "NONE",
		1: "CREATE_ARTICLE",
		2: "READ_ARTICLE",
		3: "UPDATE_ARTICLE",
		4: "DELETE_ARTICLE",
	}
	Permissions_value = map[string]int32{
		"NONE":           0,
		"CREATE_ARTICLE": 1,
		"READ_ARTICLE":   2,
		"UPDATE_ARTICLE": 3,
		"DELETE_ARTICLE": 4,
	}
)

func (x Permissions) Enum() *Permissions {
	p := new(Permissions)
	*p = x
	return p
}

func (x Permissions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permissions) Descriptor() protoreflect.EnumDescriptor {
	return file_blog_proto_enumTypes[0].Descriptor()
}

func (Permissions) Type() protoreflect.EnumType {
	return &file_blog_proto_enumTypes[0]
}

func (x Permissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permissions.Descriptor instead.
func (Permissions) EnumDescriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions         []Permissions `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=services.blog.v1.Permissions" json:"permissions,omitempty"`
	Optional            bool          `protobuf:"varint,2,opt,name=optional,proto3" json:"optional,omitempty"`
	ValidatePermissions bool          `protobuf:"varint,3,opt,name=validate_permissions,json=validatePermissions,proto3" json:"validate_permissions,omitempty"`
	Captcha             bool          `protobuf:"varint,4,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetPermissions() []Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Permission) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *Permission) GetValidatePermissions() bool {
	if x != nil {
		return x.ValidatePermissions
	}
	return false
}

func (x *Permission) GetCaptcha() bool {
	if x != nil {
		return x.Captcha
	}
	return false
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required
	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Article) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// require
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var file_blog_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Permission)(nil),
		Field:         1335954,
		Name:          "services.blog.v1.permission",
		Tag:           "bytes,1335954,opt,name=permission",
		Filename:      "blog.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional services.blog.v1.Permission permission = 1335954;
	E_Permission = &file_blog_proto_extTypes[0]
)

var File_blog_proto protoreflect.FileDescriptor

var file_blog_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12,
	0x31, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x52, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x05, 0x18, 0xf4, 0x03, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0xbf, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x65, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x41, 0x52, 0x54,
	0x49, 0x43, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x10, 0x04, 0x32, 0xb5,
	0x02, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x22, 0x3e, 0x92, 0x41, 0x1f, 0x0a, 0x07, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c,
	0x45, 0x12, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x92, 0xa9, 0x8c, 0x05, 0x03, 0x0a, 0x01, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x22, 0x42, 0x92, 0x41, 0x1f, 0x0a, 0x07, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45,
	0x12, 0x14, 0x67, 0x65, 0x74, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x92, 0xa9, 0x8c, 0x05, 0x05, 0x0a, 0x01, 0x00, 0x10, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x5e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0xc5, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xec, 0x01, 0x92, 0x41, 0xc6, 0x01, 0x12, 0x42, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x67, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2d, 0x0a,
	0x06, 0x61, 0x6d, 0x6e, 0x72, 0x61, 0x68, 0x12, 0x12, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x61, 0x6d, 0x6e, 0x72, 0x61, 0x68, 0x2e, 0x78, 0x79, 0x7a, 0x1a, 0x0f, 0x74, 0x65, 0x73,
	0x74, 0x40, 0x61, 0x6d, 0x6e, 0x72, 0x61, 0x68, 0x2e, 0x78, 0x79, 0x7a, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4a, 0x0a, 0x48, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x3e, 0x08, 0x02, 0x12, 0x29, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x00, 0x5a, 0x20, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x3b, 0x62, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_proto_rawDescOnce sync.Once
	file_blog_proto_rawDescData = file_blog_proto_rawDesc
)

func file_blog_proto_rawDescGZIP() []byte {
	file_blog_proto_rawDescOnce.Do(func() {
		file_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_proto_rawDescData)
	})
	return file_blog_proto_rawDescData
}

var file_blog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blog_proto_goTypes = []interface{}{
	(Permissions)(0),                   // 0: services.blog.v1.Permissions
	(*Permission)(nil),                 // 1: services.blog.v1.Permission
	(*CreateArticleRequest)(nil),       // 2: services.blog.v1.CreateArticleRequest
	(*Article)(nil),                    // 3: services.blog.v1.Article
	(*GetArticleRequest)(nil),          // 4: services.blog.v1.GetArticleRequest
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	(*descriptorpb.MethodOptions)(nil), // 6: google.protobuf.MethodOptions
}
var file_blog_proto_depIdxs = []int32{
	0, // 0: services.blog.v1.Permission.permissions:type_name -> services.blog.v1.Permissions
	5, // 1: services.blog.v1.Article.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: services.blog.v1.Article.updated_at:type_name -> google.protobuf.Timestamp
	6, // 3: services.blog.v1.permission:extendee -> google.protobuf.MethodOptions
	1, // 4: services.blog.v1.permission:type_name -> services.blog.v1.Permission
	2, // 5: services.blog.v1.BlogService.CreateArticle:input_type -> services.blog.v1.CreateArticleRequest
	4, // 6: services.blog.v1.BlogService.GetArticle:input_type -> services.blog.v1.GetArticleRequest
	3, // 7: services.blog.v1.BlogService.CreateArticle:output_type -> services.blog.v1.Article
	3, // 8: services.blog.v1.BlogService.GetArticle:output_type -> services.blog.v1.Article
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blog_proto_init() }
func file_blog_proto_init() {
	if File_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
			RawDescriptor: file_blog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_blog_proto_goTypes,
		DependencyIndexes: file_blog_proto_depIdxs,
		EnumInfos:         file_blog_proto_enumTypes,
		MessageInfos:      file_blog_proto_msgTypes,
		ExtensionInfos:    file_blog_proto_extTypes,
	}.Build()
	File_blog_proto = out.File
	file_blog_proto_rawDesc = nil
	file_blog_proto_goTypes = nil
	file_blog_proto_depIdxs = nil
}
