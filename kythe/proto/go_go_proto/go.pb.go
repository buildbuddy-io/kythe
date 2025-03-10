// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: kythe/proto/go.proto

package go_go_proto

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

type GoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goos       string   `protobuf:"bytes,1,opt,name=goos,proto3" json:"goos,omitempty"`
	Goarch     string   `protobuf:"bytes,2,opt,name=goarch,proto3" json:"goarch,omitempty"`
	Goroot     string   `protobuf:"bytes,3,opt,name=goroot,proto3" json:"goroot,omitempty"`
	Gopath     string   `protobuf:"bytes,4,opt,name=gopath,proto3" json:"gopath,omitempty"`
	Compiler   string   `protobuf:"bytes,5,opt,name=compiler,proto3" json:"compiler,omitempty"`
	BuildTags  []string `protobuf:"bytes,6,rep,name=build_tags,json=buildTags,proto3" json:"build_tags,omitempty"`
	CgoEnabled bool     `protobuf:"varint,7,opt,name=cgo_enabled,json=cgoEnabled,proto3" json:"cgo_enabled,omitempty"`
}

func (x *GoDetails) Reset() {
	*x = GoDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoDetails) ProtoMessage() {}

func (x *GoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoDetails.ProtoReflect.Descriptor instead.
func (*GoDetails) Descriptor() ([]byte, []int) {
	return file_kythe_proto_go_proto_rawDescGZIP(), []int{0}
}

func (x *GoDetails) GetGoos() string {
	if x != nil {
		return x.Goos
	}
	return ""
}

func (x *GoDetails) GetGoarch() string {
	if x != nil {
		return x.Goarch
	}
	return ""
}

func (x *GoDetails) GetGoroot() string {
	if x != nil {
		return x.Goroot
	}
	return ""
}

func (x *GoDetails) GetGopath() string {
	if x != nil {
		return x.Gopath
	}
	return ""
}

func (x *GoDetails) GetCompiler() string {
	if x != nil {
		return x.Compiler
	}
	return ""
}

func (x *GoDetails) GetBuildTags() []string {
	if x != nil {
		return x.BuildTags
	}
	return nil
}

func (x *GoDetails) GetCgoEnabled() bool {
	if x != nil {
		return x.CgoEnabled
	}
	return false
}

type GoPackageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportPath string `protobuf:"bytes,1,opt,name=import_path,json=importPath,proto3" json:"import_path,omitempty"`
}

func (x *GoPackageInfo) Reset() {
	*x = GoPackageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_go_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoPackageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoPackageInfo) ProtoMessage() {}

func (x *GoPackageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_go_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoPackageInfo.ProtoReflect.Descriptor instead.
func (*GoPackageInfo) Descriptor() ([]byte, []int) {
	return file_kythe_proto_go_proto_rawDescGZIP(), []int{1}
}

func (x *GoPackageInfo) GetImportPath() string {
	if x != nil {
		return x.ImportPath
	}
	return ""
}

type FlagConstructors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag []*FlagConstructor `protobuf:"bytes,1,rep,name=flag,proto3" json:"flag,omitempty"`
}

func (x *FlagConstructors) Reset() {
	*x = FlagConstructors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_go_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagConstructors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagConstructors) ProtoMessage() {}

func (x *FlagConstructors) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_go_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagConstructors.ProtoReflect.Descriptor instead.
func (*FlagConstructors) Descriptor() ([]byte, []int) {
	return file_kythe_proto_go_proto_rawDescGZIP(), []int{2}
}

func (x *FlagConstructors) GetFlag() []*FlagConstructor {
	if x != nil {
		return x.Flag
	}
	return nil
}

