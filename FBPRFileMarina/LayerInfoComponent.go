// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LayerInfoComponent struct {
	_tab flatbuffers.Table
}

func GetRootAsLayerInfoComponent(buf []byte, offset flatbuffers.UOffsetT) *LayerInfoComponent {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LayerInfoComponent{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLayerInfoComponent(buf []byte, offset flatbuffers.UOffsetT) *LayerInfoComponent {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LayerInfoComponent{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LayerInfoComponent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LayerInfoComponent) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LayerInfoComponent) Bounds(obj *Rect) *Rect {
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

func (rcv *LayerInfoComponent) ObjectId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LayerInfoComponent) Traits() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LayerInfoComponent) MutateTraits(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *LayerInfoComponent) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LayerInfoComponentStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func LayerInfoComponentAddBounds(builder *flatbuffers.Builder, bounds flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(bounds), 0)
}
func LayerInfoComponentAddObjectId(builder *flatbuffers.Builder, objectId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(objectId), 0)
}
func LayerInfoComponentAddTraits(builder *flatbuffers.Builder, traits uint32) {
	builder.PrependUint32Slot(2, traits, 0)
}
func LayerInfoComponentAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}
func LayerInfoComponentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
