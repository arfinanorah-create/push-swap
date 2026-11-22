package piles

import (
	"reflect"
	"testing"
)

// Exemple du sujet : 2 1 3 6 8 5
func TestExempleDuSujet(t *testing.T) {
	p := New([]int{2, 1, 3, 6, 8, 5})

	p.Sa()
	if !reflect.DeepEqual(p.A, []int{1, 2, 3, 6, 8, 5}) {
		t.Fatalf("après sa : %v", p.A)
	}
	p.Pb()
	p.Pb()
	p.Pb()
	if !reflect.DeepEqual(p.A, []int{6, 8, 5}) || !reflect.DeepEqual(p.B, []int{3, 2, 1}) {
		t.Fatalf("après pb pb pb : %v %v", p.A, p.B)
	}
	p.Rb()
	if !reflect.DeepEqual(p.B, []int{2, 1, 3}) {
		t.Fatalf("après rb : %v", p.B)
	}
	p.Rrr()
	if !reflect.DeepEqual(p.A, []int{5, 6, 8}) || !reflect.DeepEqual(p.B, []int{3, 2, 1}) {
		t.Fatalf("après rrr : %v %v", p.A, p.B)
	}
	p.Pa()
	p.Pa()
	p.Pa()
	if !reflect.DeepEqual(p.A, []int{1, 2, 3, 5, 6, 8}) || len(p.B) != 0 {
		t.Fatalf("à la fin : %v %v", p.A, p.B)
	}
}

func TestSsEtRr(t *testing.T) {
	p := &Piles{A: []int{1, 2, 3}, B: []int{4, 5, 6}}
	p.Ss()
	if !reflect.DeepEqual(p.A, []int{2, 1, 3}) || !reflect.DeepEqual(p.B, []int{5, 4, 6}) {
		t.Fatalf("ss : %v %v", p.A, p.B)
	}
	p.Rr()
	if !reflect.DeepEqual(p.A, []int{1, 3, 2}) || !reflect.DeepEqual(p.B, []int{4, 6, 5}) {
		t.Fatalf("rr : %v %v", p.A, p.B)
	}
}

// Les instructions sur des piles vides ou trop petites ne doivent pas planter.
func TestPilesVides(t *testing.T) {
	p := New([]int{})
	noms := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	for _, nom := range noms {
		if err := p.Appliquer(nom); err != nil {
			t.Fatal(err)
		}
	}
	if len(p.A) != 0 || len(p.B) != 0 {
		t.Fatal("les piles auraient dû rester vides")
	}
}

func TestInstructionInconnue(t *testing.T) {
	for _, nom := range []string{"", "SA", "sa ", "xx"} {
		p := New([]int{1, 2})
		if p.Appliquer(nom) == nil {
			t.Errorf("%q aurait dû donner une erreur", nom)
		}
	}
}

func TestEstTriee(t *testing.T) {
	if !(&Piles{A: []int{1, 2, 3}}).EstTriee() {
		t.Error("1 2 3 est triée")
	}
	if (&Piles{A: []int{1, 3, 2}}).EstTriee() {
		t.Error("1 3 2 n'est pas triée")
	}
	if (&Piles{A: []int{2, 3}, B: []int{1}}).EstTriee() {
		t.Error("b n'est pas vide")
	}
}
