// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: services/topic.proto

package pb

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_services_topic_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Todo      string `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	UserId    int64  `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_services_topic_proto_rawDescGZIP(), []int{1}
}

func (x *TodoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoResponse) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *TodoResponse) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *TodoResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64           `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Todo      []*TodoResponse `protobuf:"bytes,2,rep,name=todo,proto3" json:"todo,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_topic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_services_topic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_services_topic_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Response) GetTodo() []*TodoResponse {
	if x != nil {
		return x.Todo
	}
	return nil
}

var File_services_topic_proto protoreflect.FileDescriptor

var file_services_topic_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x2a, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x68, 0x0a, 0x0c, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x64,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x32, 0x3f, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_topic_proto_rawDescOnce sync.Once
	file_services_topic_proto_rawDescData = file_services_topic_proto_rawDesc
)

func file_services_topic_proto_rawDescGZIP() []byte {
	file_services_topic_proto_rawDescOnce.Do(func() {
		file_services_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_topic_proto_rawDescData)
	})
	return file_services_topic_proto_rawDescData
}

var file_services_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_topic_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: topic.GetRequest
	(*TodoResponse)(nil), // 1: topic.TodoResponse
	(*Response)(nil),     // 2: topic.Response
}
var file_services_topic_proto_depIdxs = []int32{
	1, // 0: topic.Response.todo:type_name -> topic.TodoResponse
	0, // 1: topic.TopicService.GetTopics:input_type -> topic.GetRequest
	2, // 2: topic.TopicService.GetTopics:output_type -> topic.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_topic_proto_init() }
func file_services_topic_proto_init() {
	if File_services_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_services_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_services_topic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_services_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_topic_proto_goTypes,
		DependencyIndexes: file_services_topic_proto_depIdxs,
		MessageInfos:      file_services_topic_proto_msgTypes,
	}.Build()
	File_services_topic_proto = out.File
	file_services_topic_proto_rawDesc = nil
	file_services_topic_proto_goTypes = nil
	file_services_topic_proto_depIdxs = nil
}
