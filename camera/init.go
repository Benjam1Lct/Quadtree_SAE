package camera

import (
	"bufio"
	"os"
	"strings"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place une camÃ©ra.
func (c *Camera) Init() {
	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	} else if configuration.Global.CameraMode == NoVoid {
		filePath := configuration.Global.FloorFile
		file, _ := os.Open(filePath)
		defer file.Close()

		// Scanner pour trouver la longueur maximale d'une ligne dans le fichier
		max := bufio.NewScanner(file)
		maxWidth := 0

		for max.Scan() {
			line := max.Text()
			parts := strings.Split(line, ",")
			if len(parts) > maxWidth {
				maxWidth = len(parts)
			}
		}

		file.Seek(0, 0)

		scanner := bufio.NewScanner(file)
		Height := 0
		for scanner.Scan() {
			Height = Height + 1
		}
		c.MapSizeX = maxWidth - 1
		c.MapSizeY = Height - 1
		c.X = configuration.Global.NumTileX / 2
		c.Y = configuration.Global.NumTileY / 2
	}

}
