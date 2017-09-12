// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/view.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha_paint "gomatcha.io/matcha/pb/paint"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuildNode struct {
	Id          int64                           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BuildId     int64                           `protobuf:"varint,2,opt,name=buildId" json:"buildId,omitempty"`
	BridgeName  string                          `protobuf:"bytes,3,opt,name=bridgeName" json:"bridgeName,omitempty"`
	BridgeValue *google_protobuf.Any            `protobuf:"bytes,4,opt,name=bridgeValue" json:"bridgeValue,omitempty"`
	Values      map[string]*google_protobuf.Any `protobuf:"bytes,5,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Children    []int64                         `protobuf:"varint,6,rep,packed,name=children" json:"children,omitempty"`
	AltIds      map[int64]int64                 `protobuf:"bytes,7,rep,name=altIds" json:"altIds,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *BuildNode) Reset()                    { *m = BuildNode{} }
func (m *BuildNode) String() string            { return proto.CompactTextString(m) }
func (*BuildNode) ProtoMessage()               {}
func (*BuildNode) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BuildNode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BuildNode) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *BuildNode) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *BuildNode) GetBridgeValue() *google_protobuf.Any {
	if m != nil {
		return m.BridgeValue
	}
	return nil
}

func (m *BuildNode) GetValues() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *BuildNode) GetChildren() []int64 {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *BuildNode) GetAltIds() map[int64]int64 {
	if m != nil {
		return m.AltIds
	}
	return nil
}

type LayoutPaintNode struct {
	Id       int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LayoutId int64 `protobuf:"varint,2,opt,name=layoutId" json:"layoutId,omitempty"`
	PaintId  int64 `protobuf:"varint,3,opt,name=paintId" json:"paintId,omitempty"`
	// matcha.layout.Guide layoutGuide = 4;
	// Guide
	Minx       float64             `protobuf:"fixed64,4,opt,name=minx" json:"minx,omitempty"`
	Miny       float64             `protobuf:"fixed64,5,opt,name=miny" json:"miny,omitempty"`
	Maxx       float64             `protobuf:"fixed64,6,opt,name=maxx" json:"maxx,omitempty"`
	Maxy       float64             `protobuf:"fixed64,7,opt,name=maxy" json:"maxy,omitempty"`
	ZIndex     int64               `protobuf:"varint,8,opt,name=zIndex" json:"zIndex,omitempty"`
	ChildOrder []int64             `protobuf:"varint,9,rep,packed,name=childOrder" json:"childOrder,omitempty"`
	PaintStyle *matcha_paint.Style `protobuf:"bytes,10,opt,name=paintStyle" json:"paintStyle,omitempty"`
}

func (m *LayoutPaintNode) Reset()                    { *m = LayoutPaintNode{} }
func (m *LayoutPaintNode) String() string            { return proto.CompactTextString(m) }
func (*LayoutPaintNode) ProtoMessage()               {}
func (*LayoutPaintNode) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *LayoutPaintNode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LayoutPaintNode) GetLayoutId() int64 {
	if m != nil {
		return m.LayoutId
	}
	return 0
}

func (m *LayoutPaintNode) GetPaintId() int64 {
	if m != nil {
		return m.PaintId
	}
	return 0
}

func (m *LayoutPaintNode) GetMinx() float64 {
	if m != nil {
		return m.Minx
	}
	return 0
}

func (m *LayoutPaintNode) GetMiny() float64 {
	if m != nil {
		return m.Miny
	}
	return 0
}

func (m *LayoutPaintNode) GetMaxx() float64 {
	if m != nil {
		return m.Maxx
	}
	return 0
}

func (m *LayoutPaintNode) GetMaxy() float64 {
	if m != nil {
		return m.Maxy
	}
	return 0
}

func (m *LayoutPaintNode) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func (m *LayoutPaintNode) GetChildOrder() []int64 {
	if m != nil {
		return m.ChildOrder
	}
	return nil
}

func (m *LayoutPaintNode) GetPaintStyle() *matcha_paint.Style {
	if m != nil {
		return m.PaintStyle
	}
	return nil
}

