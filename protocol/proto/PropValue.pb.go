// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: PropValue.proto

package proto

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

type PropValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Val  int64  `protobuf:"varint,4,opt,name=val,proto3" json:"val,omitempty"`
	// Types that are assignable to Value:
	//
	//	*PropValue_Ival
	//	*PropValue_Fval
	Value isPropValue_Value `protobuf_oneof:"value"`
}

func (x *PropValue) Reset() {
	*x = PropValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PropValue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropValue) ProtoMessage() {}

func (x *PropValue) ProtoReflect() protoreflect.Message {
	mi := &file_PropValue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropValue.ProtoReflect.Descriptor instead.
func (*PropValue) Descriptor() ([]byte, []int) {
	return file_PropValue_proto_rawDescGZIP(), []int{0}
}

func (x *PropValue) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PropValue) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

func (m *PropValue) GetValue() isPropValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *PropValue) GetIval() int64 {
	if x, ok := x.GetValue().(*PropValue_Ival); ok {
		return x.Ival
	}
	return 0
}

func (x *PropValue) GetFval() float32 {
	if x, ok := x.GetValue().(*PropValue_Fval); ok {
		return x.Fval
	}
	return 0
}

type isPropValue_Value interface {
	isPropValue_Value()
}

type PropValue_Ival struct {
	Ival int64 `protobuf:"varint,2,opt,name=ival,proto3,oneof"`
}

type PropValue_Fval struct {
	Fval float32 `protobuf:"fixed32,3,opt,name=fval,proto3,oneof"`
}

func (*PropValue_Ival) isPropValue_Value() {}

func (*PropValue_Fval) isPropValue_Value() {}

var File_PropValue_proto protoreflect.FileDescriptor

var file_PropValue_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x70,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x04, 0x69,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x69, 0x76, 0x61,
	0x6c, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x00, 0x52, 0x04, 0x66, 0x76, 0x61, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropValue_proto_rawDescOnce sync.Once
	file_PropValue_proto_rawDescData = file_PropValue_proto_rawDesc
)

func file_PropValue_proto_rawDescGZIP() []byte {
	file_PropValue_proto_rawDescOnce.Do(func() {
		file_PropValue_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropValue_proto_rawDescData)
	})
	return file_PropValue_proto_rawDescData
}

var file_PropValue_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PropValue_proto_goTypes = []interface{}{
	(*PropValue)(nil), // 0: proto.PropValue
}
var file_PropValue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PropValue_proto_init() }
func file_PropValue_proto_init() {
	if File_PropValue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PropValue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropValue); i {
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
	file_PropValue_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PropValue_Ival)(nil),
		(*PropValue_Fval)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PropValue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropValue_proto_goTypes,
		DependencyIndexes: file_PropValue_proto_depIdxs,
		MessageInfos:      file_PropValue_proto_msgTypes,
	}.Build()
	File_PropValue_proto = out.File
	file_PropValue_proto_rawDesc = nil
	file_PropValue_proto_goTypes = nil
	file_PropValue_proto_depIdxs = nil
}