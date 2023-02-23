/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proto_wrapper_api

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TripleRequestWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hessian4
	// json
	SerializeType string   `protobuf:"bytes,1,opt,name=serializeType,proto3" json:"serializeType,omitempty"`
	Args          [][]byte `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	ArgTypes      []string `protobuf:"bytes,3,rep,name=argTypes,proto3" json:"argTypes,omitempty"`
}

func (x *TripleRequestWrapper) Reset() {
	*x = TripleRequestWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_triple_wrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripleRequestWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripleRequestWrapper) ProtoMessage() {}

func (x *TripleRequestWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_triple_wrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripleRequestWrapper.ProtoReflect.Descriptor instead.
func (*TripleRequestWrapper) Descriptor() ([]byte, []int) {
	return file_triple_wrapper_proto_rawDescGZIP(), []int{0}
}

func (x *TripleRequestWrapper) GetSerializeType() string {
	if x != nil {
		return x.SerializeType
	}
	return ""
}

func (x *TripleRequestWrapper) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *TripleRequestWrapper) GetArgTypes() []string {
	if x != nil {
		return x.ArgTypes
	}
	return nil
}

type TripleResponseWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerializeType string `protobuf:"bytes,1,opt,name=serializeType,proto3" json:"serializeType,omitempty"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TripleResponseWrapper) Reset() {
	*x = TripleResponseWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_triple_wrapper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripleResponseWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripleResponseWrapper) ProtoMessage() {}

func (x *TripleResponseWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_triple_wrapper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripleResponseWrapper.ProtoReflect.Descriptor instead.
func (*TripleResponseWrapper) Descriptor() ([]byte, []int) {
	return file_triple_wrapper_proto_rawDescGZIP(), []int{1}
}

func (x *TripleResponseWrapper) GetSerializeType() string {
	if x != nil {
		return x.SerializeType
	}
	return ""
}

func (x *TripleResponseWrapper) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TripleResponseWrapper) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_triple_wrapper_proto protoreflect.FileDescriptor

var file_triple_wrapper_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x14, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x15, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_triple_wrapper_proto_rawDescOnce sync.Once
	file_triple_wrapper_proto_rawDescData = file_triple_wrapper_proto_rawDesc
)

func file_triple_wrapper_proto_rawDescGZIP() []byte {
	file_triple_wrapper_proto_rawDescOnce.Do(func() {
		file_triple_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_triple_wrapper_proto_rawDescData)
	})
	return file_triple_wrapper_proto_rawDescData
}

var file_triple_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_triple_wrapper_proto_goTypes = []interface{}{
	(*TripleRequestWrapper)(nil),  // 0: proto.TripleRequestWrapper
	(*TripleResponseWrapper)(nil), // 1: proto.TripleResponseWrapper
}
var file_triple_wrapper_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_triple_wrapper_proto_init() }
func file_triple_wrapper_proto_init() {
	if File_triple_wrapper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_triple_wrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripleRequestWrapper); i {
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
		file_triple_wrapper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripleResponseWrapper); i {
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
			RawDescriptor: file_triple_wrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_triple_wrapper_proto_goTypes,
		DependencyIndexes: file_triple_wrapper_proto_depIdxs,
		MessageInfos:      file_triple_wrapper_proto_msgTypes,
	}.Build()
	File_triple_wrapper_proto = out.File
	file_triple_wrapper_proto_rawDesc = nil
	file_triple_wrapper_proto_goTypes = nil
	file_triple_wrapper_proto_depIdxs = nil
}
