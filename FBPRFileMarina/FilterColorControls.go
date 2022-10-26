// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FilterColorControls struct {
	_tab flatbuffers.Table
}

func GetRootAsFilterColorControls(buf []byte, offset flatbuffers.UOffsetT) *FilterColorControls {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FilterColorControls{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFilterColorControls(buf []byte, offset flatbuffers.UOffsetT) *FilterColorControls {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FilterColorControls{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FilterColorControls) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FilterColorControls) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FilterColorControls) Saturation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 1.0
}

func (rcv *FilterColorControls) MutateSaturation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *FilterColorControls) Brightness() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *FilterColorControls) MutateBrightness(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *FilterColorControls) Contrast() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *FilterColorControls) MutateContrast(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *FilterColorControls) Hue() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *FilterColorControls) MutateHue(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func FilterColorControlsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FilterColorControlsAddSaturation(builder *flatbuffers.Builder, saturation float32) {
	builder.PrependFloat32Slot(0, saturation, 1.0)
}
func FilterColorControlsAddBrightness(builder *flatbuffers.Builder, brightness float32) {
	builder.PrependFloat32Slot(1, brightness, 0.0)
}
func FilterColorControlsAddContrast(builder *flatbuffers.Builder, contrast float32) {
	builder.PrependFloat32Slot(2, contrast, 0.0)
}
func FilterColorControlsAddHue(builder *flatbuffers.Builder, hue float32) {
	builder.PrependFloat32Slot(3, hue, 0.0)
}
func FilterColorControlsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
