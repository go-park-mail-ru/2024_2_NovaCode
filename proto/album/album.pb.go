// protoc --go_out=proto/album/. --go-grpc_out=proto/album/.
// proto/album/album.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: proto/album/album.proto

package albumService

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ReleaseDate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Image       string               `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	ArtistId    uint64               `protobuf:"varint,5,opt,name=artist_id,json=artistId,proto3" json:"artist_id,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	mi := &file_proto_album_album_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_proto_album_album_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_proto_album_album_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Album) GetReleaseDate() *timestamp.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *Album) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Album) GetArtistId() uint64 {
	if x != nil {
		return x.ArtistId
	}
	return 0
}

func (x *Album) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Album) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FindByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindByIDRequest) Reset() {
	*x = FindByIDRequest{}
	mi := &file_proto_album_album_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIDRequest) ProtoMessage() {}

func (x *FindByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_album_album_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIDRequest.ProtoReflect.Descriptor instead.
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_album_album_proto_rawDescGZIP(), []int{1}
}

func (x *FindByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *FindByIDResponse) Reset() {
	*x = FindByIDResponse{}
	mi := &file_proto_album_album_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIDResponse) ProtoMessage() {}

func (x *FindByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_album_album_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIDResponse.ProtoReflect.Descriptor instead.
func (*FindByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_album_album_proto_rawDescGZIP(), []int{2}
}

func (x *FindByIDResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

var File_proto_album_album_proto protoreflect.FileDescriptor

var file_proto_album_album_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x21,
	0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3d, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x32, 0x59, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e,
	0x3b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_album_album_proto_rawDescOnce sync.Once
	file_proto_album_album_proto_rawDescData = file_proto_album_album_proto_rawDesc
)

func file_proto_album_album_proto_rawDescGZIP() []byte {
	file_proto_album_album_proto_rawDescOnce.Do(func() {
		file_proto_album_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_album_album_proto_rawDescData)
	})
	return file_proto_album_album_proto_rawDescData
}

var file_proto_album_album_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_album_album_proto_goTypes = []any{
	(*Album)(nil),               // 0: albumService.Album
	(*FindByIDRequest)(nil),     // 1: albumService.FindByIDRequest
	(*FindByIDResponse)(nil),    // 2: albumService.FindByIDResponse
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_album_album_proto_depIdxs = []int32{
	3, // 0: albumService.Album.release_date:type_name -> google.protobuf.Timestamp
	3, // 1: albumService.Album.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: albumService.Album.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: albumService.FindByIDResponse.album:type_name -> albumService.Album
	1, // 4: albumService.AlbumService.FindByID:input_type -> albumService.FindByIDRequest
	2, // 5: albumService.AlbumService.FindByID:output_type -> albumService.FindByIDResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_album_album_proto_init() }
func file_proto_album_album_proto_init() {
	if File_proto_album_album_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_album_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_album_album_proto_goTypes,
		DependencyIndexes: file_proto_album_album_proto_depIdxs,
		MessageInfos:      file_proto_album_album_proto_msgTypes,
	}.Build()
	File_proto_album_album_proto = out.File
	file_proto_album_album_proto_rawDesc = nil
	file_proto_album_album_proto_goTypes = nil
	file_proto_album_album_proto_depIdxs = nil
}
