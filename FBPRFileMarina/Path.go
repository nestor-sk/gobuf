// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Path struct {
	_tab flatbuffers.Table
}

func GetRootAsPath(buf []byte, offset flatbuffers.UOffsetT) *Path {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Path{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPath(buf []byte, offset flatbuffers.UOffsetT) *Path {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Path{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Path) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Path) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Path) Points(obj *Point, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 8
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Path) PointsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Path) Operations(j int) PathOperation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return PathOperation(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *Path) OperationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Path) MutateOperations(j int, n PathOperation) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

func PathStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PathAddPoints(builder *flatbuffers.Builder, points flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(points), 0)
}
func PathStartPointsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 4)
}
func PathAddOperations(builder *flatbuffers.Builder, operations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(operations), 0)
}
func PathStartOperationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PathEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
