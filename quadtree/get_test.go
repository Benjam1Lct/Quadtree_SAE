package quadtree

import (
	"reflect"
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func Test_simple_get(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	}

	input_topleftnode := node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	input_toprightnode := node{
		topLeftX:        2,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	input_bottomleftnode := node{
		topLeftX:        0,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	input_bottomrightnode := node{
		topLeftX:        2,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	output_root := Quadtree{
		width:  4,
		height: 4,
		root: &node{
			topLeftX:        0,
			topLeftY:        0,
			width:           4,
			height:          4,
			content:         0,
			topLeftNode:     &input_topleftnode,
			topRightNode:    &input_toprightnode,
			bottomLeftNode:  &input_bottomleftnode,
			bottomRightNode: &input_bottomrightnode,
		},
	}

	topLeftX := 0
	topLeftY := 0

	output_contentHolder := output_getcontent(topLeftX, topLeftY, floorContent)

	contentHolder := make([][]int, (configuration.Global.NumTileY))

	output_root.GetContent(topLeftX, topLeftY, contentHolder)

	if !are2DArraysEqual(output_contentHolder, contentHolder) {
		t.Errorf("Cas de test Getcontent_simpleArray échoué. Attendu : %v, Obtenu : %v", output_contentHolder, contentHolder)
	}
}

func Test_complexe_get(t *testing.T) {
	floorContent := [][]int{
		{1, 2, 2},
		{1, 2, 2},
		{4, 4, 4},
		{3, 4, 4},
	}

	output_topleftnode := node{
		topLeftX:        0,
		topLeftY:        0,
		width:           1,
		height:          2,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	output_toprightnode := node{
		topLeftX:        1,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	outputbottomleft_topleftnode := node{
		topLeftX:        0,
		topLeftY:        2,
		width:           0,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	outputbottomleft_toprightnode := node{
		topLeftX:        0,
		topLeftY:        2,
		width:           1,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	outputbottomleft_bottomleftnode := node{
		topLeftX:        0,
		topLeftY:        3,
		width:           0,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	outputbottomleft_bottomrighttnode := node{
		topLeftX:        0,
		topLeftY:        3,
		width:           1,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	output_bottomleftnode := node{
		topLeftX:        0,
		topLeftY:        2,
		width:           1,
		height:          2,
		content:         0,
		topLeftNode:     &outputbottomleft_topleftnode,
		topRightNode:    &outputbottomleft_toprightnode,
		bottomLeftNode:  &outputbottomleft_bottomleftnode,
		bottomRightNode: &outputbottomleft_bottomrighttnode,
	}
	output_bottomrightnode := node{
		topLeftX:        1,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	noderoot := &node{
		topLeftX:        0,
		topLeftY:        0,
		width:           len(floorContent[0]),
		height:          len(floorContent),
		content:         0,
		topLeftNode:     &output_topleftnode,
		topRightNode:    &output_toprightnode,
		bottomLeftNode:  &output_bottomleftnode,
		bottomRightNode: &output_bottomrightnode,
	}

	output_root := Quadtree{
		width:  len(floorContent[0]),
		height: len(floorContent),
		root:   noderoot,
	}

	topLeftX := 0
	topLeftY := 0

	output_contentHolder := output_getcontent(topLeftX, topLeftY, floorContent)

	contentHolder := make([][]int, (configuration.Global.NumTileY))

	output_root.GetContent(topLeftX, topLeftY, contentHolder)

	if !are2DArraysEqual(output_contentHolder, contentHolder) {
		t.Errorf("Cas de test Getcontent_simpleArray échoué. Attendu : %v, Obtenu : %v", output_contentHolder, contentHolder)
	}
}

func output_getcontent(camXPos, camYPos int, fullContent [][]int) (result [][]int) {
	// Calcul des indices de la zone d'intérêt dans le tableau d'origine
	inter_x := []int{camXPos - configuration.Global.NumTileX/2, (camXPos + configuration.Global.NumTileX/2) + configuration.Global.NumTileX%2}
	inter_y := []int{camYPos - configuration.Global.NumTileY/2, (camYPos + configuration.Global.NumTileY/2) + configuration.Global.NumTileY%2}

	// Réinitialiser le contenu actuel
	content := make([][]int, (configuration.Global.NumTileY))

	for i := inter_y[0]; i < inter_y[1]; i++ {
		if i < 0 || i >= len(fullContent) {
			for x := 0; x < configuration.Global.NumTileX; x++ {
				// Ajouter des cases vides pour les lignes en dehors des limites de f.fullContent
				content[i-inter_y[0]] = append(content[i-inter_y[0]], -1)
			}
			continue
		}

		for j := inter_x[0]; j < inter_x[1]; j++ {
			if j < 0 || j >= len(fullContent[i]) {
				// Ajouter des cases vides pour les colonnes en dehors des limites de f.fullContent
				content[i-inter_y[0]] = append(content[i-inter_y[0]], -1)
			} else {
				content[i-inter_y[0]] = append(content[i-inter_y[0]], fullContent[i][j])
			}
		}
	}
	return content
}

func are2DArraysEqual(arr1, arr2 [][]int) bool {
	return reflect.DeepEqual(arr1, arr2)
}
