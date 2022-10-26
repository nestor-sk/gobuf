// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DrawingContentRadialGradient struct {
	_tab flatbuffers.Table
}

func GetRootAsDrawingContentRadialGradient(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentRadialGradient {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DrawingContentRadialGradient{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDrawingContentRadialGradient(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentRadialGradient {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DrawingContentRadialGradient{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DrawingContentRadialGradient) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DrawingContentRadialGradient) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DrawingContentRadialGradient) Center(obj *Point) *Point {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Point)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DrawingContentRadialGradient) Radius() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DrawingContentRadialGradient) MutateRadius(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *DrawingContentRadialGradient) Locations(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *DrawingContentRadialGradient) LocationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DrawingContentRadialGradient) MutateLocations(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *DrawingContentRadialGradient) Colors(obj *RGBAColor, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DrawingContentRadialGradient) ColorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DrawingContentRadialGradient) EllipseTransform(obj *Matrix) *Matrix {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Matrix)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DrawingContentRadialGradientStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DrawingContentRadialGradientAddCenter(builder *flatbuffers.Builder, center flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(center), 0)
}
func DrawingContentRadialGradientAddRadius(builder *flatbuffers.Builder, radius float32) {
	builder.PrependFloat32Slot(1, radius, 0.0)
}
func DrawingContentRadialGradientAddLocations(builder *flatbuffers.Builder, locations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(locations), 0)
}
func DrawingContentRadialGradientStartLocationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DrawingContentRadialGradientAddColors(builder *flatbuffers.Builder, colors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(colors), 0)
}
func DrawingContentRadialGradientStartColorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 4)
}
func DrawingContentRadialGradientAddEllipseTransform(builder *flatbuffers.Builder, ellipseTransform flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(ellipseTransform), 0)
}
func DrawingContentRadialGradientEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}