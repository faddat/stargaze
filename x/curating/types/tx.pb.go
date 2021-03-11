// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/curating/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgPostResponse struct {
}

func (m *MsgPostResponse) Reset()         { *m = MsgPostResponse{} }
func (m *MsgPostResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPostResponse) ProtoMessage()    {}
func (*MsgPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_307f601eca63d732, []int{0}
}
func (m *MsgPostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostResponse.Merge(m, src)
}
func (m *MsgPostResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostResponse proto.InternalMessageInfo

type MsgUpvoteResponse struct {
}

func (m *MsgUpvoteResponse) Reset()         { *m = MsgUpvoteResponse{} }
func (m *MsgUpvoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpvoteResponse) ProtoMessage()    {}
func (*MsgUpvoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_307f601eca63d732, []int{1}
}
func (m *MsgUpvoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpvoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpvoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpvoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpvoteResponse.Merge(m, src)
}
func (m *MsgUpvoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpvoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpvoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpvoteResponse proto.InternalMessageInfo

type MsgPost struct {
	VendorID        uint32  `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id" yaml:"vendor_id"`
	PostID          string  `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id" yaml:"post_id"`
	Creator         string  `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	RewardAccount   string  `protobuf:"bytes,4,opt,name=reward_account,json=rewardAccount,proto3" json:"reward_account,omitempty" yaml:"reward_account"`
	Body            string  `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
	ChainID         string  `protobuf:"bytes,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id" yaml:"chain_id"`
	ContractAddress string  `protobuf:"bytes,7,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	Metadata        string  `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty" yaml:"metadata"`
	ParentID        *PostID `protobuf:"bytes,9,opt,name=parent_id,json=parentId,proto3,customtype=PostID" json:"parent_id" yaml:"parent_id"`
}

func (m *MsgPost) Reset()         { *m = MsgPost{} }
func (m *MsgPost) String() string { return proto.CompactTextString(m) }
func (*MsgPost) ProtoMessage()    {}
func (*MsgPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_307f601eca63d732, []int{2}
}
func (m *MsgPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPost.Merge(m, src)
}
func (m *MsgPost) XXX_Size() int {
	return m.Size()
}
func (m *MsgPost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPost proto.InternalMessageInfo

func (m *MsgPost) GetVendorID() uint32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgPost) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPost) GetRewardAccount() string {
	if m != nil {
		return m.RewardAccount
	}
	return ""
}

func (m *MsgPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgPost) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *MsgPost) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgPost) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type MsgUpvote struct {
	VendorID      uint32 `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id" yaml:"vendor_id"`
	PostID        string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id" yaml:"post_id"`
	Curator       string `protobuf:"bytes,3,opt,name=curator,proto3" json:"curator,omitempty" yaml:"curator"`
	RewardAccount string `protobuf:"bytes,4,opt,name=reward_account,json=rewardAccount,proto3" json:"reward_account,omitempty" yaml:"reward_account"`
	VoteNum       int32  `protobuf:"varint,5,opt,name=vote_num,json=voteNum,proto3" json:"vote_num,omitempty"`
}

func (m *MsgUpvote) Reset()         { *m = MsgUpvote{} }
func (m *MsgUpvote) String() string { return proto.CompactTextString(m) }
func (*MsgUpvote) ProtoMessage()    {}
func (*MsgUpvote) Descriptor() ([]byte, []int) {
	return fileDescriptor_307f601eca63d732, []int{3}
}
func (m *MsgUpvote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpvote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpvote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpvote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpvote.Merge(m, src)
}
func (m *MsgUpvote) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpvote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpvote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpvote proto.InternalMessageInfo

func (m *MsgUpvote) GetVendorID() uint32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgUpvote) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgUpvote) GetCurator() string {
	if m != nil {
		return m.Curator
	}
	return ""
}

func (m *MsgUpvote) GetRewardAccount() string {
	if m != nil {
		return m.RewardAccount
	}
	return ""
}

