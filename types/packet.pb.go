// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icq/v1/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/cometbft/cometbft/abci/types"
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

// InterchainQueryPacketData is comprised of raw query.
type InterchainQueryPacketData struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// optional memo
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *InterchainQueryPacketData) Reset()         { *m = InterchainQueryPacketData{} }
func (m *InterchainQueryPacketData) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketData) ProtoMessage()    {}
func (*InterchainQueryPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{0}
}
func (m *InterchainQueryPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketData.Merge(m, src)
}
func (m *InterchainQueryPacketData) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketData proto.InternalMessageInfo

func (m *InterchainQueryPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterchainQueryPacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// InterchainQueryPacketAck is comprised of an ABCI query response with non-deterministic fields left empty (e.g. Codespace, Log, Info and ...).
type InterchainQueryPacketAck struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *InterchainQueryPacketAck) Reset()         { *m = InterchainQueryPacketAck{} }
func (m *InterchainQueryPacketAck) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketAck) ProtoMessage()    {}
func (*InterchainQueryPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{1}
}
func (m *InterchainQueryPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketAck.Merge(m, src)
}
func (m *InterchainQueryPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketAck proto.InternalMessageInfo

func (m *InterchainQueryPacketAck) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// CosmosQuery contains a list of tendermint ABCI query requests. It should be used when sending queries to an SDK host chain.
type CosmosQuery struct {
	Requests []types.RequestQuery `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *CosmosQuery) Reset()         { *m = CosmosQuery{} }
func (m *CosmosQuery) String() string { return proto.CompactTextString(m) }
func (*CosmosQuery) ProtoMessage()    {}
func (*CosmosQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{2}
}
func (m *CosmosQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosQuery.Merge(m, src)
}
func (m *CosmosQuery) XXX_Size() int {
	return m.Size()
}
func (m *CosmosQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosQuery proto.InternalMessageInfo

func (m *CosmosQuery) GetRequests() []types.RequestQuery {
	if m != nil {
		return m.Requests
	}
	return nil
}

// CosmosResponse contains a list of tendermint ABCI query responses. It should be used when receiving responses from an SDK host chain.
type CosmosResponse struct {
	Responses []types.ResponseQuery `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses"`
}

func (m *CosmosResponse) Reset()         { *m = CosmosResponse{} }
func (m *CosmosResponse) String() string { return proto.CompactTextString(m) }
func (*CosmosResponse) ProtoMessage()    {}
func (*CosmosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13b1ec1d226ce757, []int{3}
}
func (m *CosmosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosResponse.Merge(m, src)
}
func (m *CosmosResponse) XXX_Size() int {
	return m.Size()
}
func (m *CosmosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosResponse proto.InternalMessageInfo

func (m *CosmosResponse) GetResponses() []types.ResponseQuery {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*InterchainQueryPacketData)(nil), "icq.v1.InterchainQueryPacketData")
	proto.RegisterType((*InterchainQueryPacketAck)(nil), "icq.v1.InterchainQueryPacketAck")
	proto.RegisterType((*CosmosQuery)(nil), "icq.v1.CosmosQuery")
	proto.RegisterType((*CosmosResponse)(nil), "icq.v1.CosmosResponse")
}

func init() { proto.RegisterFile("icq/v1/packet.proto", fileDescriptor_13b1ec1d226ce757) }

