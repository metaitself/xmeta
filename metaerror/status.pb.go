// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: metaerror/status.proto

package metaerror

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

type MetaError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status   int32             `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg      string            `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Cause    string            `protobuf:"bytes,4,opt,name=cause,proto3" json:"cause,omitempty"`
	Reason   string            `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetaError) Reset() {
	*x = MetaError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaerror_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaError) ProtoMessage() {}

func (x *MetaError) ProtoReflect() protoreflect.Message {
	mi := &file_metaerror_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaError.ProtoReflect.Descriptor instead.
func (*MetaError) Descriptor() ([]byte, []int) {
	return file_metaerror_status_proto_rawDescGZIP(), []int{0}
}

func (x *MetaError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MetaError) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MetaError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MetaError) GetCause() string {
	if x != nil {
		return x.Cause
	}
	return ""
}

func (x *MetaError) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *MetaError) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_metaerror_status_proto protoreflect.FileDescriptor

var file_metaerror_status_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0xf4, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metaerror_status_proto_rawDescOnce sync.Once
	file_metaerror_status_proto_rawDescData = file_metaerror_status_proto_rawDesc
)

func file_metaerror_status_proto_rawDescGZIP() []byte {
	file_metaerror_status_proto_rawDescOnce.Do(func() {
		file_metaerror_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_metaerror_status_proto_rawDescData)
	})
	return file_metaerror_status_proto_rawDescData
}

var file_metaerror_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_metaerror_status_proto_goTypes = []interface{}{
	(*MetaError)(nil), // 0: metaerror.MetaError
	nil,               // 1: metaerror.MetaError.MetadataEntry
}
var file_metaerror_status_proto_depIdxs = []int32{
	1, // 0: metaerror.MetaError.metadata:type_name -> metaerror.MetaError.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metaerror_status_proto_init() }
func file_metaerror_status_proto_init() {
	if File_metaerror_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metaerror_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaError); i {
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
			RawDescriptor: file_metaerror_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metaerror_status_proto_goTypes,
		DependencyIndexes: file_metaerror_status_proto_depIdxs,
		MessageInfos:      file_metaerror_status_proto_msgTypes,
	}.Build()
	File_metaerror_status_proto = out.File
	file_metaerror_status_proto_rawDesc = nil
	file_metaerror_status_proto_goTypes = nil
	file_metaerror_status_proto_depIdxs = nil
}
