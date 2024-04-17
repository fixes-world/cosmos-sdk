// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/defaults/lockup/lockup.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Period defines a length of time and amount of coins that will be lock.
type Period struct {
	// Period duration
	Length time.Duration                            `protobuf:"bytes,1,opt,name=length,proto3,stdduration" json:"length"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b466256e1a079c, []int{0}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetLength() time.Duration {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*Period)(nil), "cosmos.accounts.defaults.lockup.Period")
}

func init() {
	proto.RegisterFile("cosmos/accounts/defaults/lockup/lockup.proto", fileDescriptor_79b466256e1a079c)
}

var fileDescriptor_79b466256e1a079c = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x90, 0x32, 0x04, 0x18, 0xa8, 0x18, 0x4a, 0x07, 0xa7, 0x62, 0xaa, 0x2a, 0x6a,
	0xab, 0x70, 0x01, 0x54, 0x10, 0xac, 0x88, 0x91, 0x05, 0x39, 0x8e, 0xeb, 0x5a, 0x4d, 0xf2, 0xaa,
	0xda, 0x41, 0xf4, 0x16, 0x8c, 0x88, 0x13, 0x20, 0xa6, 0x5e, 0x02, 0xa9, 0x63, 0x47, 0x26, 0x8a,
	0x9a, 0xa1, 0xd7, 0x40, 0xb1, 0x9d, 0x91, 0xc5, 0xef, 0x59, 0xfe, 0xbf, 0xf7, 0xbf, 0x5f, 0x8e,
	0xce, 0x39, 0xe8, 0x1c, 0x34, 0x65, 0x9c, 0x43, 0x59, 0x18, 0x4d, 0x53, 0x31, 0x66, 0x65, 0x66,
	0x34, 0xcd, 0x80, 0x4f, 0xcb, 0x99, 0x2f, 0x64, 0x36, 0x07, 0x03, 0xad, 0xd8, 0xa9, 0x49, 0xa3,
	0x26, 0x8d, 0x9a, 0x38, 0x59, 0xe7, 0x98, 0xe5, 0xaa, 0x00, 0x6a, 0x4f, 0xc7, 0x74, 0xb0, 0x77,
	0x48, 0x98, 0x16, 0xf4, 0x79, 0x98, 0x08, 0xc3, 0x86, 0x94, 0x83, 0x2a, 0xfc, 0xfb, 0x89, 0x04,
	0x09, 0xb6, 0xa5, 0x75, 0xd7, 0x50, 0x12, 0x40, 0x66, 0x82, 0xda, 0x5b, 0x52, 0x8e, 0x69, 0x5a,
	0xce, 0x99, 0x51, 0xe0, 0xa9, 0xb3, 0x2f, 0x14, 0x85, 0xf7, 0x62, 0xae, 0x20, 0x6d, 0x5d, 0x45,
	0x61, 0x26, 0x0a, 0x69, 0x26, 0x6d, 0xd4, 0x45, 0xbd, 0x83, 0x8b, 0x53, 0xe2, 0x58, 0xd2, 0xb0,
	0xe4, 0xc6, 0xb3, 0xa3, 0xa3, 0xd5, 0x4f, 0x1c, 0xbc, 0x6d, 0x62, 0xf4, 0xb1, 0x5b, 0xf6, 0xd1,
	0x83, 0xe7, 0x5a, 0x8b, 0x28, 0x64, 0x79, 0x1d, 0xa8, 0xbd, 0xd7, 0xdd, 0xb7, 0x13, 0x7c, 0xce,
	0x7a, 0x67, 0xe2, 0x77, 0x26, 0xd7, 0xa0, 0x8a, 0xd1, 0x6d, 0x3d, 0xe1, 0x73, 0x13, 0xf7, 0xa4,
	0x32, 0x93, 0x32, 0x21, 0x1c, 0x72, 0xea, 0x03, 0xba, 0x32, 0xd0, 0xe9, 0x94, 0x9a, 0xc5, 0x4c,
	0x68, 0x0b, 0xe8, 0xf7, 0xdd, 0xb2, 0x7f, 0x98, 0x09, 0xc9, 0xf8, 0xe2, 0xa9, 0x4e, 0xad, 0xbd,
	0xb5, 0x33, 0x1c, 0xdd, 0xad, 0xb6, 0x18, 0xad, 0xb7, 0x18, 0xfd, 0x6e, 0x31, 0x7a, 0xad, 0x70,
	0xb0, 0xae, 0x70, 0xf0, 0x5d, 0xe1, 0xe0, 0x71, 0xe0, 0xe6, 0xe9, 0x74, 0x4a, 0x14, 0xd0, 0x97,
	0xff, 0x7f, 0xc8, 0x9a, 0x25, 0xa1, 0x4d, 0x7b, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xa5,
	0x33, 0xd0, 0xd1, 0x01, 0x00, 0x00,
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLockup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Length, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Length):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLockup(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLockup(dAtA []byte, offset int, v uint64) int {
	offset -= sovLockup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Length)
	n += 1 + l + sovLockup(uint64(l))
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovLockup(uint64(l))
		}
	}
	return n
}

func sovLockup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLockup(x uint64) (n int) {
	return sovLockup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockup
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Length, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockup
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
func skipLockup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLockup
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
					return 0, ErrIntOverflowLockup
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
					return 0, ErrIntOverflowLockup
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
				return 0, ErrInvalidLengthLockup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLockup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLockup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLockup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLockup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLockup = fmt.Errorf("proto: unexpected end of group")
)