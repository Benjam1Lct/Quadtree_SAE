package floor

import (
	"fmt"
	"math/rand"
	"os"
)

// fonction permettant de créer un tableau aléatoire de largeur
// d'une valeur correspondant à la vaiable WidthRandomFloor et d'une hauteur correspondant à la variable HeightRandomFloor
func create_random_floor(width, height int) [][]int {
	terrain := make([][]int, height)
	for i := range terrain {
		terrain[i] = make([]int, width)
		for j := range terrain[i] {
			terrain[i][j] = rand.Intn(5)
		}
	}
	return terrain
}

// fonction permettant de rentrer les valeurs du tableau créé avec la fonction create_random_floor
// la fonction prend en paramètre les variables:
//   - terrain: un tableau de tableau d'entier qui correspond au tableau créer avec
//     avec la fonction create_random_floor
//   - filename: le nom du fichier contenant le terrain aléatoire (ici random_floor)
func writeTerrainToFile(terrain [][]int, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range terrain {
		for _, value := range row {
			_, err := fmt.Fprintf(file, "%d", value)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(file)
		if err != nil {
			return err
		}
	}

	return nil
}
