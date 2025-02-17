// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: audio/audio.proto

package audio

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AudioTask int32

const (
	AudioTask_TRANSCRIBE AudioTask = 0
	AudioTask_TRANSLATE  AudioTask = 1
)

// Enum value maps for AudioTask.
var (
	AudioTask_name = map[int32]string{
		0: "TRANSCRIBE",
		1: "TRANSLATE",
	}
	AudioTask_value = map[string]int32{
		"TRANSCRIBE": 0,
		"TRANSLATE":  1,
	}
)

func (x AudioTask) Enum() *AudioTask {
	p := new(AudioTask)
	*p = x
	return p
}

func (x AudioTask) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioTask) Descriptor() protoreflect.EnumDescriptor {
	return file_audio_audio_proto_enumTypes[0].Descriptor()
}

func (AudioTask) Type() protoreflect.EnumType {
	return &file_audio_audio_proto_enumTypes[0]
}

func (x AudioTask) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioTask.Descriptor instead.
func (AudioTask) EnumDescriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{0}
}

type AudioMetadata_AudioFormat int32

const (
	AudioMetadata_JSON         AudioMetadata_AudioFormat = 0
	AudioMetadata_TEXT         AudioMetadata_AudioFormat = 1
	AudioMetadata_SRT          AudioMetadata_AudioFormat = 2
	AudioMetadata_VERBOSE_JSON AudioMetadata_AudioFormat = 3
	AudioMetadata_VTT          AudioMetadata_AudioFormat = 4
)

// Enum value maps for AudioMetadata_AudioFormat.
var (
	AudioMetadata_AudioFormat_name = map[int32]string{
		0: "JSON",
		1: "TEXT",
		2: "SRT",
		3: "VERBOSE_JSON",
		4: "VTT",
	}
	AudioMetadata_AudioFormat_value = map[string]int32{
		"JSON":         0,
		"TEXT":         1,
		"SRT":          2,
		"VERBOSE_JSON": 3,
		"VTT":          4,
	}
)

func (x AudioMetadata_AudioFormat) Enum() *AudioMetadata_AudioFormat {
	p := new(AudioMetadata_AudioFormat)
	*p = x
	return p
}

func (x AudioMetadata_AudioFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioMetadata_AudioFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_audio_audio_proto_enumTypes[1].Descriptor()
}

func (AudioMetadata_AudioFormat) Type() protoreflect.EnumType {
	return &file_audio_audio_proto_enumTypes[1]
}

func (x AudioMetadata_AudioFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioMetadata_AudioFormat.Descriptor instead.
func (AudioMetadata_AudioFormat) EnumDescriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{0, 0}
}

type AudioMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt        string                    `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Temperature   float32                   `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Inputlanguage string                    `protobuf:"bytes,3,opt,name=inputlanguage,proto3" json:"inputlanguage,omitempty"`
	Format        AudioMetadata_AudioFormat `protobuf:"varint,4,opt,name=format,proto3,enum=audio.AudioMetadata_AudioFormat" json:"format,omitempty"`
}

func (x *AudioMetadata) Reset() {
	*x = AudioMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioMetadata) ProtoMessage() {}

func (x *AudioMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioMetadata.ProtoReflect.Descriptor instead.
func (*AudioMetadata) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{0}
}

func (x *AudioMetadata) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *AudioMetadata) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *AudioMetadata) GetInputlanguage() string {
	if x != nil {
		return x.Inputlanguage
	}
	return ""
}

func (x *AudioMetadata) GetFormat() AudioMetadata_AudioFormat {
	if x != nil {
		return x.Format
	}
	return AudioMetadata_JSON
}

type AudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*AudioRequest_Metadata
	//	*AudioRequest_ChunkData
	Request isAudioRequest_Request `protobuf_oneof:"request"`
}

func (x *AudioRequest) Reset() {
	*x = AudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioRequest) ProtoMessage() {}