var fileDescriptor_13b1ec1d226ce757 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xff, 0xaf, 0x2a, 0xea, 0x22, 0x86, 0xc0, 0x10, 0x8a, 0x30, 0x55, 0xa6, 0x2e,
	0xb5, 0xd5, 0x32, 0x30, 0x22, 0x5a, 0x16, 0x24, 0x84, 0x20, 0x62, 0x62, 0x73, 0xdd, 0xab, 0x34,
	0x2a, 0xb1, 0x53, 0xdb, 0x89, 0xd4, 0xb7, 0xe0, 0xb1, 0x3a, 0x76, 0x64, 0x42, 0xa8, 0x7d, 0x11,
	0x54, 0x1b, 0x28, 0x43, 0xb6, 0xa3, 0x7b, 0xcf, 0xf9, 0x74, 0x75, 0x2e, 0x3e, 0xce, 0xc4, 0x82,
	0x55, 0x03, 0x56, 0x70, 0x31, 0x07, 0x4b, 0x0b, 0xad, 0xac, 0x0a, 0x9b, 0x99, 0x58, 0xd0, 0x6a,
	0xd0, 0x39, 0x49, 0x55, 0xaa, 0xdc, 0x88, 0xed, 0x94, 0xdf, 0x76, 0xce, 0x2c, 0xc8, 0x29, 0xe8,
	0x3c, 0x93, 0x96, 0xf1, 0x89, 0xc8, 0x98, 0x5d, 0x16, 0x60, 0xfc, 0x32, 0x1e, 0xe3, 0xd3, 0x3b,
	0x69, 0x41, 0x8b, 0x19, 0xcf, 0xe4, 0x53, 0x09, 0x7a, 0xf9, 0xe8, 0xc8, 0xb7, 0xdc, 0xf2, 0x30,
	0xc4, 0x8d, 0x29, 0xb7, 0x3c, 0x42, 0x5d, 0xd4, 0x3b, 0x4c, 0x9c, 0xde, 0xcd, 0x72, 0xc8, 0x55,
	0xf4, 0xaf, 0x8b, 0x7a, 0xad, 0xc4, 0xe9, 0x98, 0xe2, 0xa8, 0x16, 0x72, 0x23, 0xe6, 0x75, 0x8c,
	0xf8, 0x01, 0xb7, 0xc7, 0xca, 0xe4, 0xca, 0x38, 0x6f, 0x78, 0x8d, 0x0f, 0x34, 0x2c, 0x4a, 0x30,
	0xd6, 0x44, 0xa8, 0xfb, 0xbf, 0xd7, 0x1e, 0x9e, 0xd3, 0xfd, 0xcd, 0x74, 0x77, 0x33, 0x4d, 0xbc,
	0xc1, 0x05, 0x46, 0x8d, 0xd5, 0xc7, 0x45, 0x90, 0xfc, 0x86, 0xe2, 0x67, 0x7c, 0xe4, 0x79, 0x09,
	0x98, 0x42, 0x49, 0x03, 0xe1, 0x08, 0xb7, 0xf4, 0xb7, 0xfe, 0x61, 0x92, 0x1a, 0xa6, 0x77, 0xfc,
	0x85, 0xee, 0x63, 0xa3, 0xfb, 0xd5, 0x86, 0xa0, 0xf5, 0x86, 0xa0, 0xcf, 0x0d, 0x41, 0x6f, 0x5b,
	0x12, 0xac, 0xb7, 0x24, 0x78, 0xdf, 0x92, 0xe0, 0x65, 0x98, 0x66, 0x76, 0x56, 0x4e, 0xa8, 0x50,
	0x39, 0x33, 0x56, 0x73, 0x99, 0xc2, 0xab, 0xaa, 0xa0, 0x5f, 0x81, 0xb4, 0xa5, 0x06, 0xc3, 0xb8,
	0x59, 0x4a, 0xd1, 0x77, 0xaf, 0xba, 0xf2, 0x75, 0x4f, 0x9a, 0xae, 0xef, 0xcb, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xc5, 0x69, 0xeb, 0xc1, 0x01, 0x00, 0x00,
}

func (m *InterchainQueryPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterchainQueryPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CosmosQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CosmosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterchainQueryPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *InterchainQueryPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *CosmosQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *CosmosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterchainQueryPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *InterchainQueryPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CosmosQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CosmosQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, types.RequestQuery{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CosmosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CosmosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, types.ResponseQuery{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
