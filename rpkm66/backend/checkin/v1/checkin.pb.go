// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: rpkm66/backend/checkin/v1/checkin.proto

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

type CheckinVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventType int32  `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
}

func (x *CheckinVerifyRequest) Reset() {
	*x = CheckinVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinVerifyRequest) ProtoMessage() {}

func (x *CheckinVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinVerifyRequest.ProtoReflect.Descriptor instead.
func (*CheckinVerifyRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_checkin_v1_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckinVerifyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CheckinVerifyRequest) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

type CheckinVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckinToken string `protobuf:"bytes,1,opt,name=checkin_token,json=checkinToken,proto3" json:"checkin_token,omitempty"`
	CheckinType  int32  `protobuf:"varint,2,opt,name=checkin_type,json=checkinType,proto3" json:"checkin_type,omitempty"`
}

func (x *CheckinVerifyResponse) Reset() {
	*x = CheckinVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinVerifyResponse) ProtoMessage() {}

func (x *CheckinVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinVerifyResponse.ProtoReflect.Descriptor instead.
func (*CheckinVerifyResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_checkin_v1_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckinVerifyResponse) GetCheckinToken() string {
	if x != nil {
		return x.CheckinToken
	}
	return ""
}

func (x *CheckinVerifyResponse) GetCheckinType() int32 {
	if x != nil {
		return x.CheckinType
	}
	return 0
}

type CheckinConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckinConfirmRequest) Reset() {
	*x = CheckinConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinConfirmRequest) ProtoMessage() {}

func (x *CheckinConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinConfirmRequest.ProtoReflect.Descriptor instead.
func (*CheckinConfirmRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_checkin_v1_checkin_proto_rawDescGZIP(), []int{2}
}

func (x *CheckinConfirmRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckinConfirmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckinConfirmResponse) Reset() {
	*x = CheckinConfirmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinConfirmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinConfirmResponse) ProtoMessage() {}

func (x *CheckinConfirmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinConfirmResponse.ProtoReflect.Descriptor instead.
func (*CheckinConfirmResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_checkin_v1_checkin_proto_rawDescGZIP(), []int{3}
}

func (x *CheckinConfirmResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_rpkm66_backend_checkin_v1_checkin_proto protoreflect.FileDescriptor

var file_rpkm66_backend_checkin_v1_checkin_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x70, 0x6b, 0x6d, 0x36,
	0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x15, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x15,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x16, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0xff, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x74, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x2f, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x30, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x1b, 0x5a, 0x19, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm66_backend_checkin_v1_checkin_proto_rawDescOnce sync.Once
	file_rpkm66_backend_checkin_v1_checkin_proto_rawDescData = file_rpkm66_backend_checkin_v1_checkin_proto_rawDesc
)

func file_rpkm66_backend_checkin_v1_checkin_proto_rawDescGZIP() []byte {
	file_rpkm66_backend_checkin_v1_checkin_proto_rawDescOnce.Do(func() {
		file_rpkm66_backend_checkin_v1_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm66_backend_checkin_v1_checkin_proto_rawDescData)
	})
	return file_rpkm66_backend_checkin_v1_checkin_proto_rawDescData
}

var file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpkm66_backend_checkin_v1_checkin_proto_goTypes = []interface{}{
	(*CheckinVerifyRequest)(nil),   // 0: rpkm66.backend.checkin.v1.CheckinVerifyRequest
	(*CheckinVerifyResponse)(nil),  // 1: rpkm66.backend.checkin.v1.CheckinVerifyResponse
	(*CheckinConfirmRequest)(nil),  // 2: rpkm66.backend.checkin.v1.CheckinConfirmRequest
	(*CheckinConfirmResponse)(nil), // 3: rpkm66.backend.checkin.v1.CheckinConfirmResponse
}
var file_rpkm66_backend_checkin_v1_checkin_proto_depIdxs = []int32{
	0, // 0: rpkm66.backend.checkin.v1.CheckinService.CheckinVerify:input_type -> rpkm66.backend.checkin.v1.CheckinVerifyRequest
	2, // 1: rpkm66.backend.checkin.v1.CheckinService.CheckinConfirm:input_type -> rpkm66.backend.checkin.v1.CheckinConfirmRequest
	1, // 2: rpkm66.backend.checkin.v1.CheckinService.CheckinVerify:output_type -> rpkm66.backend.checkin.v1.CheckinVerifyResponse
	3, // 3: rpkm66.backend.checkin.v1.CheckinService.CheckinConfirm:output_type -> rpkm66.backend.checkin.v1.CheckinConfirmResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpkm66_backend_checkin_v1_checkin_proto_init() }
func file_rpkm66_backend_checkin_v1_checkin_proto_init() {
	if File_rpkm66_backend_checkin_v1_checkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinVerifyRequest); i {
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
		file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinVerifyResponse); i {
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
		file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinConfirmRequest); i {
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
		file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinConfirmResponse); i {
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
			RawDescriptor: file_rpkm66_backend_checkin_v1_checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm66_backend_checkin_v1_checkin_proto_goTypes,
		DependencyIndexes: file_rpkm66_backend_checkin_v1_checkin_proto_depIdxs,
		MessageInfos:      file_rpkm66_backend_checkin_v1_checkin_proto_msgTypes,
	}.Build()
	File_rpkm66_backend_checkin_v1_checkin_proto = out.File
	file_rpkm66_backend_checkin_v1_checkin_proto_rawDesc = nil
	file_rpkm66_backend_checkin_v1_checkin_proto_goTypes = nil
	file_rpkm66_backend_checkin_v1_checkin_proto_depIdxs = nil
}
