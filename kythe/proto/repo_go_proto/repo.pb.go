// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: kythe/proto/repo.proto

package repo_go_proto

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

type BuildSystem int32

const (
	BuildSystem_UNKNOWN BuildSystem = 0
	BuildSystem_MAVEN   BuildSystem = 1
	BuildSystem_GRADLE  BuildSystem = 2
	BuildSystem_BAZEL   BuildSystem = 3
)

// Enum value maps for BuildSystem.
var (
	BuildSystem_name = map[int32]string{
		0: "UNKNOWN",
		1: "MAVEN",
		2: "GRADLE",
		3: "BAZEL",
	}
	BuildSystem_value = map[string]int32{
		"UNKNOWN": 0,
		"MAVEN":   1,
		"GRADLE":  2,
		"BAZEL":   3,
	}
)

func (x BuildSystem) Enum() *BuildSystem {
	p := new(BuildSystem)
	*p = x
	return p
}

func (x BuildSystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildSystem) Descriptor() protoreflect.EnumDescriptor {
	return file_kythe_proto_repo_proto_enumTypes[0].Descriptor()
}

func (BuildSystem) Type() protoreflect.EnumType {
	return &file_kythe_proto_repo_proto_enumTypes[0]
}

func (x BuildSystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildSystem.Descriptor instead.
func (BuildSystem) EnumDescriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{0}
}

type JavaOptions_Version int32

const (
	JavaOptions_UNKNOWN  JavaOptions_Version = 0
	JavaOptions_JAVA_1_8 JavaOptions_Version = 1
)

// Enum value maps for JavaOptions_Version.
var (
	JavaOptions_Version_name = map[int32]string{
		0: "UNKNOWN",
		1: "JAVA_1_8",
	}
	JavaOptions_Version_value = map[string]int32{
		"UNKNOWN":  0,
		"JAVA_1_8": 1,
	}
)

func (x JavaOptions_Version) Enum() *JavaOptions_Version {
	p := new(JavaOptions_Version)
	*p = x
	return p
}

func (x JavaOptions_Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JavaOptions_Version) Descriptor() protoreflect.EnumDescriptor {
	return file_kythe_proto_repo_proto_enumTypes[1].Descriptor()
}

func (JavaOptions_Version) Type() protoreflect.EnumType {
	return &file_kythe_proto_repo_proto_enumTypes[1]
}

func (x JavaOptions_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JavaOptions_Version.Descriptor instead.
func (JavaOptions_Version) EnumDescriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{2, 0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extractions []*ExtractionHint `protobuf:"bytes,1,rep,name=extractions,proto3" json:"extractions,omitempty"`
	Repo        string            `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetExtractions() []*ExtractionHint {
	if x != nil {
		return x.Extractions
	}
	return nil
}

func (x *Config) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

type ExtractionHint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildSystem BuildSystem         `protobuf:"varint,1,opt,name=build_system,json=buildSystem,proto3,enum=kythe.proto.repo.BuildSystem" json:"build_system,omitempty"`
	Corpus      string              `protobuf:"bytes,2,opt,name=corpus,proto3" json:"corpus,omitempty"`
	JavaOptions *JavaOptions        `protobuf:"bytes,3,opt,name=java_options,json=javaOptions,proto3" json:"java_options,omitempty"`
	Targets     []*ExtractionTarget `protobuf:"bytes,4,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *ExtractionHint) Reset() {
	*x = ExtractionHint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionHint) ProtoMessage() {}

func (x *ExtractionHint) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionHint.ProtoReflect.Descriptor instead.
func (*ExtractionHint) Descriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{1}
}

func (x *ExtractionHint) GetBuildSystem() BuildSystem {
	if x != nil {
		return x.BuildSystem
	}
	return BuildSystem_UNKNOWN
}

func (x *ExtractionHint) GetCorpus() string {
	if x != nil {
		return x.Corpus
	}
	return ""
}

func (x *ExtractionHint) GetJavaOptions() *JavaOptions {
	if x != nil {
		return x.JavaOptions
	}
	return nil
}

func (x *ExtractionHint) GetTargets() []*ExtractionTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

type JavaOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version JavaOptions_Version `protobuf:"varint,1,opt,name=version,proto3,enum=kythe.proto.repo.JavaOptions_Version" json:"version,omitempty"`
}

func (x *JavaOptions) Reset() {
	*x = JavaOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JavaOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JavaOptions) ProtoMessage() {}

func (x *JavaOptions) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JavaOptions.ProtoReflect.Descriptor instead.
func (*JavaOptions) Descriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{2}
}