type Root struct {
	LayoutPaintNodes map[int64]*LayoutPaintNode      `protobuf:"bytes,2,rep,name=layoutPaintNodes" json:"layoutPaintNodes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BuildNodes       map[int64]*BuildNode            `protobuf:"bytes,3,rep,name=buildNodes" json:"buildNodes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Middleware       map[string]*google_protobuf.Any `protobuf:"bytes,4,rep,name=middleware" json:"middleware,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Root) GetLayoutPaintNodes() map[int64]*LayoutPaintNode {
	if m != nil {
		return m.LayoutPaintNodes
	}
	return nil
}

func (m *Root) GetBuildNodes() map[int64]*BuildNode {
	if m != nil {
		return m.BuildNodes
	}
	return nil
}

func (m *Root) GetMiddleware() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Middleware
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildNode)(nil), "matcha.view.BuildNode")
	proto.RegisterType((*LayoutPaintNode)(nil), "matcha.view.LayoutPaintNode")
	proto.RegisterType((*Root)(nil), "matcha.view.Root")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/view.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xe2, 0x2e, 0x5b, 0x4f, 0x10, 0x9b, 0xcc, 0x98, 0x4c, 0x84, 0x50, 0xa9, 0x84, 0x56,
	0x21, 0x94, 0x4a, 0x9b, 0x84, 0xd8, 0xee, 0x56, 0x89, 0x8b, 0x4a, 0x6c, 0xab, 0x52, 0xb1, 0x0b,
	0xee, 0x9c, 0xd9, 0x74, 0x16, 0x69, 0x52, 0xa5, 0xe9, 0x9a, 0xf0, 0x1c, 0x3c, 0x01, 0xf7, 0xbc,
	0x1b, 0x8f, 0x80, 0x7c, 0xf2, 0x83, 0xdb, 0x65, 0xdc, 0x70, 0x13, 0xf9, 0x7c, 0x3e, 0xe7, 0xf3,
	0xf1, 0xf9, 0x3e, 0x07, 0xde, 0xcc, 0x92, 0x39, 0xcf, 0x6e, 0xef, 0xb8, 0xaf, 0x92, 0x61, 0xb9,
	0x1a, 0x2e, 0xc2, 0xe1, 0xbd, 0x92, 0x6b, 0xfc, 0xf8, 0x8b, 0x34, 0xc9, 0x12, 0xea, 0x56, 0x49,
	0x1a, 0xf2, 0x8e, 0xdb, 0x6b, 0x16, 0x5c, 0xc5, 0x59, 0xf9, 0x2d, 0xab, 0xbc, 0x17, 0xb3, 0x24,
	0x99, 0x45, 0x72, 0x88, 0x51, 0xb8, 0xfa, 0x3a, 0xe4, 0x71, 0x51, 0x6e, 0xf5, 0x7f, 0x11, 0xe8,
	0x8e, 0x56, 0x2a, 0x12, 0x57, 0x89, 0x90, 0xf4, 0x29, 0xd8, 0x4a, 0x30, 0xab, 0x67, 0x0d, 0x48,
	0x60, 0x2b, 0x41, 0x19, 0xec, 0x86, 0x7a, 0x73, 0x2c, 0x98, 0x8d, 0x60, 0x1d, 0xd2, 0x57, 0x00,
	0x61, 0xaa, 0xc4, 0x4c, 0x5e, 0xf1, 0xb9, 0x64, 0xa4, 0x67, 0x0d, 0xba, 0x81, 0x81, 0xd0, 0xf7,
	0xe0, 0x96, 0xd1, 0x0d, 0x8f, 0x56, 0x92, 0x75, 0x7a, 0xd6, 0xc0, 0x3d, 0x39, 0xf4, 0xcb, 0x46,
	0xfc, 0xba, 0x11, 0xff, 0x22, 0x2e, 0x02, 0x33, 0x91, 0x9e, 0x83, 0x73, 0xaf, 0x17, 0x4b, 0xb6,
	0xd3, 0x23, 0x03, 0xf7, 0xa4, 0xef, 0x1b, 0x37, 0xf6, 0x9b, 0x4e, 0x7d, 0xcc, 0x5e, 0x7e, 0x8c,
	0xb3, 0xb4, 0x08, 0xaa, 0x0a, 0xea, 0xc1, 0xde, 0xed, 0x9d, 0x8a, 0x44, 0x2a, 0x63, 0xe6, 0xf4,
	0xc8, 0x80, 0x04, 0x4d, 0xac, 0x79, 0x79, 0x94, 0x8d, 0xc5, 0x92, 0xed, 0xfe, 0x93, 0xf7, 0x02,
	0x93, 0x2a, 0xde, 0xb2, 0xc2, 0xbb, 0x06, 0xd7, 0x38, 0x8e, 0x1e, 0x00, 0xf9, 0x26, 0x0b, 0x9c,
	0x52, 0x37, 0xd0, 0x4b, 0xfa, 0x16, 0x76, 0xb0, 0x05, 0x1c, 0xd2, 0x63, 0xd7, 0x2c, 0x53, 0xce,
	0xed, 0x0f, 0x96, 0x77, 0x06, 0xae, 0x71, 0x8e, 0x49, 0x48, 0x4a, 0xc2, 0x43, 0x93, 0x90, 0x18,
	0xa5, 0xfd, 0x1f, 0x36, 0xec, 0x7f, 0xe2, 0x45, 0xb2, 0xca, 0x26, 0x5a, 0xe0, 0x56, 0xd5, 0x3c,
	0xd8, 0x8b, 0x30, 0xa5, 0x91, 0xad, 0x89, 0xb5, 0xa2, 0xe8, 0x8c, 0xb1, 0x40, 0xd1, 0x48, 0x50,
	0x87, 0x94, 0x42, 0x67, 0xae, 0xe2, 0x1c, 0xa5, 0xb2, 0x02, 0x5c, 0x57, 0x58, 0xc1, 0x76, 0x1a,
	0xac, 0x40, 0x8c, 0xe7, 0x39, 0x73, 0x2a, 0x8c, 0xe7, 0x79, 0x85, 0x15, 0x6c, 0xb7, 0xc1, 0x0a,
	0x7a, 0x04, 0xce, 0xf7, 0x71, 0x2c, 0x64, 0xce, 0xf6, 0xf0, 0xa0, 0x2a, 0xd2, 0xce, 0x41, 0x55,
	0xae, 0x53, 0x21, 0x53, 0xd6, 0x45, 0x9d, 0x0c, 0x84, 0x9e, 0x02, 0x60, 0x4b, 0xd3, 0xac, 0x88,
	0x24, 0x03, 0x9c, 0xe8, 0xb3, 0x5a, 0xad, 0xd2, 0xd5, 0xb8, 0x15, 0x18, 0x69, 0xfd, 0xdf, 0x04,
	0x3a, 0x41, 0x92, 0x64, 0x74, 0x0a, 0x07, 0xd1, 0xe6, 0x78, 0x96, 0xcc, 0x46, 0xc5, 0x8f, 0x37,
	0x14, 0xd7, 0xc9, 0xfe, 0xd6, 0x20, 0x2b, 0xd9, 0x1f, 0x10, 0xd0, 0x0b, 0x80, 0xb0, 0x76, 0xc8,
	0x92, 0x11, 0xa4, 0x7b, 0xfd, 0x90, 0xae, 0x71, 0x51, 0x45, 0x64, 0x14, 0x69, 0x8a, 0xb9, 0x12,
	0x22, 0x92, 0x6b, 0x9e, 0xea, 0xe7, 0xf0, 0x08, 0xc5, 0x65, 0x93, 0x53, 0x51, 0xfc, 0x2d, 0xf2,
	0x38, 0x3c, 0x6f, 0x6d, 0xb8, 0xc5, 0x3f, 0x27, 0x9b, 0x86, 0x7c, 0xb9, 0x71, 0xd0, 0x16, 0x89,
	0x69, 0xcc, 0xcf, 0xb0, 0xbf, 0x75, 0x89, 0x16, 0xf2, 0x77, 0x9b, 0xe4, 0x47, 0xed, 0x2f, 0xc9,
	0xa4, 0x9d, 0xc2, 0xfe, 0xd6, 0xc5, 0xfe, 0xff, 0x11, 0x8d, 0xce, 0xc0, 0x53, 0x89, 0xdf, 0xfc,
	0x02, 0x6b, 0x83, 0x84, 0xd8, 0xc7, 0xc8, 0x99, 0x84, 0x37, 0x4a, 0xae, 0xbf, 0x74, 0x74, 0xf4,
	0xd3, 0x7e, 0x72, 0x89, 0xdb, 0x1a, 0x9a, 0x8c, 0x42, 0x07, 0x39, 0x4f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xbe, 0xdf, 0xed, 0xc5, 0x75, 0x05, 0x00, 0x00,
}
