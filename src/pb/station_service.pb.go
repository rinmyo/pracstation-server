// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: station_service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RefreshStationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*RefreshStationResponse_Signal
	//	*RefreshStationResponse_Turnout
	//	*RefreshStationResponse_Section
	Value isRefreshStationResponse_Value `protobuf_oneof:"value"`
}

func (x *RefreshStationResponse) Reset() {
	*x = RefreshStationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshStationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshStationResponse) ProtoMessage() {}

func (x *RefreshStationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshStationResponse.ProtoReflect.Descriptor instead.
func (*RefreshStationResponse) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{0}
}

func (m *RefreshStationResponse) GetValue() isRefreshStationResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *RefreshStationResponse) GetSignal() *Signal {
	if x, ok := x.GetValue().(*RefreshStationResponse_Signal); ok {
		return x.Signal
	}
	return nil
}

func (x *RefreshStationResponse) GetTurnout() *Turnout {
	if x, ok := x.GetValue().(*RefreshStationResponse_Turnout); ok {
		return x.Turnout
	}
	return nil
}

func (x *RefreshStationResponse) GetSection() *Section {
	if x, ok := x.GetValue().(*RefreshStationResponse_Section); ok {
		return x.Section
	}
	return nil
}

type isRefreshStationResponse_Value interface {
	isRefreshStationResponse_Value()
}

type RefreshStationResponse_Signal struct {
	Signal *Signal `protobuf:"bytes,1,opt,name=signal,proto3,oneof"`
}

type RefreshStationResponse_Turnout struct {
	Turnout *Turnout `protobuf:"bytes,2,opt,name=turnout,proto3,oneof"`
}

type RefreshStationResponse_Section struct {
	Section *Section `protobuf:"bytes,3,opt,name=section,proto3,oneof"`
}

func (*RefreshStationResponse_Signal) isRefreshStationResponse_Value() {}

func (*RefreshStationResponse_Turnout) isRefreshStationResponse_Value() {}

func (*RefreshStationResponse_Section) isRefreshStationResponse_Value() {}

type SearchRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buttons *ButtonList `protobuf:"bytes,1,opt,name=buttons,proto3" json:"buttons,omitempty"`
}

func (x *SearchRouteRequest) Reset() {
	*x = SearchRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRouteRequest) ProtoMessage() {}

func (x *SearchRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRouteRequest.ProtoReflect.Descriptor instead.
func (*SearchRouteRequest) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRouteRequest) GetButtons() *ButtonList {
	if x != nil {
		return x.Buttons
	}
	return nil
}

type SearchRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *SearchRouteResponse) Reset() {
	*x = SearchRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRouteResponse) ProtoMessage() {}

func (x *SearchRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRouteResponse.ProtoReflect.Descriptor instead.
func (*SearchRouteResponse) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRouteResponse) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type NewRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *NewRouteRequest) Reset() {
	*x = NewRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRouteRequest) ProtoMessage() {}

func (x *NewRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRouteRequest.ProtoReflect.Descriptor instead.
func (*NewRouteRequest) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{3}
}

func (x *NewRouteRequest) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

type CancelRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *CancelRouteRequest) Reset() {
	*x = CancelRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRouteRequest) ProtoMessage() {}

func (x *CancelRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRouteRequest.ProtoReflect.Descriptor instead.
func (*CancelRouteRequest) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{4}
}

func (x *CancelRouteRequest) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

var File_station_service_proto protoreflect.FileDescriptor

var file_station_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65,
	0x74, 0x1a, 0x19, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x74,
	0x75, 0x72, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x74, 0x75, 0x72, 0x6e,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x61, 0x63,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x75, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x74, 0x75, 0x72, 0x6e, 0x6f, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x44, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x62, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x64, 0x32, 0xb9, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x20, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65,
	0x77, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x24,
	0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6c, 0x79, 0x63, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x61, 0x63, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x50, 0x01, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_station_service_proto_rawDescOnce sync.Once
	file_station_service_proto_rawDescData = file_station_service_proto_rawDesc
)

