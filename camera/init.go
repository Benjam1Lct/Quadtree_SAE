package camera

import (
	"bufio"
	"os"
	"strings"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise la caméra en fonction de la configuration du jeu.
func (c *Camera) Init() {
	// Vérifie le mode de la caméra défini dans la configuration globale
	if configuration.Global.CameraMode == Static {
		// En mode caméra statique, place la caméra au centre de l'écran
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	} else if configuration.Global.CameraMode == NoVoid {
		// En mode NoVoid, ajuste la caméra en fonction des données du fichier de terrain

		// Ouvre le fichier de terrain spécifié dans la configuration globale
		filePath := configuration.Global.FloorFile
		file, _ := os.Open(filePath)
		defer file.Close()

		// Scanner pour trouver la longueur maximale d'une ligne dans le fichier
		format := bufio.NewScanner(file)
		maxWidth := 0
		Height := 0
		newFormat := false

		// Analyse du fichier pour déterminer le format
		for format.Scan() {
			line := format.Text()
			for _, chara := range line {
				if chara == ',' {
					newFormat = true
				}
			}
		}
		file.Seek(0, 0)

		// Détermine la largeur et la hauteur de la carte en fonction du format du fichier
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

		// Ajuste les dimensions de la caméra et la place au centre de la carte
		c.MapSizeX = maxWidth - 1
		c.MapSizeY = Height - 1
		c.X = configuration.Global.NumTileX / 2
		c.Y = configuration.Global.NumTileY / 2
	}
}
