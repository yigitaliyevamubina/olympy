// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_service/streaming_service.proto

package _

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StreamEventRequest struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventRequest) Reset()         { *m = StreamEventRequest{} }
func (m *StreamEventRequest) String() string { return proto.CompactTextString(m) }
func (*StreamEventRequest) ProtoMessage()    {}
func (*StreamEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_759800fa5674cc87, []int{0}
}
func (m *StreamEventRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamEventRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventRequest.Merge(m, src)
}
func (m *StreamEventRequest) XXX_Size() int {
	return m.Size()
}
func (m *StreamEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventRequest proto.InternalMessageInfo

func (m *StreamEventRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *StreamEventRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type StreamEventResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventResponse) Reset()         { *m = StreamEventResponse{} }
func (m *StreamEventResponse) String() string { return proto.CompactTextString(m) }
func (*StreamEventResponse) ProtoMessage()    {}
func (*StreamEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_759800fa5674cc87, []int{1}
}
func (m *StreamEventResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamEventResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventResponse.Merge(m, src)
}
func (m *StreamEventResponse) XXX_Size() int {
	return m.Size()
}
func (m *StreamEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventResponse proto.InternalMessageInfo

func (m *StreamEventResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamEventRequest)(nil), "streaming_service.StreamEventRequest")
	proto.RegisterType((*StreamEventResponse)(nil), "streaming_service.StreamEventResponse")
}

func init() {
	proto.RegisterFile("stream_service/streaming_service.proto", fileDescriptor_759800fa5674cc87)
}

var fileDescriptor_759800fa5674cc87 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x87, 0x70, 0x33, 0xf3, 0xd2,
	0x61, 0x22, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x18, 0x12, 0x4a, 0xce, 0x5c, 0x42,
	0xc1, 0x60, 0x41, 0xd7, 0xb2, 0xd4, 0xbc, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x49, 0x2e, 0x8e, 0x54, 0x10, 0x3f, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88,
	0x1d, 0xcc, 0xf7, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x02, 0x0b,
	0x83, 0xd9, 0x4a, 0xfa, 0x5c, 0xc2, 0x28, 0x86, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49,
	0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0xc2, 0x0c, 0x81, 0x72, 0x8d, 0x0a, 0xb8,
	0x04, 0x82, 0x61, 0x4e, 0x09, 0x86, 0xb8, 0x44, 0x28, 0x86, 0x8b, 0x1b, 0xc9, 0x10, 0x21, 0x55,
	0x3d, 0x4c, 0x5f, 0x60, 0xba, 0x54, 0x4a, 0x8d, 0x90, 0x32, 0x88, 0x5b, 0x9c, 0x84, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2,
	0x18, 0xf5, 0x92, 0xd8, 0xc0, 0xc1, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xfe, 0x53,
	0x3a, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingServiceClient interface {
	StreamEvent(ctx context.Context, in *StreamEventRequest, opts ...grpc.CallOption) (*StreamEventResponse, error)
}

type streamingServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamingServiceClient(cc *grpc.ClientConn) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) StreamEvent(ctx context.Context, in *StreamEventRequest, opts ...grpc.CallOption) (*StreamEventResponse, error) {
	out := new(StreamEventResponse)
	err := c.cc.Invoke(ctx, "/streaming_service.StreamingService/StreamEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamingServiceServer is the server API for StreamingService service.
type StreamingServiceServer interface {
	StreamEvent(context.Context, *StreamEventRequest) (*StreamEventResponse, error)
}

// UnimplementedStreamingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingServiceServer struct {
}

func (*UnimplementedStreamingServiceServer) StreamEvent(ctx context.Context, req *StreamEventRequest) (*StreamEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamEvent not implemented")
}

func RegisterStreamingServiceServer(s *grpc.Server, srv StreamingServiceServer) {
	s.RegisterService(&_StreamingService_serviceDesc, srv)
}

func _StreamingService_StreamEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingServiceServer).StreamEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streaming_service.StreamingService/StreamEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingServiceServer).StreamEvent(ctx, req.(*StreamEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "streaming_service.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamEvent",
			Handler:    _StreamingService_StreamEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream_service/streaming_service.proto",
}

func (m *StreamEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEventRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEventRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintStreamingService(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EventId) > 0 {
		i -= len(m.EventId)
		copy(dAtA[i:], m.EventId)
		i = encodeVarintStreamingService(dAtA, i, uint64(len(m.EventId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StreamEventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEventResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEventResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintStreamingService(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreamingService(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreamingService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EventId)
	if l > 0 {
		n += 1 + l + sovStreamingService(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovStreamingService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamEventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovStreamingService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStreamingService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreamingService(x uint64) (n int) {
	return sovStreamingService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamingService
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
			return fmt.Errorf("proto: StreamEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingService
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
				return ErrInvalidLengthStreamingService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingService
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
				return ErrInvalidLengthStreamingService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreamingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamingService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamEventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamingService
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
			return fmt.Errorf("proto: StreamEventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamEventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingService
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
				return ErrInvalidLengthStreamingService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreamingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamingService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStreamingService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreamingService
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
					return 0, ErrIntOverflowStreamingService
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
					return 0, ErrIntOverflowStreamingService
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
				return 0, ErrInvalidLengthStreamingService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreamingService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreamingService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreamingService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreamingService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreamingService = fmt.Errorf("proto: unexpected end of group")
)
