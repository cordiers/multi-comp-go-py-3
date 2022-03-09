// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/provisioning/blockstorage/agent.proto

package blockstorage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateEmptyBlockStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmptyBlockStorageRequest) Reset()         { *m = CreateEmptyBlockStorageRequest{} }
func (m *CreateEmptyBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyBlockStorageRequest) ProtoMessage()    {}
func (*CreateEmptyBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{0}
}

func (m *CreateEmptyBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Unmarshal(m, b)
}
func (m *CreateEmptyBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Marshal(b, m, deterministic)
}
func (m *CreateEmptyBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyBlockStorageRequest.Merge(m, src)
}
func (m *CreateEmptyBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Size(m)
}
func (m *CreateEmptyBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyBlockStorageRequest proto.InternalMessageInfo

func (m *CreateEmptyBlockStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEmptyBlockStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type CreateEmptyBlockStorageResponse struct {
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmptyBlockStorageResponse) Reset()         { *m = CreateEmptyBlockStorageResponse{} }
func (m *CreateEmptyBlockStorageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyBlockStorageResponse) ProtoMessage()    {}
func (*CreateEmptyBlockStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{1}
}

func (m *CreateEmptyBlockStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Unmarshal(m, b)
}
func (m *CreateEmptyBlockStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Marshal(b, m, deterministic)
}
func (m *CreateEmptyBlockStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyBlockStorageResponse.Merge(m, src)
}
func (m *CreateEmptyBlockStorageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Size(m)
}
func (m *CreateEmptyBlockStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyBlockStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyBlockStorageResponse proto.InternalMessageInfo

func (m *CreateEmptyBlockStorageResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FetchBlockStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	SourceUrl            string   `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockStorageRequest) Reset()         { *m = FetchBlockStorageRequest{} }
func (m *FetchBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*FetchBlockStorageRequest) ProtoMessage()    {}
func (*FetchBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{2}
}

func (m *FetchBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockStorageRequest.Unmarshal(m, b)
}
func (m *FetchBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockStorageRequest.Marshal(b, m, deterministic)
}
func (m *FetchBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockStorageRequest.Merge(m, src)
}
func (m *FetchBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_FetchBlockStorageRequest.Size(m)
}
func (m *FetchBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockStorageRequest proto.InternalMessageInfo

func (m *FetchBlockStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FetchBlockStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *FetchBlockStorageRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

type FetchBlockStorageResponse struct {
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockStorageResponse) Reset()         { *m = FetchBlockStorageResponse{} }
func (m *FetchBlockStorageResponse) String() string { return proto.CompactTextString(m) }
func (*FetchBlockStorageResponse) ProtoMessage()    {}
func (*FetchBlockStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{3}
}

func (m *FetchBlockStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockStorageResponse.Unmarshal(m, b)
}
func (m *FetchBlockStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockStorageResponse.Marshal(b, m, deterministic)
}
func (m *FetchBlockStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockStorageResponse.Merge(m, src)
}
func (m *FetchBlockStorageResponse) XXX_Size() int {
	return xxx_messageInfo_FetchBlockStorageResponse.Size(m)
}
func (m *FetchBlockStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockStorageResponse proto.InternalMessageInfo

func (m *FetchBlockStorageResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DeleteBlockStorageRequest struct {
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlockStorageRequest) Reset()         { *m = DeleteBlockStorageRequest{} }
func (m *DeleteBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockStorageRequest) ProtoMessage()    {}
func (*DeleteBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{4}
}

func (m *DeleteBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlockStorageRequest.Unmarshal(m, b)
}
func (m *DeleteBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlockStorageRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlockStorageRequest.Merge(m, src)
}
func (m *DeleteBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlockStorageRequest.Size(m)
}
func (m *DeleteBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlockStorageRequest proto.InternalMessageInfo

func (m *DeleteBlockStorageRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ResizeBlockStorageRequest struct {
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeBlockStorageRequest) Reset()         { *m = ResizeBlockStorageRequest{} }
func (m *ResizeBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*ResizeBlockStorageRequest) ProtoMessage()    {}
func (*ResizeBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8967c0e56d6b7ff, []int{5}
}

func (m *ResizeBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeBlockStorageRequest.Unmarshal(m, b)
}
func (m *ResizeBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeBlockStorageRequest.Marshal(b, m, deterministic)
}
func (m *ResizeBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeBlockStorageRequest.Merge(m, src)
}
func (m *ResizeBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_ResizeBlockStorageRequest.Size(m)
}
func (m *ResizeBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeBlockStorageRequest proto.InternalMessageInfo

func (m *ResizeBlockStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *ResizeBlockStorageRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateEmptyBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.blockstorage.CreateEmptyBlockStorageRequest")
	proto.RegisterType((*CreateEmptyBlockStorageResponse)(nil), "n0stack.internal.n0core.provisioning.blockstorage.CreateEmptyBlockStorageResponse")
	proto.RegisterType((*FetchBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.blockstorage.FetchBlockStorageRequest")
	proto.RegisterType((*FetchBlockStorageResponse)(nil), "n0stack.internal.n0core.provisioning.blockstorage.FetchBlockStorageResponse")
	proto.RegisterType((*DeleteBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.blockstorage.DeleteBlockStorageRequest")
	proto.RegisterType((*ResizeBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.blockstorage.ResizeBlockStorageRequest")
}

func init() {
	proto.RegisterFile("pkg/api/provisioning/blockstorage/agent.proto", fileDescriptor_e8967c0e56d6b7ff)
}

var fileDescriptor_e8967c0e56d6b7ff = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x4f, 0xdb, 0x30,
	0x14, 0xc7, 0x97, 0xad, 0x9d, 0x54, 0xdf, 0x66, 0x4d, 0x5b, 0x9a, 0x69, 0x5b, 0x95, 0x53, 0x2f,
	0xb3, 0xbb, 0x4d, 0x3b, 0xed, 0x44, 0xa1, 0x1c, 0x80, 0x03, 0xa4, 0xe2, 0xc2, 0x05, 0x39, 0xd6,
	0x23, 0xb5, 0x9a, 0xda, 0xc1, 0x76, 0x2a, 0x15, 0xf1, 0x19, 0xf8, 0x2c, 0x7c, 0x06, 0x3e, 0x19,
	0x4a, 0x5c, 0x50, 0xa4, 0x24, 0x45, 0x94, 0x9e, 0x62, 0x5b, 0xfa, 0xff, 0x7f, 0x4f, 0xef, 0xfd,
	0x5f, 0xd0, 0xaf, 0x6c, 0x9e, 0x50, 0x96, 0x09, 0x9a, 0x69, 0xb5, 0x14, 0x46, 0x28, 0x29, 0x64,
	0x42, 0xe3, 0x54, 0xf1, 0xb9, 0xb1, 0x4a, 0xb3, 0x04, 0x28, 0x4b, 0x40, 0x5a, 0x92, 0x69, 0x65,
	0x15, 0xfe, 0x2d, 0x47, 0xc6, 0x32, 0x3e, 0x27, 0x42, 0x5a, 0xd0, 0x92, 0xa5, 0x44, 0x8e, 0xb8,
	0xd2, 0x40, 0xaa, 0x72, 0x52, 0x95, 0x07, 0xdf, 0x12, 0xa5, 0x92, 0x14, 0x68, 0x69, 0x10, 0xe7,
	0x57, 0x14, 0x16, 0x99, 0x5d, 0x39, 0xbf, 0xf0, 0x08, 0xfd, 0xd8, 0xd7, 0xc0, 0x2c, 0x4c, 0x8a,
	0xc7, 0x71, 0xa1, 0x9b, 0x3a, 0x5d, 0x04, 0xd7, 0x39, 0x18, 0x8b, 0x31, 0xea, 0x48, 0xb6, 0x00,
	0xdf, 0x1b, 0x78, 0xc3, 0x5e, 0x54, 0x9e, 0xf1, 0x67, 0xd4, 0x8d, 0x57, 0x16, 0x8c, 0xff, 0x7e,
	0xe0, 0x0d, 0x3b, 0x91, 0xbb, 0x84, 0xff, 0xd0, 0xcf, 0x56, 0x2f, 0x93, 0x29, 0x69, 0xa0, 0x30,
	0xcb, 0x98, 0x9d, 0xf9, 0x1f, 0x9c, 0x59, 0x71, 0x0e, 0x39, 0xf2, 0x0f, 0xc1, 0xf2, 0xd9, 0x9b,
	0xe0, 0xf8, 0x3b, 0x42, 0x46, 0xe5, 0x9a, 0xc3, 0x65, 0xae, 0xd3, 0xb5, 0x7f, 0xcf, 0xbd, 0x9c,
	0xeb, 0x34, 0xa4, 0xa8, 0xdf, 0x00, 0xd9, 0x50, 0x15, 0x45, 0xfd, 0x03, 0x48, 0xc1, 0x42, 0x4b,
	0x59, 0x35, 0xc1, 0x04, 0xf5, 0x23, 0x30, 0xe2, 0xa6, 0x51, 0xd0, 0x5c, 0x73, 0x83, 0xcd, 0x9f,
	0xbb, 0x2e, 0xf2, 0xab, 0x0e, 0x7b, 0xc5, 0xf0, 0xa7, 0xa0, 0x97, 0x82, 0x03, 0x7e, 0xf0, 0xd0,
	0xd7, 0x96, 0x16, 0xe3, 0x33, 0xf2, 0xea, 0x68, 0x90, 0xcd, 0xa3, 0x0f, 0xa2, 0x5d, 0x5a, 0xba,
	0x5e, 0x87, 0xef, 0xf0, 0xbd, 0x87, 0x3e, 0xd5, 0x66, 0x81, 0x8f, 0xb7, 0x60, 0xb5, 0xc5, 0x26,
	0x38, 0xd9, 0x8d, 0xd9, 0x73, 0xc9, 0xb7, 0x08, 0xd7, 0xc3, 0x80, 0xb7, 0xa1, 0xb4, 0x66, 0x2a,
	0xf8, 0x42, 0xdc, 0x9e, 0x92, 0xa7, 0x3d, 0x25, 0x65, 0x13, 0x1d, 0xbd, 0x9e, 0xac, 0xad, 0xe8,
	0xad, 0x01, 0x6d, 0xa7, 0x8f, 0xa3, 0x8b, 0xd3, 0x44, 0xd8, 0x59, 0x1e, 0x13, 0xae, 0x16, 0x74,
	0xcd, 0xac, 0x7c, 0x0b, 0x24, 0x7d, 0xf1, 0x27, 0xf6, 0xbf, 0x7a, 0x89, 0x3f, 0x96, 0x94, 0xbf,
	0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x74, 0x0d, 0x17, 0xfd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockStorageAgentServiceClient is the client API for BlockStorageAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockStorageAgentServiceClient interface {
	CreateEmptyBlockStorage(ctx context.Context, in *CreateEmptyBlockStorageRequest, opts ...grpc.CallOption) (*CreateEmptyBlockStorageResponse, error)
	FetchBlockStorage(ctx context.Context, in *FetchBlockStorageRequest, opts ...grpc.CallOption) (*FetchBlockStorageResponse, error)
	DeleteBlockStorage(ctx context.Context, in *DeleteBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ResizeBlockStorage(ctx context.Context, in *ResizeBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type blockStorageAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockStorageAgentServiceClient(cc *grpc.ClientConn) BlockStorageAgentServiceClient {
	return &blockStorageAgentServiceClient{cc}
}

func (c *blockStorageAgentServiceClient) CreateEmptyBlockStorage(ctx context.Context, in *CreateEmptyBlockStorageRequest, opts ...grpc.CallOption) (*CreateEmptyBlockStorageResponse, error) {
	out := new(CreateEmptyBlockStorageResponse)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/CreateEmptyBlockStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStorageAgentServiceClient) FetchBlockStorage(ctx context.Context, in *FetchBlockStorageRequest, opts ...grpc.CallOption) (*FetchBlockStorageResponse, error) {
	out := new(FetchBlockStorageResponse)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/FetchBlockStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStorageAgentServiceClient) DeleteBlockStorage(ctx context.Context, in *DeleteBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/DeleteBlockStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStorageAgentServiceClient) ResizeBlockStorage(ctx context.Context, in *ResizeBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/ResizeBlockStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockStorageAgentServiceServer is the server API for BlockStorageAgentService service.
type BlockStorageAgentServiceServer interface {
	CreateEmptyBlockStorage(context.Context, *CreateEmptyBlockStorageRequest) (*CreateEmptyBlockStorageResponse, error)
	FetchBlockStorage(context.Context, *FetchBlockStorageRequest) (*FetchBlockStorageResponse, error)
	DeleteBlockStorage(context.Context, *DeleteBlockStorageRequest) (*empty.Empty, error)
	ResizeBlockStorage(context.Context, *ResizeBlockStorageRequest) (*empty.Empty, error)
}

// UnimplementedBlockStorageAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockStorageAgentServiceServer struct {
}

func (*UnimplementedBlockStorageAgentServiceServer) CreateEmptyBlockStorage(ctx context.Context, req *CreateEmptyBlockStorageRequest) (*CreateEmptyBlockStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmptyBlockStorage not implemented")
}
func (*UnimplementedBlockStorageAgentServiceServer) FetchBlockStorage(ctx context.Context, req *FetchBlockStorageRequest) (*FetchBlockStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBlockStorage not implemented")
}
func (*UnimplementedBlockStorageAgentServiceServer) DeleteBlockStorage(ctx context.Context, req *DeleteBlockStorageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlockStorage not implemented")
}
func (*UnimplementedBlockStorageAgentServiceServer) ResizeBlockStorage(ctx context.Context, req *ResizeBlockStorageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeBlockStorage not implemented")
}

func RegisterBlockStorageAgentServiceServer(s *grpc.Server, srv BlockStorageAgentServiceServer) {
	s.RegisterService(&_BlockStorageAgentService_serviceDesc, srv)
}

func _BlockStorageAgentService_CreateEmptyBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmptyBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).CreateEmptyBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/CreateEmptyBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).CreateEmptyBlockStorage(ctx, req.(*CreateEmptyBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStorageAgentService_FetchBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).FetchBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/FetchBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).FetchBlockStorage(ctx, req.(*FetchBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStorageAgentService_DeleteBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).DeleteBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/DeleteBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).DeleteBlockStorage(ctx, req.(*DeleteBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStorageAgentService_ResizeBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).ResizeBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService/ResizeBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).ResizeBlockStorage(ctx, req.(*ResizeBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockStorageAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.internal.n0core.provisioning.blockstorage.BlockStorageAgentService",
	HandlerType: (*BlockStorageAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmptyBlockStorage",
			Handler:    _BlockStorageAgentService_CreateEmptyBlockStorage_Handler,
		},
		{
			MethodName: "FetchBlockStorage",
			Handler:    _BlockStorageAgentService_FetchBlockStorage_Handler,
		},
		{
			MethodName: "DeleteBlockStorage",
			Handler:    _BlockStorageAgentService_DeleteBlockStorage_Handler,
		},
		{
			MethodName: "ResizeBlockStorage",
			Handler:    _BlockStorageAgentService_ResizeBlockStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/provisioning/blockstorage/agent.proto",
}
