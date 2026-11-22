package tri

import (
	"math/rand"
	"testing"

	"push-swap/piles"
)

// verifier rejoue les instructions et contrôle que a est triée et b vide.
func verifier(t *testing.T, nombres []int, instructions []string) {
	p := piles.New(append([]int{}, nombres...))
	for _, instruction := range instructions {
		if err := p.Appliquer(instruction); err != nil {
			t.Fatal(err)
		}
	}
	if !p.EstTriee() {
		t.Fatalf("pas trié pour %v : a=%v b=%v", nombres, p.A, p.B)
	}
}

func TestDejaTrie(t *testing.T) {
	listes := [][]int{{}, {1}, {1, 2}, {1, 2, 3, 4, 5, 6}}
	for _, liste := range listes {
		if len(Trier(liste)) != 0 {
			t.Errorf("%v : aucune instruction attendue", liste)
		}
	}
}

func TestTrois(t *testing.T) {
	// toutes les façons de mélanger 1 2 3
	listes := [][]int{{1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	for _, liste := range listes {
		instructions := Trier(liste)
		verifier(t, liste, instructions)
		if len(instructions) > 2 {
			t.Errorf("%v : %d instructions (max 2)", liste, len(instructions))
		}
	}
}

func TestPetitsEtGrandsAleatoires(t *testing.T) {
	for _, taille := range []int{2, 4, 5, 6, 7, 10, 19, 20, 21, 50, 100, 500} {
		for essai := 0; essai < 20; essai++ {
			nombres := rand.Perm(taille * 10)[:taille]
			verifier(t, nombres, Trier(nombres))
		}
	}
}

func TestNombresNegatifs(t *testing.T) {
	nombres := []int{5, -3, 0, -100, 42, 7, -8}
	verifier(t, nombres, Trier(nombres))
}

// Valeurs données par le cours (séances 4 et 5).
func TestValeursDuCours(t *testing.T) {
	cas := []struct {
		nombres []int
		lignes  int
	}{
		{[]int{5, 2, 9, 1, 7}, 10},
		{[]int{4, 3, 2, 1}, 5},
		{[]int{40, 10, 50, 20, 30}, 6},
		{[]int{-3, 10, 0, 7, -1, 2}, 17},
		{[]int{18, 4, 22, 9, 1, 14, 25, 7, 12, 3, 20, 16, 5, 24, 11, 8, 19, 2, 15, 23, 6, 13, 21, 10, 17}, 116},
	}
	for _, c := range cas {
		if obtenu := len(Trier(c.nombres)); obtenu != c.lignes {
			t.Errorf("%v : %d lignes, le cours annonce %d", c.nombres, obtenu, c.lignes)
		}
	}
}

// Exemple de l'exercice 2-3 de la séance 5 : phase 1 puis phase 2.
func TestPhasesDuCours(t *testing.T) {
	p := piles.New([]int{4, 1, 5, 0, 3, 2})
	pousserParBlocs(p, 2)
	if len(p.Operations) != 12 || p.B[0] != 3 || p.B[5] != 5 {
		t.Errorf("phase 1 : %v %v", p.Operations, p.B)
	}
	remonterLesPlusGrands(p)
	if len(p.Operations) != 21 || !p.EstTriee() {
		t.Errorf("phase 2 : %v %v %v", p.Operations, p.A, p.B)
	}
}
