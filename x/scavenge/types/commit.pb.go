// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/commit.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Commit struct {
	Index                 string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	SolutionHash          string `protobuf:"bytes,2,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	SolutionScavengerHash string `protobuf:"bytes,3,opt,name=solutionScavengerHash,proto3" json:"solutionScavengerHash,omitempty"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc1321306004c1c, []int{0}
}
func (m *Commit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return m.Size()
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Commit) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *Commit) GetSolutionScavengerHash() string {
	if m != nil {
		return m.SolutionScavengerHash
	}
	return ""
}

func init() {
	proto.RegisterType((*Commit)(nil), "irclausen.scavenge.scavenge.Commit")
}

func init() { proto.RegisterFile("scavenge/commit.proto", fileDescriptor_7dc1321306004c1c) }

var fileDescriptor_7dc1321306004c1c = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xce, 0x2c, 0x4a, 0xce, 0x49, 0x2c, 0x2d, 0x4e, 0xcd, 0xd3, 0x83, 0x29, 0x80,
	0x33, 0x94, 0x2a, 0xb8, 0xd8, 0x9c, 0xc1, 0x8a, 0x85, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52,
	0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x25, 0x2e, 0x9e, 0xe2, 0xfc,
	0x9c, 0xd2, 0x92, 0xcc, 0xfc, 0x3c, 0x8f, 0xc4, 0xe2, 0x0c, 0x09, 0x26, 0xb0, 0x24, 0x8a, 0x98,
	0x90, 0x09, 0x97, 0x28, 0x8c, 0x1f, 0x0c, 0x35, 0xb7, 0x08, 0xac, 0x98, 0x19, 0xac, 0x18, 0xbb,
	0xa4, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa5, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0xdd, 0xae, 0x0f, 0xf7, 0x5c, 0x05,
	0x82, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x89, 0xc7, 0xc0, 0xd5, 0x00, 0x01, 0x00, 0x00,
}

func (m *Commit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SolutionScavengerHash) > 0 {
		i -= len(m.SolutionScavengerHash)
		copy(dAtA[i:], m.SolutionScavengerHash)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.SolutionScavengerHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommit(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.SolutionScavengerHash)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	return n
}

func sovCommit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommit(x uint64) (n int) {
	return sovCommit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
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
			return fmt.Errorf("proto: Commit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
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
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
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
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionScavengerHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
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
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionScavengerHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
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
func skipCommit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommit
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
					return 0, ErrIntOverflowCommit
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
					return 0, ErrIntOverflowCommit
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
				return 0, ErrInvalidLengthCommit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommit = fmt.Errorf("proto: unexpected end of group")
)