func (x *JavaOptions) GetVersion() JavaOptions_Version {
	if x != nil {
		return x.Version
	}
	return JavaOptions_UNKNOWN
}

type ExtractionTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path              string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IndividualTargets []string `protobuf:"bytes,2,rep,name=individual_targets,json=individualTargets,proto3" json:"individual_targets,omitempty"`
}

func (x *ExtractionTarget) Reset() {
	*x = ExtractionTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionTarget) ProtoMessage() {}

func (x *ExtractionTarget) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionTarget.ProtoReflect.Descriptor instead.
func (*ExtractionTarget) Descriptor() ([]byte, []int) {
	return file_kythe_proto_repo_proto_rawDescGZIP(), []int{3}
}

func (x *ExtractionTarget) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ExtractionTarget) GetIndividualTargets() []string {
	if x != nil {
		return x.IndividualTargets
	}
	return nil
}

var File_kythe_proto_repo_proto protoreflect.FileDescriptor

var file_kythe_proto_repo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x60, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x22, 0xea, 0x01, 0x0a,
	0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x12,
	0x40, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x6a, 0x61, 0x76,
	0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x4a, 0x61, 0x76, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x6a, 0x61, 0x76, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b,
	0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0x74, 0x0a, 0x0b, 0x4a, 0x61, 0x76,
	0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4a, 0x61, 0x76,
	0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x41, 0x56, 0x41, 0x5f, 0x31, 0x5f, 0x38, 0x10, 0x01, 0x22,
	0x55, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x64, 0x69, 0x76,
	0x69, 0x64, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2a, 0x3c, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x56, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x47, 0x52, 0x41, 0x44, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x5a,
	0x45, 0x4c, 0x10, 0x03, 0x42, 0x24, 0x5a, 0x22, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x69, 0x6f,
	0x2f, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kythe_proto_repo_proto_rawDescOnce sync.Once
	file_kythe_proto_repo_proto_rawDescData = file_kythe_proto_repo_proto_rawDesc
)

func file_kythe_proto_repo_proto_rawDescGZIP() []byte {
	file_kythe_proto_repo_proto_rawDescOnce.Do(func() {
		file_kythe_proto_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_repo_proto_rawDescData)
	})
	return file_kythe_proto_repo_proto_rawDescData
}

var file_kythe_proto_repo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kythe_proto_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kythe_proto_repo_proto_goTypes = []interface{}{
	(BuildSystem)(0),         // 0: kythe.proto.repo.BuildSystem
	(JavaOptions_Version)(0), // 1: kythe.proto.repo.JavaOptions.Version
	(*Config)(nil),           // 2: kythe.proto.repo.Config
	(*ExtractionHint)(nil),   // 3: kythe.proto.repo.ExtractionHint
	(*JavaOptions)(nil),      // 4: kythe.proto.repo.JavaOptions
	(*ExtractionTarget)(nil), // 5: kythe.proto.repo.ExtractionTarget
}
var file_kythe_proto_repo_proto_depIdxs = []int32{
	3, // 0: kythe.proto.repo.Config.extractions:type_name -> kythe.proto.repo.ExtractionHint
	0, // 1: kythe.proto.repo.ExtractionHint.build_system:type_name -> kythe.proto.repo.BuildSystem
	4, // 2: kythe.proto.repo.ExtractionHint.java_options:type_name -> kythe.proto.repo.JavaOptions
	5, // 3: kythe.proto.repo.ExtractionHint.targets:type_name -> kythe.proto.repo.ExtractionTarget
	1, // 4: kythe.proto.repo.JavaOptions.version:type_name -> kythe.proto.repo.JavaOptions.Version
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kythe_proto_repo_proto_init() }
func file_kythe_proto_repo_proto_init() {
	if File_kythe_proto_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_kythe_proto_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionHint); i {
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
		file_kythe_proto_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JavaOptions); i {
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
		file_kythe_proto_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionTarget); i {
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
			RawDescriptor: file_kythe_proto_repo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kythe_proto_repo_proto_goTypes,
		DependencyIndexes: file_kythe_proto_repo_proto_depIdxs,
		EnumInfos:         file_kythe_proto_repo_proto_enumTypes,
		MessageInfos:      file_kythe_proto_repo_proto_msgTypes,
	}.Build()
	File_kythe_proto_repo_proto = out.File
	file_kythe_proto_repo_proto_rawDesc = nil
	file_kythe_proto_repo_proto_goTypes = nil
	file_kythe_proto_repo_proto_depIdxs = nil
}
