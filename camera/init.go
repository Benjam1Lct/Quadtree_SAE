package camera

import (
	"bufio"
	"fmt"
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
		format := bufio.NewScanner(file)
		maxWidth := 0
		Height := 0
		newFormat := false

		for format.Scan() {
			line := format.Text()
			for _, chara := range line {
				if chara == ',' {
					newFormat = true
				}
			}
		}

		file.Seek(0, 0)

		if newFormat {
			max := bufio.NewScanner(file)
			for max.Scan() {
				line := max.Text()
				parts := strings.Split(line, ",")
				if len(parts) > maxWidth {
					maxWidth = len(parts)
				}
			}
			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				Height = Height + 1
			}

		} else {
			max := bufio.NewScanner(file)
			for max.Scan() {
				line := max.Text()
				if len(line) > maxWidth {
					maxWidth = len(line)
				}
			}
			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				Height = Height + 1
			}
		}

		fmt.Println(maxWidth, Height)

		c.MapSizeX = maxWidth - 1
		c.MapSizeY = Height - 1
		c.X = configuration.Global.NumTileX / 2
		c.Y = configuration.Global.NumTileY / 2
	}

}
