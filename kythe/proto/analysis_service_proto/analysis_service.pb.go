// Code generated by protoc-gen-gogo.
// source: kythe/proto/analysis_service.proto
// DO NOT EDIT!

/*
	Package analysis_service_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/analysis_service.proto

	It has these top-level messages:
*/
package analysis_service_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "kythe.io/third_party/proto/any_proto"
import kythe_proto2 "kythe.io/kythe/proto/analysis_proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for CompilationAnalyzer service

type CompilationAnalyzerClient interface {
	// Analyze is the main entry point for the analysis driver to send work to the
	// analyzer.  The analysis may produce many outputs which will be streamed as
	// framed AnalysisOutput messages.
	//
	// A driver may choose to retry analyses that return RPC errors.  It should
	// not retry analyses that are reported as finished unless it is necessary to
	// recover from an external production issue.
	Analyze(ctx context.Context, in *kythe_proto2.AnalysisRequest, opts ...grpc.CallOption) (CompilationAnalyzer_AnalyzeClient, error)
}

type compilationAnalyzerClient struct {
	cc *grpc.ClientConn
}

func NewCompilationAnalyzerClient(cc *grpc.ClientConn) CompilationAnalyzerClient {
	return &compilationAnalyzerClient{cc}
}

func (c *compilationAnalyzerClient) Analyze(ctx context.Context, in *kythe_proto2.AnalysisRequest, opts ...grpc.CallOption) (CompilationAnalyzer_AnalyzeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CompilationAnalyzer_serviceDesc.Streams[0], c.cc, "/kythe.proto.CompilationAnalyzer/Analyze", opts...)
	if err != nil {
		return nil, err
	}
	x := &compilationAnalyzerAnalyzeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompilationAnalyzer_AnalyzeClient interface {
	Recv() (*kythe_proto2.AnalysisOutput, error)
	grpc.ClientStream
}

type compilationAnalyzerAnalyzeClient struct {
	grpc.ClientStream
}

