// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ShapeRect struct {
	_tab flatbuffers.Table
}

func GetRootAsShapeRect(buf []byte, offset flatbuffers.UOffsetT) *ShapeRect {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ShapeRect{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsShapeRect(buf []byte, offset flatbuffers.UOffsetT) *ShapeRect {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ShapeRect{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ShapeRect) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ShapeRect) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ShapeRect) Rect(obj *Rect) *Rect {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
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

func ShapeRectStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ShapeRectAddRect(builder *flatbuffers.Builder, rect flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(rect), 0)
}
func ShapeRectEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}