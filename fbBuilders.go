package main

import (
	"fmt"
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
	PR.ImageAddHeight(builder, image.Height())
	prImage := PR.ImageEnd(builder)
	return prImage
}

func buildNodeChildrenVector(builder *flatbuffers.Builder, node *PR.Node) flatbuffers.UOffsetT {
	prChildren := []flatbuffers.UOffsetT{}
	numberOfChildren := node.ChildrenLength()
	for i := 0; i < numberOfChildren; i++ {
		child := new(PR.Node)
		node.Children(child, i)
		prChild := buildNode(builder, child)
		prChildren = append(prChildren, prChild)
	}

	PR.NodeStartChildrenVector(builder, numberOfChildren)
	for _, prChild := range prChildren {
		builder.PrependUOffsetT(prChild)
	}
	return builder.EndVector(numberOfChildren)
}

func buildNode(builder *flatbuffers.Builder, node *PR.Node) flatbuffers.UOffsetT {
	children := buildNodeChildrenVector(builder, node)
	name := builder.CreateByteString(node.Name())
	enabled := node.Enabled()
	// drawing, drawingBuilt := buildDrawing(builder, node.Drawing(nil))
	layerInfo, layerInfoBuilt := buildLayerInfo(builder, node.LayerInfo(nil))
	transform, transformBuilt := buildTransform(builder, node.Transform(nil))
	mask, maskBuilt := buildMask(builder, node.Mask(nil))
	extent, extentBuilt := buildCachedExtent(builder, node.CachedExtent(nil))
	blend, blendBuilt := buildBlend(builder, node.Blend(nil))
	surface, surfaceBuilt := buildSurface(builder, node.Surface(nil))
	bgFilter, bgFilterBuilt := buildBackgroundFilter(builder, node.BackgroundFilter(nil))

	PR.NodeStart(builder)
	PR.NodeAddChildren(builder, children)
	PR.NodeAddEnabled(builder, enabled)
	PR.NodeAddName(builder, name)
	// if drawingBuilt {
	// 	PR.NodeAddDrawing(builder, drawing)
	// }
	if layerInfoBuilt {
		PR.NodeAddLayerInfo(builder, layerInfo)
	}
	if transformBuilt {
		PR.NodeAddTransform(builder, transform)
	}
	if maskBuilt {
		PR.NodeAddMask(builder, mask)
	}
	if blendBuilt {
		PR.NodeAddBlend(builder, blend)
	}
	if surfaceBuilt {
		PR.NodeAddSurface(builder, surface)
	}
	if bgFilterBuilt {
		PR.NodeAddBackgroundFilter(builder, bgFilter)
	}
	if extentBuilt {
		PR.NodeAddCachedExtent(builder, extent)
	}
	return PR.NodeEnd(builder)
}

func buildLayerInfo(builder *flatbuffers.Builder, layerInfo *PR.LayerInfoComponent) (flatbuffers.UOffsetT, bool) {
	if layerInfo == nil {
		return 0, false
	}
	objectId := builder.CreateByteString(layerInfo.ObjectId())
	name := builder.CreateByteString(layerInfo.Name())
	traits := layerInfo.Traits()
	rect := buildRect(builder, layerInfo.Bounds(nil))

	PR.LayerInfoComponentStart(builder)
	PR.LayerInfoComponentAddBounds(builder, rect)
	PR.LayerInfoComponentAddObjectId(builder, objectId)
	PR.LayerInfoComponentAddTraits(builder, traits)
	PR.LayerInfoComponentAddName(builder, name)
	return PR.LayerInfoComponentEnd(builder), true
}

func buildTransform(builder *flatbuffers.Builder, transform *PR.TransformComponent) (flatbuffers.UOffsetT, bool) {
	if transform == nil {
		return 0, false
	}
	matrix := buildMatrix(builder, transform.Matrix(nil))
	PR.TransformComponentStart(builder)
	PR.TransformComponentAddMatrix(builder, matrix)
	return PR.TransformComponentEnd(builder), true
}

func buildCachedExtent(builder *flatbuffers.Builder, extent *PR.CachedExtentComponent) (flatbuffers.UOffsetT, bool) {
	if extent == nil {
		return 0, false
	}
	rect := buildRect(builder, extent.Rect(nil))
	PR.CachedExtentComponentStart(builder)
	PR.CachedExtentComponentAddRect(builder, rect)
	return PR.CachedExtentComponentEnd(builder), true
}

