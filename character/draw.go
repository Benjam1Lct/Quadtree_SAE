package character

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des charactéristiques du personnage (position, orientation,
// étape d'animation, etc) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).

func (c Character) Draw(screen *ebiten.Image, camX, camY int) {

	xShift := 0
	yShift := 0
	switch c.orientation {
	case orientedDown:
		yShift = c.shift
	case orientedUp:
		yShift = -c.shift
	case orientedLeft:
		xShift = -c.shift
	case orientedRight:
		xShift = c.shift
	}

	xTileForDisplay := c.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := c.Y - camY + configuration.Global.ScreenCenterTileY
	xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if c.moving {
		shiftX += c.animationStep * configuration.Global.TileSize
	}
	shiftY := c.orientation * configuration.Global.TileSize

	screen.DrawImage(assets.CharacterImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)

	if configuration.Global.Teleport && c.tp.enterX != -1 {
		xTileForDisplay := c.tp.enterX - camX + configuration.Global.ScreenCenterTileX
		yTileForDisplay := c.tp.enterY - camY + configuration.Global.ScreenCenterTileY
		xPos := (xTileForDisplay)*configuration.Global.TileSize + configuration.Global.TileSize/2
		yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(xPos), float64(yPos))

		screen.DrawImage(assets.FloorImage.SubImage(
			image.Rect((c.animationFlag-1)*configuration.Global.TileSize, 31*configuration.Global.TileSize, c.animationFlag*configuration.Global.TileSize, 32*configuration.Global.TileSize),
		).(*ebiten.Image), op)

		if c.tp.endX != -1 {
			xTileForDisplay = c.tp.endX - camX + configuration.Global.ScreenCenterTileX
			yTileForDisplay = c.tp.endY - camY + configuration.Global.ScreenCenterTileY
			xPos = (xTileForDisplay)*configuration.Global.TileSize + configuration.Global.TileSize/2
			yPos = (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2

			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(xPos), float64(yPos))

			screen.DrawImage(assets.FloorImage.SubImage(
				image.Rect((c.animationFlag-1)*configuration.Global.TileSize, 30*configuration.Global.TileSize, c.animationFlag*configuration.Global.TileSize, 31*configuration.Global.TileSize),
			).(*ebiten.Image), op)
		}
	}
}
