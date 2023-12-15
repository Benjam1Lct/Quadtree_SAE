package floor

import (
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func TestUpdateFromFileFloorSize(t *testing.T) {
	var f Floor
	f.Init() // You may need to initialize the floor before testing
	f.updateFromFileFloor(0, 0)

	if len(f.content) != configuration.Global.NumTileY {
		t.Fail()
	}

	for i := 0; i < len(f.content); i++ {
		if len(f.content[i]) != configuration.Global.NumTileX {
			t.Fail()
		}
	}
}

func TestUpdateFromFileFloorOut(t *testing.T) {
	var f Floor
	f.Init() // You may need to initialize the floor before testing
	f.updateFromFileFloor(0, 0)

	for i := 0; i < len(f.content); i++ {
		for j := 0; j < len(f.content[i]); j++ {
			if f.content[i][j] == -1 {
				if 0+len(f.content)/2 <= len(f.fullContent) {
					t.Fail()
				} else if 0-len(f.content)/2 >= 0 {
					t.Fail()
				} else if 0+len(f.content[i])/2 <= len(f.fullContent[i]) {
					t.Fail()
				} else if 0-len(f.content[i])/2 >= 0 {
					t.Fail()
				}
			}
		}
	}
}
