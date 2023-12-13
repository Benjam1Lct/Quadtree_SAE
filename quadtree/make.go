package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
    height := len(floorContent)
    width := len(floorContent[0])

    root := createNode(0, 0, width, height, floorContent)

    quadtree := &Quadtree{
        width: width,
        height: height,
        root: root,
    }
    return quadtree
}


func createNode(topLeftX, topLeftY, width, height int, floorContent [][]int) {
    currentNode := &node{
        topLeftX: topLeftX,
        topLeftY: topLeftY,
        width:    width,
        height:   height,
    }
    if width == 1 && height == 1 {
        currentNode.content = floorContent[topLeftY][topLeftX]
    } else {
        halfwidth := width/2
        halfheight := height/2
        currentNode.topLeftNode = createNode(topLeftX, topLeftY, halfWidth, halfHeight, floorContent)
        currentNode.topRightNode = createNode(topLeftX+halfWidth, topLeftY, halfWidth, halfHeight, floorContent)
        currentNode.bottomLeftNode = createNode(topLeftX, topLeftY+halfHeight, halfWidth, halfHeight, floorContent)
        currentNode.bottomRightNode = createNode(topLeftX+halfWidth, topLeftY+halfHeight, halfWidth, halfHeight, floorContent)
    }
}