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
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1 || f.content[relativeYPos-1][relativeXPos] == 4
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1 || f.content[relativeYPos][relativeXPos+1] == 4
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1 || f.content[relativeYPos+1][relativeXPos] == 4
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1 || f.content[relativeYPos][relativeXPos-1] == 4
	} else {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}

	return blocking
}
