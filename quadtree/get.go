package quadtree

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	for i := 0; i < configuration.Global.NumTileX; i++ {
		for j := 0; j < configuration.Global.NumTileY; j++ {
			absX := topLeftX + i
			absY := topLeftY + j
			if absX < q.width && absY < q.height && absX >= 0 && absY >= 0 {
				contentHolder[j][i] = ReturnFloor(absX, absY, q.root)
			} else {
				contentHolder[j][i] = -1
			}
		}
	}
}

func ReturnFloor(absX, absY int, n *node) int {
	//case feuille
	if n.topLeftNode == nil {
		return n.content
	}

	//determination coos carré suivant
	halfWidth := n.width / 2
	halfHeight := n.height / 2
	centerX := n.topLeftX + halfWidth
	centerY := n.topLeftY + halfHeight

	if absX < centerX {
		if absY < centerY {
			return ReturnFloor(absX, absY, n.topLeftNode)
		} else {
			return ReturnFloor(absX, absY, n.bottomLeftNode)
		}
	} else {
		if absY < centerY {
			return ReturnFloor(absX, absY, n.topRightNode)
		} else {
			return ReturnFloor(absX, absY, n.bottomRightNode)
		}
	}
}
