package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"bufio"
	"os"
	"strconv"
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

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	filePath := fileName

    file, err := os.Open(filePath)
    if err != nil {
        return floorContent
    }
	
    defer file.Close() 

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var tab []int = make([]int, 0, len(line))
		for i:=0; i<len(line); i++ {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				break
			}
			tab = append(tab, num)
		}
        floorContent = append(floorContent,tab)
    }
	return floorContent
}
