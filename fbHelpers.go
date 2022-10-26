package main

import (
	PR "workspace/FBPRFileMarina"

	flatbuffers "github.com/google/flatbuffers/go"
)

func bufferBytes(images []PR.Image, artboard *PR.Node, version []byte) []byte {
	fileId := []byte("PRMR")
	builder := flatbuffers.NewBuilder(512)
	prPresentation := buildPresentation(builder, images, artboard, version)
	builder.FinishWithFileIdentifier(prPresentation, fileId)
	return builder.FinishedBytes()
}

func buildPresentation(builder *flatbuffers.Builder, images []PR.Image, artboard *PR.Node, version []byte) flatbuffers.UOffsetT {

	numberOfImages := len(images)
	prImages := []flatbuffers.UOffsetT{}
	for _, image := range images {
		prImage := buildImage(builder, &image)
		prImages = append(prImages, prImage)
	}
	PR.PresentationStartImagesVector(builder, numberOfImages)
	for _, prImage := range prImages {
		builder.PrependUOffsetT(prImage)
	}
	prImagesVector := builder.EndVector(numberOfImages)
	prNode := buildNode(builder, artboard)
	prVersion := builder.CreateByteString(version)

	PR.PresentationStart(builder)
	PR.PresentationAddImages(builder, prImagesVector)
	PR.PresentationAddRoot(builder, prNode)
	PR.PresentationAddSchemaVersion(builder, prVersion)
	prPesentation := PR.PresentationEnd(builder)

	return prPesentation
}

func buildImage(builder *flatbuffers.Builder, image *PR.Image) flatbuffers.UOffsetT {
	uuid := builder.CreateByteString(image.Uuid())
	format := builder.CreateByteString(image.Format())
	PR.ImageStart(builder)
	PR.ImageAddUuid(builder, uuid)
	PR.ImageAddFormat(builder, format)
	PR.ImageAddWidth(builder, image.Width())
	PR.ImageAddWidth(builder, image.Height())
	prImage := PR.ImageEnd(builder)
	return prImage
}

func buildNode(builder *flatbuffers.Builder, node *PR.Node) flatbuffers.UOffsetT {
	name := builder.CreateByteString(node.Name())
	enabled := node.Enabled()
	layerInfo := buildLayerInfo(builder, node.LayerInfo(nil))

	PR.NodeStartChildrenVector(builder, 0)
	children := builder.EndVector(0)

	transform := buildTransform(builder, node.Transform(nil))
	mask := buildMask(builder, node.Mask(nil))
	extent := buildCachedExtent(builder, node.CachedExtent(nil))

	PR.NodeStart(builder)
	PR.NodeAddChildren(builder, children)
	PR.NodeAddEnabled(builder, enabled)
	PR.NodeAddName(builder, name)
	//Drawing
	PR.NodeAddLayerInfo(builder, layerInfo)
	PR.NodeAddTransform(builder, transform)
	PR.NodeAddMask(builder, mask)
	//blend
	//surface
	//bgfilter
	PR.NodeAddCachedExtent(builder, extent)
	return PR.NodeEnd(builder)
}

func buildLayerInfo(builder *flatbuffers.Builder, layerInfo *PR.LayerInfoComponent) flatbuffers.UOffsetT {
	objectId := builder.CreateByteString(layerInfo.ObjectId())
	name := builder.CreateByteString(layerInfo.Name())
	traits := layerInfo.Traits()
	rect := buildRect(builder, layerInfo.Bounds(nil))

	PR.LayerInfoComponentStart(builder)
	PR.LayerInfoComponentAddBounds(builder, rect)
	PR.LayerInfoComponentAddObjectId(builder, objectId)
	PR.LayerInfoComponentAddTraits(builder, traits)
	PR.LayerInfoComponentAddName(builder, name)
	return PR.LayerInfoComponentEnd(builder)
}

func buildTransform(builder *flatbuffers.Builder, transform *PR.TransformComponent) flatbuffers.UOffsetT {
	matrix := buildMatrix(builder, transform.Matrix(nil))
	PR.TransformComponentStart(builder)
	PR.TransformComponentAddMatrix(builder, matrix)
	return PR.TransformComponentEnd(builder)
}