func (x *compilationAnalyzerAnalyzeClient) Recv() (*kythe_proto2.AnalysisOutput, error) {
	m := new(kythe_proto2.AnalysisOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CompilationAnalyzer service

type CompilationAnalyzerServer interface {
	// Analyze is the main entry point for the analysis driver to send work to the
	// analyzer.  The analysis may produce many outputs which will be streamed as
	// framed AnalysisOutput messages.
	//
	// A driver may choose to retry analyses that return RPC errors.  It should
	// not retry analyses that are reported as finished unless it is necessary to
	// recover from an external production issue.
	Analyze(*kythe_proto2.AnalysisRequest, CompilationAnalyzer_AnalyzeServer) error
}

func RegisterCompilationAnalyzerServer(s *grpc.Server, srv CompilationAnalyzerServer) {
	s.RegisterService(&_CompilationAnalyzer_serviceDesc, srv)
}

func _CompilationAnalyzer_Analyze_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(kythe_proto2.AnalysisRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompilationAnalyzerServer).Analyze(m, &compilationAnalyzerAnalyzeServer{stream})
}

type CompilationAnalyzer_AnalyzeServer interface {
	Send(*kythe_proto2.AnalysisOutput) error
	grpc.ServerStream
}

type compilationAnalyzerAnalyzeServer struct {
	grpc.ServerStream
}

func (x *compilationAnalyzerAnalyzeServer) Send(m *kythe_proto2.AnalysisOutput) error {
	return x.ServerStream.SendMsg(m)
}

var _CompilationAnalyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.CompilationAnalyzer",
	HandlerType: (*CompilationAnalyzerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Analyze",
			Handler:       _CompilationAnalyzer_Analyze_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for FileDataService service

type FileDataServiceClient interface {
	// Get returns the contents of one or more files needed for analysis.  It is
	// the server's responsibility to do any caching necessary to make this
	// perform well, so that an analyzer does not need to implement its own
	// caches unless it is doing something unusual.
	//
	// For each distinct path/digest pair in the request, the server must return
	// exactly one response.  The order of the responses is arbitrary.
	//
	// For each requested file, one or both of the path and digest fields must be
	// nonempty, otherwise an error is returned.  It is not an error for there to
	// be no requested files, however.
	Get(ctx context.Context, in *kythe_proto2.FilesRequest, opts ...grpc.CallOption) (FileDataService_GetClient, error)
}

type fileDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileDataServiceClient(cc *grpc.ClientConn) FileDataServiceClient {
	return &fileDataServiceClient{cc}
}

func (c *fileDataServiceClient) Get(ctx context.Context, in *kythe_proto2.FilesRequest, opts ...grpc.CallOption) (FileDataService_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FileDataService_serviceDesc.Streams[0], c.cc, "/kythe.proto.FileDataService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileDataServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileDataService_GetClient interface {
	Recv() (*kythe_proto2.FileData, error)
	grpc.ClientStream
}

type fileDataServiceGetClient struct {
	grpc.ClientStream
}

func (x *fileDataServiceGetClient) Recv() (*kythe_proto2.FileData, error) {
	m := new(kythe_proto2.FileData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FileDataService service

type FileDataServiceServer interface {
	// Get returns the contents of one or more files needed for analysis.  It is
	// the server's responsibility to do any caching necessary to make this
	// perform well, so that an analyzer does not need to implement its own
	// caches unless it is doing something unusual.
	//
	// For each distinct path/digest pair in the request, the server must return
	// exactly one response.  The order of the responses is arbitrary.
	//
	// For each requested file, one or both of the path and digest fields must be
	// nonempty, otherwise an error is returned.  It is not an error for there to
	// be no requested files, however.
	Get(*kythe_proto2.FilesRequest, FileDataService_GetServer) error
}

func RegisterFileDataServiceServer(s *grpc.Server, srv FileDataServiceServer) {
	s.RegisterService(&_FileDataService_serviceDesc, srv)
}

func _FileDataService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(kythe_proto2.FilesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileDataServiceServer).Get(m, &fileDataServiceGetServer{stream})
}

type FileDataService_GetServer interface {
	Send(*kythe_proto2.FileData) error
	grpc.ServerStream
}

type fileDataServiceGetServer struct {
	grpc.ServerStream
}

func (x *fileDataServiceGetServer) Send(m *kythe_proto2.FileData) error {
	return x.ServerStream.SendMsg(m)
}

var _FileDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.FileDataService",
	HandlerType: (*FileDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _FileDataService_Get_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptorAnalysisService = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xae, 0x2c, 0xc9,
	0x48, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcc, 0x4b, 0xcc, 0xa9, 0x2c, 0xce, 0x2c,
	0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x0b, 0x0b, 0x71, 0x83, 0xd5, 0x40,
	0x38, 0x52, 0x92, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0x50, 0x1d, 0x49, 0xa5, 0x69, 0xfa, 0x89, 0x79,
	0x95, 0x50, 0x29, 0x29, 0x6c, 0x66, 0x41, 0xe4, 0x8c, 0xe2, 0xb9, 0x84, 0x9d, 0xf3, 0x73, 0x0b,
	0x32, 0x73, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x1c, 0x41, 0x92, 0x55, 0xa9, 0x45, 0x42, 0x1e, 0x5c,
	0xec, 0x50, 0xb6, 0x90, 0x8c, 0x1e, 0x92, 0x35, 0x7a, 0x8e, 0x50, 0xed, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x52, 0xd2, 0x58, 0x65, 0xfd, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0x94, 0x18, 0x0c,
	0x18, 0x8d, 0xfc, 0xb8, 0xf8, 0xdd, 0x32, 0x73, 0x52, 0x5d, 0x12, 0x4b, 0x12, 0x83, 0x21, 0xae,
	0x17, 0xb2, 0xe6, 0x62, 0x76, 0x4f, 0x2d, 0x11, 0x92, 0x44, 0xd1, 0x0a, 0x52, 0x04, 0x37, 0x55,
	0x14, 0x43, 0x0a, 0xa4, 0x1f, 0x64, 0x9e, 0x93, 0xe1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x03, 0x97, 0x7c, 0x72, 0x7e, 0xae, 0x1e,
	0xc4, 0xf7, 0x7a, 0x29, 0xa9, 0x65, 0x25, 0xf9, 0xf9, 0x39, 0xc5, 0xc8, 0xfa, 0x93, 0xd8, 0xc0,
	0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x12, 0x94, 0xdf, 0xe1, 0x54, 0x01, 0x00, 0x00,
}