func buildMask(builder *flatbuffers.Builder, mask *PR.MaskComponent) (flatbuffers.UOffsetT, bool) {
	if mask == nil {
		return 0, false
	}
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
	return PR.MaskComponentEnd(builder), true
}

func buildBlend(builder *flatbuffers.Builder, blend *PR.BlendComponent) (flatbuffers.UOffsetT, bool) {
	if blend == nil {
		return 0, false
	}
	PR.BlendComponentStart(builder)
	PR.BlendComponentAddMode(builder, blend.Mode())
	PR.BlendComponentAddOpacity(builder, blend.Opacity())
	return PR.BlendComponentEnd(builder), true
}

func buildSurface(builder *flatbuffers.Builder, surface *PR.SurfaceComponent) (flatbuffers.UOffsetT, bool) {
	if surface == nil {
		return 0, false
	}
	filerData := new(flatbuffers.Table)
	surface.Filter(filerData)
	filter := buildFilter(builder, surface.FilterType(), filerData)

	PR.SurfaceComponentStart(builder)
	PR.SurfaceComponentAddShouldBlendSourceOver(builder, surface.ShouldBlendSourceOver())
	PR.SurfaceComponentAddFilterType(builder, surface.FilterType())
	PR.SurfaceComponentAddFilter(builder, filter)
	return PR.SurfaceComponentEnd(builder), true
}

func buildBackgroundFilter(builder *flatbuffers.Builder, bgfilter *PR.BackgroundFilterComponent) (flatbuffers.UOffsetT, bool) {
	if bgfilter == nil {
		return 0, false
	}
	shapeData := new(flatbuffers.Table)
	bgfilter.Shape(shapeData)
	shape := buildShape(builder, bgfilter.ShapeType(), shapeData)
	image := buildImageDrawing(builder, bgfilter.Image(nil))

	PR.BackgroundFilterComponentStart(builder)
	PR.BackgroundFilterComponentAddShapeType(builder, bgfilter.ShapeType())
	PR.BackgroundFilterComponentAddShape(builder, shape)
	PR.BackgroundFilterComponentAddImage(builder, image)
	PR.BackgroundFilterComponentAddRadius(builder, bgfilter.Radius())
	PR.BackgroundFilterComponentAddSaturation(builder, bgfilter.Saturation())
	return PR.BackgroundFilterComponentEnd(builder), true
}

