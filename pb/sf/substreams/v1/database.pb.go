// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: sf/substreams/v1/database.proto

package pbsubstreams

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TableChange_Operation int32

const (
	TableChange_UNSET  TableChange_Operation = 0
	TableChange_CREATE TableChange_Operation = 1
	TableChange_UPDATE TableChange_Operation = 2
	TableChange_DELETE TableChange_Operation = 3
)

// Enum value maps for TableChange_Operation.
var (
	TableChange_Operation_name = map[int32]string{
		0: "UNSET",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	TableChange_Operation_value = map[string]int32{
		"UNSET":  0,
		"CREATE": 1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x TableChange_Operation) Enum() *TableChange_Operation {
	p := new(TableChange_Operation)
	*p = x
	return p
}

func (x TableChange_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableChange_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_substreams_v1_database_proto_enumTypes[0].Descriptor()
}

func (TableChange_Operation) Type() protoreflect.EnumType {
	return &file_sf_substreams_v1_database_proto_enumTypes[0]
}

func (x TableChange_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableChange_Operation.Descriptor instead.
func (TableChange_Operation) EnumDescriptor() ([]byte, []int) {
	return file_sf_substreams_v1_database_proto_rawDescGZIP(), []int{1, 0}
}

type DatabaseChanges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableChanges []*TableChange `protobuf:"bytes,1,rep,name=tableChanges,proto3" json:"tableChanges,omitempty"`
}

func (x *DatabaseChanges) Reset() {
	*x = DatabaseChanges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseChanges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseChanges) ProtoMessage() {}

func (x *DatabaseChanges) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseChanges.ProtoReflect.Descriptor instead.
func (*DatabaseChanges) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseChanges) GetTableChanges() []*TableChange {
	if x != nil {
		return x.TableChanges
	}
	return nil
}

type TableChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table     string                `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Pk        string                `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	BlockNum  uint64                `protobuf:"varint,3,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	Ordinal   uint64                `protobuf:"varint,4,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Operation TableChange_Operation `protobuf:"varint,5,opt,name=operation,proto3,enum=sf.substreams.v1.TableChange_Operation" json:"operation,omitempty"`
	Fields    []*Field              `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *TableChange) Reset() {
	*x = TableChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableChange) ProtoMessage() {}

func (x *TableChange) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableChange.ProtoReflect.Descriptor instead.
func (*TableChange) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *TableChange) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *TableChange) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *TableChange) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *TableChange) GetOrdinal() uint64 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *TableChange) GetOperation() TableChange_Operation {
	if x != nil {
		return x.Operation
	}
	return TableChange_UNSET
}

func (x *TableChange) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NewValue string `protobuf:"bytes,2,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	OldValue string `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_database_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *Field) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

var File_sf_substreams_v1_database_proto protoreflect.FileDescriptor

var file_sf_substreams_v1_database_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x0b, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x66, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x3a,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0x55, 0x0a, 0x05, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sf_substreams_v1_database_proto_rawDescOnce sync.Once
	file_sf_substreams_v1_database_proto_rawDescData = file_sf_substreams_v1_database_proto_rawDesc
)

func file_sf_substreams_v1_database_proto_rawDescGZIP() []byte {
	file_sf_substreams_v1_database_proto_rawDescOnce.Do(func() {
		file_sf_substreams_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_v1_database_proto_rawDescData)
	})
	return file_sf_substreams_v1_database_proto_rawDescData
}

var file_sf_substreams_v1_database_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sf_substreams_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sf_substreams_v1_database_proto_goTypes = []interface{}{
	(TableChange_Operation)(0), // 0: sf.substreams.v1.TableChange.Operation
	(*DatabaseChanges)(nil),    // 1: sf.substreams.v1.DatabaseChanges
	(*TableChange)(nil),        // 2: sf.substreams.v1.TableChange
	(*Field)(nil),              // 3: sf.substreams.v1.Field
}
var file_sf_substreams_v1_database_proto_depIdxs = []int32{
	2, // 0: sf.substreams.v1.DatabaseChanges.tableChanges:type_name -> sf.substreams.v1.TableChange
	0, // 1: sf.substreams.v1.TableChange.operation:type_name -> sf.substreams.v1.TableChange.Operation
	3, // 2: sf.substreams.v1.TableChange.fields:type_name -> sf.substreams.v1.Field
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sf_substreams_v1_database_proto_init() }
func file_sf_substreams_v1_database_proto_init() {
	if File_sf_substreams_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseChanges); i {
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
		file_sf_substreams_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableChange); i {
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
		file_sf_substreams_v1_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
			RawDescriptor: file_sf_substreams_v1_database_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_v1_database_proto_goTypes,
		DependencyIndexes: file_sf_substreams_v1_database_proto_depIdxs,
		EnumInfos:         file_sf_substreams_v1_database_proto_enumTypes,
		MessageInfos:      file_sf_substreams_v1_database_proto_msgTypes,
	}.Build()
	File_sf_substreams_v1_database_proto = out.File
	file_sf_substreams_v1_database_proto_rawDesc = nil
	file_sf_substreams_v1_database_proto_goTypes = nil
	file_sf_substreams_v1_database_proto_depIdxs = nil
}
