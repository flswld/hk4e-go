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
// source: EvtGrapplingHookMove.proto

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

type EvtGrapplingHookMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acceleration                 float32  `protobuf:"fixed32,1,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
	AnimatorStateIdList          []uint32 `protobuf:"varint,2,rep,packed,name=animator_state_id_list,json=animatorStateIdList,proto3" json:"animator_state_id_list,omitempty"`
	EntityId                     uint32   `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	NeedSetIsInAir               bool     `protobuf:"varint,13,opt,name=need_set_is_in_air,json=needSetIsInAir,proto3" json:"need_set_is_in_air,omitempty"`
	Speed                        float32  `protobuf:"fixed32,12,opt,name=speed,proto3" json:"speed,omitempty"`
	MaxSpeed                     float32  `protobuf:"fixed32,8,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty"`
	CheckAnimatorStateOnExitOnly bool     `protobuf:"varint,11,opt,name=check_animator_state_on_exit_only,json=checkAnimatorStateOnExitOnly,proto3" json:"check_animator_state_on_exit_only,omitempty"`
	OverrideCollider             string   `protobuf:"bytes,14,opt,name=override_collider,json=overrideCollider,proto3" json:"override_collider,omitempty"`
	TargetPos                    *Vector  `protobuf:"bytes,10,opt,name=target_pos,json=targetPos,proto3" json:"target_pos,omitempty"`
}

func (x *EvtGrapplingHookMove) Reset() {
	*x = EvtGrapplingHookMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtGrapplingHookMove_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtGrapplingHookMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtGrapplingHookMove) ProtoMessage() {}

func (x *EvtGrapplingHookMove) ProtoReflect() protoreflect.Message {
	mi := &file_EvtGrapplingHookMove_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtGrapplingHookMove.ProtoReflect.Descriptor instead.
func (*EvtGrapplingHookMove) Descriptor() ([]byte, []int) {
	return file_EvtGrapplingHookMove_proto_rawDescGZIP(), []int{0}
}

func (x *EvtGrapplingHookMove) GetAcceleration() float32 {
	if x != nil {
		return x.Acceleration
	}
	return 0
}

func (x *EvtGrapplingHookMove) GetAnimatorStateIdList() []uint32 {
	if x != nil {
		return x.AnimatorStateIdList
	}
	return nil
}

func (x *EvtGrapplingHookMove) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtGrapplingHookMove) GetNeedSetIsInAir() bool {
	if x != nil {
		return x.NeedSetIsInAir
	}
	return false
}

func (x *EvtGrapplingHookMove) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *EvtGrapplingHookMove) GetMaxSpeed() float32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *EvtGrapplingHookMove) GetCheckAnimatorStateOnExitOnly() bool {
	if x != nil {
		return x.CheckAnimatorStateOnExitOnly
	}
	return false
}

func (x *EvtGrapplingHookMove) GetOverrideCollider() string {
	if x != nil {
		return x.OverrideCollider
	}
	return ""
}

func (x *EvtGrapplingHookMove) GetTargetPos() *Vector {
	if x != nil {
		return x.TargetPos
	}
	return nil
}

var File_EvtGrapplingHookMove_proto protoreflect.FileDescriptor

var file_EvtGrapplingHookMove_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x74, 0x47, 0x72, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x48, 0x6f,
	0x6f, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x14, 0x45, 0x76, 0x74, 0x47, 0x72, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x48, 0x6f, 0x6f, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33,
	0x0a, 0x16, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x13,
	0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x12, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x61, 0x69, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e, 0x65,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x49, 0x73, 0x49, 0x6e, 0x41, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x47, 0x0a, 0x21, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x45, 0x78, 0x69, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtGrapplingHookMove_proto_rawDescOnce sync.Once
	file_EvtGrapplingHookMove_proto_rawDescData = file_EvtGrapplingHookMove_proto_rawDesc
)

func file_EvtGrapplingHookMove_proto_rawDescGZIP() []byte {
	file_EvtGrapplingHookMove_proto_rawDescOnce.Do(func() {
		file_EvtGrapplingHookMove_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtGrapplingHookMove_proto_rawDescData)
	})
	return file_EvtGrapplingHookMove_proto_rawDescData
}

var file_EvtGrapplingHookMove_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtGrapplingHookMove_proto_goTypes = []interface{}{
	(*EvtGrapplingHookMove)(nil), // 0: proto.EvtGrapplingHookMove
	(*Vector)(nil),               // 1: proto.Vector
}
var file_EvtGrapplingHookMove_proto_depIdxs = []int32{
	1, // 0: proto.EvtGrapplingHookMove.target_pos:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EvtGrapplingHookMove_proto_init() }
func file_EvtGrapplingHookMove_proto_init() {
	if File_EvtGrapplingHookMove_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtGrapplingHookMove_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtGrapplingHookMove); i {
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
			RawDescriptor: file_EvtGrapplingHookMove_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtGrapplingHookMove_proto_goTypes,
		DependencyIndexes: file_EvtGrapplingHookMove_proto_depIdxs,
		MessageInfos:      file_EvtGrapplingHookMove_proto_msgTypes,
	}.Build()
	File_EvtGrapplingHookMove_proto = out.File
	file_EvtGrapplingHookMove_proto_rawDesc = nil
	file_EvtGrapplingHookMove_proto_goTypes = nil
	file_EvtGrapplingHookMove_proto_depIdxs = nil
}