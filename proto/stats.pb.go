// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: stats.proto

package systat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Battery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatteryLoad float64 `protobuf:"fixed64,1,opt,name=batteryLoad,proto3" json:"batteryLoad,omitempty"`
	State       string  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Battery) Reset() {
	*x = Battery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Battery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Battery) ProtoMessage() {}

func (x *Battery) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Battery.ProtoReflect.Descriptor instead.
func (*Battery) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Battery) GetBatteryLoad() float64 {
	if x != nil {
		return x.BatteryLoad
	}
	return 0
}

func (x *Battery) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type Cpu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuLoad float64 `protobuf:"fixed64,1,opt,name=cpuLoad,proto3" json:"cpuLoad,omitempty"`
}

func (x *Cpu) Reset() {
	*x = Cpu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cpu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cpu) ProtoMessage() {}

func (x *Cpu) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cpu.ProtoReflect.Descriptor instead.
func (*Cpu) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{1}
}

func (x *Cpu) GetCpuLoad() float64 {
	if x != nil {
		return x.CpuLoad
	}
	return 0
}

type BatteriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batteries []*Battery `protobuf:"bytes,1,rep,name=batteries,proto3" json:"batteries,omitempty"`
}

func (x *BatteriesResponse) Reset() {
	*x = BatteriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatteriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatteriesResponse) ProtoMessage() {}

func (x *BatteriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatteriesResponse.ProtoReflect.Descriptor instead.
func (*BatteriesResponse) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{2}
}

func (x *BatteriesResponse) GetBatteries() []*Battery {
	if x != nil {
		return x.Batteries
	}
	return nil
}

type CpusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpus []*Cpu `protobuf:"bytes,1,rep,name=cpus,proto3" json:"cpus,omitempty"`
}

func (x *CpusResponse) Reset() {
	*x = CpusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpusResponse) ProtoMessage() {}

func (x *CpusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpusResponse.ProtoReflect.Descriptor instead.
func (*CpusResponse) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{3}
}

func (x *CpusResponse) GetCpus() []*Cpu {
	if x != nil {
		return x.Cpus
	}
	return nil
}

var File_stats_proto protoreflect.FileDescriptor

var file_stats_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x07, 0x42, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x4c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a,
	0x03, 0x43, 0x70, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x22, 0x3b,
	0x0a, 0x11, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x52, 0x09, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x43,
	0x70, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x63,
	0x70, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x43, 0x70, 0x75, 0x52,
	0x04, 0x63, 0x70, 0x75, 0x73, 0x32, 0x7d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x34,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x70, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0d, 0x2e, 0x43, 0x70, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x68, 0x65, 0x6c, 0x61, 0x62, 0x6f, 0x76, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_proto_rawDescOnce sync.Once
	file_stats_proto_rawDescData = file_stats_proto_rawDesc
)

func file_stats_proto_rawDescGZIP() []byte {
	file_stats_proto_rawDescOnce.Do(func() {
		file_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_proto_rawDescData)
	})
	return file_stats_proto_rawDescData
}

var file_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stats_proto_goTypes = []interface{}{
	(*Battery)(nil),           // 0: Battery
	(*Cpu)(nil),               // 1: Cpu
	(*BatteriesResponse)(nil), // 2: BatteriesResponse
	(*CpusResponse)(nil),      // 3: CpusResponse
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_stats_proto_depIdxs = []int32{
	0, // 0: BatteriesResponse.batteries:type_name -> Battery
	1, // 1: CpusResponse.cpus:type_name -> Cpu
	4, // 2: Stats.GetBatteries:input_type -> google.protobuf.Empty
	4, // 3: Stats.GetCpus:input_type -> google.protobuf.Empty
	2, // 4: Stats.GetBatteries:output_type -> BatteriesResponse
	3, // 5: Stats.GetCpus:output_type -> CpusResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stats_proto_init() }
func file_stats_proto_init() {
	if File_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Battery); i {
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
		file_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cpu); i {
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
		file_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatteriesResponse); i {
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
		file_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpusResponse); i {
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
			RawDescriptor: file_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stats_proto_goTypes,
		DependencyIndexes: file_stats_proto_depIdxs,
		MessageInfos:      file_stats_proto_msgTypes,
	}.Build()
	File_stats_proto = out.File
	file_stats_proto_rawDesc = nil
	file_stats_proto_goTypes = nil
	file_stats_proto_depIdxs = nil
}