func file_station_service_proto_rawDescGZIP() []byte {
	file_station_service_proto_rawDescOnce.Do(func() {
		file_station_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_station_service_proto_rawDescData)
	})
	return file_station_service_proto_rawDescData
}

var file_station_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_station_service_proto_goTypes = []interface{}{
	(*RefreshStationResponse)(nil), // 0: prac.net.RefreshStationResponse
	(*SearchRouteRequest)(nil),     // 1: prac.net.SearchRouteRequest
	(*SearchRouteResponse)(nil),    // 2: prac.net.SearchRouteResponse
	(*NewRouteRequest)(nil),        // 3: prac.net.NewRouteRequest
	(*CancelRouteRequest)(nil),     // 4: prac.net.CancelRouteRequest
	(*Signal)(nil),                 // 5: prac.net.Signal
	(*Turnout)(nil),                // 6: prac.net.Turnout
	(*Section)(nil),                // 7: prac.net.Section
	(*ButtonList)(nil),             // 8: prac.net.ButtonList
	(*Route)(nil),                  // 9: prac.net.Route
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_station_service_proto_depIdxs = []int32{
	5,  // 0: prac.net.RefreshStationResponse.signal:type_name -> prac.net.Signal
	6,  // 1: prac.net.RefreshStationResponse.turnout:type_name -> prac.net.Turnout
	7,  // 2: prac.net.RefreshStationResponse.section:type_name -> prac.net.Section
	8,  // 3: prac.net.SearchRouteRequest.buttons:type_name -> prac.net.ButtonList
	9,  // 4: prac.net.SearchRouteResponse.route:type_name -> prac.net.Route
	10, // 5: prac.net.StationService.RefreshStation:input_type -> google.protobuf.Empty
	1,  // 6: prac.net.StationService.SearchRoute:input_type -> prac.net.SearchRouteRequest
	3,  // 7: prac.net.StationService.CreateRoute:input_type -> prac.net.NewRouteRequest
	4,  // 8: prac.net.StationService.CancelRoute:input_type -> prac.net.CancelRouteRequest
	0,  // 9: prac.net.StationService.RefreshStation:output_type -> prac.net.RefreshStationResponse
	2,  // 10: prac.net.StationService.SearchRoute:output_type -> prac.net.SearchRouteResponse
	10, // 11: prac.net.StationService.CreateRoute:output_type -> google.protobuf.Empty
	10, // 12: prac.net.StationService.CancelRoute:output_type -> google.protobuf.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_station_service_proto_init() }
func file_station_service_proto_init() {
	if File_station_service_proto != nil {
		return
	}
	file_button_list_message_proto_init()
	file_path_message_proto_init()
	file_signal_message_proto_init()
	file_section_message_proto_init()
	file_turnout_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_station_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshStationResponse); i {
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
		file_station_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRouteRequest); i {
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
		file_station_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRouteResponse); i {
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
		file_station_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRouteRequest); i {
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
		file_station_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRouteRequest); i {
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
	file_station_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RefreshStationResponse_Signal)(nil),
		(*RefreshStationResponse_Turnout)(nil),
		(*RefreshStationResponse_Section)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_station_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_station_service_proto_goTypes,
		DependencyIndexes: file_station_service_proto_depIdxs,
		MessageInfos:      file_station_service_proto_msgTypes,
	}.Build()
	File_station_service_proto = out.File
	file_station_service_proto_rawDesc = nil
	file_station_service_proto_goTypes = nil
	file_station_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StationServiceClient is the client API for StationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StationServiceClient interface {
	RefreshStation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StationService_RefreshStationClient, error)
	SearchRoute(ctx context.Context, in *SearchRouteRequest, opts ...grpc.CallOption) (*SearchRouteResponse, error)
	CreateRoute(ctx context.Context, in *NewRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelRoute(ctx context.Context, in *CancelRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type stationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStationServiceClient(cc grpc.ClientConnInterface) StationServiceClient {
	return &stationServiceClient{cc}
}

func (c *stationServiceClient) RefreshStation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StationService_RefreshStationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StationService_serviceDesc.Streams[0], "/prac.net.StationService/RefreshStation", opts...)
	if err != nil {
		return nil, err
	}
	x := &stationServiceRefreshStationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StationService_RefreshStationClient interface {
	Recv() (*RefreshStationResponse, error)
	grpc.ClientStream
}

type stationServiceRefreshStationClient struct {
	grpc.ClientStream
}

func (x *stationServiceRefreshStationClient) Recv() (*RefreshStationResponse, error) {
	m := new(RefreshStationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stationServiceClient) SearchRoute(ctx context.Context, in *SearchRouteRequest, opts ...grpc.CallOption) (*SearchRouteResponse, error) {
	out := new(SearchRouteResponse)
	err := c.cc.Invoke(ctx, "/prac.net.StationService/SearchRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationServiceClient) CreateRoute(ctx context.Context, in *NewRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/prac.net.StationService/CreateRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationServiceClient) CancelRoute(ctx context.Context, in *CancelRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/prac.net.StationService/CancelRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StationServiceServer is the server API for StationService service.
type StationServiceServer interface {
	RefreshStation(*emptypb.Empty, StationService_RefreshStationServer) error
	SearchRoute(context.Context, *SearchRouteRequest) (*SearchRouteResponse, error)
	CreateRoute(context.Context, *NewRouteRequest) (*emptypb.Empty, error)
	CancelRoute(context.Context, *CancelRouteRequest) (*emptypb.Empty, error)
}

// UnimplementedStationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStationServiceServer struct {
}

func (*UnimplementedStationServiceServer) RefreshStation(*emptypb.Empty, StationService_RefreshStationServer) error {
	return status.Errorf(codes.Unimplemented, "method RefreshStation not implemented")
}
func (*UnimplementedStationServiceServer) SearchRoute(context.Context, *SearchRouteRequest) (*SearchRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRoute not implemented")
}
func (*UnimplementedStationServiceServer) CreateRoute(context.Context, *NewRouteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoute not implemented")
}
func (*UnimplementedStationServiceServer) CancelRoute(context.Context, *CancelRouteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRoute not implemented")
}

func RegisterStationServiceServer(s *grpc.Server, srv StationServiceServer) {
	s.RegisterService(&_StationService_serviceDesc, srv)
}

func _StationService_RefreshStation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StationServiceServer).RefreshStation(m, &stationServiceRefreshStationServer{stream})
}

type StationService_RefreshStationServer interface {
	Send(*RefreshStationResponse) error
	grpc.ServerStream
}

type stationServiceRefreshStationServer struct {
	grpc.ServerStream
}

func (x *stationServiceRefreshStationServer) Send(m *RefreshStationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StationService_SearchRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationServiceServer).SearchRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prac.net.StationService/SearchRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationServiceServer).SearchRoute(ctx, req.(*SearchRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StationService_CreateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationServiceServer).CreateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prac.net.StationService/CreateRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationServiceServer).CreateRoute(ctx, req.(*NewRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StationService_CancelRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationServiceServer).CancelRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prac.net.StationService/CancelRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationServiceServer).CancelRoute(ctx, req.(*CancelRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prac.net.StationService",
	HandlerType: (*StationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchRoute",
			Handler:    _StationService_SearchRoute_Handler,
		},
		{
			MethodName: "CreateRoute",
			Handler:    _StationService_CreateRoute_Handler,
		},
		{
			MethodName: "CancelRoute",
			Handler:    _StationService_CancelRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RefreshStation",
			Handler:       _StationService_RefreshStation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "station_service.proto",
}
