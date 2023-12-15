package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		Entrées :
		q : un quadtree
		topLeftX : coordonnée X du coin en haut à gauche
		topLeftY : coordonnée Y du coin en haut à gauche
		contentHolder : tableau qui sera utilisé pour afficher le sol

		Fonction qui récupère un quadtree et qui remplit contentHolder à partir de ce quadtree
	*/

	/*Parcours les coordonnées par rapport à la taille maximum de l'affichage
	Et on regarde si chaque case appartient au tableau
	Si c'est le cas alors on va chercher le content de cette case dans l'arbre*/
	for i := 0; i < configuration.Global.NumTileX; i++ {
		for j := 0; j < configuration.Global.NumTileY; j++ {
			X := topLeftX + i
			Y := topLeftY + j
			if X < q.width && Y < q.height && X >= 0 && Y >= 0 { /*On s'assure que la case est dans le tableau*/
				contentHolder[j][i] = SearchContent(X, Y, q.root)
			} else {
				contentHolder[j][i] = -1
			}
		}
	}
}

func SearchContent(X, Y int, n *node) int {
	/*
		Entrées :
		X et Y : des coordonnées
		n : un noeud du quadtree

		Sortie :
		un entier qui représente le content

		Fonction qui parcours un quadtree de manière récursive pour trouver le content de la case du tableau
	*/
	if n.topLeftNode == nil {
		return n.content
	}

	halfWidth := n.width / 2
	halfHeight := n.height / 2
	centerX := n.topLeftX + halfWidth
	centerY := n.topLeftY + halfHeight

	if X < centerX { /*On compare le centre aux coordonnées pour savoir dans quel noeud on veut aller*/
		if Y < centerY {
			return SearchContent(X, Y, n.topLeftNode)
		} else {
			return SearchContent(X, Y, n.bottomLeftNode)
		}
	} else {
		if Y < centerY {
			return SearchContent(X, Y, n.topRightNode)
		} else {
			return SearchContent(X, Y, n.bottomRightNode)
		}
	}
}
