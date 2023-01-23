// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/goeasyops/goeasyops.proto
// DO NOT EDIT!

/*
Package goeasyops is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/goeasyops/goeasyops.proto

It has these top-level messages:
	InContext
	ImmutableContext
	MutableContext
	CTXRoutingTags
	GRPCErrorList
	GRPCError
*/
package goeasyops

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "golang.conradwood.net/apis/auth"
import _ "golang.conradwood.net/apis/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// this is transported within the Context of a call. This proto is on-purpose very brief and static so that old versions of go-easyops can handle it
type InContext struct {
	ImCtx *ImmutableContext `protobuf:"bytes,1,opt,name=ImCtx" json:"ImCtx,omitempty"`
	MCtx  *MutableContext   `protobuf:"bytes,2,opt,name=MCtx" json:"MCtx,omitempty"`
}

func (m *InContext) Reset()                    { *m = InContext{} }
func (m *InContext) String() string            { return proto.CompactTextString(m) }
func (*InContext) ProtoMessage()               {}
func (*InContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InContext) GetImCtx() *ImmutableContext {
	if m != nil {
		return m.ImCtx
	}
	return nil
}

func (m *InContext) GetMCtx() *MutableContext {
	if m != nil {
		return m.MCtx
	}
	return nil
}

// this must not be changed throughout its lifetime. furthermore, go-easyops _must_ transport this as-is (preserving unknown fields)
type ImmutableContext struct {
	RequestID      string              `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
	CreatorService *auth.SignedUser    `protobuf:"bytes,2,opt,name=CreatorService" json:"CreatorService,omitempty"`
	User           *auth.SignedUser    `protobuf:"bytes,3,opt,name=User" json:"User,omitempty"`
	SudoUser       *auth.SignedUser    `protobuf:"bytes,4,opt,name=SudoUser" json:"SudoUser,omitempty"`
	Session        *auth.SignedSession `protobuf:"bytes,5,opt,name=Session" json:"Session,omitempty"`
}

func (m *ImmutableContext) Reset()                    { *m = ImmutableContext{} }
func (m *ImmutableContext) String() string            { return proto.CompactTextString(m) }
func (*ImmutableContext) ProtoMessage()               {}
func (*ImmutableContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImmutableContext) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *ImmutableContext) GetCreatorService() *auth.SignedUser {
	if m != nil {
		return m.CreatorService
	}
	return nil
}

func (m *ImmutableContext) GetUser() *auth.SignedUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ImmutableContext) GetSudoUser() *auth.SignedUser {
	if m != nil {
		return m.SudoUser
	}
	return nil
}

func (m *ImmutableContext) GetSession() *auth.SignedSession {
	if m != nil {
		return m.Session
	}
	return nil
}

// this may change. fields are not guaranteed to be preserved
type MutableContext struct {
	CallingService *auth.SignedUser `protobuf:"bytes,1,opt,name=CallingService" json:"CallingService,omitempty"`
	Debug          bool             `protobuf:"varint,2,opt,name=Debug" json:"Debug,omitempty"`
	Trace          bool             `protobuf:"varint,3,opt,name=Trace" json:"Trace,omitempty"`
	Tags           *CTXRoutingTags  `protobuf:"bytes,4,opt,name=Tags" json:"Tags,omitempty"`
}

func (m *MutableContext) Reset()                    { *m = MutableContext{} }
func (m *MutableContext) String() string            { return proto.CompactTextString(m) }
func (*MutableContext) ProtoMessage()               {}
func (*MutableContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MutableContext) GetCallingService() *auth.SignedUser {
	if m != nil {
		return m.CallingService
	}
	return nil
}

func (m *MutableContext) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *MutableContext) GetTrace() bool {
	if m != nil {
		return m.Trace
	}
	return false
}

func (m *MutableContext) GetTags() *CTXRoutingTags {
	if m != nil {
		return m.Tags
	}
	return nil
}

// within the context, routing tags are transported
type CTXRoutingTags struct {
	Tags            map[string]string `protobuf:"bytes,1,rep,name=Tags" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FallbackToPlain bool              `protobuf:"varint,2,opt,name=FallbackToPlain" json:"FallbackToPlain,omitempty"`
	Propagate       bool              `protobuf:"varint,3,opt,name=Propagate" json:"Propagate,omitempty"`
}

func (m *CTXRoutingTags) Reset()                    { *m = CTXRoutingTags{} }
func (m *CTXRoutingTags) String() string            { return proto.CompactTextString(m) }
func (*CTXRoutingTags) ProtoMessage()               {}
func (*CTXRoutingTags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CTXRoutingTags) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CTXRoutingTags) GetFallbackToPlain() bool {
	if m != nil {
		return m.FallbackToPlain
	}
	return false
}

