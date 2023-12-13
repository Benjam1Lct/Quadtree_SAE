import (
	"fmt"
	"testing" 
)

type Quadtree struct {
	width, height int
	root          *node
}

// node représente un nœud d'arbre quaternaire. Les champs sont :
//   - topLeftX, topLeftY : les coordonnées (en cases) de la case
//     située en haut à gauche de la zone du terrain représentée
//     par ce nœud.
//   - width, height :  la taille en cases de la zone représentée
//     par ce nœud.
//   - content : le type de terrain de la zone représentée par ce
//     nœud (seulement s'il s'agit d'une feuille).
//   - xxxNode : Une représentation de la partie xxx de la zone
//     représentée par ce nœud, différent de nil si et seulement
//     si le nœud actuel n'est pas une feuille.
type node struct {
	topLeftX, topLeftY int
	width, height      int
	content            int
	topLeftNode        *node
	topRightNode       *node
	bottomLeftNode     *node
	bottomRightNode    *node
}
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

floorContent := [][]int{
	{0, 1, 1, 1},
	{0, 0, 0, 1},
	{1, 1, 1, 1},
	{1, 1, 0, 0},
}

func main() {}
fmt.Println(MakeFromArray(floorContent))