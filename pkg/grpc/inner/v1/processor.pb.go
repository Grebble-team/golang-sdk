// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: inner/v1/processor.proto

package v1

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

type FlowExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowName   string `protobuf:"bytes,1,opt,name=flowName,proto3" json:"flowName,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Attributes string `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *FlowExecuteRequest) Reset() {
	*x = FlowExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inner_v1_processor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowExecuteRequest) ProtoMessage() {}

func (x *FlowExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inner_v1_processor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowExecuteRequest.ProtoReflect.Descriptor instead.
func (*FlowExecuteRequest) Descriptor() ([]byte, []int) {
	return file_inner_v1_processor_proto_rawDescGZIP(), []int{0}
}

func (x *FlowExecuteRequest) GetFlowName() string {
	if x != nil {
		return x.FlowName
	}
	return ""
}

func (x *FlowExecuteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FlowExecuteRequest) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

type FlowExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content    string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Attributes string `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	StreamEnd  bool   `protobuf:"varint,3,opt,name=streamEnd,proto3" json:"streamEnd,omitempty"`
}

func (x *FlowExecuteResponse) Reset() {
	*x = FlowExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inner_v1_processor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowExecuteResponse) ProtoMessage() {}

func (x *FlowExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inner_v1_processor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowExecuteResponse.ProtoReflect.Descriptor instead.
func (*FlowExecuteResponse) Descriptor() ([]byte, []int) {
	return file_inner_v1_processor_proto_rawDescGZIP(), []int{1}
}

func (x *FlowExecuteResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FlowExecuteResponse) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

func (x *FlowExecuteResponse) GetStreamEnd() bool {
	if x != nil {
		return x.StreamEnd
	}
	return false
}

var File_inner_v1_processor_proto protoreflect.FileDescriptor

var file_inner_v1_processor_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x65, 0x62,
	0x62, 0x6c, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x22, 0x6a, 0x0a, 0x12, 0x46, 0x6c,
	0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x13, 0x46, 0x6c, 0x6f, 0x77, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x6e, 0x64, 0x32, 0x63, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x56, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e,
	0x67, 0x72, 0x65, 0x62, 0x62, 0x6c, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x62, 0x62, 0x6c, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x62, 0x62, 0x6c, 0x65,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inner_v1_processor_proto_rawDescOnce sync.Once
	file_inner_v1_processor_proto_rawDescData = file_inner_v1_processor_proto_rawDesc
)

func file_inner_v1_processor_proto_rawDescGZIP() []byte {
	file_inner_v1_processor_proto_rawDescOnce.Do(func() {
		file_inner_v1_processor_proto_rawDescData = protoimpl.X.CompressGZIP(file_inner_v1_processor_proto_rawDescData)
	})
	return file_inner_v1_processor_proto_rawDescData
}

var file_inner_v1_processor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inner_v1_processor_proto_goTypes = []interface{}{
	(*FlowExecuteRequest)(nil),  // 0: grebbleFlow.v1.FlowExecuteRequest
	(*FlowExecuteResponse)(nil), // 1: grebbleFlow.v1.FlowExecuteResponse
}
var file_inner_v1_processor_proto_depIdxs = []int32{
	0, // 0: grebbleFlow.v1.Processor.Execute:input_type -> grebbleFlow.v1.FlowExecuteRequest
	1, // 1: grebbleFlow.v1.Processor.Execute:output_type -> grebbleFlow.v1.FlowExecuteResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inner_v1_processor_proto_init() }
func file_inner_v1_processor_proto_init() {
	if File_inner_v1_processor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inner_v1_processor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowExecuteRequest); i {
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
		file_inner_v1_processor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowExecuteResponse); i {
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
			RawDescriptor: file_inner_v1_processor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inner_v1_processor_proto_goTypes,
		DependencyIndexes: file_inner_v1_processor_proto_depIdxs,
		MessageInfos:      file_inner_v1_processor_proto_msgTypes,
	}.Build()
	File_inner_v1_processor_proto = out.File
	file_inner_v1_processor_proto_rawDesc = nil
	file_inner_v1_processor_proto_goTypes = nil
	file_inner_v1_processor_proto_depIdxs = nil
}
