package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool) {

	// Gestion de l'animation du personnage
	if configuration.Global.Animation {
		const animationDelay = 8 // Augmentez ou diminuez cette valeur pour ajuster la vitesse

		c.animationCounter += 1

		// Logique pour changer l'état de l'animation
		if c.animationCounter >= animationDelay {
			if c.animationFlag == 1 {
				c.animationFlag = 2
			} else if c.animationFlag == 2 {
				c.animationFlag = 3
			} else if c.animationFlag == 3 {
				c.animationFlag = 1
			}
			c.animationCounter = 0
		}
	}

	// Si le personnage ne bouge pas actuellement
	if !c.moving {
		// Gestion des touches pour démarrer le mouvement
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
			}
		}
	} else {
		// Si le personnage est en mouvement
		c.animationFrameCount++

		// Logique pour avancer le personnage avec animation
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep

			// Logique pour terminer le mouvement et mettre à jour les coordonnées
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
	}

	// Gestion des téléporteurs
	if configuration.Global.Teleport {
		if ebiten.IsKeyPressed(ebiten.KeyT) && c.tp.Tpress {
			c.tp.create_teleport(c.X, c.Y)
			c.tp.Tpress = false
		}
		if (c.X == c.tp.enterX && c.Y == c.tp.enterY) && !c.tp.onPortal && c.tp.endX != -1 {
			c.X, c.Y = c.tp.endX, c.tp.endY
			c.tp.onPortal = true
		} else if (c.X == c.tp.endX && c.Y == c.tp.endY) && !c.tp.onPortal {
			c.X, c.Y = c.tp.enterX, c.tp.enterY
			c.tp.onPortal = true
		}
		if (c.X != c.tp.enterX || c.Y != c.tp.enterY) && (c.X != c.tp.endX || c.Y != c.tp.endY) && c.tp.onPortal {
			c.tp.onPortal = false
			c.tp.Tpress = true
		}
	}
}
