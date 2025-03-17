package httprt

import (
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *HTTPConfig            `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	AllowHttp     bool                   `protobuf:"varint,2,opt,name=allow_http,json=allowHttp,proto3" json:"allow_http,omitempty"`
	H2PoolSize    int32                  `protobuf:"varint,3,opt,name=h2_pool_size,json=h2PoolSize,proto3" json:"h2_pool_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_request_roundtripper_httprt_config_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConfig) GetHttp() *HTTPConfig {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *ClientConfig) GetAllowHttp() bool {
	if x != nil {
		return x.AllowHttp
	}
	return false
}

func (x *ClientConfig) GetH2PoolSize() int32 {
	if x != nil {
		return x.H2PoolSize
	}
	return 0
}

type ServerConfig struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Http                 *HTTPConfig            `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	NoDecodingSessionTag bool                   `protobuf:"varint,2,opt,name=no_decoding_session_tag,json=noDecodingSessionTag,proto3" json:"no_decoding_session_tag,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_request_roundtripper_httprt_config_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfig) GetHttp() *HTTPConfig {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *ServerConfig) GetNoDecodingSessionTag() bool {
	if x != nil {
		return x.NoDecodingSessionTag
	}
	return false
}

type HTTPConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	UrlPrefix     string                 `protobuf:"bytes,2,opt,name=urlPrefix,proto3" json:"urlPrefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HTTPConfig) Reset() {
	*x = HTTPConfig{}
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HTTPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPConfig) ProtoMessage() {}

func (x *HTTPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPConfig.ProtoReflect.Descriptor instead.
func (*HTTPConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_request_roundtripper_httprt_config_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HTTPConfig) GetUrlPrefix() string {
	if x != nil {
		return x.UrlPrefix
	}
	return ""
}

var File_transport_internet_request_roundtripper_httprt_config_proto protoreflect.FileDescriptor

var file_transport_internet_request_roundtripper_httprt_config_proto_rawDesc = string([]byte{
	0x0a, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x72, 0x74,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x72, 0x74, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x59, 0x0a, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x72, 0x74, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x48, 0x74, 0x74, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x32, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x32, 0x50,
	0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x33, 0x82, 0xb5, 0x18, 0x2f, 0x0a, 0x25, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x68, 0x74, 0x74, 0x70, 0x72, 0x74, 0x22, 0xd5, 0x01, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x59, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x72, 0x74, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x6f, 0x5f, 0x64,
	0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x6e, 0x6f, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x3a,
	0x33, 0x82, 0xb5, 0x18, 0x2f, 0x0a, 0x25, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x06, 0x68, 0x74,
	0x74, 0x70, 0x72, 0x74, 0x22, 0x3e, 0x0a, 0x0a, 0x48, 0x54, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x72, 0x6c, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x72, 0x6c, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x42, 0xcc, 0x01, 0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x72, 0x74, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x72, 0x74, 0xaa, 0x02, 0x39, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e,
	0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_transport_internet_request_roundtripper_httprt_config_proto_rawDescOnce sync.Once
	file_transport_internet_request_roundtripper_httprt_config_proto_rawDescData []byte
)

func file_transport_internet_request_roundtripper_httprt_config_proto_rawDescGZIP() []byte {
	file_transport_internet_request_roundtripper_httprt_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_request_roundtripper_httprt_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_request_roundtripper_httprt_config_proto_rawDesc), len(file_transport_internet_request_roundtripper_httprt_config_proto_rawDesc)))
	})
	return file_transport_internet_request_roundtripper_httprt_config_proto_rawDescData
}

var file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transport_internet_request_roundtripper_httprt_config_proto_goTypes = []any{
	(*ClientConfig)(nil), // 0: v2ray.core.transport.internet.request.roundtripper.httprt.ClientConfig
	(*ServerConfig)(nil), // 1: v2ray.core.transport.internet.request.roundtripper.httprt.ServerConfig
	(*HTTPConfig)(nil),   // 2: v2ray.core.transport.internet.request.roundtripper.httprt.HTTPConfig
}
var file_transport_internet_request_roundtripper_httprt_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.transport.internet.request.roundtripper.httprt.ClientConfig.http:type_name -> v2ray.core.transport.internet.request.roundtripper.httprt.HTTPConfig
	2, // 1: v2ray.core.transport.internet.request.roundtripper.httprt.ServerConfig.http:type_name -> v2ray.core.transport.internet.request.roundtripper.httprt.HTTPConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_internet_request_roundtripper_httprt_config_proto_init() }
func file_transport_internet_request_roundtripper_httprt_config_proto_init() {
	if File_transport_internet_request_roundtripper_httprt_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_request_roundtripper_httprt_config_proto_rawDesc), len(file_transport_internet_request_roundtripper_httprt_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_request_roundtripper_httprt_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_request_roundtripper_httprt_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_request_roundtripper_httprt_config_proto_msgTypes,
	}.Build()
	File_transport_internet_request_roundtripper_httprt_config_proto = out.File
	file_transport_internet_request_roundtripper_httprt_config_proto_goTypes = nil
	file_transport_internet_request_roundtripper_httprt_config_proto_depIdxs = nil
}
