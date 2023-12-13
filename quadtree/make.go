package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) Quadtree {
	return Quadtree{
		width:  len(floorContent[0]),
		height: len(floorContent),
		root:   CreateQuadTree(floorContent, 0, 0, len(floorContent[0]), len(floorContent)),
	}
}

func CreateQuadTree(data [][]int, x, y, width, height int) *node {
	noeud := &node{topLeftX: x, topLeftY: y, width: width, height: height}

	// Si toute la zone est du même type de terrain, n'a pas d'enfants
	if SameFloor(data, x, y, width, height) {
		noeud.content = data[y][x]
		return noeud
	}

	halfWidth := width / 2
	halfHeight := height / 2

	// Ajouter des enfants seulement si la zone n'est pas homogène
	noeud.topLeftNode = CreateQuadTree(data, x, y, halfWidth, halfHeight)
	noeud.topRightNode = CreateQuadTree(data, x+halfWidth, y, width-halfWidth, halfHeight)
	noeud.bottomLeftNode = CreateQuadTree(data, x, y+halfHeight, halfWidth, height-halfHeight)
	noeud.bottomRightNode = CreateQuadTree(data, x+halfWidth, y+halfHeight, width-halfWidth, height-halfHeight)
	return noeud
}

func SameFloor(data [][]int, x, y, width, height int) bool {
	kindFloor := data[y][x]
	for i := y; i < y+height; i++ {
		for j := x; j < x+width; j++ {
			if data[i][j] != kindFloor {
				return false
			}
		}
	}
	return true
}