type FlagConstructor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PkgPath                string  `protobuf:"bytes,1,opt,name=pkg_path,json=pkgPath,proto3" json:"pkg_path,omitempty"`
	FuncName               string  `protobuf:"bytes,2,opt,name=func_name,json=funcName,proto3" json:"func_name,omitempty"`
	NameArgPosition        uint32  `protobuf:"varint,3,opt,name=name_arg_position,json=nameArgPosition,proto3" json:"name_arg_position,omitempty"`
	DescriptionArgPosition uint32  `protobuf:"varint,4,opt,name=description_arg_position,json=descriptionArgPosition,proto3" json:"description_arg_position,omitempty"`
	VarArgPosition         *uint32 `protobuf:"varint,5,opt,name=var_arg_position,json=varArgPosition,proto3,oneof" json:"var_arg_position,omitempty"`
}

func (x *FlagConstructor) Reset() {
	*x = FlagConstructor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_go_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagConstructor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagConstructor) ProtoMessage() {}

func (x *FlagConstructor) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_go_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagConstructor.ProtoReflect.Descriptor instead.
func (*FlagConstructor) Descriptor() ([]byte, []int) {
	return file_kythe_proto_go_proto_rawDescGZIP(), []int{3}
}

func (x *FlagConstructor) GetPkgPath() string {
	if x != nil {
		return x.PkgPath
	}
	return ""
}

func (x *FlagConstructor) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *FlagConstructor) GetNameArgPosition() uint32 {
	if x != nil {
		return x.NameArgPosition
	}
	return 0
}

func (x *FlagConstructor) GetDescriptionArgPosition() uint32 {
	if x != nil {
		return x.DescriptionArgPosition
	}
	return 0
}

func (x *FlagConstructor) GetVarArgPosition() uint32 {
	if x != nil && x.VarArgPosition != nil {
		return *x.VarArgPosition
	}
	return 0
}

var File_kythe_proto_go_proto protoreflect.FileDescriptor

var file_kythe_proto_go_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x09, 0x47, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x67, 0x6f, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x6f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x6f, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x67, 0x6f, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63,
	0x67, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x30, 0x0a, 0x0d, 0x47, 0x6f, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x10, 0x46,
	0x6c, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x30, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x61, 0x67,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x22, 0xf3, 0x01, 0x0a, 0x0f, 0x46, 0x6c, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6b, 0x67, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x72,
	0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x10, 0x76, 0x61, 0x72, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x0e, 0x76, 0x61, 0x72, 0x41, 0x72, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x43, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x6b,
	0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x20, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kythe_proto_go_proto_rawDescOnce sync.Once
	file_kythe_proto_go_proto_rawDescData = file_kythe_proto_go_proto_rawDesc
)

func file_kythe_proto_go_proto_rawDescGZIP() []byte {
	file_kythe_proto_go_proto_rawDescOnce.Do(func() {
		file_kythe_proto_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_go_proto_rawDescData)
	})
	return file_kythe_proto_go_proto_rawDescData
}

var file_kythe_proto_go_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kythe_proto_go_proto_goTypes = []interface{}{
	(*GoDetails)(nil),        // 0: kythe.proto.GoDetails
	(*GoPackageInfo)(nil),    // 1: kythe.proto.GoPackageInfo
	(*FlagConstructors)(nil), // 2: kythe.proto.FlagConstructors
	(*FlagConstructor)(nil),  // 3: kythe.proto.FlagConstructor
}
var file_kythe_proto_go_proto_depIdxs = []int32{
	3, // 0: kythe.proto.FlagConstructors.flag:type_name -> kythe.proto.FlagConstructor
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kythe_proto_go_proto_init() }
func file_kythe_proto_go_proto_init() {
	if File_kythe_proto_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoDetails); i {
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
		file_kythe_proto_go_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoPackageInfo); i {
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
		file_kythe_proto_go_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagConstructors); i {
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
		file_kythe_proto_go_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagConstructor); i {
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
	file_kythe_proto_go_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kythe_proto_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kythe_proto_go_proto_goTypes,
		DependencyIndexes: file_kythe_proto_go_proto_depIdxs,
		MessageInfos:      file_kythe_proto_go_proto_msgTypes,
	}.Build()
	File_kythe_proto_go_proto = out.File
	file_kythe_proto_go_proto_rawDesc = nil
	file_kythe_proto_go_proto_goTypes = nil
	file_kythe_proto_go_proto_depIdxs = nil
}
