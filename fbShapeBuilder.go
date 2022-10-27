package main

import (
	PR "workspace/FBPRFileMarina"

	flatbuffers "github.com/google/flatbuffers/go"
)

func buildShape(builder *flatbuffers.Builder, shapeType PR.Shape, shapeData *flatbuffers.Table) flatbuffers.UOffsetT {
	switch shapeType {
	case PR.ShapeShapeRect:
		shape := new(PR.ShapeRect)
		shape.Init(shapeData.Bytes, shapeData.Pos)
		return buildShapeRect(builder, shape)
	case PR.ShapeShapeEllipse:
		shape := new(PR.ShapeEllipse)
		shape.Init(shapeData.Bytes, shapeData.Pos)
		return buildShapeEllipse(builder, shape)
	case PR.ShapeShapePath:
		shape := new(PR.ShapePath)
		shape.Init(shapeData.Bytes, shapeData.Pos)
		return buildShapePath(builder, shape)
	}
	panic("Unhandled shape type")
}

func buildShapePath(builder *flatbuffers.Builder, shape *PR.ShapePath) flatbuffers.UOffsetT {
	path := buildPath(builder, shape.Path(nil))
	PR.ShapePathStart(builder)
	PR.ShapePathAddPath(builder, path)
	PR.ShapePathAddFillRule(builder, shape.FillRule())
	return PR.ShapePathEnd(builder)
}

func buildShapeRect(builder *flatbuffers.Builder, shape *PR.ShapeRect) flatbuffers.UOffsetT {
	rect := buildRect(builder, shape.Rect(nil))
	PR.ShapeRectStart(builder)
	PR.ShapeRectAddRect(builder, rect)
	return PR.ShapeRectEnd(builder)
}

func buildShapeEllipse(builder *flatbuffers.Builder, shape *PR.ShapeEllipse) flatbuffers.UOffsetT {
	rect := buildRect(builder, shape.Rect(nil))
	PR.ShapeRectStart(builder)
	PR.ShapeRectAddRect(builder, rect)
	return PR.ShapeRectEnd(builder)
}
