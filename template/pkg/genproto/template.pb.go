// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: template.proto

package template

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

type TemplateType int32

const (
	TemplateType_SMS   TemplateType = 0
	TemplateType_EMAIL TemplateType = 1
)

// Enum value maps for TemplateType.
var (
	TemplateType_name = map[int32]string{
		0: "SMS",
		1: "EMAIL",
	}
	TemplateType_value = map[string]int32{
		"SMS":   0,
		"EMAIL": 1,
	}
)

func (x TemplateType) Enum() *TemplateType {
	p := new(TemplateType)
	*p = x
	return p
}

func (x TemplateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TemplateType) Descriptor() protoreflect.EnumDescriptor {
	return file_template_proto_enumTypes[0].Descriptor()
}

func (TemplateType) Type() protoreflect.EnumType {
	return &file_template_proto_enumTypes[0]
}

func (x TemplateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TemplateType.Descriptor instead.
func (TemplateType) EnumDescriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template string       `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Type     TemplateType `protobuf:"varint,2,opt,name=type,proto3,enum=template.TemplateType" json:"type,omitempty"`
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTemplateRequest) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *CreateTemplateRequest) GetType() TemplateType {
	if x != nil {
		return x.Type
	}
	return TemplateType_SMS
}

type CreateTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTemplateResponse) Reset() {
	*x = CreateTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateResponse) ProtoMessage() {}

func (x *CreateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateResponse.ProtoReflect.Descriptor instead.
func (*CreateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTemplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTemplateByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTemplateByIDRequest) Reset() {
	*x = GetTemplateByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateByIDRequest) ProtoMessage() {}

func (x *GetTemplateByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateByIDRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateByIDRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{2}
}

func (x *GetTemplateByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTemplateByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTemplateByIDResponse) Reset() {
	*x = GetTemplateByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateByIDResponse) ProtoMessage() {}

func (x *GetTemplateByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateByIDResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateByIDResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetTemplateByIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteTemplateByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTemplateByIDRequest) Reset() {
	*x = DeleteTemplateByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateByIDRequest) ProtoMessage() {}

func (x *DeleteTemplateByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateByIDRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTemplateByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTemplateByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTemplateByIDResponse) Reset() {
	*x = DeleteTemplateByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateByIDResponse) ProtoMessage() {}

func (x *DeleteTemplateByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteTemplateByIDResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTemplateByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_template_proto protoreflect.FileDescriptor

var file_template_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x5f, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x53,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0x9e, 0x02,
	0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x23, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a,
	0x5a, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x63, 0x2f, 0x67, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_template_proto_rawDescOnce sync.Once
	file_template_proto_rawDescData = file_template_proto_rawDesc
)

func file_template_proto_rawDescGZIP() []byte {
	file_template_proto_rawDescOnce.Do(func() {
		file_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_proto_rawDescData)
	})
	return file_template_proto_rawDescData
}

var file_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_template_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_template_proto_goTypes = []interface{}{
	(TemplateType)(0),                  // 0: template.TemplateType
	(*CreateTemplateRequest)(nil),      // 1: template.CreateTemplateRequest
	(*CreateTemplateResponse)(nil),     // 2: template.CreateTemplateResponse
	(*GetTemplateByIDRequest)(nil),     // 3: template.GetTemplateByIDRequest
	(*GetTemplateByIDResponse)(nil),    // 4: template.GetTemplateByIDResponse
	(*DeleteTemplateByIDRequest)(nil),  // 5: template.DeleteTemplateByIDRequest
	(*DeleteTemplateByIDResponse)(nil), // 6: template.DeleteTemplateByIDResponse
}
var file_template_proto_depIdxs = []int32{
	0, // 0: template.CreateTemplateRequest.type:type_name -> template.TemplateType
	1, // 1: template.Template.CreateTemplate:input_type -> template.CreateTemplateRequest
	3, // 2: template.Template.GetTemplateByID:input_type -> template.GetTemplateByIDRequest
	5, // 3: template.Template.DeleteTemplateByID:input_type -> template.DeleteTemplateByIDRequest
	2, // 4: template.Template.CreateTemplate:output_type -> template.CreateTemplateResponse
	4, // 5: template.Template.GetTemplateByID:output_type -> template.GetTemplateByIDResponse
	6, // 6: template.Template.DeleteTemplateByID:output_type -> template.DeleteTemplateByIDResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_template_proto_init() }
func file_template_proto_init() {
	if File_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateResponse); i {
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
		file_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateByIDRequest); i {
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
		file_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateByIDResponse); i {
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
		file_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateByIDRequest); i {
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
		file_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateByIDResponse); i {
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
			RawDescriptor: file_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_proto_goTypes,
		DependencyIndexes: file_template_proto_depIdxs,
		EnumInfos:         file_template_proto_enumTypes,
		MessageInfos:      file_template_proto_msgTypes,
	}.Build()
	File_template_proto = out.File
	file_template_proto_rawDesc = nil
	file_template_proto_goTypes = nil
	file_template_proto_depIdxs = nil
}
