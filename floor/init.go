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
		parts := strings.Split(line, ",")
		fmt.Println(line)
		if len(parts) > maxLength {
			maxLength = len(parts)
		}
	}
	fmt.Println(maxLength)
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
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
		} else if len(tab) < maxLength {
			for i := len(tab); i < maxLength; i++ {
				tab = append(tab, -1)
			}
		}
		// Afficher le tableau résultant
		fmt.Println("Tableau de nombres :", tab)
		floorContent = append(floorContent, tab)
	}

	// Créer un nouveau tableau avec la même structure
	newFloorContent := make([][]int, len(floorContent))

	// Copier chaque slice intérieure individuellement
	for i, innerSlice := range floorContent {
		newFloorContent[i] = make([]int, len(innerSlice))
		copy(newFloorContent[i], innerSlice)
	}

	if configuration.Global.EnhanceFloor {

		for i := 0; i < len(floorContent); i++ {
			for j := 0; j < len(floorContent[i]); j++ {
				if i == 0 && j == 0 {
					if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 198
						} else {
							newFloorContent[i][j] = 4
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 100
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 {
						newFloorContent[i][j] = 7
					}
				} else if i == len(floorContent)-1 && j == len(floorContent[i])-1 {
					if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 231
						} else {
							newFloorContent[i][j] = 70
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 102
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 71
					}
				} else if i == 0 {
					if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j-1] == 41 && floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 197
						} else {
							newFloorContent[i][j] = 5
						}

					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j-1] == 41 {
							newFloorContent[i][j] = 199
						} else {
							newFloorContent[i][j] = 6
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 198
						} else {
							newFloorContent[i][j] = 4
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 {
						newFloorContent[i][j] = 39
					}
				} else if j == 0 {
					if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 {
						if floorContent[i+1][j+1] == 41 && floorContent[i-1][j+1] == 41 {
							newFloorContent[i][j] = 196
						} else {
							newFloorContent[i][j] = 36
						}
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 {
						if floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 198
						} else {
							newFloorContent[i][j] = 4
						}
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 && floorContent[i][j+1] == 41 {
						if floorContent[i-1][j+1] == 41 {
							newFloorContent[i][j] = 230
						} else {
							newFloorContent[i][j] = 68
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 39
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 {
						newFloorContent[i][j] = 39
					}
				} else if i == len(floorContent)-1 {
					if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i][j+1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i-1][j+1] == 41 && floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 228
						} else {
							newFloorContent[i][j] = 69
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i+1][j-1] == 41 {
							newFloorContent[i][j] = 199
						} else {
							newFloorContent[i][j] = 6
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 198
						} else {
							newFloorContent[i][j] = 4
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 39
					}
				} else if j == len(floorContent[i])-1 {
					if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j-1] == 41 {
						if floorContent[i+1][j-1] == 41 && floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 229
						} else {
							newFloorContent[i][j] = 38
						}
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j-1] == 41 {
						if floorContent[i+1][j-1] == 41 {
							newFloorContent[i][j] = 199
						} else {
							newFloorContent[i][j] = 6
						}
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 && floorContent[i][j-1] == 41 {
						if floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 231
						} else {
							newFloorContent[i][j] = 70
						}
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 39
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 102
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 71
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 {
						newFloorContent[i][j] = 7
					}
				} else {
					if floorContent[i][j] == 41 && floorContent[i][j-1] != 41 && floorContent[i-1][j] != 41 && floorContent[i+1][j] != 41 && floorContent[i][j+1] != 41 {
						newFloorContent[i][j] = 103
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j+1] == 41 && floorContent[i+1][j-1] == 41 && floorContent[i-1][j-1] == 41 && floorContent[i-1][j+1] == 41 {
						newFloorContent[i][j] = 41
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 37
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i][j+1] == 41 {
						if floorContent[i-1][j+1] == 41 && floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 228
						} else {
							newFloorContent[i][j] = 69
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i+1][j] == 41 && floorContent[i][j+1] == 41 {
						if floorContent[i+1][j-1] == 41 && floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 197
						} else {
							newFloorContent[i][j] = 5
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j-1] == 41 && floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 229
						} else {
							newFloorContent[i][j] = 38
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i-1][j] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j+1] == 41 && floorContent[i-1][j+1] == 41 {
							newFloorContent[i][j] = 196
						} else {
							newFloorContent[i][j] = 36
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j-1] == 41 {
							newFloorContent[i][j] = 199
						} else {
							newFloorContent[i][j] = 6
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i-1][j-1] == 41 {
							newFloorContent[i][j] = 231
						} else {
							newFloorContent[i][j] = 70
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i+1][j] == 41 {
						if floorContent[i+1][j+1] == 41 {
							newFloorContent[i][j] = 198
						} else {
							newFloorContent[i][j] = 4
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i-1][j] == 41 {
						if floorContent[i-1][j+1] == 41 {
							newFloorContent[i][j] = 230
						} else {
							newFloorContent[i][j] = 68
						}
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 101
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 39
					} else if floorContent[i][j] == 41 && floorContent[i-1][j] == 41 {
						newFloorContent[i][j] = 7
					} else if floorContent[i][j] == 41 && floorContent[i+1][j] == 41 {
						newFloorContent[i][j] = 71
					} else if floorContent[i][j] == 41 && floorContent[i][j-1] == 41 {
						newFloorContent[i][j] = 102
					} else if floorContent[i][j] == 41 && floorContent[i][j+1] == 41 {
						newFloorContent[i][j] = 100
					}
				}
			}
		}
	}

	fmt.Println(newFloorContent)
	return newFloorContent
}
