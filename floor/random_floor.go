package floor

import (
	"fmt"
	"math/rand"
	"os"
)

func create_random_floor(width, height int) [][]int {
	terrain := make([][]int, height)
	for i := range terrain {
		terrain[i] = make([]int, width)
		for j := range terrain[i] {
			terrain[i][j] = rand.Intn(4) + 1
		}
	}
	return terrain
}

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
