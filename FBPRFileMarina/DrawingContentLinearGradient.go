// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DrawingContentLinearGradient struct {
	_tab flatbuffers.Table
}

func GetRootAsDrawingContentLinearGradient(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentLinearGradient {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DrawingContentLinearGradient{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDrawingContentLinearGradient(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentLinearGradient {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DrawingContentLinearGradient{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DrawingContentLinearGradient) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DrawingContentLinearGradient) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DrawingContentLinearGradient) Start(obj *Point) *Point {
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

func (rcv *DrawingContentLinearGradient) End(obj *Point) *Point {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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

func (rcv *DrawingContentLinearGradient) Locations(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *DrawingContentLinearGradient) LocationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DrawingContentLinearGradient) MutateLocations(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *DrawingContentLinearGradient) Colors(obj *RGBAColor, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DrawingContentLinearGradient) ColorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DrawingContentLinearGradientStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DrawingContentLinearGradientAddStart(builder *flatbuffers.Builder, start flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(start), 0)
}
func DrawingContentLinearGradientAddEnd(builder *flatbuffers.Builder, end flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(end), 0)
}
func DrawingContentLinearGradientAddLocations(builder *flatbuffers.Builder, locations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(locations), 0)
}
func DrawingContentLinearGradientStartLocationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DrawingContentLinearGradientAddColors(builder *flatbuffers.Builder, colors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(colors), 0)
}
func DrawingContentLinearGradientStartColorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 4)
}
func DrawingContentLinearGradientEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
