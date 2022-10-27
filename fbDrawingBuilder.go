package main

import (
	PR "workspace/FBPRFileMarina"

	flatbuffers "github.com/google/flatbuffers/go"
)

func buildDrawing(builder *flatbuffers.Builder, drawing *PR.DrawingComponent) (flatbuffers.UOffsetT, bool) {
	if drawing == nil {
		return 0, false
	}
	contentType := drawing.ContentType()
	var content flatbuffers.UOffsetT
	if contentType != PR.DrawingContentNONE {
		drawingContent := new(flatbuffers.Table)
		drawing.Content(drawingContent)
		content = buildDrawingContent(builder, contentType, drawingContent)
	}
	mask := buildDrawingMask(builder, drawing.Mask(nil))

	PR.DrawingComponentStart(builder)
	PR.DrawingComponentAddContentType(builder, drawing.ContentType())
	if contentType != PR.DrawingContentNONE {
		PR.DrawingComponentAddContent(builder, content)
	}
	PR.DrawingComponentAddMask(builder, mask)
	return PR.DrawingComponentEnd(builder), true
}

func buildDrawingMask(builder *flatbuffers.Builder, mask *PR.DrawingMask) flatbuffers.UOffsetT {

	prKnockoutRect := mask.KnockoutRect(nil)
	var knockoutRect flatbuffers.UOffsetT
	if prKnockoutRect != nil {
		knockoutRect = buildRect(builder, prKnockoutRect)
	}
	shapeType := mask.ShapeType()
	var shape flatbuffers.UOffsetT
	if shapeType != PR.ShapeNONE {
		shapeData := new(flatbuffers.Table)
		mask.Shape(shapeData)
		shape = buildShape(builder, shapeType, shapeData)
	}
	strokedPath := buildPath(builder, mask.StrokedPath(nil))
	stroke := buildStroke(builder, mask.Stroke(nil))

	PR.DrawingMaskStart(builder)
	PR.DrawingMaskAddMode(builder, mask.Mode())
	PR.DrawingMaskAddShapeType(builder, shapeType)
	if shapeType != PR.ShapeNONE {
		PR.DrawingMaskAddShape(builder, shape)
	}
	PR.DrawingMaskAddStroke(builder, stroke)
	PR.DrawingMaskAddStrokedPath(builder, strokedPath)
	if prKnockoutRect != nil {
		PR.DrawingMaskAddKnockoutRect(builder, knockoutRect)
	}
	return PR.DrawingMaskEnd(builder)
}

func buildDrawingContent(builder *flatbuffers.Builder, contentType PR.DrawingContent, contentData *flatbuffers.Table) flatbuffers.UOffsetT {
	switch contentType {
	case PR.DrawingContentDrawingContentColor:
		content := new(PR.DrawingContentColor)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentColor(builder, content)
	case PR.DrawingContentDrawingContentLinearGradient:
		content := new(PR.DrawingContentLinearGradient)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentLinearGradient(builder, content)
	case PR.DrawingContentDrawingContentAngularGradient:
		content := new(PR.DrawingContentAngularGradient)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentAngularGradient(builder, content)
	case PR.DrawingContentDrawingContentRadialGradient:
		content := new(PR.DrawingContentRadialGradient)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentRadialGradient(builder, content)
	case PR.DrawingContentDrawingContentImage:
		content := new(PR.DrawingContentImage)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentImage(builder, content)
	case PR.DrawingContentDrawingContentText:
		content := new(PR.DrawingContentText)
		content.Init(contentData.Bytes, contentData.Pos)
		return buildContentText(builder, content)
	}
	panic("Unhandled content type")
}

func buildContentColor(builder *flatbuffers.Builder, content *PR.DrawingContentColor) flatbuffers.UOffsetT {
	prColor := content.Color(nil)
	color := PR.CreateRGBAColor(builder, prColor.R(), prColor.G(), prColor.B(), prColor.A())
	PR.DrawingContentColorStart(builder)
	PR.DrawingContentColorAddColor(builder, color)
	return PR.DrawingContentColorEnd(builder)
}

func buildContentLinearGradient(builder *flatbuffers.Builder, content *PR.DrawingContentLinearGradient) flatbuffers.UOffsetT {
	start := PR.CreatePoint(builder, content.Start(nil).X(), content.Start(nil).X())
	end := PR.CreatePoint(builder, content.End(nil).X(), content.End(nil).X())

	numberOfLocs := content.LocationsLength()
	PR.DrawingContentLinearGradientStartLocationsVector(builder, numberOfLocs)
	for i := 0; i < numberOfLocs; i++ {
		loc := content.Locations(i)
		builder.PrependFloat32(loc)
	}
	prLocsVector := builder.EndVector(numberOfLocs)

	prColors := []flatbuffers.UOffsetT{}
	numberOfColors := content.ColorsLength()
	for i := 0; i < numberOfColors; i++ {
		color := new(PR.RGBAColor)
		content.Colors(color, i)
		prColor := PR.CreateRGBAColor(builder, color.R(), color.G(), color.B(), color.A())
		prColors = append(prColors, prColor)
	}

	PR.DrawingContentLinearGradientStartColorsVector(builder, numberOfColors)
	for _, prColor := range prColors {
		builder.PrependUOffsetT(prColor)
	}
	prColorsVector := builder.EndVector(numberOfColors)

	PR.DrawingContentLinearGradientStart(builder)
	PR.DrawingContentLinearGradientAddStart(builder, start)
	PR.DrawingContentLinearGradientAddEnd(builder, end)
	PR.DrawingContentLinearGradientAddLocations(builder, prLocsVector)
	PR.DrawingContentLinearGradientAddColors(builder, prColorsVector)
	return PR.DrawingContentLinearGradientEnd(builder)
}

