package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// updateFromFileFloor met à jour le contenu du sol à partir d'un tableau lu depuis un fichier.
// Il prend en compte la position de la caméra (camXPos, camYPos) pour extraire la partie pertinente du tableau.
// Les cases en dehors des limites du tableau d'origine sont remplies avec -1.
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	// Calcul des indices de la zone d'intérêt dans le tableau d'origine
	inter_x := []int{camXPos - configuration.Global.NumTileX/2, (camXPos + configuration.Global.NumTileX/2) + configuration.Global.NumTileX%2}
	inter_y := []int{camYPos - configuration.Global.NumTileY/2, (camYPos + configuration.Global.NumTileY/2) + configuration.Global.NumTileY%2}

	// Réinitialiser le contenu actuel
	f.content = make([][]int, (configuration.Global.NumTileY))

	for i := inter_y[0]; i < inter_y[1]; i++ {
		if i < 0 || i >= len(f.fullContent) {
			for x := 0; x < configuration.Global.NumTileX; x++ {
				// Ajouter des cases vides pour les lignes en dehors des limites de f.fullContent
				f.content[i-inter_y[0]] = append(f.content[i-inter_y[0]], -1)
			}
			continue
		}

		for j := inter_x[0]; j < inter_x[1]; j++ {
			if j < 0 || j >= len(f.fullContent[i]) {
				// Ajouter des cases vides pour les colonnes en dehors des limites de f.fullContent
				f.content[i-inter_y[0]] = append(f.content[i-inter_y[0]], -1)
			} else {
				f.content[i-inter_y[0]] = append(f.content[i-inter_y[0]], f.fullContent[i][j])
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
