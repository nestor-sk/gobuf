package main

import (
	"fmt"
	"os"
	PR "workspace/FBPRFileMarina"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1] //PR file path
	buf, err := os.ReadFile(filePath)
	check(err)
	// fmt.Print(string(buf))
	presentation := PR.GetRootAsPresentation(buf, 0)
	images := getImages(presentation)
	version := presentation.SchemaVersion()
	node := presentation.Root(nil) //this is the page
	artboards := findArtboards(node)
	for i, artboard := range artboards {
		createFile(images, artboard, version, i)
	}
}

func createFile(images []PR.Image, artboard *PR.Node, version []byte, index int) {
	bytes := bufferBytes(images, artboard, version)
	fileName := fmt.Sprintf("%v.sketchpresentation", index)
	err := os.WriteFile(fileName, bytes, 0644)
	check(err)
}

func getImages(root *PR.Presentation) []PR.Image {
	numberOfImages := root.ImagesLength()
	images := []PR.Image{}
	for i := 0; i < numberOfImages; i++ {
		image := new(PR.Image)
		root.Images(image, i)
		images = append(images, *image)
	}
	return images
}

func treeWalk(node *PR.Node) {
	fmt.Println(string(node.Name()))
	for i := 0; i < node.ChildrenLength(); i++ {
		child := new(PR.Node)
		node.Children(child, i)
		treeWalk(child)
	}
}

func findArtboards(node *PR.Node) []*PR.Node {
	artboards := []*PR.Node{}

	for i := 0; i < node.ChildrenLength(); i++ {
		child := new(PR.Node)
		node.Children(child, i)
		layerInfo := child.LayerInfo(nil)
		traits := layerInfo.Traits()
		if traits&(1<<13) != 0 {
			artboards = append(artboards, child)
		}
	}

	return artboards
}
