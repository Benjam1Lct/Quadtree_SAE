package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY)
	case NoVoid:
		c.updateNoVoid(characterPosX, characterPosY)
	case MovieCamera:
		c.updateMovieCamera(characterPosX, characterPosY)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int) {
	c.X = characterPosX
	c.Y = characterPosY
}

func (c *Camera) updateNoVoid(characterPosX, characterPosY int) {
	halfNumTileX := configuration.Global.NumTileX / 2
	halfNumTileY := configuration.Global.NumTileY / 2
	c.X = clamp(characterPosX, halfNumTileX, c.MapSizeX-halfNumTileX)
	c.Y = clamp(characterPosY, halfNumTileY, c.MapSizeY-halfNumTileY)
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func (c *Camera) updateMovieCamera(characterPosX, characterPosY int) {
	spaceBeforeMove := 2

	diffCamCharX := characterPosX - c.X
	diffCamCharY := characterPosY - c.Y

	if diffCamCharX > spaceBeforeMove || diffCamCharX < -spaceBeforeMove {
		c.X = characterPosX
		c.Y = characterPosY
	}

	if diffCamCharY > spaceBeforeMove || diffCamCharY < -spaceBeforeMove {
		c.X = characterPosX
		c.Y = characterPosY
	}
}
