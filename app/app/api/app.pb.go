// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: app/app/api/app.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FilUsdtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Interval  string `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Limit     string `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	StartTime string `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   string `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *FilUsdtRequest) Reset() {
	*x = FilUsdtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_api_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilUsdtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilUsdtRequest) ProtoMessage() {}

func (x *FilUsdtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_api_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilUsdtRequest.ProtoReflect.Descriptor instead.
func (*FilUsdtRequest) Descriptor() ([]byte, []int) {
	return file_app_app_api_app_proto_rawDescGZIP(), []int{0}
}

func (x *FilUsdtRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *FilUsdtRequest) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *FilUsdtRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *FilUsdtRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *FilUsdtRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type FilUsdtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataListK []*FilUsdtReply_ListK `protobuf:"bytes,1,rep,name=dataListK,proto3" json:"dataListK,omitempty"`
}

func (x *FilUsdtReply) Reset() {
	*x = FilUsdtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_api_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilUsdtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilUsdtReply) ProtoMessage() {}

func (x *FilUsdtReply) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_api_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilUsdtReply.ProtoReflect.Descriptor instead.
func (*FilUsdtReply) Descriptor() ([]byte, []int) {
	return file_app_app_api_app_proto_rawDescGZIP(), []int{1}
}

func (x *FilUsdtReply) GetDataListK() []*FilUsdtReply_ListK {
	if x != nil {
		return x.DataListK
	}
	return nil
}

type FilUsdtReply_ListK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime           int64   `protobuf:"varint,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime             int64   `protobuf:"varint,2,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	StartPrice          float64 `protobuf:"fixed64,3,opt,name=StartPrice,proto3" json:"StartPrice,omitempty"`
	TopPrice            float64 `protobuf:"fixed64,4,opt,name=TopPrice,proto3" json:"TopPrice,omitempty"`
	LowPrice            float64 `protobuf:"fixed64,5,opt,name=LowPrice,proto3" json:"LowPrice,omitempty"`
	EndPrice            float64 `protobuf:"fixed64,6,opt,name=EndPrice,proto3" json:"EndPrice,omitempty"`
	DealTotalAmount     float64 `protobuf:"fixed64,7,opt,name=DealTotalAmount,proto3" json:"DealTotalAmount,omitempty"`
	DealAmount          float64 `protobuf:"fixed64,8,opt,name=DealAmount,proto3" json:"DealAmount,omitempty"`
	DealTotal           int64   `protobuf:"varint,9,opt,name=DealTotal,proto3" json:"DealTotal,omitempty"`
	DealSelfTotalAmount float64 `protobuf:"fixed64,10,opt,name=DealSelfTotalAmount,proto3" json:"DealSelfTotalAmount,omitempty"`
	DealSelfAmount      float64 `protobuf:"fixed64,11,opt,name=DealSelfAmount,proto3" json:"DealSelfAmount,omitempty"`
}

func (x *FilUsdtReply_ListK) Reset() {
	*x = FilUsdtReply_ListK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_api_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilUsdtReply_ListK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilUsdtReply_ListK) ProtoMessage() {}

func (x *FilUsdtReply_ListK) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_api_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilUsdtReply_ListK.ProtoReflect.Descriptor instead.
func (*FilUsdtReply_ListK) Descriptor() ([]byte, []int) {
	return file_app_app_api_app_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FilUsdtReply_ListK) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetStartPrice() float64 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetTopPrice() float64 {
	if x != nil {
		return x.TopPrice
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetLowPrice() float64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetEndPrice() float64 {
	if x != nil {
		return x.EndPrice
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetDealTotalAmount() float64 {
	if x != nil {
		return x.DealTotalAmount
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetDealAmount() float64 {
	if x != nil {
		return x.DealAmount
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetDealTotal() int64 {
	if x != nil {
		return x.DealTotal
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetDealSelfTotalAmount() float64 {
	if x != nil {
		return x.DealSelfTotalAmount
	}
	return 0
}

func (x *FilUsdtReply_ListK) GetDealSelfAmount() float64 {
	if x != nil {
		return x.DealSelfAmount
	}
	return 0
}

var File_app_app_api_app_proto protoreflect.FileDescriptor

var file_app_app_api_app_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x55, 0x73, 0x64, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xbf, 0x03, 0x0a, 0x0c, 0x46, 0x69, 0x6c,
	0x55, 0x73, 0x64, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x55, 0x73, 0x64, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x5f, 0x6b, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x1a, 0xf6, 0x02, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x5f, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x54, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x45, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x45, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65, 0x61, 0x6c,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x44, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x30, 0x0a, 0x13, 0x44, 0x65, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x44,
	0x65, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x44, 0x65, 0x61, 0x6c,
	0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5a, 0x0a, 0x03, 0x41, 0x70,
	0x70, 0x12, 0x53, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x55, 0x73, 0x64, 0x74, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x55, 0x73, 0x64, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x55, 0x73, 0x64, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x69,
	0x6c, 0x5f, 0x75, 0x73, 0x64, 0x74, 0x42, 0x11, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a,
	0x08, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_app_api_app_proto_rawDescOnce sync.Once
	file_app_app_api_app_proto_rawDescData = file_app_app_api_app_proto_rawDesc
)

func file_app_app_api_app_proto_rawDescGZIP() []byte {
	file_app_app_api_app_proto_rawDescOnce.Do(func() {
		file_app_app_api_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_app_api_app_proto_rawDescData)
	})
	return file_app_app_api_app_proto_rawDescData
}

var file_app_app_api_app_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_app_api_app_proto_goTypes = []interface{}{
	(*FilUsdtRequest)(nil),     // 0: api.FilUsdtRequest
	(*FilUsdtReply)(nil),       // 1: api.FilUsdtReply
	(*FilUsdtReply_ListK)(nil), // 2: api.FilUsdtReply.List_k
}
var file_app_app_api_app_proto_depIdxs = []int32{
	2, // 0: api.FilUsdtReply.dataListK:type_name -> api.FilUsdtReply.List_k
	0, // 1: api.App.FilUsdt:input_type -> api.FilUsdtRequest
	1, // 2: api.App.FilUsdt:output_type -> api.FilUsdtReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_app_api_app_proto_init() }
func file_app_app_api_app_proto_init() {
	if File_app_app_api_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_app_api_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilUsdtRequest); i {
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
		file_app_app_api_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilUsdtReply); i {
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
		file_app_app_api_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilUsdtReply_ListK); i {
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
			RawDescriptor: file_app_app_api_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_app_api_app_proto_goTypes,
		DependencyIndexes: file_app_app_api_app_proto_depIdxs,
		MessageInfos:      file_app_app_api_app_proto_msgTypes,
	}.Build()
	File_app_app_api_app_proto = out.File
	file_app_app_api_app_proto_rawDesc = nil
	file_app_app_api_app_proto_goTypes = nil
	file_app_app_api_app_proto_depIdxs = nil
}