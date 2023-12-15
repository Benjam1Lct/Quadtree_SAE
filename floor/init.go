package floor

import (
	"bufio"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case fromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case quadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
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