func (m *CTXRoutingTags) GetPropagate() bool {
	if m != nil {
		return m.Propagate
	}
	return false
}

// a single error can only hold a single proto
type GRPCErrorList struct {
	Errors []*GRPCError `protobuf:"bytes,1,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *GRPCErrorList) Reset()                    { *m = GRPCErrorList{} }
func (m *GRPCErrorList) String() string            { return proto.CompactTextString(m) }
func (*GRPCErrorList) ProtoMessage()               {}
func (*GRPCErrorList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GRPCErrorList) GetErrors() []*GRPCError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type GRPCError struct {
	UserMessage string `protobuf:"bytes,1,opt,name=UserMessage" json:"UserMessage,omitempty"`
	LogMessage  string `protobuf:"bytes,2,opt,name=LogMessage" json:"LogMessage,omitempty"`
	MethodName  string `protobuf:"bytes,3,opt,name=MethodName" json:"MethodName,omitempty"`
	ServiceName string `protobuf:"bytes,4,opt,name=ServiceName" json:"ServiceName,omitempty"`
}

func (m *GRPCError) Reset()                    { *m = GRPCError{} }
func (m *GRPCError) String() string            { return proto.CompactTextString(m) }
func (*GRPCError) ProtoMessage()               {}
func (*GRPCError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GRPCError) GetUserMessage() string {
	if m != nil {
		return m.UserMessage
	}
	return ""
}

func (m *GRPCError) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

func (m *GRPCError) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *GRPCError) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func init() {
	proto.RegisterType((*InContext)(nil), "goeasyops.InContext")
	proto.RegisterType((*ImmutableContext)(nil), "goeasyops.ImmutableContext")
	proto.RegisterType((*MutableContext)(nil), "goeasyops.MutableContext")
	proto.RegisterType((*CTXRoutingTags)(nil), "goeasyops.CTXRoutingTags")
	proto.RegisterType((*GRPCErrorList)(nil), "goeasyops.GRPCErrorList")
	proto.RegisterType((*GRPCError)(nil), "goeasyops.GRPCError")
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/goeasyops/goeasyops.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xad, 0x1d, 0xcd, 0xa9, 0x28, 0x95, 0xd9, 0x45, 0x29, 0x08, 0x4a, 0xe0, 0xa2,
	0x42, 0x5b, 0x26, 0xc6, 0xc5, 0x26, 0x10, 0x37, 0x74, 0x03, 0x55, 0x5a, 0x51, 0xe5, 0x16, 0x89,
	0x5b, 0xb7, 0xb5, 0xbc, 0x68, 0x89, 0x5d, 0x6c, 0x67, 0xac, 0x4f, 0xc1, 0x53, 0xf0, 0x32, 0xbc,
	0x0b, 0xef, 0x80, 0xfc, 0x91, 0x34, 0xad, 0xe8, 0x6e, 0x92, 0xf8, 0x7f, 0x7e, 0xe7, 0xe3, 0x9f,
	0x93, 0xc0, 0x87, 0xa5, 0x14, 0x5a, 0xa8, 0x13, 0x26, 0x52, 0xc2, 0x59, 0x3c, 0x17, 0x5c, 0x92,
	0xc5, 0x4f, 0x21, 0x16, 0x31, 0xa7, 0xfa, 0x84, 0x2c, 0x13, 0x13, 0xa2, 0x44, 0xad, 0xc4, 0xb2,
	0xf2, 0x14, 0xdb, 0x2c, 0x14, 0x96, 0x42, 0xf7, 0xcd, 0x3d, 0x05, 0x48, 0xae, 0xaf, 0xed, 0xc5,
	0xa5, 0x75, 0xe3, 0x7b, 0xd8, 0xb9, 0xc8, 0x32, 0xc1, 0xfd, 0xcd, 0xf1, 0x51, 0x06, 0xe1, 0x90,
	0x0f, 0x04, 0xd7, 0xf4, 0x4e, 0xa3, 0xb7, 0x50, 0x1f, 0x66, 0x03, 0x7d, 0xd7, 0x09, 0x7a, 0x41,
	0xbf, 0x79, 0xfa, 0x34, 0x5e, 0x0f, 0x35, 0xcc, 0xb2, 0x5c, 0x93, 0x59, 0x4a, 0x3d, 0x8b, 0x1d,
	0x89, 0x8e, 0xa1, 0x36, 0x32, 0x19, 0x7b, 0x36, 0xe3, 0x49, 0x25, 0x63, 0xb4, 0xc9, 0x5b, 0x2c,
	0xfa, 0x1b, 0x40, 0x7b, 0xbb, 0x14, 0x7a, 0x06, 0x21, 0xa6, 0x3f, 0x72, 0xaa, 0xf4, 0xf0, 0xc2,
	0xb6, 0x0e, 0xf1, 0x5a, 0x40, 0xe7, 0xd0, 0x1a, 0x48, 0x4a, 0xb4, 0x90, 0x13, 0x2a, 0x6f, 0x93,
	0x39, 0xf5, 0xbd, 0xda, 0xb1, 0xb5, 0x3d, 0x49, 0x18, 0xa7, 0x8b, 0x6f, 0x8a, 0x4a, 0xbc, 0xc5,
	0xa1, 0xd7, 0x50, 0x33, 0x7a, 0x67, 0x7f, 0x07, 0x6f, 0xa3, 0xe8, 0x08, 0x1a, 0x93, 0x7c, 0x21,
	0x2c, 0x59, 0xdb, 0x41, 0x96, 0x04, 0x3a, 0x86, 0x07, 0x13, 0xaa, 0x54, 0x22, 0x78, 0xa7, 0x6e,
	0xe1, 0xc7, 0x55, 0xd8, 0x87, 0x70, 0xc1, 0x44, 0xbf, 0x03, 0x68, 0x6d, 0xbe, 0x08, 0xeb, 0x87,
	0xa4, 0x69, 0xc2, 0x59, 0xe1, 0x27, 0xd8, 0xe9, 0x67, 0x83, 0x43, 0x87, 0x50, 0xbf, 0xa0, 0xb3,
	0x9c, 0xd9, 0x17, 0xd0, 0xc0, 0xee, 0x60, 0xd4, 0xa9, 0x24, 0x73, 0x6a, 0x6d, 0x36, 0xb0, 0x3b,
	0x98, 0xbd, 0x4c, 0x09, 0x53, 0xde, 0x51, 0x75, 0x2f, 0x83, 0xe9, 0x77, 0x2c, 0x72, 0x9d, 0x70,
	0x66, 0x00, 0x6c, 0xb1, 0xe8, 0x4f, 0x00, 0xad, 0xcd, 0x00, 0x3a, 0xf3, 0x15, 0x82, 0xde, 0x7e,
	0xbf, 0x79, 0xfa, 0x6a, 0x67, 0x85, 0xd8, 0x5c, 0x2e, 0xb9, 0x96, 0x2b, 0x57, 0x0b, 0xf5, 0xe1,
	0xd1, 0x67, 0x92, 0xa6, 0x33, 0x32, 0xbf, 0x99, 0x8a, 0x71, 0x4a, 0x12, 0xee, 0x07, 0xde, 0x96,
	0xcd, 0xe2, 0xc7, 0x52, 0x2c, 0x09, 0x23, 0xba, 0x18, 0x7f, 0x2d, 0x74, 0xcf, 0x20, 0x2c, 0x4b,
	0xa3, 0x36, 0xec, 0xdf, 0xd0, 0x95, 0xff, 0x3a, 0xcc, 0xa3, 0xf1, 0x7d, 0x4b, 0xd2, 0xdc, 0x7d,
	0x0e, 0x21, 0x76, 0x87, 0xf7, 0x7b, 0xe7, 0x41, 0xf4, 0x11, 0x1e, 0x7e, 0xc1, 0xe3, 0xc1, 0xa5,
	0x94, 0x42, 0x5e, 0x25, 0x4a, 0xa3, 0x23, 0x38, 0xb0, 0x87, 0xc2, 0xcc, 0x61, 0xc5, 0x4c, 0x49,
	0x62, 0xcf, 0x44, 0xbf, 0x02, 0x08, 0x4b, 0x15, 0xf5, 0xa0, 0x69, 0x96, 0x31, 0xa2, 0x4a, 0x11,
	0x46, 0xfd, 0x00, 0x55, 0x09, 0x3d, 0x07, 0xb8, 0x12, 0xac, 0x00, 0xdc, 0x34, 0x15, 0xc5, 0xc4,
	0x47, 0x54, 0x5f, 0x8b, 0xc5, 0x57, 0x92, 0x39, 0x9b, 0x21, 0xae, 0x28, 0xa6, 0x83, 0xdf, 0xb0,
	0x05, 0x6a, 0xae, 0x43, 0x45, 0xfa, 0xf4, 0x12, 0x5e, 0xfc, 0xf7, 0xb7, 0x5e, 0xdb, 0x98, 0x1d,
	0xd8, 0xdf, 0xf9, 0xdd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x1c, 0xb8, 0x8f, 0x74, 0x04,
	0x00, 0x00,
}