func (m *MsgUpvote) GetVoteNum() int32 {
	if m != nil {
		return m.VoteNum
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgPostResponse)(nil), "stargaze.curating.v1beta1.MsgPostResponse")
	proto.RegisterType((*MsgUpvoteResponse)(nil), "stargaze.curating.v1beta1.MsgUpvoteResponse")
	proto.RegisterType((*MsgPost)(nil), "stargaze.curating.v1beta1.MsgPost")
	proto.RegisterType((*MsgUpvote)(nil), "stargaze.curating.v1beta1.MsgUpvote")
}

func init() {
	proto.RegisterFile("stargaze/curating/v1beta1/tx.proto", fileDescriptor_307f601eca63d732)
}

var fileDescriptor_307f601eca63d732 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6a, 0xd5, 0x40,
	0x14, 0xc7, 0x9b, 0xf6, 0xf6, 0x26, 0x77, 0xa4, 0xbd, 0xed, 0x54, 0x31, 0xad, 0x90, 0x29, 0xa3,
	0x60, 0x91, 0x36, 0xa1, 0x76, 0x57, 0x37, 0x6d, 0x15, 0xe1, 0x0a, 0x57, 0x4a, 0xc0, 0x22, 0x22,
	0x5c, 0xe6, 0x26, 0x43, 0x7a, 0xa1, 0xc9, 0x84, 0xcc, 0xa4, 0x1f, 0x3e, 0x85, 0x8f, 0xe0, 0xda,
	0x57, 0xf0, 0x05, 0xba, 0xec, 0x52, 0x5c, 0x04, 0x49, 0x37, 0xe2, 0x32, 0x4f, 0x20, 0x33, 0x93,
	0xe4, 0xa2, 0xf8, 0xb5, 0x70, 0xe1, 0xee, 0xe4, 0x7f, 0x7e, 0xff, 0x93, 0x4c, 0xce, 0x39, 0x03,
	0x30, 0x17, 0x24, 0x8b, 0xc8, 0x1b, 0xea, 0x05, 0x79, 0x46, 0xc4, 0x24, 0x89, 0xbc, 0xd3, 0xed,
	0x31, 0x15, 0x64, 0xdb, 0x13, 0xe7, 0x6e, 0x9a, 0x31, 0xc1, 0xe0, 0x6a, 0xc3, 0xb8, 0x0d, 0xe3,
	0xd6, 0xcc, 0xda, 0xcd, 0x88, 0x45, 0x4c, 0x51, 0x9e, 0x8c, 0xb4, 0x01, 0x2f, 0x83, 0xfe, 0x90,
	0x47, 0x87, 0x8c, 0x0b, 0x9f, 0xf2, 0x94, 0x25, 0x9c, 0xe2, 0x15, 0xb0, 0x3c, 0xe4, 0xd1, 0x8b,
	0xf4, 0x94, 0x09, 0xda, 0x8a, 0x57, 0x1d, 0x60, 0xd6, 0x20, 0x7c, 0x06, 0x7a, 0xa7, 0x34, 0x09,
	0x59, 0x36, 0x9a, 0x84, 0xb6, 0xb1, 0x6e, 0x6c, 0x2c, 0x1c, 0x6c, 0x95, 0x05, 0xb2, 0x8e, 0x94,
	0x38, 0x78, 0xf2, 0xb5, 0x40, 0x53, 0xa0, 0x2a, 0xd0, 0xd2, 0x05, 0x89, 0x4f, 0x76, 0x71, 0x2b,
	0x61, 0xdf, 0xd2, 0xf1, 0x20, 0x84, 0x7b, 0xc0, 0x4c, 0x19, 0x17, 0xb2, 0xd2, 0xec, 0xba, 0xb1,
	0xd1, 0x3b, 0xb8, 0x5f, 0x16, 0xa8, 0x2b, 0x5f, 0xa3, 0xea, 0x34, 0xc9, 0xaa, 0x40, 0x8b, 0xba,
	0x4a, 0x2d, 0x60, 0xbf, 0x2b, 0xa3, 0x41, 0x08, 0x37, 0x81, 0x19, 0x64, 0x94, 0x08, 0x96, 0xd9,
	0x73, 0xaa, 0x02, 0x9c, 0xc2, 0x75, 0x02, 0xfb, 0x0d, 0x02, 0xf7, 0xc0, 0x62, 0x46, 0xcf, 0x48,
	0x16, 0x8e, 0x48, 0x10, 0xb0, 0x3c, 0x11, 0x76, 0x47, 0x99, 0x56, 0xab, 0x02, 0xdd, 0xd2, 0xa6,
	0xef, 0xf3, 0xd8, 0x5f, 0xd0, 0xc2, 0xbe, 0x7e, 0x86, 0x77, 0x41, 0x67, 0xcc, 0xc2, 0x0b, 0x7b,
	0x5e, 0xf9, 0xfa, 0x55, 0x81, 0x6e, 0x68, 0x9f, 0x54, 0xb1, 0xaf, 0x92, 0x70, 0x00, 0xac, 0xe0,
	0x98, 0x4c, 0x12, 0x79, 0xae, 0xae, 0x02, 0xdd, 0xcb, 0x02, 0x19, 0x65, 0x81, 0xcc, 0xc7, 0x52,
	0x57, 0x87, 0x6b, 0x91, 0xaa, 0x40, 0xfd, 0xfa, 0x83, 0x6b, 0x45, 0x7e, 0xb1, 0xe2, 0x42, 0xf8,
	0x14, 0x2c, 0x05, 0x2c, 0x11, 0x19, 0x09, 0xc4, 0x88, 0x84, 0x61, 0x46, 0x39, 0xb7, 0x4d, 0x55,
	0xf2, 0x4e, 0x55, 0xa0, 0xdb, 0xb5, 0xef, 0x07, 0x02, 0xfb, 0xfd, 0x46, 0xda, 0xd7, 0x0a, 0xf4,
	0x80, 0x15, 0x53, 0x41, 0x42, 0x22, 0x88, 0x6d, 0x29, 0xff, 0xca, 0xf4, 0xbd, 0x4d, 0x06, 0xfb,
	0x2d, 0x04, 0x5f, 0x82, 0x5e, 0x4a, 0x32, 0x9a, 0xa8, 0xe6, 0xf4, 0x94, 0xe3, 0x91, 0x3c, 0xc4,
	0xa7, 0xb6, 0x41, 0xb2, 0xe9, 0x87, 0x0a, 0xd1, 0x4d, 0x6f, 0xf1, 0x69, 0xd3, 0x5b, 0x09, 0xfb,
	0x96, 0x8e, 0x07, 0xe1, 0x6e, 0xe7, 0xcb, 0x3b, 0x64, 0xe0, 0xf7, 0xb3, 0xa0, 0xd7, 0x0e, 0xda,
	0x7f, 0x38, 0x54, 0x72, 0x81, 0x7e, 0x3a, 0x54, 0x3a, 0x21, 0x5b, 0xa4, 0xa3, 0x7f, 0x30, 0x54,
	0xab, 0xc0, 0x92, 0x7f, 0x61, 0x94, 0xe4, 0xb1, 0x1a, 0xac, 0x79, 0xdf, 0x94, 0xcf, 0xcf, 0xf3,
	0x58, 0xff, 0xac, 0x87, 0x1f, 0x0c, 0x30, 0x37, 0xe4, 0x11, 0x3c, 0x02, 0x1d, 0xb5, 0x83, 0xd8,
	0xfd, 0xe5, 0xa6, 0xbb, 0xf5, 0x9e, 0xae, 0x3d, 0xf8, 0x33, 0xd3, 0xec, 0x37, 0x7c, 0x0d, 0xba,
	0x75, 0x23, 0xee, 0xfd, 0xde, 0xa5, 0xa9, 0xb5, 0xcd, 0xbf, 0xa1, 0x9a, 0xea, 0x07, 0xc3, 0xcb,
	0xd2, 0x31, 0xae, 0x4a, 0xc7, 0xf8, 0x5c, 0x3a, 0xc6, 0xdb, 0x6b, 0x67, 0xe6, 0xea, 0xda, 0x99,
	0xf9, 0x78, 0xed, 0xcc, 0xbc, 0xda, 0x89, 0x26, 0xe2, 0x38, 0x1f, 0xbb, 0x01, 0x8b, 0xbd, 0x34,
	0x1f, 0x9f, 0x4c, 0x82, 0x2d, 0x72, 0x46, 0x39, 0x8b, 0xa9, 0xd7, 0x5e, 0x77, 0xe7, 0xd3, 0x0b,
	0x4f, 0x5c, 0xa4, 0x94, 0x8f, 0xbb, 0xea, 0xee, 0xda, 0xf9, 0x16, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0xcf, 0x8f, 0xe8, 0x12, 0x05, 0x00, 0x00,
}

