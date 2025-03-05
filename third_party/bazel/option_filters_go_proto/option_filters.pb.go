// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: third_party/bazel/src/main/protobuf/option_filters.proto

package option_filters_go_proto

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

type OptionEffectTag int32

const (
	OptionEffectTag_UNKNOWN                             OptionEffectTag = 0
	OptionEffectTag_NO_OP                               OptionEffectTag = 1
	OptionEffectTag_LOSES_INCREMENTAL_STATE             OptionEffectTag = 2
	OptionEffectTag_CHANGES_INPUTS                      OptionEffectTag = 3
	OptionEffectTag_AFFECTS_OUTPUTS                     OptionEffectTag = 4
	OptionEffectTag_BUILD_FILE_SEMANTICS                OptionEffectTag = 5
	OptionEffectTag_BAZEL_INTERNAL_CONFIGURATION        OptionEffectTag = 6
	OptionEffectTag_LOADING_AND_ANALYSIS                OptionEffectTag = 7
	OptionEffectTag_EXECUTION                           OptionEffectTag = 8
	OptionEffectTag_HOST_MACHINE_RESOURCE_OPTIMIZATIONS OptionEffectTag = 9
	OptionEffectTag_EAGERNESS_TO_EXIT                   OptionEffectTag = 10
	OptionEffectTag_BAZEL_MONITORING                    OptionEffectTag = 11
	OptionEffectTag_TERMINAL_OUTPUT                     OptionEffectTag = 12
	OptionEffectTag_ACTION_COMMAND_LINES                OptionEffectTag = 13
	OptionEffectTag_TEST_RUNNER                         OptionEffectTag = 14
)

// Enum value maps for OptionEffectTag.
var (
	OptionEffectTag_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "NO_OP",
		2:  "LOSES_INCREMENTAL_STATE",
		3:  "CHANGES_INPUTS",
		4:  "AFFECTS_OUTPUTS",
		5:  "BUILD_FILE_SEMANTICS",
		6:  "BAZEL_INTERNAL_CONFIGURATION",
		7:  "LOADING_AND_ANALYSIS",
		8:  "EXECUTION",
		9:  "HOST_MACHINE_RESOURCE_OPTIMIZATIONS",
		10: "EAGERNESS_TO_EXIT",
		11: "BAZEL_MONITORING",
		12: "TERMINAL_OUTPUT",
		13: "ACTION_COMMAND_LINES",
		14: "TEST_RUNNER",
	}
	OptionEffectTag_value = map[string]int32{
		"UNKNOWN":                             0,
		"NO_OP":                               1,
		"LOSES_INCREMENTAL_STATE":             2,
		"CHANGES_INPUTS":                      3,
		"AFFECTS_OUTPUTS":                     4,
		"BUILD_FILE_SEMANTICS":                5,
		"BAZEL_INTERNAL_CONFIGURATION":        6,
		"LOADING_AND_ANALYSIS":                7,
		"EXECUTION":                           8,
		"HOST_MACHINE_RESOURCE_OPTIMIZATIONS": 9,
		"EAGERNESS_TO_EXIT":                   10,
		"BAZEL_MONITORING":                    11,
		"TERMINAL_OUTPUT":                     12,
		"ACTION_COMMAND_LINES":                13,
		"TEST_RUNNER":                         14,
	}
)

func (x OptionEffectTag) Enum() *OptionEffectTag {
	p := new(OptionEffectTag)
	*p = x
	return p
}

func (x OptionEffectTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionEffectTag) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes[0].Descriptor()
}

func (OptionEffectTag) Type() protoreflect.EnumType {
	return &file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes[0]
}

