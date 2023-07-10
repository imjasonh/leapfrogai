// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: audio/audio.proto

package audio

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AudioClient is the client API for Audio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioClient interface {
	Translate(ctx context.Context, opts ...grpc.CallOption) (Audio_TranslateClient, error)
	Transcribe(ctx context.Context, opts ...grpc.CallOption) (Audio_TranscribeClient, error)
}

type audioClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioClient(cc grpc.ClientConnInterface) AudioClient {
	return &audioClient{cc}
}

func (c *audioClient) Translate(ctx context.Context, opts ...grpc.CallOption) (Audio_TranslateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Audio_ServiceDesc.Streams[0], "/audio.Audio/Translate", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioTranslateClient{stream}
	return x, nil
}

type Audio_TranslateClient interface {
	Send(*AudioRequest) error
	CloseAndRecv() (*AudioResponse, error)
	grpc.ClientStream
}

type audioTranslateClient struct {
	grpc.ClientStream
}

func (x *audioTranslateClient) Send(m *AudioRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioTranslateClient) CloseAndRecv() (*AudioResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AudioResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *audioClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (Audio_TranscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Audio_ServiceDesc.Streams[1], "/audio.Audio/Transcribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioTranscribeClient{stream}
	return x, nil
}

type Audio_TranscribeClient interface {
	Send(*AudioRequest) error
	CloseAndRecv() (*AudioResponse, error)
	grpc.ClientStream
}

type audioTranscribeClient struct {
	grpc.ClientStream
}

func (x *audioTranscribeClient) Send(m *AudioRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioTranscribeClient) CloseAndRecv() (*AudioResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AudioResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioServer is the server API for Audio service.
// All implementations must embed UnimplementedAudioServer
// for forward compatibility
type AudioServer interface {
	Translate(Audio_TranslateServer) error
	Transcribe(Audio_TranscribeServer) error
	mustEmbedUnimplementedAudioServer()
}

// UnimplementedAudioServer must be embedded to have forward compatible implementations.
type UnimplementedAudioServer struct {
}

func (UnimplementedAudioServer) Translate(Audio_TranslateServer) error {
	return status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedAudioServer) Transcribe(Audio_TranscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
func (UnimplementedAudioServer) mustEmbedUnimplementedAudioServer() {}

// UnsafeAudioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServer will
// result in compilation errors.
type UnsafeAudioServer interface {
	mustEmbedUnimplementedAudioServer()
}

func RegisterAudioServer(s grpc.ServiceRegistrar, srv AudioServer) {
	s.RegisterService(&Audio_ServiceDesc, srv)
}

func _Audio_Translate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioServer).Translate(&audioTranslateServer{stream})
}

type Audio_TranslateServer interface {
	SendAndClose(*AudioResponse) error
	Recv() (*AudioRequest, error)
	grpc.ServerStream
}

type audioTranslateServer struct {
	grpc.ServerStream
}

func (x *audioTranslateServer) SendAndClose(m *AudioResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioTranslateServer) Recv() (*AudioRequest, error) {
	m := new(AudioRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Audio_Transcribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioServer).Transcribe(&audioTranscribeServer{stream})
}

type Audio_TranscribeServer interface {
	SendAndClose(*AudioResponse) error
	Recv() (*AudioRequest, error)
	grpc.ServerStream
}

type audioTranscribeServer struct {
	grpc.ServerStream
}

func (x *audioTranscribeServer) SendAndClose(m *AudioResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioTranscribeServer) Recv() (*AudioRequest, error) {
	m := new(AudioRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Audio_ServiceDesc is the grpc.ServiceDesc for Audio service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Audio_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audio.Audio",
	HandlerType: (*AudioServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Translate",
			Handler:       _Audio_Translate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Transcribe",
			Handler:       _Audio_Transcribe_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "audio/audio.proto",
}
