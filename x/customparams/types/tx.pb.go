// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/customparams/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgUpdateStakingParams struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// staking_params holds the parameters related to the staking module.
	StakingParams StakingParams `protobuf:"bytes,2,opt,name=staking_params,json=stakingParams,proto3" json:"staking_params"`
}

func (m *MsgUpdateStakingParams) Reset()         { *m = MsgUpdateStakingParams{} }
func (m *MsgUpdateStakingParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateStakingParams) ProtoMessage()    {}
func (*MsgUpdateStakingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9f2c8294c3378c0, []int{0}
}
func (m *MsgUpdateStakingParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateStakingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateStakingParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateStakingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateStakingParams.Merge(m, src)
}
func (m *MsgUpdateStakingParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateStakingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateStakingParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateStakingParams proto.InternalMessageInfo

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9f2c8294c3378c0, []int{1}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return m.Size()
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateStakingParams)(nil), "coreum.customparams.v1.MsgUpdateStakingParams")
	proto.RegisterType((*EmptyResponse)(nil), "coreum.customparams.v1.EmptyResponse")
}

func init() { proto.RegisterFile("coreum/customparams/v1/tx.proto", fileDescriptor_c9f2c8294c3378c0) }

var fileDescriptor_c9f2c8294c3378c0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x2f, 0x4a,
	0x2d, 0xcd, 0xd5, 0x4f, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x2d, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6,
	0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83, 0x28, 0xd0,
	0x43, 0x56, 0xa0, 0x57, 0x66, 0x28, 0x25, 0x98, 0x98, 0x9b, 0x99, 0x97, 0xaf, 0x0f, 0x26, 0x21,
	0x4a, 0xa5, 0x94, 0x71, 0x98, 0x05, 0xd5, 0x04, 0x51, 0x24, 0x9e, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f,
	0xac, 0x9f, 0x5b, 0x9c, 0x0e, 0x92, 0xcb, 0x2d, 0x4e, 0x87, 0x4a, 0x48, 0x42, 0x24, 0xe2, 0xc1,
	0x3c, 0x7d, 0x08, 0x07, 0x2a, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x11, 0x07, 0xb1, 0x20, 0xa2,
	0x4a, 0x77, 0x18, 0xb9, 0xc4, 0x7c, 0x8b, 0xd3, 0x43, 0x0b, 0x52, 0x12, 0x4b, 0x52, 0x83, 0x4b,
	0x12, 0xb3, 0x33, 0xf3, 0xd2, 0x03, 0xc0, 0x56, 0x09, 0x99, 0x71, 0x71, 0x26, 0x96, 0x96, 0x64,
	0xe4, 0x17, 0x65, 0x96, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x49, 0x5c, 0xda, 0xa2,
	0x2b, 0x02, 0x35, 0xd5, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33, 0x2f,
	0x3d, 0x08, 0xa1, 0x54, 0x28, 0x88, 0x8b, 0xaf, 0x18, 0x62, 0x50, 0x3c, 0xc4, 0xd1, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xaa, 0x7a, 0xd8, 0x43, 0x41, 0x0f, 0xc5, 0x5a, 0x27, 0x96, 0x13,
	0xf7, 0xe4, 0x19, 0x82, 0x78, 0x8b, 0x91, 0x05, 0xad, 0xcc, 0x9a, 0x9e, 0x6f, 0xd0, 0x42, 0xd8,
	0xd1, 0xf5, 0x7c, 0x83, 0x96, 0x32, 0x4a, 0x08, 0x61, 0xf7, 0x83, 0x12, 0x3f, 0x17, 0xaf, 0x6b,
	0x6e, 0x41, 0x49, 0x65, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x51, 0x15, 0x17, 0xb3,
	0x6f, 0x71, 0xba, 0x50, 0x0e, 0x97, 0x30, 0x36, 0x2f, 0xeb, 0xe1, 0x72, 0x22, 0x76, 0xe3, 0xa5,
	0x70, 0x7a, 0x09, 0xc5, 0x52, 0x29, 0xd6, 0x86, 0xe7, 0x1b, 0xb4, 0x18, 0x9d, 0x22, 0x4e, 0x3c,
	0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xab, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x67, 0xb0, 0xa9, 0x6e, 0xf9, 0xa5, 0x79,
	0x29, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xfa, 0xd0, 0x44, 0x51, 0x66, 0xa6, 0x5f, 0x81, 0x9a, 0x32,
	0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x91, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xe6, 0xfa, 0x15, 0x89, 0x02, 0x00, 0x00,
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
	// UpdateStakingParams is a governance operation that sets the staking parameter.
	// NOTE: all parameters must be provided.
	UpdateStakingParams(ctx context.Context, in *MsgUpdateStakingParams, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateStakingParams(ctx context.Context, in *MsgUpdateStakingParams, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/coreum.customparams.v1.Msg/UpdateStakingParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateStakingParams is a governance operation that sets the staking parameter.
	// NOTE: all parameters must be provided.
	UpdateStakingParams(context.Context, *MsgUpdateStakingParams) (*EmptyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateStakingParams(ctx context.Context, req *MsgUpdateStakingParams) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStakingParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateStakingParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateStakingParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateStakingParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreum.customparams.v1.Msg/UpdateStakingParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateStakingParams(ctx, req.(*MsgUpdateStakingParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreum.customparams.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStakingParams",
			Handler:    _Msg_UpdateStakingParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coreum/customparams/v1/tx.proto",
}

func (m *MsgUpdateStakingParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateStakingParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateStakingParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.StakingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmptyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgUpdateStakingParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.StakingParams.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *EmptyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateStakingParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateStakingParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateStakingParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
