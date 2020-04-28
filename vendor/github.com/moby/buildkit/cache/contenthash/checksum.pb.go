// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checksum.proto

package contenthash

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"
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

type CacheRecordType int32

const (
	CacheRecordTypeFile      CacheRecordType = 0
	CacheRecordTypeDir       CacheRecordType = 1
	CacheRecordTypeDirHeader CacheRecordType = 2
	CacheRecordTypeSymlink   CacheRecordType = 3
)

var CacheRecordType_name = map[int32]string{
	0: "FILE",
	1: "DIR",
	2: "DIR_HEADER",
	3: "SYMLINK",
}

var CacheRecordType_value = map[string]int32{
	"FILE":       0,
	"DIR":        1,
	"DIR_HEADER": 2,
	"SYMLINK":    3,
}

func (x CacheRecordType) String() string {
	return proto.EnumName(CacheRecordType_name, int32(x))
}

func (CacheRecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_843938c28b799986, []int{0}
}

type CacheRecord struct {
	Digest   github_com_opencontainers_go_digest.Digest `protobuf:"bytes,1,opt,name=digest,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"digest"`
	Type     CacheRecordType                            `protobuf:"varint,2,opt,name=type,proto3,enum=contenthash.CacheRecordType" json:"type,omitempty"`
	Linkname string                                     `protobuf:"bytes,3,opt,name=linkname,proto3" json:"linkname,omitempty"`
}

func (m *CacheRecord) Reset()         { *m = CacheRecord{} }
func (m *CacheRecord) String() string { return proto.CompactTextString(m) }
func (*CacheRecord) ProtoMessage()    {}
func (*CacheRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_843938c28b799986, []int{0}
}
func (m *CacheRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheRecord.Merge(m, src)
}
func (m *CacheRecord) XXX_Size() int {
	return m.Size()
}
func (m *CacheRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CacheRecord proto.InternalMessageInfo

func (m *CacheRecord) GetType() CacheRecordType {
	if m != nil {
		return m.Type
	}
	return CacheRecordTypeFile
}

func (m *CacheRecord) GetLinkname() string {
	if m != nil {
		return m.Linkname
	}
	return ""
}

type CacheRecordWithPath struct {
	Path   string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Record *CacheRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (m *CacheRecordWithPath) Reset()         { *m = CacheRecordWithPath{} }
func (m *CacheRecordWithPath) String() string { return proto.CompactTextString(m) }
func (*CacheRecordWithPath) ProtoMessage()    {}
func (*CacheRecordWithPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_843938c28b799986, []int{1}
}
func (m *CacheRecordWithPath) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheRecordWithPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheRecordWithPath.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheRecordWithPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheRecordWithPath.Merge(m, src)
}
func (m *CacheRecordWithPath) XXX_Size() int {
	return m.Size()
}
func (m *CacheRecordWithPath) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheRecordWithPath.DiscardUnknown(m)
}

var xxx_messageInfo_CacheRecordWithPath proto.InternalMessageInfo

