// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: statistic.proto

package statistic

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

type CountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A          int32   `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B          int32   `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	WinCount   float64 `protobuf:"fixed64,3,opt,name=win_count,json=winCount,proto3" json:"win_count,omitempty"`
	TotalCount float64 `protobuf:"fixed64,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CountData) Reset() {
	*x = CountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountData) ProtoMessage() {}

func (x *CountData) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountData.ProtoReflect.Descriptor instead.
func (*CountData) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{0}
}

func (x *CountData) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CountData) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *CountData) GetWinCount() float64 {
	if x != nil {
		return x.WinCount
	}
	return 0
}

func (x *CountData) GetTotalCount() float64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type StatisticData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*CountData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *StatisticData) Reset() {
	*x = StatisticData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticData) ProtoMessage() {}

func (x *StatisticData) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticData.ProtoReflect.Descriptor instead.
func (*StatisticData) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{1}
}

func (x *StatisticData) GetData() []*CountData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_statistic_proto protoreflect.FileDescriptor

var file_statistic_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x66, 0x0a, 0x0a,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x3b,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_statistic_proto_rawDescOnce sync.Once
	file_statistic_proto_rawDescData = file_statistic_proto_rawDesc
)

func file_statistic_proto_rawDescGZIP() []byte {
	file_statistic_proto_rawDescOnce.Do(func() {
		file_statistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistic_proto_rawDescData)
	})
	return file_statistic_proto_rawDescData
}

var file_statistic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_statistic_proto_goTypes = []interface{}{
	(*CountData)(nil),     // 0: statistic.count_data
	(*StatisticData)(nil), // 1: statistic.statistic_data
}
var file_statistic_proto_depIdxs = []int32{
	0, // 0: statistic.statistic_data.data:type_name -> statistic.count_data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_statistic_proto_init() }
func file_statistic_proto_init() {
	if File_statistic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountData); i {
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
		file_statistic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticData); i {
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
			RawDescriptor: file_statistic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_statistic_proto_goTypes,
		DependencyIndexes: file_statistic_proto_depIdxs,
		MessageInfos:      file_statistic_proto_msgTypes,
	}.Build()
	File_statistic_proto = out.File
	file_statistic_proto_rawDesc = nil
	file_statistic_proto_goTypes = nil
	file_statistic_proto_depIdxs = nil
}
