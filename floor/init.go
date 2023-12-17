package floor

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"

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

	for max.Scan() {
		line := max.Text()
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
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
				tab = append(tab, num)
			} else {
				tab = append(tab, -1)
			}
		}
		floorContent = append(floorContent, tab)
	}
	return floorContent
}
