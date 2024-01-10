package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	// Extension waterblock
	// l'extension permet de bloquer le personnage lorsqu'il rencontre une case qui est de l'eau
	// la méthode permet déjà de bloquer le personnage sur les cases vides (-1), ainsi
	// nous avons juste à faire en sorte que le personnage soit aussi bloqué sur les cases d'eaux (ici 4).

	if configuration.Global.WaterBlocked {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1 || calcFloor(f.content[relativeYPos-1][relativeXPos])
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1 || calcFloor(f.content[relativeYPos][relativeXPos+1])
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1 || calcFloor(f.content[relativeYPos+1][relativeXPos])
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1 || calcFloor(f.content[relativeYPos][relativeXPos-1])
	} else {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}
	return blocking
}

func calcFloor(content int) (result bool) {
	// Liste des numéros de cases d'eau
	listeCasesEau := []int{406, 20, 21, 22, 23, 52, 53, 54, 55, 84, 85, 86, 87, 116, 117, 118, 119, 148, 149, 150, 151, 180, 181, 182, 183, 212, 213, 214, 215, 244, 245, 246, 247, 276, 277, 278, 279, 308, 309, 310, 311, 340, 341, 342, 343, 372, 373, 374, 375}

	// Numéro de la case actuelle
	numeroCaseActuelle := content

	// Vérifier si la case actuelle est une case d'eau
	var estCaseDEau bool
	for _, numero := range listeCasesEau {
		if numero == numeroCaseActuelle {
			estCaseDEau = true
			break
		}
	}

	// Bloquer le déplacement si c'est une case d'eau
	if estCaseDEau {
		return true
	} else {
		return false
	}

}
