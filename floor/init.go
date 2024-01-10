package floor

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}
	if configuration.Global.RandomFloor {
		random_floor := create_random_floor(configuration.Global.WidthRandomFloor, configuration.Global.HeightRandomFloor)
		writeTerrainToFile(random_floor, "../floor-files/random_floor")
		terrain := readFloorFromFile("../floor-files/random_floor")
		if configuration.Global.WaterBlocked {
			if configuration.Global.CameraMode == 0 && terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] == 4 {
				terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = rand.Intn(4)
			} else if configuration.Global.CameraMode == 1 && terrain[0][0] == 4 {
				terrain[0][0] = rand.Intn(4)
			}
		}
		f.quadtreeContent = quadtree.MakeFromArray(terrain)
	} else {
		switch configuration.Global.FloorKind {
		case fromFileFloor:
			f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
			if configuration.Global.WaterBlocked {
				if configuration.Global.CameraMode == 0 && f.fullContent[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] == 4 {
					f.fullContent[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = rand.Intn(4)
				} else if configuration.Global.CameraMode == 1 && f.fullContent[0][0] == 4 {
					f.fullContent[0][0] = rand.Intn(4)
				}
			}
		case quadTreeFloor:
			terrain_quadtree := readFloorFromFile(configuration.Global.FloorFile)
			if configuration.Global.WaterBlocked {
				if configuration.Global.CameraMode == 0 && terrain_quadtree[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] == 4 {
					terrain_quadtree[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = rand.Intn(4)
				} else if configuration.Global.CameraMode == 1 && terrain_quadtree[0][0] == 4 {
					terrain_quadtree[0][0] = rand.Intn(4)
				}
			}
			f.quadtreeContent = quadtree.MakeFromArray(terrain_quadtree)
		case SphereWorld:
			f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
		}

	}
}

func adjustTile(floorContent, newFloorContent [][]int, i, j int) {
	if floorContent[i][j] == 41 {
		if i > 0 && j > 0 && i < len(floorContent)-1 && j < len(floorContent[i])-1 {
			if floorContent[i][j-1] != 41 && floorContent[i-1][j] != 41 && floorContent[i+1][j] != 41 && floorContent[i][j+1] != 41 {
				newFloorContent[i][j] = 103
			} else if floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j+1] == 41 && floorContent[i+1][j-1] == 41 && floorContent[i-1][j-1] == 41 && floorContent[i-1][j+1] == 41 {
				newFloorContent[i][j] = 41
			} else if floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 {
				newFloorContent[i][j] = 37
			} else {

			}
		}
	}
}

func updateFloor(floorContent, newFloorContent [][]int) {
	for i := 0; i < len(floorContent); i++ {
		for j := 0; j < len(floorContent[i]); j++ {
			if i == 0 && j == 0 {

			} else if i == len(floorContent)-1 && j == len(floorContent[i])-1 {

			} else if i == 0 {

			} else if j == 0 {

			} else if i == len(floorContent)-1 {

			} else if j == len(floorContent[i])-1 {

			} else {
				adjustTile(floorContent, newFloorContent, i, j)
			}
		}
	}
}

// readFloorFromFile lit le contenu d'un fichier représentant un terrain
// et le stocke dans un tableau 2D. Les lignes plus courtes sont remplies avec -1
// pour obtenir un tableau rectangulaire.
func readFloorFromFile(fileName string) (floorContent [][]int) {
	filePath := fileName
	file, err := os.Open(filePath)
	if err != nil {
		return floorContent
	}
	defer file.Close()

	max := bufio.NewScanner(file)
	maxLength := 0
	newFormat := false

	for max.Scan() {
		line := max.Text()
		parts := strings.Split(line, ",")
		if len(parts) > maxLength {
			maxLength = len(parts)
		}
		for _, chara := range line {
			if chara == ',' {
				newFormat = true
			}
		}
	}
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	if newFormat {
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ",")

			// Créer un tableau pour stocker les nombres
			var tab []int

			// Convertir chaque partie en entier et ajouter au tableau
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					// Gérer l'erreur, par exemple, imprimer un message
					fmt.Println("Erreur de conversion en entier:", err)
					break
				}
				tab = append(tab, num)
			}

			if len(line) == 0 {
				continue
			} else if line == "newformat" {
				continue
			} else if len(tab) < maxLength {
				for i := len(tab); i < maxLength; i++ {
					tab = append(tab, -1)
				}
			}
			floorContent = append(floorContent, tab)
		}

	} else {
		max := bufio.NewScanner(file)
		maxLength := 0

		// Trouver la longueur maximale
		for max.Scan() {
			line := max.Text()
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}
		file.Seek(0, 0)

		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}
			var tab []int = make([]int, 0, maxLength)
			for i := 0; i < maxLength; i++ {
				if i < len(line) {
					num, err := strconv.Atoi(string(line[i]))
					if err != nil {
						break
					}
					if num == 0 {
						tab = append(tab, 33)
					} else if num == 1 {
						tab = append(tab, 41)
					} else if num == 2 {
						tab = append(tab, 61)
					} else if num == 3 {
						tab = append(tab, 657)
					} else if num == 4 {
						tab = append(tab, 406)
					}
				} else {
					// Dans le cas où la ligne est plus courte que la longueur maximale
					tab = append(tab, -1)
				}
			}
			floorContent = append(floorContent, tab)
		}
	}

	// Créer un nouveau tableau avec la même structure
	newFloorContent := make([][]int, len(floorContent))

	// Copier chaque slice intérieure individuellement
	for i, innerSlice := range floorContent {
		newFloorContent[i] = make([]int, len(innerSlice))
		copy(newFloorContent[i], innerSlice)
	}

	if configuration.Global.EnhanceFloor {
		updateFloor(floorContent, newFloorContent)
	}

	return newFloorContent
}
