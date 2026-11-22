package arguments

import (
	"reflect"
	"testing"
)

func TestNombresValides(t *testing.T) {
	obtenu, err := LireNombres([]string{"2 1 3 6 5 8"})
	if err != nil || !reflect.DeepEqual(obtenu, []int{2, 1, 3, 6, 5, 8}) {
		t.Errorf("un seul argument : %v %v", obtenu, err)
	}
	obtenu, err = LireNombres([]string{"2", "1", "-3"})
	if err != nil || !reflect.DeepEqual(obtenu, []int{2, 1, -3}) {
		t.Errorf("plusieurs arguments : %v %v", obtenu, err)
	}
}

func TestErreurs(t *testing.T) {
	// pas un nombre entier
	for _, args := range [][]string{{"0 one 2 3"}, {"1.5"}, {"99999999999999999999"}} {
		if _, err := LireNombres(args); err == nil {
			t.Errorf("%v aurait dû donner une erreur", args)
		}
	}
}

func TestDoublons(t *testing.T) {
	if VerifierDoublons([]int{1, 2, 3}) != nil {
		t.Error("pas de doublon ici")
	}
	if VerifierDoublons([]int{1, 2, 1}) == nil {
		t.Error("1 est en double")
	}
}
