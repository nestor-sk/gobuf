// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DrawingContentColor struct {
	_tab flatbuffers.Table
}

func GetRootAsDrawingContentColor(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentColor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DrawingContentColor{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDrawingContentColor(buf []byte, offset flatbuffers.UOffsetT) *DrawingContentColor {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DrawingContentColor{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DrawingContentColor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DrawingContentColor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DrawingContentColor) Color(obj *RGBAColor) *RGBAColor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(RGBAColor)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DrawingContentColorStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DrawingContentColorAddColor(builder *flatbuffers.Builder, color flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(color), 0)
}
func DrawingContentColorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}