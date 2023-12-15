package floor

import (
	"reflect"
	"testing"
)

// teste pour vérifier si le fichier est bien que le fichier est bien un fichier texte
// et donc que notre fonction retourne bien un tableau de tableau d'entier
func Test_file(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	if reflect.TypeOf(tab).Kind() == reflect.Slice {
		elemType := reflect.TypeOf(tab).Elem()
		if elemType.Kind() != reflect.Slice || elemType.Elem().Kind() != reflect.Int {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func Test_empty_alpha(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	for i := 0; i < len(tab); i++ {
		if len(tab[i]) == 0 {
			t.Fail()
		}
	}
}

// teste pour vérifier si les lignes du fichiers texte ont bien toutes le même nombre de caractères
func Test_line_size(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	size := len(tab[0])
	for i := 0; i < len(tab); i++ {
		if len(tab[i]) != size {
			t.Fail()
		}
	}
}
