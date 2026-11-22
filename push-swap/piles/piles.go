// Package piles contient les deux piles (a et b) et les 11 instructions.
// C'est le SEUL endroit où les piles sont modifiées : push-swap et
// checker utilisent exactement le même code.
package piles

import "fmt"

// Piles contient nos deux piles. Le sommet d'une pile est l'indice 0.
// Operations garde, dans l'ordre, le nom de chaque instruction exécutée
// (push-swap s'en sert pour afficher le résultat).
type Piles struct {
	A          []int
	B          []int
	Operations []string
}

// New crée les piles : a reçoit les nombres, b est vide.
func New(nombres []int) *Piles {
	return &Piles{A: nombres, B: []int{}}
}

// EstTrie renvoie true si la liste est en ordre croissant.
func EstTrie(liste []int) bool {
	for i := 0; i < len(liste)-1; i++ {
		if liste[i] > liste[i+1] {
			return false
		}
	}
	return true
}

// EstTriee renvoie true si a est triée ET b est vide.
func (p *Piles) EstTriee() bool {
	return EstTrie(p.A) && len(p.B) == 0
}

// ---------- Petites fonctions qui travaillent sur UNE seule pile ----------

// echanger échange les 2 premiers éléments (rien s'il y en a moins de 2).
func echanger(pile []int) {
	if len(pile) < 2 {
		return
	}
	pile[0], pile[1] = pile[1], pile[0]
}

// tourner : le premier élément devient le dernier.
func tourner(pile []int) []int {
	if len(pile) < 2 {
		return pile
	}
	nouvelle := []int{}
	nouvelle = append(nouvelle, pile[1:]...)
	nouvelle = append(nouvelle, pile[0])
	return nouvelle
}

// tournerInverse : le dernier élément devient le premier.
func tournerInverse(pile []int) []int {
	if len(pile) < 2 {
		return pile
	}
	nouvelle := []int{pile[len(pile)-1]}
	nouvelle = append(nouvelle, pile[:len(pile)-1]...)
	return nouvelle
}

// ---------- Les 11 instructions ----------

// Sa échange les 2 premiers éléments de a.
func (p *Piles) Sa() {
	if len(p.A) < 2 {
		return
	}
	echanger(p.A)
	p.Operations = append(p.Operations, "sa")
}

// Sb échange les 2 premiers éléments de b.
func (p *Piles) Sb() {
	if len(p.B) < 2 {
		return
	}
	echanger(p.B)
	p.Operations = append(p.Operations, "sb")
}

// Ss fait sa et sb en même temps.
func (p *Piles) Ss() {
	echanger(p.A)
	echanger(p.B)
	p.Operations = append(p.Operations, "ss")
}

// Pa prend le sommet de b et le met au sommet de a.
func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

// Pb prend le sommet de a et le met au sommet de b.
func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

// Ra : le premier élément de a devient le dernier.
func (p *Piles) Ra() {
	if len(p.A) < 2 {
		return
	}
	p.A = tourner(p.A)
	p.Operations = append(p.Operations, "ra")
}

// Rb : le premier élément de b devient le dernier.
func (p *Piles) Rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = tourner(p.B)
	p.Operations = append(p.Operations, "rb")
}

// Rr fait ra et rb en même temps.
func (p *Piles) Rr() {
	p.A = tourner(p.A)
	p.B = tourner(p.B)
	p.Operations = append(p.Operations, "rr")
}

// Rra : le dernier élément de a devient le premier.
func (p *Piles) Rra() {
	if len(p.A) < 2 {
		return
	}
	p.A = tournerInverse(p.A)
	p.Operations = append(p.Operations, "rra")
}

// Rrb : le dernier élément de b devient le premier.
func (p *Piles) Rrb() {
	if len(p.B) < 2 {
		return
	}
	p.B = tournerInverse(p.B)
	p.Operations = append(p.Operations, "rrb")
}

// Rrr fait rra et rrb en même temps.
func (p *Piles) Rrr() {
	p.A = tournerInverse(p.A)
	p.B = tournerInverse(p.B)
	p.Operations = append(p.Operations, "rrr")
}

// Appliquer exécute l'instruction dont on donne le nom (ex : "sa").
// Elle renvoie une erreur si l'instruction n'existe pas.
func (p *Piles) Appliquer(nom string) error {
	switch nom {
	case "sa":
		p.Sa()
	case "sb":
		p.Sb()
	case "ss":
		p.Ss()
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "ra":
		p.Ra()
	case "rb":
		p.Rb()
	case "rr":
		p.Rr()
	case "rra":
		p.Rra()
	case "rrb":
		p.Rrb()
	case "rrr":
		p.Rrr()
	default:
		return fmt.Errorf("instruction inconnue : %s", nom)
	}
	return nil
}