func (x *AudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioRequest.ProtoReflect.Descriptor instead.
func (*AudioRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{1}
}

func (m *AudioRequest) GetRequest() isAudioRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *AudioRequest) GetMetadata() *AudioMetadata {
	if x, ok := x.GetRequest().(*AudioRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *AudioRequest) GetChunkData() []byte {
	if x, ok := x.GetRequest().(*AudioRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isAudioRequest_Request interface {
	isAudioRequest_Request()
}

type AudioRequest_Metadata struct {
	Metadata *AudioMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type AudioRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*AudioRequest_Metadata) isAudioRequest_Request() {}

func (*AudioRequest_ChunkData) isAudioRequest_Request() {}

type AudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     AudioTask                `protobuf:"varint,1,opt,name=task,proto3,enum=audio.AudioTask" json:"task,omitempty"`
	Language string                   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Duration float64                  `protobuf:"fixed64,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Segments []*AudioResponse_Segment `protobuf:"bytes,4,rep,name=segments,proto3" json:"segments,omitempty"`
	Text     string                   `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AudioResponse) Reset() {
	*x = AudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioResponse) ProtoMessage() {}

func (x *AudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioResponse.ProtoReflect.Descriptor instead.
func (*AudioResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{2}
}

func (x *AudioResponse) GetTask() AudioTask {
	if x != nil {
		return x.Task
	}
	return AudioTask_TRANSCRIBE
}

func (x *AudioResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *AudioResponse) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AudioResponse) GetSegments() []*AudioResponse_Segment {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *AudioResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AudioResponse_Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seek             int32   `protobuf:"varint,2,opt,name=seek,proto3" json:"seek,omitempty"`
	Start            float64 `protobuf:"fixed64,3,opt,name=start,proto3" json:"start,omitempty"`
	End              float64 `protobuf:"fixed64,4,opt,name=end,proto3" json:"end,omitempty"`
	Text             string  `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Tokens           []int32 `protobuf:"varint,6,rep,packed,name=tokens,proto3" json:"tokens,omitempty"`
	Temperature      float64 `protobuf:"fixed64,7,opt,name=temperature,proto3" json:"temperature,omitempty"`
	AvgLogprob       float64 `protobuf:"fixed64,8,opt,name=avg_logprob,json=avgLogprob,proto3" json:"avg_logprob,omitempty"`
	CompressionRatio float64 `protobuf:"fixed64,9,opt,name=compression_ratio,json=compressionRatio,proto3" json:"compression_ratio,omitempty"`
	NoSpeechProb     float64 `protobuf:"fixed64,10,opt,name=no_speech_prob,json=noSpeechProb,proto3" json:"no_speech_prob,omitempty"`
	Transient        bool    `protobuf:"varint,11,opt,name=transient,proto3" json:"transient,omitempty"`
}

func (x *AudioResponse_Segment) Reset() {
	*x = AudioResponse_Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioResponse_Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioResponse_Segment) ProtoMessage() {}

func (x *AudioResponse_Segment) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioResponse_Segment.ProtoReflect.Descriptor instead.
func (*AudioResponse_Segment) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AudioResponse_Segment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AudioResponse_Segment) GetSeek() int32 {
	if x != nil {
		return x.Seek
	}
	return 0
}

func (x *AudioResponse_Segment) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *AudioResponse_Segment) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *AudioResponse_Segment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AudioResponse_Segment) GetTokens() []int32 {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *AudioResponse_Segment) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *AudioResponse_Segment) GetAvgLogprob() float64 {
	if x != nil {
		return x.AvgLogprob
	}
	return 0
}

func (x *AudioResponse_Segment) GetCompressionRatio() float64 {
	if x != nil {
		return x.CompressionRatio
	}
	return 0
}

func (x *AudioResponse_Segment) GetNoSpeechProb() float64 {
	if x != nil {
		return x.NoSpeechProb
	}
	return 0
}

func (x *AudioResponse_Segment) GetTransient() bool {
	if x != nil {
		return x.Transient
	}
	return false
}

var File_audio_audio_proto protoreflect.FileDescriptor

var file_audio_audio_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x0d, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x45, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x52, 0x54,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x5f, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x54, 0x54, 0x10, 0x04, 0x22, 0x6e, 0x0a,
	0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf3, 0x03,
	0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x1a, 0xb5, 0x02, 0x0a, 0x07,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x65, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x62,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x76, 0x67, 0x4c, 0x6f, 0x67, 0x70, 0x72,
	0x6f, 0x62, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12,
	0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x70, 0x72, 0x6f,
	0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6e, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x50, 0x72, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x65, 0x6e, 0x74, 0x2a, 0x2a, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x01, 0x32,
	0x7c, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x13, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x73, 0x65, 0x75, 0x6e, 0x69, 0x63, 0x6f, 0x72, 0x6e, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x70,
	0x66, 0x72, 0x6f, 0x67, 0x61, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audio_audio_proto_rawDescOnce sync.Once
	file_audio_audio_proto_rawDescData = file_audio_audio_proto_rawDesc
)

func file_audio_audio_proto_rawDescGZIP() []byte {
	file_audio_audio_proto_rawDescOnce.Do(func() {
		file_audio_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_audio_audio_proto_rawDescData)
	})
	return file_audio_audio_proto_rawDescData
}

var file_audio_audio_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_audio_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_audio_audio_proto_goTypes = []interface{}{
	(AudioTask)(0),                 // 0: audio.AudioTask
	(AudioMetadata_AudioFormat)(0), // 1: audio.AudioMetadata.AudioFormat
	(*AudioMetadata)(nil),          // 2: audio.AudioMetadata
	(*AudioRequest)(nil),           // 3: audio.AudioRequest
	(*AudioResponse)(nil),          // 4: audio.AudioResponse
	(*AudioResponse_Segment)(nil),  // 5: audio.AudioResponse.Segment
}
var file_audio_audio_proto_depIdxs = []int32{
	1, // 0: audio.AudioMetadata.format:type_name -> audio.AudioMetadata.AudioFormat
	2, // 1: audio.AudioRequest.metadata:type_name -> audio.AudioMetadata
	0, // 2: audio.AudioResponse.task:type_name -> audio.AudioTask
	5, // 3: audio.AudioResponse.segments:type_name -> audio.AudioResponse.Segment
	3, // 4: audio.Audio.Translate:input_type -> audio.AudioRequest
	3, // 5: audio.Audio.Transcribe:input_type -> audio.AudioRequest
	4, // 6: audio.Audio.Translate:output_type -> audio.AudioResponse
	4, // 7: audio.Audio.Transcribe:output_type -> audio.AudioResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_audio_audio_proto_init() }
func file_audio_audio_proto_init() {
	if File_audio_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audio_audio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_audio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_audio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_audio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioResponse_Segment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_audio_audio_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AudioRequest_Metadata)(nil),
		(*AudioRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audio_audio_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audio_audio_proto_goTypes,
		DependencyIndexes: file_audio_audio_proto_depIdxs,
		EnumInfos:         file_audio_audio_proto_enumTypes,
		MessageInfos:      file_audio_audio_proto_msgTypes,
	}.Build()
	File_audio_audio_proto = out.File
	file_audio_audio_proto_rawDesc = nil
	file_audio_audio_proto_goTypes = nil
	file_audio_audio_proto_depIdxs = nil
}
