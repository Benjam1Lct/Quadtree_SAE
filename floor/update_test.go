package floor 

import (
	"testing"
	"reflect"
)

func Test_size(t *testing.T) {
	content := updateFromFileFloor(camXPos, camYPos) 
	if len(content)!= Global.NumTileY {
		t.Fail()
	}
	for i:=0;i<len(content);i++ {
		if content[i]!= Global.NumTileX {
			t.Fail()
		}
	}
}

func Test_out(t *testing.T) {
	content := updateFromFileFloor(camXPos, camYPos) 
	for i:=0; i<len(content); i++ {
		for j:=0; j<len(content[i]); j++ {
			if content[i][j] == -1 {
				if camYPos+len(content)/2 <= len(f.fullContent) {
					t.Fail()
				}else if camYPos-len(content)/2 >= 0 {
					t.Fail()
				}else if camXPos+len(content[i])/2 <= len(f.fullContent[i]){
					t.Fail()
				}else if camXPos-len(content[i])/2 >= 0 {
					t.Fail()
				}
			}
		}
	}
}

