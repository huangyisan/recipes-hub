// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: recipes.proto

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

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageContent   string `protobuf:"bytes,3,opt,name=imageContent,proto3" json:"imageContent,omitempty"`
	Instructions   string `protobuf:"bytes,4,opt,name=instructions,proto3" json:"instructions,omitempty"`
	CookingTime    int64  `protobuf:"varint,5,opt,name=cookingTime,proto3" json:"cookingTime,omitempty"`
	Difficulty     string `protobuf:"bytes,6,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	RecipeType     string `protobuf:"bytes,7,opt,name=recipeType,proto3" json:"recipeType,omitempty"`
	CreationDate   string `protobuf:"bytes,8,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastUpdated    string `protobuf:"bytes,9,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	LastCookedDate string `protobuf:"bytes,10,opt,name=lastCookedDate,proto3" json:"lastCookedDate,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_recipes_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Recipe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipe) GetImageContent() string {
	if x != nil {
		return x.ImageContent
	}
	return ""
}

func (x *Recipe) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *Recipe) GetCookingTime() int64 {
	if x != nil {
		return x.CookingTime
	}
	return 0
}

func (x *Recipe) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Recipe) GetRecipeType() string {
	if x != nil {
		return x.RecipeType
	}
	return ""
}

func (x *Recipe) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *Recipe) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *Recipe) GetLastCookedDate() string {
	if x != nil {
		return x.LastCookedDate
	}
	return ""
}

type RecipeCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageContent     []byte `protobuf:"bytes,2,opt,name=image_content,json=imageContent,proto3" json:"image_content,omitempty"`
	Instructions     string `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	CookingTime      int64  `protobuf:"varint,4,opt,name=cookingTime,proto3" json:"cookingTime,omitempty"`
	Difficulty       string `protobuf:"bytes,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	RecipeType       string `protobuf:"bytes,6,opt,name=recipeType,proto3" json:"recipeType,omitempty"`
	CreationDate     string `protobuf:"bytes,7,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	ImageName        string `protobuf:"bytes,8,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	ImageContentType string `protobuf:"bytes,9,opt,name=image_content_type,json=imageContentType,proto3" json:"image_content_type,omitempty"`
}

func (x *RecipeCreateReq) Reset() {
	*x = RecipeCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeCreateReq) ProtoMessage() {}

func (x *RecipeCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeCreateReq.ProtoReflect.Descriptor instead.
func (*RecipeCreateReq) Descriptor() ([]byte, []int) {
	return file_recipes_proto_rawDescGZIP(), []int{1}
}

func (x *RecipeCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecipeCreateReq) GetImageContent() []byte {
	if x != nil {
		return x.ImageContent
	}
	return nil
}

func (x *RecipeCreateReq) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *RecipeCreateReq) GetCookingTime() int64 {
	if x != nil {
		return x.CookingTime
	}
	return 0
}

func (x *RecipeCreateReq) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *RecipeCreateReq) GetRecipeType() string {
	if x != nil {
		return x.RecipeType
	}
	return ""
}

func (x *RecipeCreateReq) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *RecipeCreateReq) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *RecipeCreateReq) GetImageContentType() string {
	if x != nil {
		return x.ImageContentType
	}
	return ""
}

type RecipeCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe *Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *RecipeCreateResp) Reset() {
	*x = RecipeCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeCreateResp) ProtoMessage() {}

func (x *RecipeCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeCreateResp.ProtoReflect.Descriptor instead.
func (*RecipeCreateResp) Descriptor() ([]byte, []int) {
	return file_recipes_proto_rawDescGZIP(), []int{2}
}

func (x *RecipeCreateResp) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

var File_recipes_proto protoreflect.FileDescriptor

var file_recipes_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xc4, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x43, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x0f, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x36,
	0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x32, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recipes_proto_rawDescOnce sync.Once
	file_recipes_proto_rawDescData = file_recipes_proto_rawDesc
)

func file_recipes_proto_rawDescGZIP() []byte {
	file_recipes_proto_rawDescOnce.Do(func() {
		file_recipes_proto_rawDescData = protoimpl.X.CompressGZIP(file_recipes_proto_rawDescData)
	})
	return file_recipes_proto_rawDescData
}

var file_recipes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_recipes_proto_goTypes = []interface{}{
	(*Recipe)(nil),           // 0: pb.Recipe
	(*RecipeCreateReq)(nil),  // 1: pb.RecipeCreateReq
	(*RecipeCreateResp)(nil), // 2: pb.RecipeCreateResp
}
var file_recipes_proto_depIdxs = []int32{
	0, // 0: pb.RecipeCreateResp.recipe:type_name -> pb.Recipe
	1, // 1: pb.recipe.recipeCreate:input_type -> pb.RecipeCreateReq
	2, // 2: pb.recipe.recipeCreate:output_type -> pb.RecipeCreateResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recipes_proto_init() }
func file_recipes_proto_init() {
	if File_recipes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recipes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_recipes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeCreateReq); i {
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
		file_recipes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeCreateResp); i {
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
			RawDescriptor: file_recipes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recipes_proto_goTypes,
		DependencyIndexes: file_recipes_proto_depIdxs,
		MessageInfos:      file_recipes_proto_msgTypes,
	}.Build()
	File_recipes_proto = out.File
	file_recipes_proto_rawDesc = nil
	file_recipes_proto_goTypes = nil
	file_recipes_proto_depIdxs = nil
}
