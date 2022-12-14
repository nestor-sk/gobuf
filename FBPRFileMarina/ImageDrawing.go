// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ImageDrawing struct {
	_tab flatbuffers.Table
}

func GetRootAsImageDrawing(buf []byte, offset flatbuffers.UOffsetT) *ImageDrawing {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ImageDrawing{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsImageDrawing(buf []byte, offset flatbuffers.UOffsetT) *ImageDrawing {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ImageDrawing{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ImageDrawing) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ImageDrawing) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ImageDrawing) ImageIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ImageDrawing) MutateImageIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ImageDrawing) Rect(obj *Rect) *Rect {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Rect)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ImageDrawing) InterpolationQuality() InterpolationQuality {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return InterpolationQuality(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 4
}

func (rcv *ImageDrawing) MutateInterpolationQuality(n InterpolationQuality) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func ImageDrawingStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ImageDrawingAddImageIndex(builder *flatbuffers.Builder, imageIndex uint32) {
	builder.PrependUint32Slot(0, imageIndex, 0)
}
func ImageDrawingAddRect(builder *flatbuffers.Builder, rect flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(rect), 0)
}
func ImageDrawingAddInterpolationQuality(builder *flatbuffers.Builder, interpolationQuality InterpolationQuality) {
	builder.PrependInt8Slot(2, int8(interpolationQuality), 4)
}
func ImageDrawingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
