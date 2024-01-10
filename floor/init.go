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
// Cette fonction initialise le contenu d'un étage (Floor) en fonction des paramètres de configuration.
// Possibilité de créer en aléatoire ou à partir d'un fichier avec ou sans quadtree
func (f *Floor) Init() {
	// Initialisation de la grille de contenu du floor avec des dimensions spécifiées dans la configuration globale
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	// Si la création du floor est aléatoire
	if configuration.Global.RandomFloor {
		f.initRandomFloor()
	} else {
		f.initNonRandomFloor()
	}
}

// initRandomFloor initialise un floor aléatoire
func (f *Floor) initRandomFloor() {
	randomFloor := create_random_floor(configuration.Global.WidthRandomFloor, configuration.Global.HeightRandomFloor)
	writeTerrainToFile(randomFloor, "../floor-files/random_floor")

	// Lecture du floor aléatoire depuis le fichier
	terrain := readFloorFromFile("../floor-files/random_floor")

	// Sauvegarde du floor aléatoire avec la date actuelle comme nom de fichier
	if configuration.Global.SaveRandomFloor {
		f.saveRandomFloor(terrain)
	}

	// Mise à jour de la position centrale ou [0][0] pour éviter au joueur de spawner sur de l'eau
	if configuration.Global.WaterBlocked {
		f.updateWaterBlock(terrain)
	}

	// Création d'un quadtree à partir du floor aléatoire
	f.quadtreeContent = quadtree.MakeFromArray(terrain)
}

// initNonRandomFloor initialise un floor non aléatoire en fonction de la configuration globale
func (f *Floor) initNonRandomFloor() {
	switch configuration.Global.FloorKind {
	case fromFileFloor:
		f.initFromFileFloor()
	case quadTreeFloor:
		f.initQuadTreeFloor()
	}
}

// initFromFileFloor initialise un floor à partir d'un fichier
func (f *Floor) initFromFileFloor() {
	// Lecture du floor depuis un fichier spécifié dans la configuration globale
	f.fullContent = readFloorFromFile(configuration.Global.FloorFile)

	// Mise à jour de la position centrale ou [0][0] pour éviter au joueur de spawner sur de l'eau
	if configuration.Global.WaterBlocked {
		f.updateWaterBlock(f.fullContent)
	}

}

// initQuadTreeFloor initialise un floor à partir d'un quadtree
func (f *Floor) initQuadTreeFloor() {
	// Lecture du floor depuis un fichier spécifié dans la configuration globale
	terrainQuadtree := readFloorFromFile(configuration.Global.FloorFile)

	// Mise à jour de la position centrale ou [0][0] pour éviter au joueur de spawner sur de l'eau
	if configuration.Global.WaterBlocked {
		f.updateWaterBlock(terrainQuadtree)
	}

	// Création d'un quadtree à partir du floor lu depuis le fichier
	f.quadtreeContent = quadtree.MakeFromArray(terrainQuadtree)
}

// saveRandomFloor sauvegarde un floor aléatoire avec la date actuelle comme nom de fichier
func (f *Floor) saveRandomFloor(terrain [][]int) {
	// Création du dossier "save" s'il n'existe pas
	saveDir := "saveFloors"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.Mkdir(saveDir, os.ModePerm)
	}

	// Ajout de la date actuelle au nom de fichier
	currentDateTime := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("%s_%s.txt", "randomFloor", currentDateTime)
	filePath := fmt.Sprintf("%s/%s", saveDir, fileName)

	// Création et ouverture du fichier
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer file.Close()

	// Écriture des lignes dans le fichier
	for _, sublist := range terrain {
		line := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sublist)), ","), "[]")
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
	}

	// Affichage du message de réussite
	fmt.Printf("Fichier '%s' créé avec succès dans le dossier '%s'.\n", fileName, saveDir)
}

// updateWaterBlock met à jour la position centrale ou [0][0] pour éviter au joueur de spawner sur de l'eau
func (f *Floor) updateWaterBlock(terrain [][]int) {
	if configuration.Global.CameraMode == 0 && calcFloor(terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY]) {
		terrain[configuration.Global.ScreenCenterTileX][configuration.Global.ScreenCenterTileY] = 33
	} else if (configuration.Global.CameraMode == 1 || configuration.Global.CameraMode == 2) && calcFloor(terrain[0][0]) {
		terrain[0][0] = 33
	}
}

