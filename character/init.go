package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init() {
	c.animationStep = 1

	if configuration.Global.CameraMode == camera.Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}

	c.animationFlag = 1

	c.tp = Init_Teleport()

	configuration.Global.LstTeleport = [][]int{{-1, -1}, {-1, -1}}
}
