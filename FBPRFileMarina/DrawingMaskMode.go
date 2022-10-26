// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FBPRFileMarina

import "strconv"

type DrawingMaskMode int8

const (
	DrawingMaskModeFill     DrawingMaskMode = 0
	DrawingMaskModeStroke   DrawingMaskMode = 1
	DrawingMaskModeKnockout DrawingMaskMode = 2
)

var EnumNamesDrawingMaskMode = map[DrawingMaskMode]string{
	DrawingMaskModeFill:     "Fill",
	DrawingMaskModeStroke:   "Stroke",
	DrawingMaskModeKnockout: "Knockout",
}

var EnumValuesDrawingMaskMode = map[string]DrawingMaskMode{
	"Fill":     DrawingMaskModeFill,
	"Stroke":   DrawingMaskModeStroke,
	"Knockout": DrawingMaskModeKnockout,
}

func (v DrawingMaskMode) String() string {
	if s, ok := EnumNamesDrawingMaskMode[v]; ok {
		return s
	}
	return "DrawingMaskMode(" + strconv.FormatInt(int64(v), 10) + ")"
}
