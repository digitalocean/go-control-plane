// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/certs.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Certificates struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	CaCert               []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SerialNumber         string                  `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	SubjectAltNames      []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	DaysUntilExpiration  uint64                  `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	ValidFrom            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ExpirationTime       *timestamp.Timestamp    `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *timestamp.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v3.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v3.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v3.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v3.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v3/certs.proto", fileDescriptor_51228a64c985b2dc) }

var fileDescriptor_51228a64c985b2dc = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x25, 0xd3, 0x52, 0xd4, 0xdb, 0x32, 0x23, 0x0c, 0x48, 0x51, 0x41, 0x4c, 0x27, 0xa0, 0x21,
	0x42, 0x90, 0x48, 0xed, 0x8a, 0xb2, 0x40, 0xd3, 0x02, 0x62, 0x35, 0xaa, 0x02, 0xac, 0xa3, 0xdb,
	0xc4, 0x6d, 0x8d, 0x12, 0x3b, 0xb2, 0x9d, 0xd0, 0x6e, 0x59, 0xf1, 0x0d, 0xfc, 0x07, 0x9f, 0xc0,
	0x7f, 0x21, 0x3b, 0xf3, 0x68, 0x87, 0x8a, 0xb2, 0xb3, 0xcf, 0xb9, 0xe7, 0x9e, 0xfb, 0x82, 0x1e,
	0xe5, 0x95, 0x58, 0x87, 0x98, 0xe6, 0x8c, 0x87, 0xd5, 0x30, 0x4c, 0xa8, 0xd4, 0x2a, 0x28, 0xa4,
	0xd0, 0x82, 0x1c, 0x5a, 0x2e, 0xb0, 0x5c, 0x50, 0x0d, 0x7b, 0xc7, 0x0b, 0x21, 0x16, 0x19, 0x0d,
	0x2d, 0x3b, 0x2b, 0xe7, 0xa1, 0x66, 0x39, 0x55, 0x1a, 0xf3, 0xa2, 0x16, 0xf4, 0x4e, 0xca, 0xb4,
	0xc0, 0x10, 0x39, 0x17, 0x1a, 0x35, 0x13, 0x5c, 0x85, 0x15, 0x95, 0x8a, 0x09, 0xce, 0xf8, 0xa2,
	0x0e, 0xf1, 0x56, 0xd0, 0x9d, 0x50, 0xa9, 0xd9, 0x9c, 0x25, 0xa8, 0xa9, 0x22, 0x6f, 0xa1, 0x9b,
	0x6c, 0xfc, 0x5d, 0xa7, 0xdf, 0xf0, 0x3b, 0x83, 0x47, 0xc1, 0xb6, 0x75, 0xb0, 0xa1, 0x89, 0xb6,
	0x04, 0xa3, 0xe7, 0x3f, 0x7f, 0xff, 0x78, 0xe2, 0x41, 0x7f, 0x4b, 0x30, 0xc0, 0xac, 0x58, 0xe2,
	0xa6, 0x4a, 0x79, 0xbf, 0x1c, 0xe8, 0x6c, 0x00, 0xe4, 0x0d, 0xdc, 0x49, 0x30, 0x36, 0xb9, 0x2e,
	0x4c, 0xbd, 0x7f, 0x98, 0xbe, 0xa3, 0x1a, 0x59, 0xa6, 0xa2, 0x56, 0x82, 0x06, 0x25, 0x67, 0x00,
	0x46, 0x19, 0x27, 0x4b, 0x64, 0xdc, 0x3d, 0xf8, 0x6f, 0x7d, 0xdb, 0xa8, 0x26, 0x46, 0x34, 0x3a,
	0x35, 0x85, 0x9f, 0xc0, 0xf1, 0x9e, 0xc2, 0xbd, 0xef, 0x0d, 0x20, 0x7f, 0x67, 0x22, 0x04, 0x9a,
	0x05, 0xea, 0xa5, 0xeb, 0xf4, 0x1d, 0xbf, 0x1d, 0xd9, 0x37, 0x79, 0x0a, 0x77, 0x15, 0x95, 0x0c,
	0xb3, 0x98, 0x97, 0xf9, 0x8c, 0x4a, 0xf7, 0xc0, 0x92, 0xdd, 0x1a, 0x3c, 0xb7, 0x18, 0x99, 0xc2,
	0x3d, 0x55, 0xce, 0xbe, 0xd2, 0x44, 0xc7, 0x98, 0xe9, 0x98, 0x63, 0x4e, 0x95, 0xdb, 0xb0, 0x1d,
	0x3c, 0xbb, 0xd9, 0xc1, 0xa7, 0x3a, 0xf0, 0x2c, 0xd3, 0x54, 0x72, 0xd4, 0xf4, 0x1c, 0x73, 0x1a,
	0x1d, 0xa9, 0x2b, 0xd4, 0xfc, 0x15, 0x19, 0xc0, 0xc3, 0x14, 0xd7, 0x2a, 0x2e, 0xb9, 0x66, 0x59,
	0x4c, 0x57, 0x05, 0x93, 0x76, 0xfd, 0x6e, 0xb3, 0xef, 0xf8, 0xcd, 0xe8, 0xbe, 0x21, 0xbf, 0x18,
	0xee, 0xfd, 0x15, 0x45, 0x5e, 0x03, 0x54, 0x98, 0xb1, 0x34, 0x9e, 0x4b, 0x91, 0xbb, 0xb7, 0xfb,
	0x8e, 0xdf, 0x19, 0xf4, 0x82, 0xfa, 0xc0, 0x82, 0xcb, 0x03, 0x0b, 0x3e, 0x5f, 0x1e, 0x58, 0xd4,
	0xb6, 0xd1, 0x1f, 0xa4, 0xc8, 0xc9, 0x04, 0x8e, 0xae, 0x3d, 0x62, 0x73, 0x83, 0x6e, 0x6b, 0xaf,
	0xfe, 0xf0, 0x5a, 0x62, 0xc0, 0xd1, 0x2b, 0x33, 0x7d, 0x1f, 0x4e, 0xf7, 0x4c, 0xff, 0x62, 0xda,
	0xde, 0x37, 0x78, 0xb0, 0x6b, 0x16, 0x84, 0x40, 0x23, 0xe5, 0xaa, 0x5e, 0xc2, 0xc7, 0x5b, 0x91,
	0xf9, 0x18, 0xac, 0x94, 0xac, 0x9e, 0xbd, 0xc1, 0x4a, 0xc9, 0x46, 0xa1, 0xb1, 0x7b, 0x01, 0xfe,
	0x2e, 0xbb, 0x5d, 0x89, 0xc7, 0x2d, 0x68, 0x9a, 0xcd, 0x8c, 0x5f, 0xc2, 0x63, 0x26, 0xea, 0xb5,
	0x14, 0x52, 0xac, 0xd6, 0x37, 0x36, 0x34, 0x06, 0x53, 0xac, 0x9a, 0x9a, 0x86, 0xa7, 0xce, 0xac,
	0x65, 0x3b, 0x1f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x4b, 0xa8, 0xbd, 0xd6, 0x03, 0x00,
	0x00,
}