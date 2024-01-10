package floor

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

func adjustTilesWood(floorContent, newFloorContent [][]int, i, j int) {
	if i == 0 && j == 0 {
		if floorContent[i+1][j] != floorContent[i][j] && floorContent[i][j+1] != floorContent[i][j] {
			newFloorContent[i][j] = 755
		} else if floorContent[i+1][j] == floorContent[i][j] && floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = 754
		} else if floorContent[i+1][j] == floorContent[i][j] {
			newFloorContent[i][j] = 720
		} else if floorContent[i][j+1] == floorContent[i][j] {
			newFloorContent[i][j] = 721
		}
	} else if i == 0 && j == len(floorContent[i])-1 {

	} else if i == len(floorContent)-1 && j == 0 {

	} else if i == len(floorContent)-1 && j == len(floorContent[i])-1 {

	} else if i == 0 {

	} else if j == 0 {

	} else if i == len(floorContent)-1 {

	} else if j == len(floorContent[i])-1 {

	} else {

	}
}
