package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	//	Entrée :
	//	floorContent : Un tableau qui contient des entiers qui sont les content du floor
	//	Sortie :
	//	q : un Quadtree qui stock les données du floor
	//	Fonction qui récupère les données d'un tableau pour en faire un quadtree

	if len(floorContent) <= 0 {
		return q
	}

	height := len(floorContent)
	width := len(floorContent[0])

	root := createNode(0, 0, width, height, floorContent) //On crée le noeud principale à l'aide de createNode

	quadtree := Quadtree{
		width:  width,
		height: height,
		root:   root,
	}
	return quadtree
}

func createNode(topLeftX, topLeftY, width, height int, floorContent [][]int) *node {
	//	Entrées :
	//	topLeftX : coordonnée X du coin en haut à gauche du noeud
	//	topLeftY : coordonnée Y du coin en haut à gauche du neoud
	//	width : la largeur du noeud
	//	height : la hauteur du noeud
	//	floorContent : un tableau qui contient des entiers qui sont les content du floor
	//	Sorties :
	//	un noeud
	//	Fonction récursive qui crée des noeud enfants dans un quadtree si besoin
	currentNode := &node{
		topLeftX: topLeftX,
		topLeftY: topLeftY,
		width:    width,
		height:   height,
	}
	//On test si une zone a le meme content
	sameContent := true
	content := floorContent[topLeftY][topLeftX]
	for i := topLeftY; i < topLeftY+height; i++ {
		for j := topLeftX; j < topLeftX+width; j++ {
			if floorContent[i][j] != content {
				sameContent = false
			}
		}
	}
	//si oui le content du noeud actuel prend la valeur du tableau
	if sameContent {
		currentNode.content = floorContent[topLeftY][topLeftX]
		return currentNode
	} //sinon, on recrée des noeud enfants que l'on parcours aussi
	halfWidth := width / 2
	halfHeight := height / 2
	//ici, les calculs du type "width-halfWidth" sont fait car halfWidth et halfHeight sont arrondi à l'entier inférieur (11 / 2 = 5)
	//et donc d'où "width-halfWidth" car on veut 6 sinon ça marche pas
	currentNode.topLeftNode = createNode(topLeftX, topLeftY, halfWidth, halfHeight, floorContent)
	currentNode.topRightNode = createNode(topLeftX+halfWidth, topLeftY, width-halfWidth, halfHeight, floorContent)
	currentNode.bottomLeftNode = createNode(topLeftX, topLeftY+halfHeight, halfWidth, height-halfHeight, floorContent)
	currentNode.bottomRightNode = createNode(topLeftX+halfWidth, topLeftY+halfHeight, width-halfWidth, height-halfHeight, floorContent)
	return currentNode

}