// adjustTiles ajuste le contenu de la tuile en pierre ou sable ou en eau en fonction de son emplacement dans la grille.
// elle verifier les tuile adjacente et determine quelle texture améliorer définir
// le sprite a selection est determiner par un somme entre un coeficient fixe et la valeur du floorcontent ou l'on se situe
// cela permet de determiner dynamiquement les texture en partant du principe qu'une meme texture de virage par exemple pour l'eau et le sable se trouve a la meme distance que la texture par defaut
// une variable typeFloor est presente afin de determiner les bonne coordonnées avec les decalage necessaires car la texture par defaut de l'eau n'est pas au meme endroid que les autres
func adjustTiles(floorContent, newFloorContent [][]int, i, j, typeFloor int) {
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
				newFloorContent[i][j] = floorContent[i][j] + 27 - typeFloor
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

// updateFloor met à jour le contenu du floor en fonction des règles spécifiées.
// La fonction parcourt chaque élément du tableau floorContent et applique des modifications
// selon les valeurs spécifiques rencontrées.
// Les règles sont les suivantes :
//   - Si la valeur est 41 ou 61 pour le sable et la pierre, ajuster les carreaux en appelant la fonction adjustTiles avec un décalage de 0.
//   - Si la valeur est 406 pour l'eau, ajuster les carreaux en appelant la fonction adjustTiles avec un décalage de 349.
//
// Paramètres :
//   - floorContent : tableau représentant le contenu actuel du floor, ce tableau sert de reference et n'est pas modifier a l'appel de la fonction adjustTiles
//   - newFloorContent : tableau représentant le nouveau contenu du floor après les modifications.
//
// La fonction adjustTiles est appelée pour effectuer les ajustements spécifiques sur les carreaux.
func updateFloor(floorContent, newFloorContent [][]int) {
	for i := 0; i < len(floorContent); i++ {
		for j := 0; j < len(floorContent[i]); j++ {
			switch floorContent[i][j] {
			case 41, 61:
				adjustTiles(floorContent, newFloorContent, i, j, 0)
			case 406:
				adjustTiles(floorContent, newFloorContent, i, j, 349)
			}
		}
	}
}

// readFloorFromFile lit le contenu d'un fichier représentant un terrain
// et le stocke dans un tableau 2D. Les lignes plus courtes sont remplies avec -1
// pour obtenir un tableau rectangulaire.
// Du au changement d'assets pour le floor les coordonnées des premmiere textures ne sont plus les memes
// une nouvelle boucle est donc presente afin de pouvoir lire les anciens fichiers et attribuer les nouvelle valeurs de floor afin qu'il puisse fonctionné et etre affiché
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// Ouvrir le fichier spécifié
	filePath := fileName
	file, err := os.Open(filePath)
	if err != nil {
		return floorContent
	}
	defer file.Close()

	// Utiliser un scanner pour trouver la longueur maximale des lignes dans le fichier
	max := bufio.NewScanner(file)
	maxLength := 0
	newFormat := false

	// Parcourir chaque ligne du fichier
	for max.Scan() {
		line := max.Text()
		parts := strings.Split(line, ",")

		// Mettre à jour la longueur maximale si la ligne actuelle est plus longue
		if len(parts) > maxLength {
			maxLength = len(parts)
		}

		// Vérifier si le format de la ligne est nouveau (contient des virgules), si ou le garder en memoire
		for _, chara := range line {
			if chara == ',' {
				newFormat = true
			}
		}
	}
	// Retourner au début du fichier
	file.Seek(0, 0)

	// Utiliser un autre scanner pour lire le contenu réel du fichier
	scanner := bufio.NewScanner(file)
	// si le format du fichier est nouveau (contient pas virgules)
	// permet de traitet le nouveau format des fichiers floor avec les nouvelles textures
	if newFormat {
		// Traitement pour le nouveau format avec des valeurs séparées par des virgules
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

			// Ignorer les lignes vides ou "newformat"
			if len(line) == 0 || line == "newformat" {
				continue
			} else if len(tab) < maxLength {
				// Remplir avec des valeurs par défaut (-1) si la ligne est plus courte que la longueur maximale
				for i := len(tab); i < maxLength; i++ {
					tab = append(tab, -1)
				}
			}
			// Ajouter la ligne au contenu du floor
			floorContent = append(floorContent, tab)
		}
		// permet de traitet l'ancien format des fichiers floor
		// afin d'afficher le floor depuis les anciens fichier floor-files
	} else {
		// Traitement pour l'ancien format sans virgules
		max := bufio.NewScanner(file)
		maxLength := 0

		// Trouver la longueur maximale
		for max.Scan() {
			line := max.Text()
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}
		// Retourner au début du fichier
		file.Seek(0, 0)

		// Lire chaque ligne du fichier
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
					// Mapper les valeurs selon certaines règles
					if num == 0 { // pour la tuile d'herbe
						tab = append(tab, 33)
					} else if num == 1 { // pour la tuile de sable
						tab = append(tab, 41)
					} else if num == 2 { // pour la tuile de pierre
						tab = append(tab, 61)
					} else if num == 3 { // pour la tuile de bois
						tab = append(tab, 657)
					} else if num == 4 { // pour la tuile d'eau
						tab = append(tab, 406)
					}
				} else {
					// Dans le cas où la ligne est plus courte que la longueur maximale, remplir avec -1
					tab = append(tab, -1)
				}
			}
			// Ajouter la ligne au contenu du floor
			floorContent = append(floorContent, tab)
		}
	}

	// Créer une nouvelle copie du tableau avec la même structure
	newFloorContent := make([][]int, len(floorContent))
	// Copier chaque sous-tableau individuellement
	for i, innerSlice := range floorContent {
		newFloorContent[i] = make([]int, len(innerSlice))
		copy(newFloorContent[i], innerSlice)
	}

	// Mettre à jour le contenu du floor si l'amélioration des textures est activée
	if configuration.Global.EnhanceFloor {
		updateFloor(floorContent, newFloorContent)
	}

	// Retourner le nouveau contenu du floor
	return newFloorContent
}
