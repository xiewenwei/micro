// Code generated by protoc-gen-go.
// source: github.com/micro/micro/run/proto/run.proto
// DO NOT EDIT!

/*
Package go_micro_run is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/micro/run/proto/run.proto

It has these top-level messages:
	Source
	Binary
	Process
	FetchRequest
	FetchResponse
	BuildRequest
	BuildResponse
	ExecRequest
	ExecResponse
	KillRequest
	KillResponse
	WaitRequest
	WaitResponse
*/
package go_micro_run

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Dir string `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Source) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Source) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

type Binary struct {
	Path   string  `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Source *Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
}

func (m *Binary) Reset()                    { *m = Binary{} }
func (m *Binary) String() string            { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()               {}
func (*Binary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Binary) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Binary) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type Process struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Binary *Binary `protobuf:"bytes,2,opt,name=binary" json:"binary,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Process) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Process) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type FetchRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *FetchRequest) Reset()                    { *m = FetchRequest{} }
func (m *FetchRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()               {}
func (*FetchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type FetchResponse struct {
	Source *Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *FetchResponse) Reset()                    { *m = FetchResponse{} }
func (m *FetchResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()               {}
func (*FetchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchResponse) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type BuildRequest struct {
	Source *Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildRequest) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type BuildResponse struct {
	Binary *Binary `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BuildResponse) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type ExecRequest struct {
	Binary *Binary `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ExecRequest) GetBinary() *Binary {
	if m != nil {
		return m.Binary
	}
	return nil
}

type ExecResponse struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ExecResponse) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type KillRequest struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *KillRequest) Reset()                    { *m = KillRequest{} }
func (m *KillRequest) String() string            { return proto.CompactTextString(m) }
func (*KillRequest) ProtoMessage()               {}
func (*KillRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *KillRequest) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type KillResponse struct {
}

func (m *KillResponse) Reset()                    { *m = KillResponse{} }
func (m *KillResponse) String() string            { return proto.CompactTextString(m) }
func (*KillResponse) ProtoMessage()               {}
func (*KillResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type WaitRequest struct {
	Process *Process `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
}

func (m *WaitRequest) Reset()                    { *m = WaitRequest{} }
func (m *WaitRequest) String() string            { return proto.CompactTextString(m) }
func (*WaitRequest) ProtoMessage()               {}
func (*WaitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *WaitRequest) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type WaitResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *WaitResponse) Reset()                    { *m = WaitResponse{} }
func (m *WaitResponse) String() string            { return proto.CompactTextString(m) }
func (*WaitResponse) ProtoMessage()               {}
func (*WaitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *WaitResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "go.micro.run.Source")
	proto.RegisterType((*Binary)(nil), "go.micro.run.Binary")
	proto.RegisterType((*Process)(nil), "go.micro.run.Process")
	proto.RegisterType((*FetchRequest)(nil), "go.micro.run.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "go.micro.run.FetchResponse")
	proto.RegisterType((*BuildRequest)(nil), "go.micro.run.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "go.micro.run.BuildResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.run.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.run.ExecResponse")
	proto.RegisterType((*KillRequest)(nil), "go.micro.run.KillRequest")
	proto.RegisterType((*KillResponse)(nil), "go.micro.run.KillResponse")
	proto.RegisterType((*WaitRequest)(nil), "go.micro.run.WaitRequest")
	proto.RegisterType((*WaitResponse)(nil), "go.micro.run.WaitResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Runtime service

type RuntimeClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error)
}

type runtimeClient struct {
	c           client.Client
	serviceName string
}

func NewRuntimeClient(serviceName string, c client.Client) RuntimeClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.run"
	}
	return &runtimeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *runtimeClient) Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Fetch", in)
	out := new(FetchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Build", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Kill", in)
	out := new(KillResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Wait", &WaitRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeWaitClient{stream}, nil
}

type Runtime_WaitClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WaitResponse, error)
}

type runtimeWaitClient struct {
	stream client.Streamer
}

func (x *runtimeWaitClient) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitClient) Recv() (*WaitResponse, error) {
	m := new(WaitResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Fetch(context.Context, *FetchRequest, *FetchResponse) error
	Build(context.Context, *BuildRequest, *BuildResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
	Kill(context.Context, *KillRequest, *KillResponse) error
	Wait(context.Context, *WaitRequest, Runtime_WaitStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Runtime{hdlr}, opts...))
}

type Runtime struct {
	RuntimeHandler
}

func (h *Runtime) Fetch(ctx context.Context, in *FetchRequest, out *FetchResponse) error {
	return h.RuntimeHandler.Fetch(ctx, in, out)
}

func (h *Runtime) Build(ctx context.Context, in *BuildRequest, out *BuildResponse) error {
	return h.RuntimeHandler.Build(ctx, in, out)
}

func (h *Runtime) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.RuntimeHandler.Exec(ctx, in, out)
}

func (h *Runtime) Kill(ctx context.Context, in *KillRequest, out *KillResponse) error {
	return h.RuntimeHandler.Kill(ctx, in, out)
}

func (h *Runtime) Wait(ctx context.Context, stream server.Streamer) error {
	m := new(WaitRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Wait(ctx, m, &runtimeWaitStream{stream})
}

type Runtime_WaitStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WaitResponse) error
}

type runtimeWaitStream struct {
	stream server.Streamer
}

func (x *runtimeWaitStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitStream) Send(m *WaitResponse) error {
	return x.stream.Send(m)
}