func (x OptionEffectTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptionEffectTag.Descriptor instead.
func (OptionEffectTag) EnumDescriptor() ([]byte, []int) {
	return file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescGZIP(), []int{0}
}

type OptionMetadataTag int32

const (
	OptionMetadataTag_EXPERIMENTAL            OptionMetadataTag = 0
	OptionMetadataTag_INCOMPATIBLE_CHANGE     OptionMetadataTag = 1
	OptionMetadataTag_DEPRECATED              OptionMetadataTag = 2
	OptionMetadataTag_HIDDEN                  OptionMetadataTag = 3
	OptionMetadataTag_INTERNAL                OptionMetadataTag = 4
	OptionMetadataTag_EXPLICIT_IN_OUTPUT_PATH OptionMetadataTag = 6
)

// Enum value maps for OptionMetadataTag.
var (
	OptionMetadataTag_name = map[int32]string{
		0: "EXPERIMENTAL",
		1: "INCOMPATIBLE_CHANGE",
		2: "DEPRECATED",
		3: "HIDDEN",
		4: "INTERNAL",
		6: "EXPLICIT_IN_OUTPUT_PATH",
	}
	OptionMetadataTag_value = map[string]int32{
		"EXPERIMENTAL":            0,
		"INCOMPATIBLE_CHANGE":     1,
		"DEPRECATED":              2,
		"HIDDEN":                  3,
		"INTERNAL":                4,
		"EXPLICIT_IN_OUTPUT_PATH": 6,
	}
)

func (x OptionMetadataTag) Enum() *OptionMetadataTag {
	p := new(OptionMetadataTag)
	*p = x
	return p
}

func (x OptionMetadataTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionMetadataTag) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes[1].Descriptor()
}

func (OptionMetadataTag) Type() protoreflect.EnumType {
	return &file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes[1]
}

func (x OptionMetadataTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptionMetadataTag.Descriptor instead.
func (OptionMetadataTag) EnumDescriptor() ([]byte, []int) {
	return file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescGZIP(), []int{1}
}

var File_third_party_bazel_src_main_protobuf_option_filters_proto protoreflect.FileDescriptor

var file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x62, 0x61,
	0x7a, 0x65, 0x6c, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2a, 0xea, 0x02, 0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x54, 0x61, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x5f, 0x4f, 0x50, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x53, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x53, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x41, 0x46, 0x46, 0x45, 0x43, 0x54, 0x53, 0x5f, 0x4f, 0x55, 0x54, 0x50,
	0x55, 0x54, 0x53, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4d, 0x41, 0x4e, 0x54, 0x49, 0x43, 0x53, 0x10, 0x05, 0x12,
	0x20, 0x0a, 0x1c, 0x42, 0x41, 0x5a, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x06, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x53, 0x49, 0x53, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x27, 0x0a, 0x23, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x41, 0x47, 0x45, 0x52, 0x4e, 0x45, 0x53, 0x53,
	0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41,
	0x5a, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0b,
	0x12, 0x13, 0x0a, 0x0f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4f, 0x55, 0x54,
	0x50, 0x55, 0x54, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x53, 0x10, 0x0d, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x0e,
	0x2a, 0xb2, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49,
	0x4d, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x45,
	0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55,
	0x54, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x10, 0x06, 0x22, 0x04, 0x08, 0x05, 0x10, 0x05, 0x2a, 0x25,
	0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x4c, 0x4c,
	0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x53, 0x42, 0x2a, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescOnce sync.Once
	file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescData = file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDesc
)

func file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescGZIP() []byte {
	file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescOnce.Do(func() {
		file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescData)
	})
	return file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDescData
}

var file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_third_party_bazel_src_main_protobuf_option_filters_proto_goTypes = []interface{}{
	(OptionEffectTag)(0),   // 0: options.OptionEffectTag
	(OptionMetadataTag)(0), // 1: options.OptionMetadataTag
}
var file_third_party_bazel_src_main_protobuf_option_filters_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_third_party_bazel_src_main_protobuf_option_filters_proto_init() }
func file_third_party_bazel_src_main_protobuf_option_filters_proto_init() {
	if File_third_party_bazel_src_main_protobuf_option_filters_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_bazel_src_main_protobuf_option_filters_proto_goTypes,
		DependencyIndexes: file_third_party_bazel_src_main_protobuf_option_filters_proto_depIdxs,
		EnumInfos:         file_third_party_bazel_src_main_protobuf_option_filters_proto_enumTypes,
	}.Build()
	File_third_party_bazel_src_main_protobuf_option_filters_proto = out.File
	file_third_party_bazel_src_main_protobuf_option_filters_proto_rawDesc = nil
	file_third_party_bazel_src_main_protobuf_option_filters_proto_goTypes = nil
	file_third_party_bazel_src_main_protobuf_option_filters_proto_depIdxs = nil
}