func buildCachedExtent(builder *flatbuffers.Builder, extent *PR.CachedExtentComponent) flatbuffers.UOffsetT {
	rect := buildRect(builder, extent.Rect(nil))
	PR.CachedExtentComponentStart(builder)
	PR.CachedExtentComponentAddRect(builder, rect)
	return PR.CachedExtentComponentEnd(builder)
}

func buildMask(builder *flatbuffers.Builder, mask *PR.MaskComponent) flatbuffers.UOffsetT {
	prKnockoutRect := mask.KnockoutRect(nil)
	var knockoutRect flatbuffers.UOffsetT
	if prKnockoutRect != nil {
		knockoutRect = buildRect(builder, prKnockoutRect)
	}
	clipShapeType := mask.ClipShapeType()
	var shape flatbuffers.UOffsetT
	if clipShapeType != PR.ShapeNONE {
		clipShape := new(flatbuffers.Table)
		mask.ClipShape(clipShape)
		shape = buildShape(builder, mask.ClipShapeType(), clipShape)
	}

	PR.MaskComponentStart(builder)
	PR.MaskComponentAddMode(builder, mask.Mode())
	PR.MaskComponentAddClipShapeType(builder, clipShapeType)
	if clipShapeType != PR.ShapeNONE {
		PR.MaskComponentAddClipShape(builder, shape)
	}
	if prKnockoutRect != nil {
		PR.MaskComponentAddKnockoutRect(builder, knockoutRect)
	}
	return PR.MaskComponentEnd(builder)
}

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

func buildPath(builder *flatbuffers.Builder, path *PR.Path) flatbuffers.UOffsetT {
	prPoints := []flatbuffers.UOffsetT{}
	numberOfPoints := path.PointsLength()
	for i := 0; i < numberOfPoints; i++ {
		point := new(PR.Point)
		path.Points(point, i)
		prPoint := PR.CreatePoint(builder, point.X(), point.Y())
		prPoints = append(prPoints, prPoint)
	}

	PR.PathStartPointsVector(builder, numberOfPoints)
	for _, prPoint := range prPoints {
		builder.PrependUOffsetT(prPoint)
	}
	prPointsVector := builder.EndVector(numberOfPoints)

	numberOfOps := path.OperationsLength()
	PR.PathStartOperationsVector(builder, numberOfOps)
	for i := 0; i < numberOfOps; i++ {
		op := path.Operations(i)
		builder.PrependInt8(int8(op)) //??
	}
	prOpsVector := builder.EndVector(numberOfOps)

	PR.PathStart(builder)
	PR.PathAddPoints(builder, prPointsVector)
	PR.PathAddPoints(builder, prOpsVector)
	return PR.PathEnd(builder)
}

//cont
//
//

func buildDrawing(builder *flatbuffers.Builder, drawing *PR.DrawingComponent) flatbuffers.UOffsetT {

	// drawing.Content()
	mask := buildDrawingMask(builder, drawing.Mask(nil))

	PR.DrawingComponentStart(builder)
	PR.DrawingComponentAddContentType(builder, drawing.ContentType())
	// PR.DrawingComponentAddContent(builder, ...)
	PR.DrawingComponentAddMask(builder, mask)
	return PR.DrawingComponentEnd(builder)
}

func buildDrawingMask(builder *flatbuffers.Builder, mask *PR.DrawingMask) flatbuffers.UOffsetT {

	PR.DrawingMaskStart(builder)

	return PR.DrawingMaskEnd(builder)
}

// table DrawingComponent {
// 	content: DrawingContent (required);
// 	mask: DrawingMask (required); //build
//   }

//   table DrawingMask {
// 	mode: DrawingMaskMode; //enum
// 	shape: Shape; //enum
// 	stroke: Stroke; //build
// 	stroked_path: Path; //build
// 	knockout_rect: Rect; //create
//   }

func buildRect(builder *flatbuffers.Builder, rect *PR.Rect) flatbuffers.UOffsetT {
	rectOrigin := rect.Origin(nil)
	rectSize := rect.Size(nil)
	return PR.CreateRect(builder,
		rectOrigin.X(),
		rectOrigin.Y(),
		rectSize.Width(),
		rectSize.Height())
}

func buildMatrix(builder *flatbuffers.Builder, matrix *PR.Matrix) flatbuffers.UOffsetT {
	return PR.CreateMatrix(builder,
		matrix.A(),
		matrix.B(),
		matrix.C(),
		matrix.D(),
		matrix.Tx(),
		matrix.Ty())
}