func init() { proto.RegisterFile("github.com/micro/micro/run/proto/run.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x93, 0x3c, 0x8d, 0xbc, 0x6b, 0x94, 0xc7, 0xe0, 0x03, 0x5f, 0xde, 0x46, 0x86, 0x2e,
	0x4a, 0x91, 0x58, 0xec, 0xb2, 0xb5, 0x52, 0xa1, 0x2d, 0xb4, 0x9b, 0x92, 0x2e, 0xba, 0xd6, 0x64,
	0xd0, 0x81, 0x98, 0x49, 0x27, 0x19, 0x68, 0xbf, 0x68, 0x3f, 0x4f, 0x99, 0x3f, 0xd2, 0x49, 0xea,
	0x42, 0xbb, 0x91, 0x31, 0x73, 0xcf, 0xef, 0xdc, 0x7b, 0x4f, 0x08, 0x9c, 0xad, 0x69, 0xb5, 0x11,
	0xab, 0x28, 0x61, 0xdb, 0xc9, 0x96, 0x26, 0x9c, 0x99, 0x5f, 0x2e, 0xf2, 0x49, 0xc1, 0x59, 0xa5,
	0x4e, 0x91, 0x3a, 0xa1, 0x60, 0xcd, 0x22, 0x75, 0x1b, 0x71, 0x91, 0xe3, 0x31, 0xf8, 0xcf, 0x4c,
	0xf0, 0x84, 0xa0, 0x3f, 0xf0, 0x4b, 0xf0, 0x6c, 0xe8, 0x8e, 0xdc, 0xd3, 0xdf, 0xb1, 0x3c, 0xca,
	0x27, 0x29, 0xe5, 0x43, 0x4f, 0x3f, 0x49, 0x29, 0xc7, 0x0f, 0xe0, 0x2f, 0x68, 0xbe, 0xe4, 0xef,
	0x08, 0x41, 0xab, 0x58, 0x56, 0x1b, 0x53, 0xae, 0xce, 0x68, 0x0c, 0x7e, 0xa9, 0x58, 0x4a, 0xd2,
	0x9d, 0x0e, 0x22, 0xdb, 0x2a, 0xd2, 0x3e, 0xb1, 0xa9, 0xc1, 0xf7, 0xd0, 0x79, 0xe2, 0x2c, 0x21,
	0x65, 0x89, 0xfa, 0xe0, 0xd1, 0xd4, 0xa0, 0x3c, 0x9a, 0x4a, 0xd0, 0x4a, 0xd9, 0xec, 0x07, 0xe9,
	0x16, 0x62, 0x53, 0x83, 0x47, 0x10, 0xdc, 0x91, 0x2a, 0xd9, 0xc4, 0xe4, 0x55, 0x90, 0xb2, 0xfa,
	0x3e, 0x08, 0x9e, 0x41, 0xcf, 0x54, 0x94, 0x05, 0xcb, 0x4b, 0x62, 0x75, 0xea, 0x1e, 0xd0, 0xe9,
	0x15, 0x04, 0x0b, 0x41, 0xb3, 0x74, 0x67, 0x70, 0x9c, 0x7a, 0x06, 0x3d, 0xa3, 0xfe, 0x32, 0x37,
	0xd3, 0xb9, 0x07, 0x4c, 0x77, 0x09, 0xdd, 0xdb, 0x37, 0x92, 0x58, 0xde, 0x47, 0x88, 0xe7, 0x10,
	0x68, 0xb1, 0xb1, 0x9e, 0x40, 0xa7, 0xd0, 0x3b, 0x37, 0xf2, 0xbf, 0x75, 0xb9, 0x09, 0x24, 0xde,
	0x55, 0xe1, 0x6b, 0xe8, 0x3e, 0xd2, 0x2c, 0xdb, 0xb9, 0x1f, 0xad, 0xef, 0x43, 0xa0, 0xf5, 0xba,
	0x01, 0xc9, 0x7b, 0x59, 0xd2, 0xea, 0xc7, 0xbc, 0x13, 0x08, 0xb4, 0xde, 0x0c, 0x34, 0x80, 0x36,
	0xe1, 0x9c, 0x71, 0x93, 0xb6, 0xfe, 0x33, 0xfd, 0xf0, 0xa0, 0x13, 0x8b, 0xbc, 0xa2, 0x5b, 0x82,
	0x16, 0xd0, 0x56, 0xd9, 0xa3, 0xb0, 0x8e, 0xb6, 0x5f, 0x99, 0xf0, 0xff, 0xde, 0x3b, 0xd3, 0xb3,
	0x23, 0x19, 0x2a, 0xc2, 0x26, 0xc3, 0x7e, 0x2b, 0x9a, 0x8c, 0x5a, 0xe6, 0xd8, 0x41, 0x73, 0x68,
	0xc9, 0x28, 0xd0, 0xbf, 0x7a, 0x99, 0x95, 0x6d, 0x18, 0xee, 0xbb, 0xb2, 0x01, 0x72, 0x95, 0x4d,
	0x80, 0x15, 0x4f, 0x13, 0x50, 0xdb, 0xbc, 0x83, 0x6e, 0xa0, 0x25, 0x77, 0xd7, 0x04, 0x58, 0x79,
	0x34, 0x01, 0xf6, 0xaa, 0xb1, 0x73, 0xee, 0xae, 0x7c, 0xf5, 0x09, 0xb9, 0xf8, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x16, 0x83, 0x49, 0x54, 0x70, 0x04, 0x00, 0x00,
}