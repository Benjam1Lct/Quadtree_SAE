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

	if configuration.Global.WaterBlocked {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1 || f.content[relativeYPos-1][relativeXPos] == 406
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1 || f.content[relativeYPos][relativeXPos+1] == 406
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1 || f.content[relativeYPos+1][relativeXPos] == 406
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1 || f.content[relativeYPos][relativeXPos-1] == 406
	} else {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}

	//fmt.Println("from blocking", blocking)
	return blocking
}
