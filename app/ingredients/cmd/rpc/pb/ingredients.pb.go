// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ingredients.proto

package __

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

type Ingredient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl    string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Ingredient) Reset() {
	*x = Ingredient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingredient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredient) ProtoMessage() {}

func (x *Ingredient) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredient.ProtoReflect.Descriptor instead.
func (*Ingredient) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{0}
}

func (x *Ingredient) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ingredient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ingredient) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Ingredient) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IngredientDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IngredientDetailReq) Reset() {
	*x = IngredientDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientDetailReq) ProtoMessage() {}

func (x *IngredientDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientDetailReq.ProtoReflect.Descriptor instead.
func (*IngredientDetailReq) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{1}
}

func (x *IngredientDetailReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IngredientDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingredient *Ingredient `protobuf:"bytes,1,opt,name=ingredient,proto3" json:"ingredient,omitempty"`
}

func (x *IngredientDetailResp) Reset() {
	*x = IngredientDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientDetailResp) ProtoMessage() {}

func (x *IngredientDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientDetailResp.ProtoReflect.Descriptor instead.
func (*IngredientDetailResp) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{2}
}

func (x *IngredientDetailResp) GetIngredient() *Ingredient {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

type IngredientListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *IngredientListReq) Reset() {
	*x = IngredientListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientListReq) ProtoMessage() {}

func (x *IngredientListReq) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientListReq.ProtoReflect.Descriptor instead.
func (*IngredientListReq) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{3}
}

func (x *IngredientListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IngredientListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type IngredientListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Ingredient `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *IngredientListResp) Reset() {
	*x = IngredientListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientListResp) ProtoMessage() {}

func (x *IngredientListResp) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientListResp.ProtoReflect.Descriptor instead.
func (*IngredientListResp) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{4}
}

func (x *IngredientListResp) GetList() []*Ingredient {
	if x != nil {
		return x.List
	}
	return nil
}

type IngredientCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageContent string `protobuf:"bytes,2,opt,name=image_content,json=imageContent,proto3" json:"image_content,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *IngredientCreateReq) Reset() {
	*x = IngredientCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientCreateReq) ProtoMessage() {}

func (x *IngredientCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientCreateReq.ProtoReflect.Descriptor instead.
func (*IngredientCreateReq) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{5}
}

func (x *IngredientCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IngredientCreateReq) GetImageContent() string {
	if x != nil {
		return x.ImageContent
	}
	return ""
}

func (x *IngredientCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IngredientCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IngredientCreateResp) Reset() {
	*x = IngredientCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingredients_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientCreateResp) ProtoMessage() {}

func (x *IngredientCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_ingredients_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientCreateResp.ProtoReflect.Descriptor instead.
func (*IngredientCreateResp) Descriptor() ([]byte, []int) {
	return file_ingredients_proto_rawDescGZIP(), []int{6}
}

var File_ingredients_proto protoreflect.FileDescriptor

var file_ingredients_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x6e, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x14, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x38, 0x0a, 0x12, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x13, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x32, 0xdb, 0x01, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ingredients_proto_rawDescOnce sync.Once
	file_ingredients_proto_rawDescData = file_ingredients_proto_rawDesc
)

func file_ingredients_proto_rawDescGZIP() []byte {
	file_ingredients_proto_rawDescOnce.Do(func() {
		file_ingredients_proto_rawDescData = protoimpl.X.CompressGZIP(file_ingredients_proto_rawDescData)
	})
	return file_ingredients_proto_rawDescData
}

var file_ingredients_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ingredients_proto_goTypes = []interface{}{
	(*Ingredient)(nil),           // 0: pb.Ingredient
	(*IngredientDetailReq)(nil),  // 1: pb.IngredientDetailReq
	(*IngredientDetailResp)(nil), // 2: pb.IngredientDetailResp
	(*IngredientListReq)(nil),    // 3: pb.IngredientListReq
	(*IngredientListResp)(nil),   // 4: pb.IngredientListResp
	(*IngredientCreateReq)(nil),  // 5: pb.IngredientCreateReq
	(*IngredientCreateResp)(nil), // 6: pb.IngredientCreateResp
}
var file_ingredients_proto_depIdxs = []int32{
	0, // 0: pb.IngredientDetailResp.ingredient:type_name -> pb.Ingredient
	0, // 1: pb.IngredientListResp.list:type_name -> pb.Ingredient
	1, // 2: pb.ingredient.ingredientDetail:input_type -> pb.IngredientDetailReq
	3, // 3: pb.ingredient.ingredientList:input_type -> pb.IngredientListReq
	5, // 4: pb.ingredient.ingredientCreate:input_type -> pb.IngredientCreateReq
	2, // 5: pb.ingredient.ingredientDetail:output_type -> pb.IngredientDetailResp
	4, // 6: pb.ingredient.ingredientList:output_type -> pb.IngredientListResp
	6, // 7: pb.ingredient.ingredientCreate:output_type -> pb.IngredientCreateResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ingredients_proto_init() }
func file_ingredients_proto_init() {
	if File_ingredients_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ingredients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingredient); i {
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
		file_ingredients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientDetailReq); i {
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
		file_ingredients_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientDetailResp); i {
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
		file_ingredients_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientListReq); i {
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
		file_ingredients_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientListResp); i {
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
		file_ingredients_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientCreateReq); i {
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
		file_ingredients_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientCreateResp); i {
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
			RawDescriptor: file_ingredients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ingredients_proto_goTypes,
		DependencyIndexes: file_ingredients_proto_depIdxs,
		MessageInfos:      file_ingredients_proto_msgTypes,
	}.Build()
	File_ingredients_proto = out.File
	file_ingredients_proto_rawDesc = nil
	file_ingredients_proto_goTypes = nil
	file_ingredients_proto_depIdxs = nil
}