func buildContentAngularGradient(builder *flatbuffers.Builder, content *PR.DrawingContentAngularGradient) flatbuffers.UOffsetT {
	center := PR.CreatePoint(builder, content.Center(nil).X(), content.Center(nil).X())

	numberOfLocs := content.LocationsLength()
	PR.DrawingContentAngularGradientStartLocationsVector(builder, numberOfLocs)
	for i := 0; i < numberOfLocs; i++ {
		loc := content.Locations(i)
		builder.PrependFloat32(loc)
	}
	prLocsVector := builder.EndVector(numberOfLocs)

	prColors := []flatbuffers.UOffsetT{}
	numberOfColors := content.ColorsLength()
	for i := 0; i < numberOfColors; i++ {
		color := new(PR.RGBAColor)
		content.Colors(color, i)
		prColor := PR.CreateRGBAColor(builder, color.R(), color.G(), color.B(), color.A())
		prColors = append(prColors, prColor)
	}

	PR.DrawingContentAngularGradientStartColorsVector(builder, numberOfColors)
	for _, prColor := range prColors {
		builder.PrependUOffsetT(prColor)
	}
	prColorsVector := builder.EndVector(numberOfColors)

	PR.DrawingContentAngularGradientStart(builder)
	PR.DrawingContentAngularGradientAddCenter(builder, center)
	PR.DrawingContentAngularGradientAddAngle(builder, content.Angle())
	PR.DrawingContentAngularGradientAddLocations(builder, prLocsVector)
	PR.DrawingContentAngularGradientAddColors(builder, prColorsVector)
	return PR.DrawingContentAngularGradientEnd(builder)
}

func buildContentRadialGradient(builder *flatbuffers.Builder, content *PR.DrawingContentRadialGradient) flatbuffers.UOffsetT {
	center := PR.CreatePoint(builder, content.Center(nil).X(), content.Center(nil).X())
	matrix := buildMatrix(builder, content.EllipseTransform(nil))

	numberOfLocs := content.LocationsLength()
	PR.DrawingContentRadialGradientStartLocationsVector(builder, numberOfLocs)
	for i := 0; i < numberOfLocs; i++ {
		loc := content.Locations(i)
		builder.PrependFloat32(loc)
	}
	prLocsVector := builder.EndVector(numberOfLocs)

	prColors := []flatbuffers.UOffsetT{}
	numberOfColors := content.ColorsLength()
	for i := 0; i < numberOfColors; i++ {
		color := new(PR.RGBAColor)
		content.Colors(color, i)
		prColor := PR.CreateRGBAColor(builder, color.R(), color.G(), color.B(), color.A())
		prColors = append(prColors, prColor)
	}

	PR.DrawingContentRadialGradientStartColorsVector(builder, numberOfColors)
	for _, prColor := range prColors {
		builder.PrependUOffsetT(prColor)
	}
	prColorsVector := builder.EndVector(numberOfColors)

	PR.DrawingContentRadialGradientStart(builder)
	PR.DrawingContentRadialGradientAddCenter(builder, center)
	PR.DrawingContentRadialGradientAddRadius(builder, content.Radius())
	PR.DrawingContentRadialGradientAddLocations(builder, prLocsVector)
	PR.DrawingContentRadialGradientAddColors(builder, prColorsVector)
	PR.DrawingContentRadialGradientAddEllipseTransform(builder, matrix)
	return PR.DrawingContentAngularGradientEnd(builder)
}

func buildContentImage(builder *flatbuffers.Builder, content *PR.DrawingContentImage) flatbuffers.UOffsetT {
	image := buildImageDrawing(builder, content.Image(nil))

	PR.DrawingContentImageStart(builder)
	PR.DrawingContentImageAddImage(builder, image)
	PR.DrawingContentImageAddTile(builder, content.Tile())
	return PR.DrawingContentImageEnd(builder)
}

func buildContentText(builder *flatbuffers.Builder, content *PR.DrawingContentText) flatbuffers.UOffsetT {
	PR.DrawingContentTextStart(builder)
	return PR.DrawingContentTextEnd(builder)
}

func buildImageDrawing(builder *flatbuffers.Builder, image *PR.ImageDrawing) flatbuffers.UOffsetT {
	rect := buildRect(builder, image.Rect(nil))

	PR.ImageDrawingStart(builder)
	PR.ImageDrawingAddImageIndex(builder, image.ImageIndex())
	PR.ImageDrawingAddRect(builder, rect)
	PR.ImageDrawingAddInterpolationQuality(builder, image.InterpolationQuality())
	return PR.ImageDrawingEnd(builder)
}
