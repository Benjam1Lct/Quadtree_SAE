package quadtree

import (
	"reflect"
	"testing"
)

func Test_simple(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	}

	expectedQuadtree := Quadtree{
		width:  4,
		height: 4,
		root: &node{
			topLeftX: 0,
			topLeftY: 0,
			width:    4,
			height:   4,
		},
	}
	output_topleftnode := node{
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
	output_toprightnode := node{
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
	output_bottomleftnode := node{
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
	output_bottomrightnode := node{
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
		width:  len(floorContent[0]),
		height: len(floorContent),
		root: &node{
			topLeftX:        0,
			topLeftY:        0,
			width:           4,
			height:          4,
			content:         0,
			topLeftNode:     &output_topleftnode,
			topRightNode:    &output_toprightnode,
			bottomLeftNode:  &output_bottomleftnode,
			bottomRightNode: &output_bottomrightnode,
		},
	}

	q := MakeFromArray(floorContent)
	if !reflect.DeepEqual(q, output_root) {
		t.Errorf("Cas de test MakeFromArray_ComplexArray échoué. Attendu : %v, Obtenu : %v", expectedQuadtree, q)
	}
}