func (m *CacheRecordWithPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CacheRecordWithPath) GetRecord() *CacheRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type CacheRecords struct {
	Paths []*CacheRecordWithPath `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *CacheRecords) Reset()         { *m = CacheRecords{} }
func (m *CacheRecords) String() string { return proto.CompactTextString(m) }
func (*CacheRecords) ProtoMessage()    {}
func (*CacheRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_843938c28b799986, []int{2}
}
func (m *CacheRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheRecords.Merge(m, src)
}
func (m *CacheRecords) XXX_Size() int {
	return m.Size()
}
func (m *CacheRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheRecords.DiscardUnknown(m)
}

var xxx_messageInfo_CacheRecords proto.InternalMessageInfo

func (m *CacheRecords) GetPaths() []*CacheRecordWithPath {
	if m != nil {
		return m.Paths
	}
	return nil
}

func init() {
	proto.RegisterEnum("contenthash.CacheRecordType", CacheRecordType_name, CacheRecordType_value)
	proto.RegisterType((*CacheRecord)(nil), "contenthash.CacheRecord")
	proto.RegisterType((*CacheRecordWithPath)(nil), "contenthash.CacheRecordWithPath")
	proto.RegisterType((*CacheRecords)(nil), "contenthash.CacheRecords")
}

func init() { proto.RegisterFile("checksum.proto", fileDescriptor_843938c28b799986) }

var fileDescriptor_843938c28b799986 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x77, 0x9a, 0x18, 0xf5, 0x8b, 0xd4, 0x30, 0x85, 0x76, 0x19, 0xca, 0x64, 0xcc, 0xc5,
	0x50, 0xec, 0xa6, 0x44, 0xf0, 0x6e, 0xdd, 0x84, 0x46, 0xab, 0xc8, 0x54, 0x10, 0xf1, 0x20, 0x9b,
	0xcd, 0xb8, 0xb3, 0xb4, 0xd9, 0x59, 0x76, 0x27, 0x87, 0xbc, 0x81, 0xec, 0xc9, 0x17, 0xd8, 0x93,
	0x82, 0xef, 0xe0, 0x5d, 0xe8, 0xb1, 0x47, 0xf1, 0x50, 0x24, 0x79, 0x11, 0xd9, 0xd9, 0x2a, 0xcb,
	0x4a, 0x4e, 0xf3, 0x7d, 0x33, 0xbf, 0xef, 0xff, 0xff, 0xcf, 0x30, 0xb0, 0xed, 0x4b, 0xe1, 0x9f,
	0xa7, 0x8b, 0xb9, 0x13, 0x27, 0x4a, 0x2b, 0xdc, 0xf6, 0x55, 0xa4, 0x45, 0xa4, 0xa5, 0x97, 0x4a,
	0x72, 0x18, 0x84, 0x5a, 0x2e, 0xa6, 0x8e, 0xaf, 0xe6, 0x83, 0x40, 0x05, 0x6a, 0x60, 0x98, 0xe9,
	0xe2, 0xa3, 0xe9, 0x4c, 0x63, 0xaa, 0x72, 0xb6, 0xf7, 0x0d, 0x41, 0xfb, 0x99, 0xe7, 0x4b, 0xc1,
	0x85, 0xaf, 0x92, 0x19, 0x7e, 0x0e, 0xad, 0x59, 0x18, 0x88, 0x54, 0xdb, 0x88, 0xa1, 0xfe, 0xdd,
	0xe3, 0xe1, 0xe5, 0x75, 0xd7, 0xfa, 0x75, 0xdd, 0x3d, 0xa8, 0xc8, 0xaa, 0x58, 0x44, 0x85, 0xa5,
	0x17, 0x46, 0x22, 0x49, 0x07, 0x81, 0x3a, 0x2c, 0x47, 0x1c, 0xd7, 0x2c, 0xfc, 0x46, 0x01, 0x1f,
	0x41, 0x53, 0x2f, 0x63, 0x61, 0x6f, 0x31, 0xd4, 0xdf, 0x1e, 0xee, 0x3b, 0x95, 0x98, 0x4e, 0xc5,
	0xf3, 0xcd, 0x32, 0x16, 0xdc, 0x90, 0x98, 0xc0, 0x9d, 0x8b, 0x30, 0x3a, 0x8f, 0xbc, 0xb9, 0xb0,
	0x1b, 0x85, 0x3f, 0xff, 0xd7, 0xf7, 0xde, 0xc3, 0x4e, 0x65, 0xe8, 0x6d, 0xa8, 0xe5, 0x6b, 0x4f,
	0x4b, 0x8c, 0xa1, 0x19, 0x7b, 0x5a, 0x96, 0x71, 0xb9, 0xa9, 0xf1, 0x11, 0xb4, 0x12, 0x43, 0x19,
	0xeb, 0xf6, 0xd0, 0xde, 0x64, 0xcd, 0x6f, 0xb8, 0xde, 0x18, 0xee, 0x55, 0xb6, 0x53, 0xfc, 0x04,
	0x6e, 0x15, 0x4a, 0xa9, 0x8d, 0x58, 0xa3, 0xdf, 0x1e, 0xb2, 0x4d, 0x02, 0x7f, 0x63, 0xf0, 0x12,
	0x3f, 0xf8, 0x81, 0xe0, 0x7e, 0xed, 0x6a, 0xf8, 0x01, 0x34, 0xc7, 0x93, 0xd3, 0x51, 0xc7, 0x22,
	0x7b, 0x59, 0xce, 0x76, 0x6a, 0xc7, 0xe3, 0xf0, 0x42, 0xe0, 0x2e, 0x34, 0xdc, 0x09, 0xef, 0x20,
	0xb2, 0x9b, 0xe5, 0x0c, 0xd7, 0x08, 0x37, 0x4c, 0xf0, 0x23, 0x00, 0x77, 0xc2, 0x3f, 0x9c, 0x8c,
	0x9e, 0xba, 0x23, 0xde, 0xd9, 0x22, 0xfb, 0x59, 0xce, 0xec, 0xff, 0xb9, 0x13, 0xe1, 0xcd, 0x44,
	0x82, 0x1f, 0xc2, 0xed, 0xb3, 0x77, 0x2f, 0x4f, 0x27, 0xaf, 0x5e, 0x74, 0x1a, 0x84, 0x64, 0x39,
	0xdb, 0xad, 0xa1, 0x67, 0xcb, 0x79, 0xf1, 0xae, 0x64, 0xef, 0xd3, 0x17, 0x6a, 0x7d, 0xff, 0x4a,
	0xeb, 0x99, 0x8f, 0xed, 0xcb, 0x15, 0x45, 0x57, 0x2b, 0x8a, 0x7e, 0xaf, 0x28, 0xfa, 0xbc, 0xa6,
	0xd6, 0xd5, 0x9a, 0x5a, 0x3f, 0xd7, 0xd4, 0x9a, 0xb6, 0xcc, 0xbf, 0x79, 0xfc, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0xd7, 0xd8, 0x37, 0x85, 0x02, 0x00, 0x00,
}

func (m *CacheRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Linkname) > 0 {
		i -= len(m.Linkname)
		copy(dAtA[i:], m.Linkname)
		i = encodeVarintChecksum(dAtA, i, uint64(len(m.Linkname)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintChecksum(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintChecksum(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CacheRecordWithPath) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheRecordWithPath) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheRecordWithPath) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Record != nil {
		{
			size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChecksum(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintChecksum(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CacheRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for iNdEx := len(m.Paths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Paths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintChecksum(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintChecksum(dAtA []byte, offset int, v uint64) int {
	offset -= sovChecksum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CacheRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovChecksum(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovChecksum(uint64(m.Type))
	}
	l = len(m.Linkname)
	if l > 0 {
		n += 1 + l + sovChecksum(uint64(l))
	}
	return n
}

func (m *CacheRecordWithPath) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovChecksum(uint64(l))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovChecksum(uint64(l))
	}
	return n
}

func (m *CacheRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovChecksum(uint64(l))
		}
	}
	return n
}

func sovChecksum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChecksum(x uint64) (n int) {
	return sovChecksum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CacheRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: CacheRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CacheRecordType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Linkname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Linkname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChecksum
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
func (m *CacheRecordWithPath) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: CacheRecordWithPath: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheRecordWithPath: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Record == nil {
				m.Record = &CacheRecord{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChecksum
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
func (m *CacheRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: CacheRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, &CacheRecordWithPath{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChecksum
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
func skipChecksum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChecksum
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
					return 0, ErrIntOverflowChecksum
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
					return 0, ErrIntOverflowChecksum
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
				return 0, ErrInvalidLengthChecksum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChecksum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChecksum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChecksum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChecksum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChecksum = fmt.Errorf("proto: unexpected end of group")
)