func buildPath(builder *flatbuffers.Builder, path *PR.Path) flatbuffers.UOffsetT {

	prPoints := []flatbuffers.UOffsetT{}
	numberOfPoints := path.PointsLength()
	for i := 0; i < numberOfPoints; i++ {
		point := new(PR.Point)
		path.Points(point, i)
		prPoint := PR.CreatePoint(builder, point.X(), point.Y())
		fmt.Printf("%v,%v\n", point.X(), point.Y())
		prPoints = append(prPoints, prPoint)
	}

	PR.PathStartPointsVector(builder, numberOfPoints)
	for i := numberOfPoints - 1; i >= 0; i-- {
		builder.PrependUOffsetT(prPoints[i])
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

func buildStroke(builder *flatbuffers.Builder, stroke *PR.Stroke) flatbuffers.UOffsetT {

	numberOfPatterns := stroke.LineDashPatternLength()
	PR.StrokeStartLineDashPatternVector(builder, numberOfPatterns)
	for i := 0; i < numberOfPatterns; i++ {
		pattern := stroke.LineDashPattern(i)
		builder.PrependFloat32(pattern)
	}
	prPatterns := builder.EndVector(numberOfPatterns)

	PR.StrokeStart(builder)
	PR.StrokeAddLineWidth(builder, stroke.LineWidth())
	PR.StrokeAddLineCap(builder, stroke.LineCap())
	PR.StrokeAddLineJoin(builder, stroke.LineJoin())
	PR.StrokeAddLineDashPhase(builder, stroke.LineDashPhase())
	PR.StrokeAddLineDashPattern(builder, prPatterns)
	PR.StrokeAddMiterLimit(builder, stroke.MiterLimit())
	return PR.StrokeEnd(builder)
}

func buildFilter(builder *flatbuffers.Builder, filterType PR.Filter, filterData *flatbuffers.Table) flatbuffers.UOffsetT {
	switch filterType {
	case PR.FilterFilterColorControls:
		filter := new(PR.FilterColorControls)
		filter.Init(filterData.Bytes, filterData.Pos)
		return buildFilterColorControls(builder, filter)
	case PR.FilterFilterGaussianBlur:
		filter := new(PR.FilterGaussianBlur)
		filter.Init(filterData.Bytes, filterData.Pos)
		return buildFilterGaussianBlur(builder, filter)
	case PR.FilterFilterMotionBlur:
		filter := new(PR.FilterMotionBlur)
		filter.Init(filterData.Bytes, filterData.Pos)
		return buildFilterMotionBlur(builder, filter)
	case PR.FilterFilterZoomBlur:
		filter := new(PR.FilterZoomBlur)
		filter.Init(filterData.Bytes, filterData.Pos)
		return buildFilterZoomBlur(builder, filter)
	case PR.FilterFilterShadow:
		filter := new(PR.FilterShadow)
		filter.Init(filterData.Bytes, filterData.Pos)
		return buildFilterShadow(builder, filter)
	}
	panic("Unhandled filter type")
}

func buildFilterColorControls(builder *flatbuffers.Builder, filter *PR.FilterColorControls) flatbuffers.UOffsetT {
	PR.FilterColorControlsStart(builder)
	PR.FilterColorControlsAddSaturation(builder, filter.Saturation())
	PR.FilterColorControlsAddBrightness(builder, filter.Brightness())
	PR.FilterColorControlsAddContrast(builder, filter.Contrast())
	PR.FilterColorControlsAddHue(builder, filter.Hue())
	return PR.FilterColorControlsEnd(builder)
}

func buildFilterGaussianBlur(builder *flatbuffers.Builder, filter *PR.FilterGaussianBlur) flatbuffers.UOffsetT {
	PR.FilterGaussianBlurStart(builder)
	PR.FilterGaussianBlurAddRadius(builder, filter.Radius())
	return PR.FilterGaussianBlurEnd(builder)
}

func buildFilterMotionBlur(builder *flatbuffers.Builder, filter *PR.FilterMotionBlur) flatbuffers.UOffsetT {
	PR.FilterMotionBlurStart(builder)
	PR.FilterMotionBlurAddRadius(builder, filter.Radius())
	PR.FilterMotionBlurAddAngle(builder, filter.Angle())
	return PR.FilterMotionBlurEnd(builder)
}

func buildFilterZoomBlur(builder *flatbuffers.Builder, filter *PR.FilterZoomBlur) flatbuffers.UOffsetT {
	center := PR.CreatePoint(builder, filter.Center(nil).X(), filter.Center(nil).Y())

	PR.FilterZoomBlurStart(builder)
	PR.FilterZoomBlurAddRadius(builder, filter.Radius())
	PR.FilterZoomBlurAddCenter(builder, center)
	return PR.FilterZoomBlurEnd(builder)
}

func buildFilterShadow(builder *flatbuffers.Builder, filter *PR.FilterShadow) flatbuffers.UOffsetT {
	shadow := buildShadow(builder, filter.Shadow(nil))
	PR.FilterShadowStart(builder)
	PR.FilterShadowAddShadow(builder, shadow)
	return PR.FilterShadowEnd(builder)
}

func buildShadow(builder *flatbuffers.Builder, shadow *PR.Shadow) flatbuffers.UOffsetT {
	PR.ShadowStart(builder)
	PR.ShadowAddRadius(builder, shadow.Radius())
	offset := PR.CreateSize(builder, shadow.Offset(nil).Width(), shadow.Offset(nil).Height())
	PR.ShadowAddOffset(builder, offset)
	PR.ShadowAddSpread(builder, shadow.Spread())
	color := PR.CreateRGBAColor(builder, shadow.Color(nil).R(), shadow.Color(nil).G(), shadow.Color(nil).B(), shadow.Color(nil).A())
	PR.ShadowAddColor(builder, color)
	return PR.ShadowEnd(builder)
}
