// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: system/v1alpha1/zone_overview.proto

package v1alpha1

import (
	_ "github.com/kumahq/kuma/api/mesh"
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

// ZoneOverview defines the projected state of a Zone.
type ZoneOverview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zone        *Zone        `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
	ZoneInsight *ZoneInsight `protobuf:"bytes,2,opt,name=zone_insight,json=zoneInsight,proto3" json:"zone_insight,omitempty"`
}

func (x *ZoneOverview) Reset() {
	*x = ZoneOverview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1alpha1_zone_overview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneOverview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneOverview) ProtoMessage() {}

func (x *ZoneOverview) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1alpha1_zone_overview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneOverview.ProtoReflect.Descriptor instead.
func (*ZoneOverview) Descriptor() ([]byte, []int) {
	return file_system_v1alpha1_zone_overview_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneOverview) GetZone() *Zone {
	if x != nil {
		return x.Zone
	}
	return nil
}

func (x *ZoneOverview) GetZoneInsight() *ZoneInsight {
	if x != nil {
		return x.ZoneInsight
	}
	return nil
}

var File_system_v1alpha1_zone_overview_proto protoreflect.FileDescriptor

var file_system_v1alpha1_zone_overview_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdc, 0x01, 0x0a, 0x0c, 0x5a, 0x6f, 0x6e, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x2e, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x44, 0x0a, 0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x56, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x16, 0x0a, 0x14,
	0x5a, 0x6f, 0x6e, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x0e, 0x12, 0x0c, 0x5a, 0x6f, 0x6e, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x08, 0x22, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x18, 0x01, 0xaa, 0x8c,
	0x89, 0xa6, 0x01, 0x02, 0x30, 0x01, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x60, 0x01, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d,
	0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1alpha1_zone_overview_proto_rawDescOnce sync.Once
	file_system_v1alpha1_zone_overview_proto_rawDescData = file_system_v1alpha1_zone_overview_proto_rawDesc
)

func file_system_v1alpha1_zone_overview_proto_rawDescGZIP() []byte {
	file_system_v1alpha1_zone_overview_proto_rawDescOnce.Do(func() {
		file_system_v1alpha1_zone_overview_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1alpha1_zone_overview_proto_rawDescData)
	})
	return file_system_v1alpha1_zone_overview_proto_rawDescData
}

var file_system_v1alpha1_zone_overview_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_system_v1alpha1_zone_overview_proto_goTypes = []interface{}{
	(*ZoneOverview)(nil), // 0: kuma.system.v1alpha1.ZoneOverview
	(*Zone)(nil),         // 1: kuma.system.v1alpha1.Zone
	(*ZoneInsight)(nil),  // 2: kuma.system.v1alpha1.ZoneInsight
}
var file_system_v1alpha1_zone_overview_proto_depIdxs = []int32{
	1, // 0: kuma.system.v1alpha1.ZoneOverview.zone:type_name -> kuma.system.v1alpha1.Zone
	2, // 1: kuma.system.v1alpha1.ZoneOverview.zone_insight:type_name -> kuma.system.v1alpha1.ZoneInsight
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_system_v1alpha1_zone_overview_proto_init() }
func file_system_v1alpha1_zone_overview_proto_init() {
	if File_system_v1alpha1_zone_overview_proto != nil {
		return
	}
	file_system_v1alpha1_zone_proto_init()
	file_system_v1alpha1_zone_insight_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_system_v1alpha1_zone_overview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneOverview); i {
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
			RawDescriptor: file_system_v1alpha1_zone_overview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_v1alpha1_zone_overview_proto_goTypes,
		DependencyIndexes: file_system_v1alpha1_zone_overview_proto_depIdxs,
		MessageInfos:      file_system_v1alpha1_zone_overview_proto_msgTypes,
	}.Build()
	File_system_v1alpha1_zone_overview_proto = out.File
	file_system_v1alpha1_zone_overview_proto_rawDesc = nil
	file_system_v1alpha1_zone_overview_proto_goTypes = nil
	file_system_v1alpha1_zone_overview_proto_depIdxs = nil
}
