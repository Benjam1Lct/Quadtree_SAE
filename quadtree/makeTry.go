package main

import "fmt"

type Quadtree struct {
	width, height int
	root          *node
}

type node struct {
	topLeftX, topLeftY int
	width, height      int
	content            int
	topLeftNode        *node
	topRightNode       *node
	bottomLeftNode     *node
	bottomRightNode    *node
}

func MakeFromArray(floorContent [][]int) (q Quadtree) {
	height := len(floorContent)
	width := len(floorContent[0])

	root := createNode(0, 0, width, height, floorContent)

	quadtree := &Quadtree{
		width:  width,
		height: height,
		root:   root,
	}
	return *quadtree
}

func createNode(topLeftX, topLeftY, width, height int, floorContent [][]int) *node {
	currentNode := &node{
		topLeftX: topLeftX,
		topLeftY: topLeftY,
		width:    width,
		height:   height,
	}
	if width == 1 && height == 1 {
		currentNode.content = floorContent[topLeftY][topLeftX]
	} else {
		halfWidth := width / 2
		halfHeight := height / 2
		currentNode.topLeftNode = createNode(topLeftX, topLeftY, halfWidth, halfHeight, floorContent)
		currentNode.topRightNode = createNode(topLeftX+halfWidth, topLeftY, halfWidth, halfHeight, floorContent)
		currentNode.bottomLeftNode = createNode(topLeftX, topLeftY+halfHeight, halfWidth, halfHeight, floorContent)
		currentNode.bottomRightNode = createNode(topLeftX+halfWidth, topLeftY+halfHeight, halfWidth, halfHeight, floorContent)
	}
	return currentNode
}

func main() {
	floorContent := [][]int{
		{0, 1, 1, 1},
		{0, 0, 0, 1},
		{1, 1, 1, 1},
		{1, 1, 0, 0}
	}
	fmt.Println(MakeFromArray(floorContent))
}
