package quadtree

import (
	"reflect"
	"testing"
)

func Test_simple_make(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
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
		t.Errorf("Cas de test MakeFromArray_ComplexArray échoué. Attendu : %v, Obtenu : %v", output_root, q)
	}
}
func Test_complexe_make(t *testing.T) {
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
	q := MakeFromArray(floorContent)
	if !reflect.DeepEqual(q, output_root) {
		t.Errorf("Cas de test MakeFromArray_ComplexArray échoué. Attendu : %v, Obtenu : %v", output_root, q)
	}
}