func (this *MsgPost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPost)
	if !ok {
		that2, ok := that.(MsgPost)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VendorID != that1.VendorID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.RewardAccount != that1.RewardAccount {
		return false
	}
	if this.Body != that1.Body {
		return false
	}
	if this.ChainID != that1.ChainID {
		return false
	}
	if this.ContractAddress != that1.ContractAddress {
		return false
	}
	if this.Metadata != that1.Metadata {
		return false
	}
	if that1.ParentID == nil {
		if this.ParentID != nil {
			return false
		}
	} else if !this.ParentID.Equal(*that1.ParentID) {
		return false
	}
	return true
}
func (this *MsgUpvote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpvote)
	if !ok {
		that2, ok := that.(MsgUpvote)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VendorID != that1.VendorID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if this.Curator != that1.Curator {
		return false
	}
	if this.RewardAccount != that1.RewardAccount {
		return false
	}
	if this.VoteNum != that1.VoteNum {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Post defines a method for sending a post
	Post(ctx context.Context, in *MsgPost, opts ...grpc.CallOption) (*MsgPostResponse, error)
	// Upvote defines a method for upvoting a post
	Upvote(ctx context.Context, in *MsgUpvote, opts ...grpc.CallOption) (*MsgUpvoteResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Post(ctx context.Context, in *MsgPost, opts ...grpc.CallOption) (*MsgPostResponse, error) {
	out := new(MsgPostResponse)
	err := c.cc.Invoke(ctx, "/stargaze.curating.v1beta1.Msg/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Upvote(ctx context.Context, in *MsgUpvote, opts ...grpc.CallOption) (*MsgUpvoteResponse, error) {
	out := new(MsgUpvoteResponse)
	err := c.cc.Invoke(ctx, "/stargaze.curating.v1beta1.Msg/Upvote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Post defines a method for sending a post
	Post(context.Context, *MsgPost) (*MsgPostResponse, error)
	// Upvote defines a method for upvoting a post
	Upvote(context.Context, *MsgUpvote) (*MsgUpvoteResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Post(ctx context.Context, req *MsgPost) (*MsgPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (*UnimplementedMsgServer) Upvote(ctx context.Context, req *MsgUpvote) (*MsgUpvoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upvote not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargaze.curating.v1beta1.Msg/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Post(ctx, req.(*MsgPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Upvote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpvote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Upvote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargaze.curating.v1beta1.Msg/Upvote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Upvote(ctx, req.(*MsgUpvote))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stargaze.curating.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Msg_Post_Handler,
		},
		{
			MethodName: "Upvote",
			Handler:    _Msg_Upvote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargaze/curating/v1beta1/tx.proto",
}

func (m *MsgPostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpvoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpvoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpvoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParentID != nil {
		{
			size := m.ParentID.Size()
			i -= size
			if _, err := m.ParentID.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RewardAccount) > 0 {
		i -= len(m.RewardAccount)
		copy(dAtA[i:], m.RewardAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RewardAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpvote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpvote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpvote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteNum != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VoteNum))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RewardAccount) > 0 {
		i -= len(m.RewardAccount)
		copy(dAtA[i:], m.RewardAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RewardAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Curator) > 0 {
		i -= len(m.Curator)
		copy(dAtA[i:], m.Curator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Curator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPostResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpvoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RewardAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ParentID != nil {
		l = m.ParentID.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpvote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Curator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RewardAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VoteNum != 0 {
		n += 1 + sovTx(uint64(m.VoteNum))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPostResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpvoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpvoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpvoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v PostID
			m.ParentID = &v
			if err := m.ParentID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpvote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpvote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpvote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Curator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Curator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteNum", wireType)
			}
			m.VoteNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteNum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
