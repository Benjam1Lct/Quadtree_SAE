package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) Quadtree {
	return Quadtree{
		width:  len(floorContent[0]),
		height: len(floorContent),
		root:   buildQuadTree(floorContent, 0, 0, len(floorContent[0]), len(floorContent)),
	}
}

func buildQuadTree(data [][]int, x, y, width, height int) *node {
	n := &node{topLeftX: x, topLeftY: y, width: width, height: height}

	// Si toute la zone est du même type de terrain, n'a pas d'enfants
	if isHomogeneous(data, x, y, width, height) {
		n.content = data[y][x]
		return n
	}

	halfWidth := width / 2
	halfHeight := height / 2

	// Ajouter des enfants seulement si la zone n'est pas homogène
	n.topLeftNode = buildQuadTree(data, x, y, halfWidth, halfHeight)
	n.topRightNode = buildQuadTree(data, x+halfWidth, y, width-halfWidth, halfHeight)
	n.bottomLeftNode = buildQuadTree(data, x, y+halfHeight, halfWidth, height-halfHeight)
	n.bottomRightNode = buildQuadTree(data, x+halfWidth, y+halfHeight, width-halfWidth, height-halfHeight)

	return n
}

func isHomogeneous(data [][]int, x, y, width, height int) bool {
	terrainType := data[y][x]
	for i := y; i < y+height; i++ {
		for j := x; j < x+width; j++ {
			if data[i][j] != terrainType {
				return false
			}
		}
	}
	return true
}
