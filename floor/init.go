package floor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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

		if configuration.Global.SaveRandomFloor {
			// Créer un dossier "save" s'il n'existe pas
			saveDir := "saveFloors"
			if _, err := os.Stat(saveDir); os.IsNotExist(err) {
				os.Mkdir(saveDir, os.ModePerm)
			}

			// Ajouter la date du jour au nom de fichier
			currentDateTime := time.Now().Format("2006-01-02_15-04-05")
			fileName := fmt.Sprintf("%s_%s.txt", "randomFloor", currentDateTime)
			filePath := fmt.Sprintf("%s/%s", saveDir, fileName)

			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier :", err)
				return
			}
			defer file.Close()

			for _, sublist := range terrain {
				line := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sublist)), ","), "[]")
				_, err := fmt.Fprintln(file, line)
				if err != nil {
					fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
					return
				}
			}

			fmt.Printf("Fichier '%s' créé avec succès dans le dossier '%s'.\n", fileName, saveDir)
		}

		if configuration.Global.WaterBlocked {
			if configuration.Global.CameraMode == 0 && calcFloor(terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY]) {
				terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = 33
			} else if (configuration.Global.CameraMode == 1 || configuration.Global.CameraMode == 2) && calcFloor(terrain[0][0]) {
				terrain[0][0] = 33
			}
		}
		f.quadtreeContent = quadtree.MakeFromArray(terrain)

	} else {
		switch configuration.Global.FloorKind {
		case fromFileFloor:
			f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
			if configuration.Global.WaterBlocked {
				if configuration.Global.CameraMode == 0 && calcFloor(f.fullContent[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY]) {
					f.fullContent[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = 33
				} else if (configuration.Global.CameraMode == 1 || configuration.Global.CameraMode == 2) && calcFloor(f.fullContent[0][0]) {
					f.fullContent[0][0] = 33
				}
			}
		case quadTreeFloor:
			terrain_quadtree := readFloorFromFile(configuration.Global.FloorFile)
			if configuration.Global.WaterBlocked {
				if configuration.Global.CameraMode == 0 && calcFloor(terrain_quadtree[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY]) {
					terrain_quadtree[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = 33
				} else if (configuration.Global.CameraMode == 1 || configuration.Global.CameraMode == 2) && calcFloor(terrain_quadtree[0][0]) {
					terrain_quadtree[0][0] = 33
				}
			}
			f.quadtreeContent = quadtree.MakeFromArray(terrain_quadtree)
		}
	}
}

func adjustTileStoneSable(floorContent, newFloorContent [][]int, i, j, typeFloor int) {
	if i == 0 && j == 0 {
		if floorContent[i][j+1] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 157 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 37 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 34 - typeFloor
		}
	} else if i == 0 && j == len(floorContent[i])-1 {
		if floorContent[i][j-1] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 158 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 35 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 34 - typeFloor
		}
	} else if i == len(floorContent)-1 && j == 0 {
		if floorContent[i][j+1] != floorContent[i][j] && floorContent[i-1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 189 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 27 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		}
	} else if i == len(floorContent)-1 && j == len(floorContent[i])-1 {
		if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 190 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 29 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		}
	} else if i == len(floorContent)-1 && j == len(floorContent[i])-1 {
		if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 190 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 29 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		}
	} else if i == 0 {
		if floorContent[i][j-1] != floorContent[i][j] && floorContent[i][j+1] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 156 - typeFloor
			} else if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 285 - typeFloor
			} else if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 286 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 36 - typeFloor
			}

		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 158 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 35 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 157 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 37 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 60 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 2 - typeFloor
		}
	} else if j == 0 {
		if floorContent[i][j+1] != floorContent[i][j] && floorContent[i-1][j] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 155 - typeFloor
			} else if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 315 - typeFloor
			} else if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 283 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 5 - typeFloor
			}
		} else if floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 157 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 37 - typeFloor
			}
		} else if floorContent[i-1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 189 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 27 - typeFloor
			}
		} else if floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 2 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 34 - typeFloor
		}
	} else if i == len(floorContent)-1 {
		if floorContent[i][j-1] != floorContent[i][j] && floorContent[i-1][j] != floorContent[i][j] && floorContent[i][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 187 - typeFloor
			} else if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 318 - typeFloor
			} else if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 317 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 28 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 190 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 29 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 189 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 27 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 60 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 2 - typeFloor
		}
	} else if j == len(floorContent[i])-1 {
		if floorContent[i][j-1] != floorContent[i][j] && floorContent[i-1][j] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j-1] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 188 - typeFloor
			} else if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 316 - typeFloor
			} else if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 284 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 3 - typeFloor
			}
		} else if floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j-1] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 158 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 35 - typeFloor
			}
		} else if floorContent[i-1][j] == floorContent[i][j] && floorContent[i][j-1] == floorContent[i][j] {
			if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 190 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 29 - typeFloor
			}
		} else if floorContent[i+1][j] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 2 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 34 - typeFloor
		}
	} else {
		if floorContent[i][j-1] != floorContent[i][j] && floorContent[i-1][j] != floorContent[i][j] && floorContent[i+1][j] != floorContent[i][j] && floorContent[i][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 62 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j]
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 4 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 222 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 221 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 253 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 254 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 93 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 94 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 126 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 125 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 91 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 92 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] && floorContent[i-1][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 124 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j+1] != floorContent[i][j] && floorContent[i+1][j-1] != floorContent[i][j] && floorContent[i-1][j-1] != floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 123 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 187 - typeFloor
			} else if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 318 - typeFloor
			} else if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 317 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 28 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 156 - typeFloor
			} else if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 285 - typeFloor
			} else if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 286 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 36 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] && floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 188 - typeFloor
			} else if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 316 - typeFloor
			} else if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 284 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 3 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] && floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 155 - typeFloor
			} else if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 315 - typeFloor
			} else if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 283 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 5 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 158 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 35 - typeFloor
			}
		} else if floorContent[i][j-1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j-1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 190 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 29 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i+1][j] == floorContent[i][j] {
			if floorContent[i+1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 157 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] - 37 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			if floorContent[i-1][j+1] == floorContent[i][j] {
				newFloorContent[i][j] = floorContent[i][j] + 189 - typeFloor
			} else {
				newFloorContent[i][j] = floorContent[i][j] + 27 - typeFloor
			}
		} else if floorContent[i][j+1] == floorContent[i][j] && floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 60 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] && floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 2 - typeFloor
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] - 34 - typeFloor
		} else if floorContent[i-1][j] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 30 - typeFloor
		} else if floorContent[i][j-1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 61 - typeFloor
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = floorContent[i][j] + 59 - typeFloor
		}

	}
}

func updateFloor(floorContent, newFloorContent [][]int) {
	for i := 0; i < len(floorContent); i++ {
		for j := 0; j < len(floorContent[i]); j++ {
			switch floorContent[i][j] {
			case 41, 61:
				adjustTileStoneSable(floorContent, newFloorContent, i, j, 0)
			case 406:
				adjustTileStoneSable(floorContent, newFloorContent, i, j, 349)
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

	fmt.Println(newFloorContent)
	return newFloorContent
}
