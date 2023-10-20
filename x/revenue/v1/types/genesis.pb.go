// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/revenue/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// params are the revenue module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// revenues is a slice of active registered contracts for fee distribution
	Revenues []Revenue `protobuf:"bytes,2,rep,name=revenues,proto3" json:"revenues"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_649d64d9c3438055, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRevenues() []Revenue {
	if m != nil {
		return m.Revenues
	}
	return nil
}

// Params defines the revenue module params
type Params struct {
	// enable_revenue defines a parameter to enable the revenue module
	EnableRevenue bool `protobuf:"varint,1,opt,name=enable_revenue,json=enableRevenue,proto3" json:"enable_revenue,omitempty"`
	// developer_shares defines the proportion of the transaction fees to be
	// distributed to the registered contract owner
	DeveloperShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_shares,json=developerShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_shares"`
	// addr_derivation_cost_create defines the cost of address derivation for
	// verifying the contract deployer at fee registration
	AddrDerivationCostCreate uint64 `protobuf:"varint,3,opt,name=addr_derivation_cost_create,json=addrDerivationCostCreate,proto3" json:"addr_derivation_cost_create,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_649d64d9c3438055, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableRevenue() bool {
	if m != nil {
		return m.EnableRevenue
	}
	return false
}

func (m *Params) GetAddrDerivationCostCreate() uint64 {
	if m != nil {
		return m.AddrDerivationCostCreate
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fury.revenue.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "fury.revenue.v1.Params")
}

func init() { proto.RegisterFile("fury/revenue/v1/genesis.proto", fileDescriptor_649d64d9c3438055) }

var fileDescriptor_649d64d9c3438055 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0xb3, 0xaf, 0x22, 0xbe, 0x6b, 0xff, 0x48, 0xe8, 0x21, 0xb5, 0x10, 0x45, 0x68, 0xc9,
	0xc5, 0x0d, 0x5a, 0xda, 0x4b, 0xe9, 0x45, 0x05, 0xaf, 0x25, 0x9e, 0xda, 0x4b, 0x58, 0x93, 0x21,
	0x86, 0x6a, 0x36, 0xec, 0xae, 0xa1, 0x3d, 0xf7, 0x0b, 0xf4, 0x63, 0xd9, 0x9b, 0xc7, 0xd2, 0x83,
	0x14, 0xfd, 0x22, 0x25, 0xbb, 0xa9, 0x48, 0xbd, 0x24, 0xc3, 0x3c, 0xcf, 0xf3, 0x9b, 0x61, 0x07,
	0xdb, 0x90, 0xcd, 0x99, 0x70, 0x39, 0x64, 0x90, 0x2c, 0xc0, 0xcd, 0xba, 0x6e, 0x04, 0x09, 0x88,
	0x58, 0x90, 0x94, 0x33, 0xc9, 0xcc, 0xba, 0xd2, 0x49, 0xa1, 0x93, 0xac, 0xdb, 0x38, 0x4c, 0xfc,
	0x8a, 0x2a, 0xd1, 0x38, 0x8b, 0x58, 0xc4, 0x54, 0xe9, 0xe6, 0x95, 0xee, 0xb6, 0xdf, 0x10, 0x3e,
	0x1a, 0x69, 0xf2, 0x58, 0x52, 0x09, 0xe6, 0x2d, 0xae, 0xa4, 0x94, 0xd3, 0xb9, 0xb0, 0x50, 0x0b,
	0x39, 0xb5, 0x9e, 0x45, 0xfe, 0x4e, 0x22, 0x0f, 0x4a, 0xef, 0x97, 0x97, 0xeb, 0xa6, 0xe1, 0x15,
	0x6e, 0xf3, 0x0e, 0x57, 0x0b, 0x8b, 0xb0, 0xfe, 0xb5, 0x4a, 0x4e, 0xad, 0x77, 0x7e, 0x98, 0xf4,
	0x74, 0x59, 0x44, 0x77, 0x81, 0xf6, 0x07, 0xc2, 0x15, 0x4d, 0x35, 0x2f, 0xf1, 0x09, 0x24, 0x74,
	0x32, 0x03, 0xbf, 0x50, 0xd5, 0x1e, 0x55, 0xef, 0x58, 0x77, 0x0b, 0x82, 0xf9, 0x88, 0xeb, 0x21,
	0x64, 0x30, 0x63, 0x29, 0x70, 0x5f, 0x4c, 0x29, 0x57, 0x63, 0x91, 0xf3, 0xbf, 0x4f, 0x72, 0xf6,
	0xd7, 0xba, 0x79, 0x15, 0xc5, 0x72, 0xba, 0x98, 0x90, 0x80, 0xcd, 0xdd, 0x80, 0x89, 0xfc, 0x6d,
	0xf4, 0xaf, 0x23, 0xc2, 0x67, 0x57, 0xbe, 0xa6, 0x20, 0xc8, 0x10, 0x02, 0xef, 0x74, 0xc7, 0x19,
	0x2b, 0x8c, 0x79, 0x8f, 0x2f, 0x68, 0x18, 0x72, 0x3f, 0x04, 0x1e, 0x67, 0x54, 0xc6, 0x2c, 0xf1,
	0x03, 0x26, 0xa4, 0x1f, 0x70, 0xa0, 0x12, 0xac, 0x52, 0x0b, 0x39, 0x65, 0xcf, 0xca, 0x2d, 0xc3,
	0x9d, 0x63, 0xc0, 0x84, 0x1c, 0x28, 0xbd, 0x3f, 0x5a, 0x6e, 0x6c, 0xb4, 0xda, 0xd8, 0xe8, 0x7b,
	0x63, 0xa3, 0xf7, 0xad, 0x6d, 0xac, 0xb6, 0xb6, 0xf1, 0xb9, 0xb5, 0x8d, 0xa7, 0xce, 0xde, 0x46,
	0xfa, 0x58, 0xfa, 0x9b, 0x75, 0x6f, 0xdc, 0x97, 0xfd, 0xc3, 0xa9, 0xe5, 0x26, 0x15, 0x75, 0xa1,
	0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x45, 0x00, 0xf6, 0x0b, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Revenues) > 0 {
		for iNdEx := len(m.Revenues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Revenues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AddrDerivationCostCreate != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AddrDerivationCostCreate))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.DeveloperShares.Size()
		i -= size
		if _, err := m.DeveloperShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableRevenue {
		i--
		if m.EnableRevenue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Revenues) > 0 {
		for _, e := range m.Revenues {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableRevenue {
		n += 2
	}
	l = m.DeveloperShares.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AddrDerivationCostCreate != 0 {
		n += 1 + sovGenesis(uint64(m.AddrDerivationCostCreate))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revenues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Revenues = append(m.Revenues, Revenue{})
			if err := m.Revenues[len(m.Revenues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableRevenue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableRevenue = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeveloperShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddrDerivationCostCreate", wireType)
			}
			m.AddrDerivationCostCreate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddrDerivationCostCreate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
